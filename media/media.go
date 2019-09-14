package media

import (
	"encoding/json"
)

type Error struct {
	Err string
}

func (m Error) Error() string {
	return m.Err
}

type Media struct {
	ID    int `json:"id"`
	Title struct {
		Romaji        string `json:"romaji"`
		English       string `json:"english"`
		Native        string `json:"native"`
		UserPreferred string `json:"userPreferred"`
	} `json:"title"`
	Synonyms        []string `json:"synonyms"`
	Type            string   `json:"type"`
	Format          string   `json:"format"`
	Status          string   `json:"status"`
	Description     string   `json:"description"`
	Season          string   `json:"season"`
	Episodes        int      `json:"episodes"`
	Chapters        int      `json:"chapters"`
	Volumes         int      `json:"volumes"`
	CountryOfOrigin string   `json:"countryOfOrigin"`
	IsLicensed      bool     `json:"isLicensed"`
	Source          string   `json:"source"`
	CoverImage      struct {
		Large string `json:"large"`
	} `json:"coverImage"`
	Genres       []string `json:"genres"`
	AverageScore int      `json:"averageScore"`
	MeanScore    int      `json:"meanScore"`
	Studios      struct {
		Edges []struct {
			IsMain bool `json:"isMain"`
			Node   struct {
				Name string `json:"name"`
			} `json:"node"`
		} `json:"edges"`
	} `json:"studios"`
	IsLocked bool `json:"isLocked"`
}

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type List struct {
	Name         string   `json:"name"`
	IsCustomList bool     `json:"isCustomList"`
	Status       string   `json:"status"`
	Entries      []*Entry `json:"entries"`
}

func (ptr *List) SortEntries() {
	for i := range ptr.Entries {
		swap := false
		for j := 0; j < len(ptr.Entries)-i-1; j++ {
			if ptr.Entries[j].Media.Title.UserPreferred > ptr.Entries[j+1].Media.Title.UserPreferred {
				temp := ptr.Entries[j]
				ptr.Entries[j] = ptr.Entries[j+1]
				ptr.Entries[j+1] = temp
				swap = true
			}
		}
		if swap == false {
			break
		}
	}
}

type CustomLists struct {
	Map map[string]bool
}

func (ptr *CustomLists) UnmarshalJSON(data []byte) error {
	ptr.Map = make(map[string]bool)
	err := json.Unmarshal(data, &ptr.Map)
	if err != nil {
		return err
	}
	return nil
}

type Entry struct {
	Title                 string
	ListEntryID           int         `json:"id"`
	Status                string      `json:"status"`
	Score                 float64     `json:"score"`
	Progress              int         `json:"progress"`
	ProgressVolumes       int         `json:"progressVolumes"`
	Repeat                int         `json:"repeat"`
	Priority              int         `json:"priority"`
	Private               bool        `json:"private"`
	Notes                 string      `json:"notes"`
	HiddenFromStatusLists bool        `json:"hiddenFromStatusLists"`
	CustomLists           CustomLists `json:"customLists"`
	StartedAt             Date        `json:"startedAt"`
	CompletedAt           Date        `json:"completedAt"`
	Media                 Media       `json:"media"`
}

func (ptr *Entry) Update() {

}

// type Update Entry

// func (a *Anime) FindListEntry() Entry {
// 	return &Anime{}
// }

// func (a *Anime) FindAnilistId() int {
// 	return 0
// }

// func (a *Anime) SetTitle(s string) {
// 	a.Title = s
// }

// func (a *Anime) GetTitle() string {
// 	return a.Title
// }

// func (a *Anime) Update() {

// }

// func (m *Manga) FindListEntry() Entry {
// 	return &Manga{}
// }

// func (m *Manga) FindAnilistId() int {
// 	return 0
// }

// func (m *Manga) GetTitle() string {
// 	return m.Title
// }

// func (m *Manga) SetTitle(s string) {
// 	m.Title = s
// }

// func (m *Manga) Update() {

// }

const (
	AnimeType = "ANIME"
	MangaType = "MANGA"
)
