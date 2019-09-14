package vlc

// import (
// 	"hikari/media"
// 	"hikari/players"
// 	"testing"
// )
//
// func TestParse(t *testing.T) {
// 	vlcTest := []struct {
// 		WTitle players.WindowTitle
// 		Expected *media.Anime
// 	}{
// 		{
// 			"VLC media player",
// 			nil,
// 		},
// 		{
// 			"[Mushin] Hajimete no Gal - 11 [OVA] My first school festival (BD 1280x720 x264 AAC).mkv - VLC media player",
// 			&media.Anime {
// 				Title: "Hajimete no Gal",
// 				Episodes: 11,
// 			},
// 		},
// 		{
// 			"[Edo] Zoku Owarimonogatari - 002 (1080).mkv - VLC media player",
// 			&media.Anime {
// 				Title: "Zoku Owarimonogatari",
// 				Episodes: 2,
// 			},
// 		},
// 		{
// 			"[35mm] Liz to Aoi Tori - Liz and the Blue Bird [1080p] [562EF58C].mkv - VLC media player",
// 			&media.Anime {
// 				Title: "Liz to Aoi Tori",
// 				Episodes: -1,
// 			},
// 		},
// 	}
//
// 	for _, v := range vlcTest {
// 		cr, _ := v.WTitle.GetMedia()
//
// 		if cr != nil {
// 			if v.Expected == nil {
// 				t.Error("Expected", v.Expected, "Got", cr)
// 			}
// 		} else if cr == nil {
// 			if v.Expected != nil {
// 				t.Error("Expected", v.Expected, "Got", cr)
// 			}
// 		} else {
// 			if cr.(*media.Anime).Title != v.Expected.Title {
// 				t.Error("Expected", v.Expected.Title, "Got", cr.(*media.Anime).Title)
// 			} else if cr.(*media.Anime).Episodes != v.Expected.Episodes {
// 				t.Error("Expected", v.Expected.Episodes, "Got", cr.(*media.Anime).Episodes)
// 			}
// 		}
// 	}
// }
