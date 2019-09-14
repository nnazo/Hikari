package settings

import (
	"github.com/therecipe/qt/core"
)

type Settings struct {
	core.QSettings
}

func (ptr *Settings) Set(setting string, value string) {
	ptr.SetValue(setting, core.NewQVariant14(value))
}

func (ptr *Settings) Get(setting string) string {
	return ptr.Value(setting, core.NewQVariant14(ptr.Default(setting))).ToString()
}

func (ptr *Settings) Erase(setting string) {
	ptr.SetValue(setting, core.NewQVariant14(""))
	ptr.Remove(setting)
}

func (ptr *Settings) Default(setting string) string {
	switch setting {
	case "theme":
		return "native"
	case "tokenData":
		return ""
	default:
		return ""
	}
}
