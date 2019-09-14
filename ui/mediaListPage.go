package ui

import (
	"hikari/lists"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type MediaListPage struct {
	widgets.QWidget
	_               func() `constructor:"init"`
	filterContainer *widgets.QWidget
	filter          *widgets.QLineEdit
	tabs            *MediaListTabs
}

func (ptr *MediaListPage) SetupUi(list *lists.MediaList) {
	ptr.tabs.SetupUi(list)
}

func (ptr *MediaListPage) init() {
	ptr.SetLayout(widgets.NewQVBoxLayout())
	ptr.Layout().SetContentsMargins(0, 0, 0, 0)

	ptr.filterContainer = widgets.NewQWidget(ptr.ParentWidget(), core.Qt__Widget)
	ptr.filter = widgets.NewQLineEdit(ptr.filterContainer)
	ptr.tabs = NewMediaListTabs(ptr)

	ptr.filter.SetFixedWidth(240)
	ptr.filter.SetPlaceholderText("Filter list")
	ptr.filter.SetEditFocus(false)
	ptr.filter.SetFrame(false)
	// ptr.filter.ConnectTextChanged(func(s string) {

	// })

	ptr.filterContainer.SetLayout(widgets.NewQHBoxLayout())
	ptr.filterContainer.Layout().SetContentsMargins(0, 0, 0, 0)
	ptr.filterContainer.Layout().SetAlign(core.Qt__AlignRight)
	ptr.filterContainer.Layout().AddWidget(ptr.filter)

	ptr.Layout().AddWidget(ptr.filterContainer)
	ptr.Layout().AddWidget(ptr.tabs)
}
