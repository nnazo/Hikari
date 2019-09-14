package players

import (
	"fmt"
	"hikari/media"
	"hikari/players/crunchyroll"
	"hikari/players/mangadex"
	"hikari/players/vlc"
	"hikari/window"
	"strings"
)

type supportedMedia struct {
	List []mediaPlayer
}

type mediaPlayer struct {
	Name string
}

var SupportedMedia supportedMedia = supportedMedia{
	List: []mediaPlayer{
		{
			Name: "Crunchyroll",
		},
		{
			Name: "VLC media player",
		},
		{
			Name: "MangaDex",
		},
	},
}

type WindowTitle string

func (w WindowTitle) isMediaPlayer() (mediaPlayer, bool) {
	for _, v := range SupportedMedia.List {
		if strings.Contains(string(w), v.Name) {
			return v, true
		}
	}
	return mediaPlayer{}, false
}

func (w WindowTitle) GetMedia() (*media.Entry, error) {
	// !strings.Contains fixes ubuntu "exit status 1" being a valid media title
	if mp, isMp := w.isMediaPlayer(); isMp && !strings.Contains(string(w), "exit status") {
		// process window title for name of media and ep # or ch/vol#
		switch mp.Name {
		case "Crunchyroll":
			return crunchyroll.Parse(string(w))
		case "VLC media player":
			return vlc.Parse(string(w))
		case "MangaDex":
			return mangadex.Parse(string(w))
		}
	}
	return nil, media.Error{"Media player not supported"}
}

func GetWindowTitles() []string {
	return window.GetWindowTitles()
}

func GetPlayingMedia() (string, *media.Entry) {
	titles := GetWindowTitles()
	for _, title := range titles {
		entry, err := WindowTitle(title).GetMedia()
		if err != nil {
			fmt.Errorf("Error returned: %v", err.Error())
		}
		if entry != nil {
			var s string
			switch entry.Media.Type {
			case media.AnimeType:
				s = fmt.Sprintf("%v - Episode %v", entry.Title, entry.Progress)
			case media.MangaType:
				s = fmt.Sprintf("%v - Vol. %v, Ch. %v", entry.Title, entry.ProgressVolumes, entry.Progress)
			}
			return s, entry
		}
	}
	return "None", nil
}
