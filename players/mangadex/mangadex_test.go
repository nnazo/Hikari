package mangadex

// import (
// 	"hikari/media"
// 	"hikari/players"
// 	"testing"
// )
//
// func TestParse(t *testing.T) {
// 	mdTest := []struct {
// 		WTitle   players.WindowTitle
// 		Expected *media.Manga
// 	}{
// 		{"Home - MangaDex - Mozilla Firefox", nil},
// 		{"Follows - MangaDex - Mozilla Firefox", nil},
// 		{"Updates - MangaDex - Mozilla Firefox", nil},
// 		{"Manga titles - MangaDex - Mozilla Firefox", nil},
// 		{"Search - MangaDex - Mozilla Firefox", nil},
// 		{"Featured - MangaDex - Mozilla Firefox", nil},
// 		{"Manga new - MangaDex - Mozilla Firefox", nil},
// 		{"Forums - MangaDex - Mozilla Firefox", nil},
// 		{"(Forum) - MangaDex - Mozilla Firefox", nil},
// 		{"(Thread) - MangaDex - Mozilla Firefox", nil},
// 		{"Groups - MangaDex - Mozilla Firefox", nil},
// 		{"Users - MangaDex - Mozilla Firefox", nil},
// 		{"Stats - MangaDex - Mozilla Firefox", nil},
// 		{"Rules - MangaDex - Mozilla Firefox", nil},
// 		{"About - MangaDex - Mozilla Firefox", nil},
// 		{"Changelog - MangaDex - Mozilla Firefox", nil},
// 		{"(User) - MangaDex - Mozilla Firefox", nil},
// 		{"(Group) - MangaDex - Mozilla Firefox", nil},
// 		{"History - MangaDex - Mozilla Firefox", nil},
// 		{"Settings - MangaDex - Mozilla Firefox", nil},
// 		{"MDList - MangaDex - Mozilla Firefox", nil},
// 		{"Social - MangaDex - Mozilla Firefox", nil},
// 		{"Messages - MangaDex - Mozilla Firefox", nil},
// 		{"Support - MangaDex - Mozilla Firefox", nil},
// 		{"MangaDex (@MangaDex) | Twitter - Mozilla Firefox", nil},
// 		{"mangadex: a scanlator friendly manga hosting site - Mozilla Firefox", nil},
// 		{"MangaDex - Mozilla Firefox", nil},
// 		{
// 			"Do Chokkyuu Kareshi x Kanojo - Vol. 3 Ch. 14 The First Step - MangaDex - Mozilla Firefox",
// 			&media.Manga {
// 				Title: "Do Chokkyuu Kareshi x Kanojo",
// 				Volumes: 3,
// 				Chapters: 14,
// 			},
// 		},
// 		{
// 			"Edens Zero - Ch. 44 The Palace of Knowledge - MangaDex - Mozilla Firefox",
// 			&media.Manga{
// 				Title: "Edens Zero",
// 				Volumes: 0,
// 				Chapters: 44,
// 			},
// 		},
// 		{
// 			"Hasumi-kun to Hasumi-san - Oneshot - MangaDex - Mozilla Firefox",
// 			&media.Manga{
// 				Title: "Hasumi-kun to Hasumi-san",
// 				Volumes: 0,
// 				Chapters: 1,
// 			},
// 		},
// 	}
//
// 	for _, v := range mdTest {
// 		md, _ := v.WTitle.GetMedia()
//
// 		if md == nil {
// 			if v.Expected != nil {
// 				t.Error("Expected", v.Expected, "Got", md)
// 			}
// 		} else if md != nil {
// 			if v.Expected == nil {
// 				t.Error("Expected", v.Expected, "Got", md)
// 			}
// 		} else {
// 			if md.(*media.Manga).Title != v.Expected.Title {
// 				t.Error("Expected", v.Expected.Title, "Got", md.(*media.Manga).Title)
// 			} else if md.(*media.Manga).Volumes != v.Expected.Volumes {
// 				t.Error("Expected", v.Expected.Volumes, "Got", md.(*media.Manga).Volumes)
// 			} else if md.(*media.Manga).Chapters != v.Expected.Chapters {
// 				t.Error("Expected", v.Expected.Chapters, "Got", md.(*media.Manga).Chapters)
// 			}
// 		}
// 	}
// }
