package main

import (
	"hikari/settings"
	"hikari/ui"
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	app.SetApplicationName("Hikari")
	app.SetOrganizationName("nazo")
	app.SetOrganizationDomain("https://hikari.page/")

	s := settings.NewSettings("nazo", "Hikari", nil)
	ui.SetAppStyle(app, s)

	window := ui.NewMainWindow(nil, 0)
	window.SetupUi(s)

	window.Show()

	app.Exec()
}
