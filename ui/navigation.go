package ui

import (
	"hikari/lists"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type Navigation struct {
	widgets.QWidget
	_       func() `constructor:"init"`
	sideBar *widgets.QToolBar

	animeList     *lists.MediaList
	mangaList     *lists.MediaList
	animeListPage *MediaListPage
	mangaListPage *MediaListPage
	currentPage   string
}

func (ptr *Navigation) SetupUi(animeList, mangaList *lists.MediaList, currentPage string) {
	ptr.currentPage = currentPage
	ptr.animeList = animeList
	ptr.mangaList = mangaList
	ptr.animeListPage.SetupUi(animeList)
	ptr.mangaListPage.SetupUi(mangaList)
}

func (ptr *Navigation) init() {
	ptr.SetLayout(widgets.NewQHBoxLayout())
	ptr.Layout().SetContentsMargins(0, 0, 0, 0)

	ptr.sideBar = widgets.NewQToolBar("Navigation", ptr)
	ptr.sideBar.SetOrientation(core.Qt__Vertical)
	// bar.SetToolButtonStyle(core.Qt__ToolButtonTextBesideIcon)
	anime := ptr.sideBar.AddAction("Anime List")
	manga := ptr.sideBar.AddAction("Manga List")
	current := ptr.sideBar.AddAction("Current Media")
	ptr.Layout().AddWidget(ptr.sideBar)

	ptr.animeListPage = NewMediaListPage(ptr, core.Qt__Widget) // ptr.animeList,
	ptr.animeListPage.SetLayout(widgets.NewQVBoxLayout())
	ptr.animeListPage.Layout().SetContentsMargins(0, 0, 0, 0)
	ptr.Layout().AddWidget(ptr.animeListPage)
	ptr.animeListPage.Hide()

	ptr.mangaListPage = NewMediaListPage(ptr, core.Qt__Widget) // ptr.mangaList,
	ptr.mangaListPage.SetLayout(widgets.NewQVBoxLayout())
	ptr.mangaListPage.Layout().SetContentsMargins(0, 0, 0, 0)
	ptr.Layout().AddWidget(ptr.mangaListPage)
	ptr.mangaListPage.Hide()

	mediaPage := NewCurrentMediaPage(ptr)
	ptr.Layout().AddWidget(mediaPage)
	mediaPage.Hide()

	// ptr.sideBar.SetToolTipDuration(0)
	// anime.ConnectHover(func() {
	// 	ptr.sideBar.SetCursor(gui.NewQCursor2(core.Qt__PointingHandCursor))
	// })

	anime.ConnectTriggered(func(bool) {
		mediaPage.Hide()
		ptr.mangaListPage.Hide()
		ptr.animeListPage.Show()
		ptr.currentPage = "animeList"
	})
	manga.ConnectTriggered(func(bool) {
		mediaPage.Hide()
		ptr.animeListPage.Hide()
		ptr.mangaListPage.Show()
		ptr.currentPage = "mangaList"
	})
	current.ConnectTriggered(func(bool) {
		ptr.animeListPage.Hide()
		ptr.mangaListPage.Hide()
		mediaPage.Show()
		ptr.currentPage = "curMedia"
	})

	switch ptr.currentPage {
	case "animeList":
		anime.Trigger()
	case "mangaList":
		manga.Trigger()
	case "curMedia":
		current.Trigger()
	default:
		anime.Trigger()
	}
}
