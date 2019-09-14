package crunchyroll

import (
	"hikari/media"
	"log"
	"strconv"
	"strings"
)

func Parse(w string) (*media.Entry, error) {
	if strings.Contains(w, "| Crunchyroll Store") {
		return nil, media.Error{"Crunchyroll page, but not episode page"}
	}

	sections := strings.Split(w, ", ")

	var title string
	var ep int
	var err error
	if len(sections) == 3 {
		titleAndEp := sections[0]
		titleAndEpSlice := strings.Split(titleAndEp, " Episode ")

		title = titleAndEpSlice[0]
		/* handles things like "Title: arc" --- but messes up things like
		the Magus Bride OVA since the ": arc" is actually the real title.
		Probably gonna have to handle this when searching user list
		if strings.Contains(title, ": ") {
			titleSlice := strings.Split(title, ": ")
			title = titleSlice[0]
		}*/

		ep, err = strconv.Atoi(titleAndEpSlice[1])
		if err != nil {
			log.Fatalln(err)
		}

		return &media.Entry{
			Title:    title,
			Progress: ep,
			Media: media.Media{
				Type: media.AnimeType,
			},
		}, nil
	}
	return nil, media.Error{"Crunchyroll page, but not episode page"}
}
