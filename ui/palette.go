package ui

import (
	"fmt"
	"hikari/settings"
	"io/ioutil"
	"strings"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func LoadQss(app *widgets.QApplication, settings *settings.Settings) {
	var theme *core.QVariant
	if settings.Contains("theme") {
		theme = settings.Value("theme", core.NewQVariant14("native"))
	} else {
		settings.SetValue("theme", core.NewQVariant14("native"))
	}

	var qss []byte
	var err error
	if theme != nil {
		if theme.ToString() == "custom" {
			qss, err = ioutil.ReadFile("../../custom.qss")
		} else {
			qss, err = ioutil.ReadFile("../../native.qss")
		}
	} else {
		qss, err = ioutil.ReadFile("../../native.qss")
	}

	if err != nil {
		fmt.Println("could not load qss")
	} else {
		str := strings.ReplaceAll(string(qss), "\r\n", " ")
		str = strings.ReplaceAll(str, "\n", "")
		str = strings.ReplaceAll(str, "  ", "")
		app.SetStyleSheet(str)
	}
}

func SetAppStyle(app *widgets.QApplication, settings *settings.Settings) {
	app.SetStyle(widgets.QStyleFactory_Create("Fusion"))
	palette := gui.NewQPalette()
	if settings.Contains("theme") {
		theme := settings.Value("theme", core.NewQVariant14("native"))
		if theme.ToString() == "custom" {
			darkPalette := palette
			darkColor := gui.NewQColor3(32, 34, 37, 1)
			white := gui.NewQColor6("white")
			disabledColor := gui.NewQColor3(127, 127, 127, 1)

			darkPalette.SetColor2(gui.QPalette__Window, darkColor)
			darkPalette.SetColor2(gui.QPalette__WindowText, white)
			darkPalette.SetColor2(gui.QPalette__Base, gui.NewQColor3(24, 26, 29, 120))
			darkPalette.SetColor2(gui.QPalette__AlternateBase, darkColor)
			darkPalette.SetColor2(gui.QPalette__ToolTipBase, white)
			darkPalette.SetColor2(gui.QPalette__ToolTipText, white)
			darkPalette.SetColor2(gui.QPalette__PlaceholderText, gui.NewQColor3(81, 81, 81, 255))
			darkPalette.SetColor2(gui.QPalette__Text, white)
			darkPalette.SetColor(gui.QPalette__Disabled, gui.QPalette__Text, disabledColor)
			darkPalette.SetColor2(gui.QPalette__Button, darkColor)
			darkPalette.SetColor2(gui.QPalette__ButtonText, white)
			darkPalette.SetColor(gui.QPalette__Disabled, gui.QPalette__ButtonText, darkColor)
			darkPalette.SetColor2(gui.QPalette__BrightText, gui.NewQColor3(255, 0, 0, 255))
			darkPalette.SetColor2(gui.QPalette__Link, gui.NewQColor3(0, 255, 0, 1))

			darkPalette.SetColor2(gui.QPalette__Highlight, gui.NewQColor3(42, 130, 218, 255))
			darkPalette.SetColor2(gui.QPalette__HighlightedText, gui.NewQColor3(0, 0, 0, 255))
			darkPalette.SetColor(gui.QPalette__Disabled, gui.QPalette__HighlightedText, disabledColor)
			// app.SetPalette(darkPalette, "")
		} else {
			app.SetPalette(gui.NewQPalette(), "")
		}
	} else {
		settings.SetValue("theme", core.NewQVariant14("native"))
		app.SetPalette(gui.NewQPalette(), "")
	}
	LoadQss(app, settings)
	app.SetPalette(palette, "")
}
