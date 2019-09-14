package settings

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_Moc_PackedString) []byte {
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type Settings_ITF interface {
	std_core.QSettings_ITF
	Settings_PTR() *Settings
}

func (ptr *Settings) Settings_PTR() *Settings {
	return ptr
}

func (ptr *Settings) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSettings_PTR().Pointer()
	}
	return nil
}

func (ptr *Settings) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSettings_PTR().SetPointer(p)
	}
}

func PointerFromSettings(ptr Settings_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Settings_PTR().Pointer()
	}
	return nil
}

func NewSettingsFromPointer(ptr unsafe.Pointer) (n *Settings) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(Settings)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *Settings:
			n = deduced

		case *std_core.QSettings:
			n = &Settings{QSettings: *deduced}

		default:
			n = new(Settings)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackSettingsc4adcb_Constructor
func callbackSettingsc4adcb_Constructor(ptr unsafe.Pointer) {
	this := NewSettingsFromPointer(ptr)
	qt.Register(ptr, this)
}

func Settings_QRegisterMetaType() int {
	return int(int32(C.Settingsc4adcb_Settingsc4adcb_QRegisterMetaType()))
}

func (ptr *Settings) QRegisterMetaType() int {
	return int(int32(C.Settingsc4adcb_Settingsc4adcb_QRegisterMetaType()))
}

func Settings_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Settingsc4adcb_Settingsc4adcb_QRegisterMetaType2(typeNameC)))
}

func (ptr *Settings) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Settingsc4adcb_Settingsc4adcb_QRegisterMetaType2(typeNameC)))
}

func Settings_QmlRegisterType() int {
	return int(int32(C.Settingsc4adcb_Settingsc4adcb_QmlRegisterType()))
}

func (ptr *Settings) QmlRegisterType() int {
	return int(int32(C.Settingsc4adcb_Settingsc4adcb_QmlRegisterType()))
}

func Settings_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.Settingsc4adcb_Settingsc4adcb_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Settings) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.Settingsc4adcb_Settingsc4adcb_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Settings) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.Settingsc4adcb___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *Settings) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.Settingsc4adcb___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *Settings) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.Settingsc4adcb___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *Settings) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Settingsc4adcb___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Settings) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Settingsc4adcb___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Settings) __findChildren_newList2() unsafe.Pointer {
	return C.Settingsc4adcb___findChildren_newList2(ptr.Pointer())
}

func (ptr *Settings) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Settingsc4adcb___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Settings) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Settingsc4adcb___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Settings) __findChildren_newList3() unsafe.Pointer {
	return C.Settingsc4adcb___findChildren_newList3(ptr.Pointer())
}

func (ptr *Settings) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Settingsc4adcb___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Settings) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Settingsc4adcb___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Settings) __findChildren_newList() unsafe.Pointer {
	return C.Settingsc4adcb___findChildren_newList(ptr.Pointer())
}

func (ptr *Settings) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Settingsc4adcb___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Settings) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Settingsc4adcb___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Settings) __children_newList() unsafe.Pointer {
	return C.Settingsc4adcb___children_newList(ptr.Pointer())
}

func NewSettings5(parent std_core.QObject_ITF) *Settings {
	tmpValue := NewSettingsFromPointer(C.Settingsc4adcb_NewSettings5(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewSettings3(format std_core.QSettings__Format, scope std_core.QSettings__Scope, organization string, application string, parent std_core.QObject_ITF) *Settings {
	var organizationC *C.char
	if organization != "" {
		organizationC = C.CString(organization)
		defer C.free(unsafe.Pointer(organizationC))
	}
	var applicationC *C.char
	if application != "" {
		applicationC = C.CString(application)
		defer C.free(unsafe.Pointer(applicationC))
	}
	tmpValue := NewSettingsFromPointer(C.Settingsc4adcb_NewSettings3(C.longlong(format), C.longlong(scope), C.struct_Moc_PackedString{data: organizationC, len: C.longlong(len(organization))}, C.struct_Moc_PackedString{data: applicationC, len: C.longlong(len(application))}, std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewSettings2(scope std_core.QSettings__Scope, organization string, application string, parent std_core.QObject_ITF) *Settings {
	var organizationC *C.char
	if organization != "" {
		organizationC = C.CString(organization)
		defer C.free(unsafe.Pointer(organizationC))
	}
	var applicationC *C.char
	if application != "" {
		applicationC = C.CString(application)
		defer C.free(unsafe.Pointer(applicationC))
	}
	tmpValue := NewSettingsFromPointer(C.Settingsc4adcb_NewSettings2(C.longlong(scope), C.struct_Moc_PackedString{data: organizationC, len: C.longlong(len(organization))}, C.struct_Moc_PackedString{data: applicationC, len: C.longlong(len(application))}, std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewSettings4(fileName string, format std_core.QSettings__Format, parent std_core.QObject_ITF) *Settings {
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	tmpValue := NewSettingsFromPointer(C.Settingsc4adcb_NewSettings4(C.struct_Moc_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, C.longlong(format), std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewSettings(organization string, application string, parent std_core.QObject_ITF) *Settings {
	var organizationC *C.char
	if organization != "" {
		organizationC = C.CString(organization)
		defer C.free(unsafe.Pointer(organizationC))
	}
	var applicationC *C.char
	if application != "" {
		applicationC = C.CString(application)
		defer C.free(unsafe.Pointer(applicationC))
	}
	tmpValue := NewSettingsFromPointer(C.Settingsc4adcb_NewSettings(C.struct_Moc_PackedString{data: organizationC, len: C.longlong(len(organization))}, C.struct_Moc_PackedString{data: applicationC, len: C.longlong(len(application))}, std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackSettingsc4adcb_DestroySettings
func callbackSettingsc4adcb_DestroySettings(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~Settings"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsFromPointer(ptr).DestroySettingsDefault()
	}
}

func (ptr *Settings) ConnectDestroySettings(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~Settings"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~Settings", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~Settings", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Settings) DisconnectDestroySettings() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~Settings")
	}
}

func (ptr *Settings) DestroySettings() {
	if ptr.Pointer() != nil {
		C.Settingsc4adcb_DestroySettings(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *Settings) DestroySettingsDefault() {
	if ptr.Pointer() != nil {
		C.Settingsc4adcb_DestroySettingsDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackSettingsc4adcb_Event
func callbackSettingsc4adcb_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSettingsFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(event)))))
}

func (ptr *Settings) EventDefault(event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.Settingsc4adcb_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackSettingsc4adcb_EventFilter
func callbackSettingsc4adcb_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSettingsFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *Settings) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.Settingsc4adcb_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackSettingsc4adcb_ChildEvent
func callbackSettingsc4adcb_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewSettingsFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *Settings) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Settingsc4adcb_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackSettingsc4adcb_ConnectNotify
func callbackSettingsc4adcb_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewSettingsFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Settings) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Settingsc4adcb_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackSettingsc4adcb_CustomEvent
func callbackSettingsc4adcb_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewSettingsFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *Settings) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Settingsc4adcb_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackSettingsc4adcb_DeleteLater
func callbackSettingsc4adcb_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *Settings) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.Settingsc4adcb_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackSettingsc4adcb_Destroyed
func callbackSettingsc4adcb_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackSettingsc4adcb_DisconnectNotify
func callbackSettingsc4adcb_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewSettingsFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Settings) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Settingsc4adcb_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackSettingsc4adcb_ObjectNameChanged
func callbackSettingsc4adcb_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackSettingsc4adcb_TimerEvent
func callbackSettingsc4adcb_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewSettingsFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *Settings) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Settingsc4adcb_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
