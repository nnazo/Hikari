package ui

import (
	"hikari/players"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type CurrentMediaPage struct {
	widgets.QTabWidget
	_          func() `constructor:"init"`
	mediaTitle *widgets.QLabel
	cover      *gui.QPixmap
	dummy      *widgets.QLabel
}

func (ptr *CurrentMediaPage) init() {
	ptr.SetLayout(widgets.NewQHBoxLayout())
	ptr.Layout().SetContentsMargins(-1, -1, -1, -1)
	ptr.Layout().SetAlign(core.Qt__AlignTop | core.Qt__AlignLeft)
	ptr.mediaTitle = widgets.NewQLabel(ptr, core.Qt__Widget)
	ptr.mediaTitle.SetWordWrap(true)
	ptr.mediaTitle.SetSizePolicy2(widgets.QSizePolicy__MinimumExpanding, widgets.QSizePolicy__MinimumExpanding)
	ptr.mediaTitle.SetAlignment(core.Qt__AlignTop | core.Qt__AlignLeft)

	go ptr.searchForMedia()

	ptr.cover = gui.NewQPixmap3(200, 300)
	ptr.dummy = widgets.NewQLabel(ptr, core.Qt__Widget)
	ptr.dummy.SetPixmap(ptr.cover)
	ptr.dummy.SetAlignment(core.Qt__AlignTop | core.Qt__AlignLeft)

	ptr.Layout().AddWidget(ptr.dummy)
	ptr.Layout().AddWidget(ptr.mediaTitle)
}

// plan: search list via https://en.wikipedia.org/wiki/Levenshtein_distance

func (ptr *CurrentMediaPage) searchForMedia() {
	var old string
	for range time.NewTicker(time.Duration(1) * time.Second).C {
		title, med := players.GetPlayingMedia()
		if old != title {
			ptr.mediaTitle.SetText(title)
			ptr.mediaTitle.UpdateGeometry()
			old = title
			if med != nil { // and is in list
				med.Update()
			}
		}
	}
}
