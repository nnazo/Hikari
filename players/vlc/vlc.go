package vlc

import (
	"hikari/media"
	"strconv"
	"strings"
)

func Parse(w string) (*media.Entry, error) {
	sections := strings.SplitAfter(string(w), "] ")
	if len(sections) > 1 {
		s2 := strings.Split(sections[1], " - ")
		title := s2[0]

		var epstr string
		if strings.Contains(s2[1], " [") {
			epstr = strings.Split(s2[1], " [")[0]
		} else if strings.Contains(s2[1], " (") {
			epstr = strings.Split(s2[1], " (")[0]
		}

		ep, err := strconv.Atoi(epstr)
		if err != nil {
			ep = -1
		}

		return &media.Entry{
			Title:    title,
			Progress: ep,
			Media: media.Media{
				Type: media.AnimeType,
			},
		}, nil
	}
	return nil, media.Error{"VLC media player is not playing anything"}
}
