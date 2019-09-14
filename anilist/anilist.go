package anilist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/browser"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"

	"hikari/media"
	"hikari/settings"
)

// add rate limiting. Maybe add requestQueue data structure
// and give the anilist struct a value of that type
const (
	maxRequestsPerMin = 90
)

type auth struct {
	config   *oauth2.Config
	finished chan bool
	state    string
	settings *settings.Settings
	Token    *oauth2.Token
}

func newAuth(settings *settings.Settings) *auth {
	cfg := &oauth2.Config{
		ClientID: "2355",
		Endpoint: oauth2.Endpoint{
			AuthURL: "https://anilist.co/api/v2/oauth/authorize",
		},
		RedirectURL: "http://localhost:8080/callback",
	}

	rand.Seed(time.Now().UnixNano())
	const alnum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRZTUVWXYZ1234567890"
	b := make([]byte, 15)
	for i := range b {
		b[i] = alnum[rand.Intn(len(alnum))]
	}
	state := string(b)

	auth := &auth{
		config:   cfg,
		finished: make(chan bool),
		state:    state,
		settings: settings,
	}

	http.HandleFunc("/", auth.handleLogin)
	http.HandleFunc("/callback", auth.handleCallback)

	return auth
}

func (ptr *auth) startHTTP() *http.Server {
	srv := &http.Server{Addr: ":8080"}
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed && err != nil {
			// might not run b/c main wont wait
			log.Fatalf("ListenAndServe(): %s", err.Error())
		}
	}()
	return srv
}

func (ptr *auth) handleLogin(w http.ResponseWriter, r *http.Request) {
	url := ptr.config.AuthCodeURL(ptr.state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (ptr *auth) handleCallback(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != ptr.state {
		fmt.Println("State is not valid")
		http.Redirect(w, r, "/oopsie", http.StatusTemporaryRedirect)
		ptr.finished <- false
		return
	}

	if r.FormValue("code") == "" {
		fmt.Println("no code returned")
		ptr.finished <- false
		return
	}

	authCode := &struct {
		Code string `json:"code"`
	}{
		Code: r.FormValue("code"),
	}
	codeJSON, _ := json.Marshal(authCode)

	req, err := http.NewRequest("POST", "https://auth.hikari.page", bytes.NewBuffer(codeJSON))
	if err != nil {
		fmt.Println("request err: ", err.Error())
		ptr.finished <- false
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ptr.finished <- false
		panic(err)
	}
	defer resp.Body.Close()
	// fmt.Println("status:", resp.Status)

	tokenJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("resp read err: ", err.Error())
		ptr.finished <- false
		return
	}

	token := &oauth2.Token{}
	err = json.Unmarshal(tokenJSON, token)
	if err != nil {
		fmt.Println("unmarshal err: ", err.Error())
		ptr.finished <- false
		return
	}

	ptr.Token = token
	fmt.Println(ptr.Token)

	b, _ := json.Marshal(ptr.Token)
	ptr.settings.Set("tokenData", string(b))
	ptr.finished <- true
}

type AniList struct {
	Auth             *auth
	Settings         *settings.Settings
	ID               int
	Username         string
	ScoreFormat      string
	AnimeCustomLists []string
	MangaCustomLists []string
}

func NewAniList(settings *settings.Settings) *AniList {
	return &AniList{
		Auth:     newAuth(settings),
		Settings: settings,
	}
}

func (ptr *AniList) Authorize() chan bool {
	ptr.ForgetToken()
	finished := make(chan bool)
	if ptr.IsAuthenticated() == false {
		return ptr.requestAuth(finished)
	}
	finished <- false
	return finished
}

func (ptr *AniList) requestAuth(status chan bool) chan bool {
	srv := ptr.Auth.startHTTP()
	_ = browser.OpenURL("http://localhost:8080/")
	go func() {
		if <-ptr.Auth.finished == true {
			status <- true
		} else {
			status <- false
		}
		if err := srv.Shutdown(context.TODO()); err != nil {
			panic(err)
		}
		// close(finished)
	}()
	// go func() {
	// 	time.Sleep(time.Duration(15) * time.Second)
	// 	if // channel is closed? or has sent something?
	// 	status <- false
	// }()
	return status
}

func (ptr *AniList) SetupUser() error {
	err := ptr.QueryUserSettings()
	if err != nil {
		return fmt.Errorf("SetupUser() QueryUserSettings() error: %v", err.Error())
	}
	return nil
}

func (ptr *AniList) ForgetToken() {
	ptr.Auth.Token = nil
	ptr.ID = 0
	ptr.Username = ""
	ptr.ScoreFormat = ""
	ptr.AnimeCustomLists = nil
	ptr.MangaCustomLists = nil
	ptr.Settings.Erase("tokenData")
}

func (ptr *AniList) LoadToken() error {
	tokenJSON := ptr.Settings.Get("tokenData")
	if tokenJSON != "" {
		token := &oauth2.Token{}
		_ = json.Unmarshal([]byte(tokenJSON), token)
		ptr.Auth.Token = token
		fmt.Println(ptr.Auth.Token)
		// id, _ := ptr.findUserID()
		// ptr.ID = id
		// fmt.Println(ptr.ID)
		return nil
	}
	return fmt.Errorf("could not load token from settings")
}

func (ptr *AniList) IsAuthenticated() bool {
	if ptr.Auth.Token != nil {
		fmt.Println("authenticated, token isn't nil")
		return true
	}
	fmt.Println("token nil, calling LoadToken()...")
	err := ptr.LoadToken()
	if err == nil {
		if ptr.IsExpired() || ptr.Auth.Token.AccessToken == "" {
			fmt.Println("token expired")
			ptr.ForgetToken()
			return false
		}
		return true
	}

	fmt.Println("token is nil and couldn't load it from settings")
	return false
	// the only way i can think of fixing bad auth token if they revoke it anilist side (while the program is running) is if i do a test query right here and check if its valid
}

func (ptr *AniList) IsExpired() bool {
	if ptr.Auth.Token != nil {
		if ptr.Auth.Token.AccessToken == "" {
			return true
		}

		expYear, expMonth, expDay := ptr.Auth.Token.Expiry.Date()
		year, month, day := time.Now().Date()
		if (expYear >= year) && (expMonth >= month) && (expDay >= day) {
			expHour, expMin, expSec := ptr.Auth.Token.Expiry.Clock()
			hour, min, sec := time.Now().Clock()
			if (expHour >= hour) && (expMin >= min) && (expSec >= sec) {
				return true
			}
		}
	}
	return false
}

func (ptr *AniList) graphqlQuery(fileName string, vars interface{}) ([]byte, error) {
	if ptr.Auth.Token != nil {
		b, err := ioutil.ReadFile(fileName)
		if err != nil {
			return nil, fmt.Errorf("error reading gql file: %v", err.Error())
		}
		q := string(b)

		var v []byte
		var query string
		q = strings.ReplaceAll(q, "\r\n", " ")
		q = strings.ReplaceAll(q, "\n", "")
		q = strings.ReplaceAll(q, "  ", "")

		if vars != nil {
			v, _ = json.Marshal(vars)
			query = fmt.Sprintf(`{"query": "%v", "variables": %v}`, q, string(v))
		} else {
			query = fmt.Sprintf(`{"query": "%v"}`, q)
		}

		// val := url.Values{"query": {query}}
		req, _ := http.NewRequest("POST", "https://graphql.anilist.co", bytes.NewBufferString(query))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("%v %v", ptr.Auth.Token.TokenType, ptr.Auth.Token.AccessToken))
		// req.PostForm = val

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return []byte{}, fmt.Errorf("query request error: %v", err.Error())
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			if resp.StatusCode == 429 {
				s := resp.Header.Get("Retry-After")
				sec, _ := strconv.Atoi(s)
				fmt.Printf("rate limit for %v seconds\n", sec)
				ch := ptr.rateLimit(sec)
				if <-ch == true {
					close(ch)
					return ptr.graphqlQuery(q, vars)
				}
			} else {
				return nil, fmt.Errorf("graphqlQuery bad status: %v", resp.Status)
			}
		}

		respBody, _ := ioutil.ReadAll(resp.Body)

		if strings.Contains(string(respBody), `"message":"Invalid token"`) {
			ptr.ForgetToken()
			return nil, fmt.Errorf("graphql response error: Invalid token sent with request")
		}
		if strings.Contains(string(respBody), `{"errors":`) {
			return nil, fmt.Errorf("graphql response error: %v", string(respBody))
		}

		return respBody, nil
	}
	return nil, fmt.Errorf("graphqlQuery error: Auth token is nil")
}

func (ptr *AniList) rateLimit(seconds int) chan bool {
	ch := make(chan bool)
	fmt.Println("rate limiting...")
	go func(ch chan bool) {
		time.Sleep(time.Duration(seconds) * time.Second)
		ch <- true
	}(ch)
	return ch
}

func (ptr *AniList) QueryUserSettings() error {
	if ptr.Auth.Token != nil {
		resp, err := ptr.graphqlQuery("../../query/UserSettings.gql", nil) //, vars)
		if err != nil {
			fmt.Printf("graphql err: %v", err.Error())
			return err
		}

		info := struct {
			Data struct {
				Viewer struct {
					ID               int    `json:"id"`
					Name             string `json:"name"`
					MediaListOptions struct {
						ScoreFormat string `json:"scoreFormat"`
						AnimeList   struct {
							CustomLists []string `json:"customLists"`
						} `json:"animeList"`
						MangaList struct {
							CustomLists []string `json:"customLists"`
						} `json:"mangaList"`
					} `json:"mediaListOptions"`
				} `json:"Viewer"`
			} `json:"data"`
		}{}

		fmt.Println(string(resp))
		err = json.Unmarshal(resp, &info)
		if err != nil {
			fmt.Println(err.Error())
			return fmt.Errorf("QueryUserSettings() unmarshal err: %v", err.Error())
		}

		fmt.Println(info)
		ptr.ID = info.Data.Viewer.ID
		ptr.Username = info.Data.Viewer.Name
		ptr.ScoreFormat = info.Data.Viewer.MediaListOptions.ScoreFormat
		ptr.AnimeCustomLists = info.Data.Viewer.MediaListOptions.AnimeList.CustomLists
		ptr.MangaCustomLists = info.Data.Viewer.MediaListOptions.MangaList.CustomLists

		fmt.Println(ptr.ID)

		return nil
	}
	return fmt.Errorf("QueryUserSettings(): Auth token is nil")
}

func (ptr *AniList) QueryMediaListCollection(mediaType string) []media.List {
	if ptr.IsAuthenticated() && (ptr.ID > 0) && (mediaType == "ANIME" || mediaType == "MANGA") {
		//vars :=

		resp, err := ptr.graphqlQuery("../../query/MediaListCollection.gql", struct {
			ID        int    `json:"id"`
			MediaType string `json:"type"`
		}{
			ID:        ptr.ID,
			MediaType: mediaType,
		})
		if err != nil {
			fmt.Println("QueryMLC() graphql err:", err.Error())
			return nil
		}

		wrap := struct {
			Data struct {
				MediaListCollection struct {
					Lists []media.List `json:"lists"`
				} `json:"MediaListCollection"`
			} `json:"data"`
		}{}

		err = json.Unmarshal(resp, &wrap)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		if wrap.Data.MediaListCollection.Lists == nil {
			fmt.Println("nil media list collection")
			return nil
		}

		return wrap.Data.MediaListCollection.Lists

		// list := make([]*media.Entry, 0)
		// if mediaType == "ANIME" {
		// 	for i := 0; i < 100; i++ {
		// 		list = append(list, &media.Entry{Title: "title", Progress: 5, Score: 5.5, Status: "Completed"})
		// 	}
		// } else if mediaType == "MANGA" {
		// 	for i := 0; i < 100; i++ {
		// 		list = append(list, &media.Entry{Title: "title", ProgressVolumes: 6, Progress: 5, Score: 5.5, Status: "Completed"})
		// 	}
		// }
		// return list
	}
	return nil
}
