package ui

import (
	"fmt"
	"hikari/anilist"
	"hikari/lists"
	"hikari/media"
	"hikari/settings"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type MainWindow struct {
	widgets.QMainWindow

	_ func() `constructor:"init"`

	centralWidget *widgets.QWidget
	menuBar       *MenuBar
	navigation    *Navigation
	user          *anilist.AniList
}

func (ptr *MainWindow) SetupUi(settings *settings.Settings) {
	ptr.user = anilist.NewAniList(settings)
	animeList := lists.NewMediaList(ptr.user, media.AnimeType)
	mangaList := lists.NewMediaList(ptr.user, media.MangaType)

	fmt.Println("initalizing token from settings...")
	err := ptr.user.LoadToken()
	if err != nil {
		fmt.Println("could not initialize token from settings...")
	} else {
		fmt.Println("successfully initialized token from settings...")
		err = ptr.user.QueryUserSettings()
		if err != nil {
			fmt.Println("error returned from QueryUserSettings():", err.Error())
		} else {
			fmt.Println("successfully loaded user settings...")
			animeList.RefreshList()
			mangaList.RefreshList()
		}
	}

	ptr.navigation.SetupUi(animeList, mangaList, "curMedia")
	ptr.menuBar.SetupUi(ptr.user, ptr.navigation)
}

func (ptr *MainWindow) init() {
	ptr.SetMinimumSize2(250, 200)
	ptr.SetWindowTitle("Hikari")

	ptr.centralWidget = widgets.NewQWidget(ptr, 0)
	layout := widgets.NewQVBoxLayout()
	layout.SetContentsMargins(0, 0, 0, 0)
	ptr.centralWidget.SetLayout(layout)
	ptr.centralWidget.SetAutoFillBackground(true)
	ptr.SetCentralWidget(ptr.centralWidget)

	ptr.navigation = NewNavigation(ptr.centralWidget, core.Qt__Widget)

	ptr.menuBar = NewMenuBar(ptr.centralWidget)

	ptr.centralWidget.Layout().AddWidget(ptr.menuBar)
	ptr.centralWidget.Layout().AddWidget(ptr.navigation)
}
