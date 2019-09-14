package ui

import (
	"hikari/lists"

	"github.com/therecipe/qt/widgets"
)

type MediaListTabs struct {
	widgets.QTabWidget
	_          func() `constructor:"init"`
	tableViews []*widgets.QTableView
	tables     []*MediaListTable
}

func (ptr *MediaListTabs) SetupUi(list *lists.MediaList) {
	if list != nil {
		for i, v := range list.Lists {
			ptr.tables = append(ptr.tables, NewMediaListTable(ptr))
			ptr.tables[i].SetupUi(list, v)
			ptr.tableViews = append(ptr.tableViews, widgets.NewQTableView(nil))
			ptr.tableViews[i].SetModel(ptr.tables[i])
			// ptr.tableViews[i].VerticalHeader().SetVisible(false)
			ptr.tableViews[i].ResizeColumnsToContents()
			ptr.tableViews[i].SetWordWrap(false)
			ptr.tableViews[i].SetSortingEnabled(true)
			ptr.tableViews[i].SetCornerButtonEnabled(false)
			// ptr.tableViews[i].SortByColumn
			ptr.AddTab(ptr.tableViews[i], v.Name)
		}
	}
}

func (ptr *MediaListTabs) init() {
	ptr.tableViews = make([]*widgets.QTableView, 0)
	ptr.tables = make([]*MediaListTable, 0)
}
