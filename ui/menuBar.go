package ui

import (
	"fmt"
	"runtime"

	"hikari/anilist"
	// "hikari/media"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type MenuBar struct {
	widgets.QMenuBar

	_ func() `constructor:"init"`

	fileMenu    *widgets.QMenu
	anilistMenu *widgets.QMenu
	toolsMenu   *widgets.QMenu

	nav  *Navigation
	user *anilist.AniList
}

func (ptr *MenuBar) SetupUi(user *anilist.AniList, nav *Navigation) {
	ptr.user = user
	ptr.nav = nav
}

func (ptr *MenuBar) init() {
	ptr.fileMenu = ptr.AddMenu2("File")
	ptr.anilistMenu = ptr.AddMenu2("AniList")
	ptr.toolsMenu = ptr.AddMenu2("Tools")

	f := ptr.fileMenu.AddAction("Exit")
	f.ConnectTriggered(func(bool) {
		core.QCoreApplication_Exit(0)
	})

	a := ptr.anilistMenu.AddAction("Refresh List")
	a.ConnectTriggered(func(bool) {
		// This should probably return an error, need to modify the method.
		ptr.nav.animeList.RefreshList()
		ptr.nav.mangaList.RefreshList()

		old := ptr.nav
		parent := ptr.nav.ParentWidget()
		parent.Layout().RemoveWidget(ptr.nav)
		ptr.nav = NewNavigation(parent, core.Qt__Widget)
		ptr.nav.SetupUi(old.animeList, old.mangaList, old.currentPage)
		parent.Layout().AddWidget(ptr.nav)
		old.sideBar.DestroyQToolBar()
		old.animeListPage.DestroyMediaListPage()
		old.mangaListPage.DestroyMediaListPage()
		old.sideBar = nil
		old.animeListPage = nil
		old.mangaListPage = nil
		old.animeList = nil
		old.mangaList = nil
		old.DestroyNavigation()
		old = nil
		runtime.GC()
	})

	t := ptr.toolsMenu.AddAction("Settings")
	t.ConnectTriggered(func(bool) {
		fmt.Println("creating settings dialog...")
		popup := NewSettingsDialog(ptr, core.Qt__Dialog)
		popup.SetupUi(ptr.user, ptr.nav.animeList, ptr.nav.mangaList)
		popup.Exec()
	})
}
