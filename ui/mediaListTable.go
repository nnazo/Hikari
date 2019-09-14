package ui

import (
	"hikari/lists"
	"hikari/media"
	"strconv"

	"github.com/therecipe/qt/core"
)

type MediaListTable struct {
	core.QAbstractTableModel
	_         func() `constructor:"init"`
	modelData media.List
	mediaType string
}

func (ptr *MediaListTable) SetupUi(list *lists.MediaList, modelData media.List) {
	ptr.mediaType = list.Type
	ptr.modelData = modelData
}

func (ptr *MediaListTable) init() {
	ptr.ConnectHeaderData(ptr.headerData)
	ptr.ConnectRowCount(ptr.rowCount)
	ptr.ConnectColumnCount(ptr.columnCount)
	ptr.ConnectData(ptr.data)
}

func (ptr *MediaListTable) headerData(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	if role != int(core.Qt__DisplayRole) || orientation == core.Qt__Vertical {
		return ptr.HeaderDataDefault(section, orientation, role)
	}

	switch ptr.mediaType {
	case media.AnimeType:
		switch section {
		case 0:
			return core.NewQVariant14("Title")
		case 1:
			return core.NewQVariant14("Progress")
		case 2:
			return core.NewQVariant14("Score")
		case 3:
			return core.NewQVariant14("Status")
		}
	case media.MangaType:
		switch section {
		case 0:
			return core.NewQVariant14("Title")
		case 1:
			return core.NewQVariant14("Progress")
		case 2:
			return core.NewQVariant14("Progress Volumes")
		case 3:
			return core.NewQVariant14("Score")
		case 4:
			return core.NewQVariant14("Status")
		}
	}
	return core.NewQVariant()
}

func (ptr *MediaListTable) rowCount(*core.QModelIndex) int {
	return len(ptr.modelData.Entries)
}

func (ptr *MediaListTable) columnCount(*core.QModelIndex) int {
	switch ptr.mediaType {
	case media.AnimeType:
		return 4
	case media.MangaType:
		return 5
	}
	return 4
}

func (ptr *MediaListTable) data(index *core.QModelIndex, role int) *core.QVariant {
	if role != int(core.Qt__DisplayRole) {
		return core.NewQVariant()
	}

	item := ptr.modelData.Entries[index.Row()]
	switch ptr.mediaType { // item.Media.Type
	case media.AnimeType:
		switch ptr.HeaderData(index.Column(), core.Qt__Horizontal, role).ToString() {
		case "Title":
			return core.NewQVariant14(item.Media.Title.UserPreferred) // item.Title
		case "Progress":
			return core.NewQVariant14(strconv.Itoa(item.Progress))
		case "Score":
			return core.NewQVariant14(strconv.FormatFloat(item.Score, 'f', 1, 64))
		case "Status":
			return core.NewQVariant14(item.Status)
		}
	case media.MangaType:
		switch ptr.HeaderData(index.Column(), core.Qt__Horizontal, role).ToString() {
		case "Title":
			return core.NewQVariant14(item.Media.Title.UserPreferred) // item.Title
		case "Progress":
			return core.NewQVariant14(strconv.Itoa(item.Progress))
		case "Progress Volumes":
			return core.NewQVariant14(strconv.Itoa(item.ProgressVolumes))
		case "Score":
			return core.NewQVariant14(strconv.FormatFloat(item.Score, 'f', 1, 64))
		case "Status":
			return core.NewQVariant14(item.Status)
		}
	}
	return core.NewQVariant()
}
