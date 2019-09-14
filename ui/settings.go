package ui

import (
	"fmt"
	"hikari/anilist"
	"hikari/lists"
	"runtime"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type SettingsDialog struct {
	widgets.QDialog

	_ func() `constructor:"init"`

	content *widgets.QWidget
	bar     *widgets.QToolBar
	general *widgets.QAction
	al      *widgets.QAction

	appTab *widgets.QWidget
	gen    *widgets.QTabWidget
	theme  *widgets.QCheckBox

	alTab    *widgets.QWidget
	account  *widgets.QTabWidget
	username *widgets.QLabel
	loginBtn *widgets.QPushButton

	btnWidget *widgets.QWidget
	okBtn     *widgets.QPushButton
	cancelBtn *widgets.QPushButton

	user      *anilist.AniList
	animeList *lists.MediaList
	mangaList *lists.MediaList

	changes map[string]string
}

func (ptr *SettingsDialog) SetupUi(user *anilist.AniList, animeList, mangaList *lists.MediaList) {
	ptr.user = user
	ptr.animeList = animeList
	ptr.mangaList = mangaList

	t := ptr.user.Settings.Get("theme")
	if t == "custom" {
		ptr.theme.SetChecked(true)
	} else {
		ptr.theme.SetChecked(false)
	}
	// user has to restart application for changes
	ptr.theme.ConnectClicked(func(bool) {
		s := ptr.user.Settings.Get("theme")
		if s == "native" {
			ptr.changes["theme"] = "custom"
		} else {
			ptr.changes["theme"] = "native"
		}
	})

	fmt.Println("setting up auth/revoke button...")
	if ptr.user.IsAuthenticated() {
		err := ptr.user.SetupUser()
		if err != nil {
			fmt.Println("could not set up user:")
		}
		ptr.loginBtn.SetText("Revoke")
		// username.SetText(a.QueryUsername())
		ptr.username.SetText(ptr.user.Username)
	} else {
		ptr.loginBtn.SetText("Authorize")
	}
	fmt.Println("done setting up auth/revoke button...")
	ptr.loginBtn.ConnectPressed(func() {
		if ptr.loginBtn.Text() == "Authorize" {
			ch := ptr.user.Authorize()
			if <-ch == true {
				err := ptr.user.SetupUser()
				if err != nil {
					fmt.Println("could not set up user:")
					ptr.loginBtn.SetText("Authorize")
					ptr.username.SetText("User setup failed")
				} else {
					ptr.loginBtn.SetText("Revoke")
					ptr.username.SetText(ptr.user.Username)

					ptr.animeList.RefreshList()
					ptr.mangaList.RefreshList()
				}
			} else {
				ptr.username.SetText("Authorization failed, try again")
			}
		} else {
			ptr.user.ForgetToken()
			ptr.loginBtn.SetText("Authorize")
			ptr.username.SetText("")
		}
	})
}

func (ptr *SettingsDialog) init() {
	ptr.changes = make(map[string]string)

	// Main layout is vertical box
	vl := widgets.NewQVBoxLayout()
	vl.SetContentsMargins(0, 0, 0, 0)
	ptr.SetWindowTitle("Settings")
	ptr.SetLayout(vl)
	ptr.SetMinimumHeight(300)

	// Add widget with horiz layout for main tabs + content
	ptr.content = widgets.NewQWidget(nil, core.Qt__Widget)
	ptr.content.SetLayout(widgets.NewQHBoxLayout())
	ptr.content.Layout().SetContentsMargins(0, 0, 0, 0)
	ptr.bar = widgets.NewQToolBar("Settings", ptr)
	ptr.bar.SetOrientation(core.Qt__Vertical)
	ptr.general = ptr.bar.AddAction("General")
	ptr.al = ptr.bar.AddAction("Account")
	ptr.content.Layout().AddWidget(ptr.bar)

	ptr.appTab = widgets.NewQWidget(nil, core.Qt__Widget)
	ptr.appTab.SetLayout(widgets.NewQVBoxLayout())
	ptr.theme = widgets.NewQCheckBox2("Custom Theme", ptr.appTab)

	ptr.gen = widgets.NewQTabWidget(nil)
	ptr.gen.AddTab(ptr.appTab, "Application")
	ptr.content.Layout().AddWidget(ptr.gen)
	ptr.gen.Hide()

	ptr.alTab = widgets.NewQWidget(nil, core.Qt__Widget)
	ptr.alTab.SetLayout(widgets.NewQVBoxLayout())
	ptr.username = widgets.NewQLabel2("", ptr.alTab, core.Qt__Widget)
	ptr.username.SetAlignment(core.Qt__AlignTop | core.Qt__AlignHCenter)
	ptr.loginBtn = widgets.NewQPushButton(ptr.alTab)

	ptr.alTab.Layout().AddWidget(ptr.loginBtn)
	ptr.alTab.Layout().AddWidget(ptr.username)

	ptr.account = widgets.NewQTabWidget(nil)
	ptr.account.AddTab(ptr.alTab, "AniList")
	ptr.content.Layout().AddWidget(ptr.account)
	ptr.account.Hide()

	ptr.general.ConnectTriggered(func(bool) {
		ptr.account.Hide()
		ptr.gen.Show()
	})
	ptr.al.ConnectTriggered(func(bool) {
		ptr.gen.Hide()
		ptr.account.Show()
	})
	ptr.general.Trigger()
	ptr.Layout().AddWidget(ptr.content)

	// Add widget with horiz layout for buttons
	ptr.btnWidget = widgets.NewQWidget(nil, core.Qt__Widget)
	ptr.btnWidget.SetLayout(widgets.NewQHBoxLayout())
	ptr.btnWidget.Layout().SetContentsMargins(10, 10, 10, 10)
	ptr.btnWidget.Layout().SetAlign(core.Qt__AlignBottom | core.Qt__AlignRight)

	ptr.okBtn = widgets.NewQPushButton2("OK", ptr.btnWidget)
	ptr.cancelBtn = widgets.NewQPushButton2("Cancel", ptr.btnWidget)
	ptr.okBtn.ConnectClicked(func(bool) {
		ptr.saveChanges()
		ptr.DestroySettingsDialog()
		ptr.content = nil
		ptr.bar = nil
		ptr.general = nil
		ptr.al = nil
		ptr.appTab = nil
		ptr.gen = nil
		ptr.theme = nil
		ptr.alTab = nil
		ptr.account = nil
		ptr.username = nil
		ptr.loginBtn = nil
		ptr.btnWidget = nil
		ptr.okBtn = nil
		ptr.cancelBtn = nil
		ptr.changes = nil
		runtime.GC()
	})
	ptr.cancelBtn.ConnectClicked(func(bool) {
		ptr.DestroySettingsDialog()
		ptr.content = nil
		ptr.bar = nil
		ptr.general = nil
		ptr.al = nil
		ptr.appTab = nil
		ptr.gen = nil
		ptr.theme = nil
		ptr.alTab = nil
		ptr.account = nil
		ptr.username = nil
		ptr.loginBtn = nil
		ptr.btnWidget = nil
		ptr.okBtn = nil
		ptr.cancelBtn = nil
		ptr.changes = nil
		runtime.GC()
	})
	ptr.btnWidget.Layout().AddWidget(ptr.okBtn)
	ptr.btnWidget.Layout().AddWidget(ptr.cancelBtn)

	ptr.Layout().AddWidget(ptr.btnWidget)
}

func (ptr *SettingsDialog) saveChanges() {
	for k, v := range ptr.changes {
		ptr.user.Settings.Set(k, v)
	}
}
