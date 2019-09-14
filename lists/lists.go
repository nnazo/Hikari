package lists

import (
	"fmt"
	"hikari/anilist"
	"hikari/media"
	"strings"
)

type MediaList struct {
	Lists []media.List
	Type  string
	user  *anilist.AniList
}

func NewMediaList(user *anilist.AniList, mediaType string) *MediaList {
	lists := make([]media.List, 0)
	var defaultLists []string
	if mediaType == media.AnimeType {
		defaultLists = []string{"Watching", "Completed", "Paused", "Dropped", "Planning"}
	} else if mediaType == media.MangaType {
		defaultLists = []string{"Reading", "Completed", "Paused", "Dropped", "Planning"}
	}

	for _, v := range defaultLists {
		lists = append(lists, media.List{
			Name:         v,
			IsCustomList: false,
			Status:       strings.ToUpper(v),
			Entries:      make([]*media.Entry, 0),
		})
	}

	list := &MediaList{
		Lists: lists,
		Type:  mediaType,
		user:  user,
	}

	return list
}

func (ptr *MediaList) RefreshList() {
	if ptr.user != nil {
		lists := ptr.user.QueryMediaListCollection(ptr.Type)
		if lists != nil {
			ptr.Lists = lists
			ptr.SortLists()
			for i := range ptr.Lists {
				ptr.Lists[i].SortEntries()
			}
		}
	}
}

func (ptr *MediaList) SortLists() {
	fmt.Println("sorting lists...")
	for i := range ptr.Lists {
		swap := false
		for j := 0; j < len(ptr.Lists)-i-1; j++ {
			if ptr.Lists[j].IsCustomList != ptr.Lists[j+1].IsCustomList {
				if !ptr.Lists[j].IsCustomList {
					temp := ptr.Lists[j]
					ptr.Lists[j] = ptr.Lists[j+1]
					ptr.Lists[j+1] = temp
					swap = true
				}
			} else {
				if ptr.Lists[j].Name < ptr.Lists[j+1].Name {
					temp := ptr.Lists[j]
					ptr.Lists[j] = ptr.Lists[j+1]
					ptr.Lists[j+1] = temp
					swap = true
				}
			}
		}
		if swap == false {
			break
		}
	}
}
