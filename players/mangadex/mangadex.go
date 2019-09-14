package mangadex

import (
	"hikari/media"
	"strconv"
	"strings"
)

func Parse(w string) (*media.Entry, error) {
	sections := strings.Split(string(w), " - ")
	if len(sections) == 4 && sections[2] == "MangaDex" {
		title := sections[0]
		chVol := strings.Split(sections[1], " ")

		var vol int
		var ch int
		for i, v := range chVol {
			// Assuming the slice has "Vol. or "Ch." there is a number following it in the slice
			if v == "Vol." {
				vol, _ = strconv.Atoi(chVol[i+1])
			} else if v == "Ch." {
				ch, _ = strconv.Atoi(chVol[i+1])
			} else if v == "Oneshot" {
				ch = 1
			}
		}

		return &media.Entry{
			Title:           title,
			ProgressVolumes: vol,
			Progress:        ch,
			Media: media.Media{
				Type: media.MangaType,
			},
		}, nil
	}
	return nil, media.Error{"MangaDex page, but not chapter page"}
}
