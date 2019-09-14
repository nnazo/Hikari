package ui

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
	std_gui "github.com/therecipe/qt/gui"
	std_widgets "github.com/therecipe/qt/widgets"
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

type MediaListTable_ITF interface {
	std_core.QAbstractTableModel_ITF
	MediaListTable_PTR() *MediaListTable
}

func (ptr *MediaListTable) MediaListTable_PTR() *MediaListTable {
	return ptr
}

func (ptr *MediaListTable) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractTableModel_PTR().Pointer()
	}
	return nil
}

func (ptr *MediaListTable) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractTableModel_PTR().SetPointer(p)
	}
}

func PointerFromMediaListTable(ptr MediaListTable_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.MediaListTable_PTR().Pointer()
	}
	return nil
}

func NewMediaListTableFromPointer(ptr unsafe.Pointer) (n *MediaListTable) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(MediaListTable)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *MediaListTable:
			n = deduced

		case *std_core.QAbstractTableModel:
			n = &MediaListTable{QAbstractTableModel: *deduced}

		default:
			n = new(MediaListTable)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *MediaListTable) Init() { this.init() }

//export callbackMediaListTable46aa22_Constructor
func callbackMediaListTable46aa22_Constructor(ptr unsafe.Pointer) {
	this := NewMediaListTableFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

func MediaListTable_QRegisterMetaType() int {
	return int(int32(C.MediaListTable46aa22_MediaListTable46aa22_QRegisterMetaType()))
}

func (ptr *MediaListTable) QRegisterMetaType() int {
	return int(int32(C.MediaListTable46aa22_MediaListTable46aa22_QRegisterMetaType()))
}

func MediaListTable_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.MediaListTable46aa22_MediaListTable46aa22_QRegisterMetaType2(typeNameC)))
}

func (ptr *MediaListTable) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.MediaListTable46aa22_MediaListTable46aa22_QRegisterMetaType2(typeNameC)))
}

func MediaListTable_QmlRegisterType() int {
	return int(int32(C.MediaListTable46aa22_MediaListTable46aa22_QmlRegisterType()))
}

func (ptr *MediaListTable) QmlRegisterType() int {
	return int(int32(C.MediaListTable46aa22_MediaListTable46aa22_QmlRegisterType()))
}

func MediaListTable_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.MediaListTable46aa22_MediaListTable46aa22_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *MediaListTable) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.MediaListTable46aa22_MediaListTable46aa22_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *MediaListTable) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.MediaListTable46aa22_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *MediaListTable) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *MediaListTable) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.MediaListTable46aa22_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *MediaListTable) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.MediaListTable46aa22_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *MediaListTable) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *MediaListTable) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.MediaListTable46aa22_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *MediaListTable) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.MediaListTable46aa22_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *MediaListTable) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *MediaListTable) ____itemData_keyList_newList() unsafe.Pointer {
	return C.MediaListTable46aa22_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *MediaListTable) __setItemData_roles_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.MediaListTable46aa22___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *MediaListTable) __setItemData_roles_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *MediaListTable) __setItemData_roles_newList() unsafe.Pointer {
	return C.MediaListTable46aa22___setItemData_roles_newList(ptr.Pointer())
}

func (ptr *MediaListTable) __setItemData_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewMediaListTableFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setItemData_roles_keyList_atList(i)
			}
			return out
		}(C.MediaListTable46aa22___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *MediaListTable) __changePersistentIndexList_from_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.MediaListTable46aa22___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *MediaListTable) __changePersistentIndexList_from_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22___changePersistentIndexList_from_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *MediaListTable) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.MediaListTable46aa22___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *MediaListTable) __changePersistentIndexList_to_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.MediaListTable46aa22___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *MediaListTable) __changePersistentIndexList_to_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22___changePersistentIndexList_to_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *MediaListTable) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.MediaListTable46aa22___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *MediaListTable) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.MediaListTable46aa22___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *MediaListTable) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *MediaListTable) __dataChanged_roles_newList() unsafe.Pointer {
	return C.MediaListTable46aa22___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *MediaListTable) __layoutAboutToBeChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.MediaListTable46aa22___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *MediaListTable) __layoutAboutToBeChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *MediaListTable) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.MediaListTable46aa22___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *MediaListTable) __layoutChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.MediaListTable46aa22___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *MediaListTable) __layoutChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22___layoutChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *MediaListTable) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.MediaListTable46aa22___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *MediaListTable) __roleNames_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.MediaListTable46aa22___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *MediaListTable) __roleNames_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22___roleNames_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *MediaListTable) __roleNames_newList() unsafe.Pointer {
	return C.MediaListTable46aa22___roleNames_newList(ptr.Pointer())
}

func (ptr *MediaListTable) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewMediaListTableFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roleNames_keyList_atList(i)
			}
			return out
		}(C.MediaListTable46aa22___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *MediaListTable) __itemData_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.MediaListTable46aa22___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *MediaListTable) __itemData_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22___itemData_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *MediaListTable) __itemData_newList() unsafe.Pointer {
	return C.MediaListTable46aa22___itemData_newList(ptr.Pointer())
}

func (ptr *MediaListTable) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewMediaListTableFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____itemData_keyList_atList(i)
			}
			return out
		}(C.MediaListTable46aa22___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *MediaListTable) __mimeData_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.MediaListTable46aa22___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *MediaListTable) __mimeData_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22___mimeData_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *MediaListTable) __mimeData_indexes_newList() unsafe.Pointer {
	return C.MediaListTable46aa22___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *MediaListTable) __match_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.MediaListTable46aa22___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *MediaListTable) __match_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22___match_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *MediaListTable) __match_newList() unsafe.Pointer {
	return C.MediaListTable46aa22___match_newList(ptr.Pointer())
}

func (ptr *MediaListTable) __persistentIndexList_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.MediaListTable46aa22___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *MediaListTable) __persistentIndexList_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22___persistentIndexList_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *MediaListTable) __persistentIndexList_newList() unsafe.Pointer {
	return C.MediaListTable46aa22___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *MediaListTable) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.MediaListTable46aa22_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *MediaListTable) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *MediaListTable) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.MediaListTable46aa22_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *MediaListTable) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.MediaListTable46aa22_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *MediaListTable) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *MediaListTable) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.MediaListTable46aa22_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *MediaListTable) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.MediaListTable46aa22___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *MediaListTable) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *MediaListTable) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.MediaListTable46aa22___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *MediaListTable) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.MediaListTable46aa22___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MediaListTable) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *MediaListTable) __findChildren_newList2() unsafe.Pointer {
	return C.MediaListTable46aa22___findChildren_newList2(ptr.Pointer())
}

func (ptr *MediaListTable) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.MediaListTable46aa22___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MediaListTable) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *MediaListTable) __findChildren_newList3() unsafe.Pointer {
	return C.MediaListTable46aa22___findChildren_newList3(ptr.Pointer())
}

func (ptr *MediaListTable) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.MediaListTable46aa22___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MediaListTable) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *MediaListTable) __findChildren_newList() unsafe.Pointer {
	return C.MediaListTable46aa22___findChildren_newList(ptr.Pointer())
}

func (ptr *MediaListTable) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.MediaListTable46aa22___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MediaListTable) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *MediaListTable) __children_newList() unsafe.Pointer {
	return C.MediaListTable46aa22___children_newList(ptr.Pointer())
}

func NewMediaListTable(parent std_core.QObject_ITF) *MediaListTable {
	tmpValue := NewMediaListTableFromPointer(C.MediaListTable46aa22_NewMediaListTable(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackMediaListTable46aa22_DestroyMediaListTable
func callbackMediaListTable46aa22_DestroyMediaListTable(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~MediaListTable"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListTableFromPointer(ptr).DestroyMediaListTableDefault()
	}
}

func (ptr *MediaListTable) ConnectDestroyMediaListTable(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~MediaListTable"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~MediaListTable", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~MediaListTable", unsafe.Pointer(&f))
		}
	}
}

func (ptr *MediaListTable) DisconnectDestroyMediaListTable() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~MediaListTable")
	}
}

func (ptr *MediaListTable) DestroyMediaListTable() {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22_DestroyMediaListTable(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *MediaListTable) DestroyMediaListTableDefault() {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22_DestroyMediaListTableDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackMediaListTable46aa22_DropMimeData
func callbackMediaListTable46aa22_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListTableFromPointer(ptr).DropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *MediaListTable) DropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.MediaListTable46aa22_DropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackMediaListTable46aa22_Index
func callbackMediaListTable46aa22_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
	}

	return std_core.PointerFromQModelIndex(NewMediaListTableFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
}

func (ptr *MediaListTable) IndexDefault(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.MediaListTable46aa22_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackMediaListTable46aa22_Sibling
func callbackMediaListTable46aa22_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
	}

	return std_core.PointerFromQModelIndex(NewMediaListTableFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
}

func (ptr *MediaListTable) SiblingDefault(row int, column int, idx std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.MediaListTable46aa22_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackMediaListTable46aa22_Flags
func callbackMediaListTable46aa22_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong((*(*func(*std_core.QModelIndex) std_core.Qt__ItemFlag)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewMediaListTableFromPointer(ptr).FlagsDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *MediaListTable) FlagsDefault(index std_core.QModelIndex_ITF) std_core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return std_core.Qt__ItemFlag(C.MediaListTable46aa22_FlagsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackMediaListTable46aa22_InsertColumns
func callbackMediaListTable46aa22_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListTableFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *MediaListTable) InsertColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.MediaListTable46aa22_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackMediaListTable46aa22_InsertRows
func callbackMediaListTable46aa22_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListTableFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *MediaListTable) InsertRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.MediaListTable46aa22_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackMediaListTable46aa22_MoveColumns
func callbackMediaListTable46aa22_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListTableFromPointer(ptr).MoveColumnsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *MediaListTable) MoveColumnsDefault(sourceParent std_core.QModelIndex_ITF, sourceColumn int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.MediaListTable46aa22_MoveColumnsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackMediaListTable46aa22_MoveRows
func callbackMediaListTable46aa22_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListTableFromPointer(ptr).MoveRowsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *MediaListTable) MoveRowsDefault(sourceParent std_core.QModelIndex_ITF, sourceRow int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.MediaListTable46aa22_MoveRowsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackMediaListTable46aa22_RemoveColumns
func callbackMediaListTable46aa22_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListTableFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *MediaListTable) RemoveColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.MediaListTable46aa22_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackMediaListTable46aa22_RemoveRows
func callbackMediaListTable46aa22_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListTableFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *MediaListTable) RemoveRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.MediaListTable46aa22_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackMediaListTable46aa22_SetData
func callbackMediaListTable46aa22_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, *std_core.QVariant, int) bool)(signal))(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListTableFromPointer(ptr).SetDataDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *MediaListTable) SetDataDefault(index std_core.QModelIndex_ITF, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.MediaListTable46aa22_SetDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackMediaListTable46aa22_SetHeaderData
func callbackMediaListTable46aa22_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, std_core.Qt__Orientation, *std_core.QVariant, int) bool)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListTableFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *MediaListTable) SetHeaderDataDefault(section int, orientation std_core.Qt__Orientation, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.MediaListTable46aa22_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackMediaListTable46aa22_SetItemData
func callbackMediaListTable46aa22_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_Moc_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, map[int]*std_core.QVariant) bool)(signal))(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewMediaListTableFromPointer(l.data)
			for i, v := range tmpList.__setItemData_roles_keyList() {
				out[v] = tmpList.__setItemData_roles_atList(v, i)
			}
			return out
		}(roles)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListTableFromPointer(ptr).SetItemDataDefault(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
		out := make(map[int]*std_core.QVariant, int(l.len))
		tmpList := NewMediaListTableFromPointer(l.data)
		for i, v := range tmpList.__setItemData_roles_keyList() {
			out[v] = tmpList.__setItemData_roles_atList(v, i)
		}
		return out
	}(roles)))))
}

func (ptr *MediaListTable) SetItemDataDefault(index std_core.QModelIndex_ITF, roles map[int]*std_core.QVariant) bool {
	if ptr.Pointer() != nil {
		return int8(C.MediaListTable46aa22_SetItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewMediaListTableFromPointer(NewMediaListTableFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackMediaListTable46aa22_Submit
func callbackMediaListTable46aa22_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListTableFromPointer(ptr).SubmitDefault())))
}

func (ptr *MediaListTable) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.MediaListTable46aa22_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackMediaListTable46aa22_ColumnsAboutToBeInserted
func callbackMediaListTable46aa22_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackMediaListTable46aa22_ColumnsAboutToBeMoved
func callbackMediaListTable46aa22_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackMediaListTable46aa22_ColumnsAboutToBeRemoved
func callbackMediaListTable46aa22_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackMediaListTable46aa22_ColumnsInserted
func callbackMediaListTable46aa22_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackMediaListTable46aa22_ColumnsMoved
func callbackMediaListTable46aa22_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackMediaListTable46aa22_ColumnsRemoved
func callbackMediaListTable46aa22_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackMediaListTable46aa22_DataChanged
func callbackMediaListTable46aa22_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		(*(*func(*std_core.QModelIndex, *std_core.QModelIndex, []int))(signal))(std_core.NewQModelIndexFromPointer(topLeft), std_core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewMediaListTableFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackMediaListTable46aa22_FetchMore
func callbackMediaListTable46aa22_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		(*(*func(*std_core.QModelIndex))(signal))(std_core.NewQModelIndexFromPointer(parent))
	} else {
		NewMediaListTableFromPointer(ptr).FetchMoreDefault(std_core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *MediaListTable) FetchMoreDefault(parent std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22_FetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))
	}
}

//export callbackMediaListTable46aa22_HeaderDataChanged
func callbackMediaListTable46aa22_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		(*(*func(std_core.Qt__Orientation, int, int))(signal))(std_core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackMediaListTable46aa22_LayoutAboutToBeChanged
func callbackMediaListTable46aa22_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		(*(*func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewMediaListTableFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackMediaListTable46aa22_LayoutChanged
func callbackMediaListTable46aa22_LayoutChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		(*(*func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewMediaListTableFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackMediaListTable46aa22_ModelAboutToBeReset
func callbackMediaListTable46aa22_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackMediaListTable46aa22_ModelReset
func callbackMediaListTable46aa22_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackMediaListTable46aa22_ResetInternalData
func callbackMediaListTable46aa22_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListTableFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *MediaListTable) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackMediaListTable46aa22_Revert
func callbackMediaListTable46aa22_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListTableFromPointer(ptr).RevertDefault()
	}
}

func (ptr *MediaListTable) RevertDefault() {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22_RevertDefault(ptr.Pointer())
	}
}

//export callbackMediaListTable46aa22_RowsAboutToBeInserted
func callbackMediaListTable46aa22_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackMediaListTable46aa22_RowsAboutToBeMoved
func callbackMediaListTable46aa22_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackMediaListTable46aa22_RowsAboutToBeRemoved
func callbackMediaListTable46aa22_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackMediaListTable46aa22_RowsInserted
func callbackMediaListTable46aa22_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackMediaListTable46aa22_RowsMoved
func callbackMediaListTable46aa22_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackMediaListTable46aa22_RowsRemoved
func callbackMediaListTable46aa22_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackMediaListTable46aa22_Sort
func callbackMediaListTable46aa22_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		(*(*func(int, std_core.Qt__SortOrder))(signal))(int(int32(column)), std_core.Qt__SortOrder(order))
	} else {
		NewMediaListTableFromPointer(ptr).SortDefault(int(int32(column)), std_core.Qt__SortOrder(order))
	}
}

func (ptr *MediaListTable) SortDefault(column int, order std_core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackMediaListTable46aa22_RoleNames
func callbackMediaListTable46aa22_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewMediaListTableFromPointer(NewMediaListTableFromPointer(nil).__roleNames_newList())
			for k, v := range (*(*func() map[int]*std_core.QByteArray)(signal))() {
				tmpList.__roleNames_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewMediaListTableFromPointer(NewMediaListTableFromPointer(nil).__roleNames_newList())
		for k, v := range NewMediaListTableFromPointer(ptr).RoleNamesDefault() {
			tmpList.__roleNames_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *MediaListTable) RoleNamesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewMediaListTableFromPointer(l.data)
			for i, v := range tmpList.__roleNames_keyList() {
				out[v] = tmpList.__roleNames_atList(v, i)
			}
			return out
		}(C.MediaListTable46aa22_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackMediaListTable46aa22_ItemData
func callbackMediaListTable46aa22_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewMediaListTableFromPointer(NewMediaListTableFromPointer(nil).__itemData_newList())
			for k, v := range (*(*func(*std_core.QModelIndex) map[int]*std_core.QVariant)(signal))(std_core.NewQModelIndexFromPointer(index)) {
				tmpList.__itemData_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewMediaListTableFromPointer(NewMediaListTableFromPointer(nil).__itemData_newList())
		for k, v := range NewMediaListTableFromPointer(ptr).ItemDataDefault(std_core.NewQModelIndexFromPointer(index)) {
			tmpList.__itemData_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *MediaListTable) ItemDataDefault(index std_core.QModelIndex_ITF) map[int]*std_core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewMediaListTableFromPointer(l.data)
			for i, v := range tmpList.__itemData_keyList() {
				out[v] = tmpList.__itemData_atList(v, i)
			}
			return out
		}(C.MediaListTable46aa22_ItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*std_core.QVariant, 0)
}

//export callbackMediaListTable46aa22_MimeData
func callbackMediaListTable46aa22_MimeData(ptr unsafe.Pointer, indexes C.struct_Moc_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return std_core.PointerFromQMimeData((*(*func([]*std_core.QModelIndex) *std_core.QMimeData)(signal))(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewMediaListTableFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__mimeData_indexes_atList(i)
			}
			return out
		}(indexes)))
	}

	return std_core.PointerFromQMimeData(NewMediaListTableFromPointer(ptr).MimeDataDefault(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
		out := make([]*std_core.QModelIndex, int(l.len))
		tmpList := NewMediaListTableFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__mimeData_indexes_atList(i)
		}
		return out
	}(indexes)))
}

func (ptr *MediaListTable) MimeDataDefault(indexes []*std_core.QModelIndex) *std_core.QMimeData {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQMimeDataFromPointer(C.MediaListTable46aa22_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewMediaListTableFromPointer(NewMediaListTableFromPointer(nil).__mimeData_indexes_newList())
			for _, v := range indexes {
				tmpList.__mimeData_indexes_setList(v)
			}
			return tmpList.Pointer()
		}()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackMediaListTable46aa22_Buddy
func callbackMediaListTable46aa22_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewMediaListTableFromPointer(ptr).BuddyDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *MediaListTable) BuddyDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.MediaListTable46aa22_BuddyDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackMediaListTable46aa22_Parent
func callbackMediaListTable46aa22_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewMediaListTableFromPointer(ptr).ParentDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *MediaListTable) ParentDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.MediaListTable46aa22_ParentDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackMediaListTable46aa22_Match
func callbackMediaListTable46aa22_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewMediaListTableFromPointer(NewMediaListTableFromPointer(nil).__match_newList())
			for _, v := range (*(*func(*std_core.QModelIndex, int, *std_core.QVariant, int, std_core.Qt__MatchFlag) []*std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewMediaListTableFromPointer(NewMediaListTableFromPointer(nil).__match_newList())
		for _, v := range NewMediaListTableFromPointer(ptr).MatchDefault(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *MediaListTable) MatchDefault(start std_core.QModelIndex_ITF, role int, value std_core.QVariant_ITF, hits int, flags std_core.Qt__MatchFlag) []*std_core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewMediaListTableFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__match_atList(i)
			}
			return out
		}(C.MediaListTable46aa22_MatchDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(start), C.int(int32(role)), std_core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbackMediaListTable46aa22_Span
func callbackMediaListTable46aa22_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return std_core.PointerFromQSize((*(*func(*std_core.QModelIndex) *std_core.QSize)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQSize(NewMediaListTableFromPointer(ptr).SpanDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *MediaListTable) SpanDefault(index std_core.QModelIndex_ITF) *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.MediaListTable46aa22_SpanDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackMediaListTable46aa22_MimeTypes
func callbackMediaListTable46aa22_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := (*(*func() []string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
	}
	tempVal := NewMediaListTableFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
}

func (ptr *MediaListTable) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.MediaListTable46aa22_MimeTypesDefault(ptr.Pointer())))
	}
	return make([]string, 0)
}

//export callbackMediaListTable46aa22_Data
func callbackMediaListTable46aa22_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(*std_core.QModelIndex, int) *std_core.QVariant)(signal))(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewMediaListTableFromPointer(ptr).DataDefault(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *MediaListTable) DataDefault(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.MediaListTable46aa22_DataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackMediaListTable46aa22_HeaderData
func callbackMediaListTable46aa22_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(int, std_core.Qt__Orientation, int) *std_core.QVariant)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewMediaListTableFromPointer(ptr).HeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *MediaListTable) HeaderDataDefault(section int, orientation std_core.Qt__Orientation, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.MediaListTable46aa22_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackMediaListTable46aa22_SupportedDragActions
func callbackMediaListTable46aa22_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewMediaListTableFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *MediaListTable) SupportedDragActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.MediaListTable46aa22_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackMediaListTable46aa22_SupportedDropActions
func callbackMediaListTable46aa22_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewMediaListTableFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *MediaListTable) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.MediaListTable46aa22_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackMediaListTable46aa22_CanDropMimeData
func callbackMediaListTable46aa22_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListTableFromPointer(ptr).CanDropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *MediaListTable) CanDropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.MediaListTable46aa22_CanDropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackMediaListTable46aa22_CanFetchMore
func callbackMediaListTable46aa22_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListTableFromPointer(ptr).CanFetchMoreDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *MediaListTable) CanFetchMoreDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.MediaListTable46aa22_CanFetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackMediaListTable46aa22_HasChildren
func callbackMediaListTable46aa22_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListTableFromPointer(ptr).HasChildrenDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *MediaListTable) HasChildrenDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.MediaListTable46aa22_HasChildrenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackMediaListTable46aa22_ColumnCount
func callbackMediaListTable46aa22_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewMediaListTableFromPointer(ptr).ColumnCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *MediaListTable) ColumnCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.MediaListTable46aa22_ColumnCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackMediaListTable46aa22_RowCount
func callbackMediaListTable46aa22_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewMediaListTableFromPointer(ptr).RowCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *MediaListTable) RowCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.MediaListTable46aa22_RowCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackMediaListTable46aa22_Event
func callbackMediaListTable46aa22_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListTableFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *MediaListTable) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.MediaListTable46aa22_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackMediaListTable46aa22_EventFilter
func callbackMediaListTable46aa22_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListTableFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *MediaListTable) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.MediaListTable46aa22_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackMediaListTable46aa22_ChildEvent
func callbackMediaListTable46aa22_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewMediaListTableFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *MediaListTable) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackMediaListTable46aa22_ConnectNotify
func callbackMediaListTable46aa22_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewMediaListTableFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *MediaListTable) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackMediaListTable46aa22_CustomEvent
func callbackMediaListTable46aa22_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewMediaListTableFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *MediaListTable) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackMediaListTable46aa22_DeleteLater
func callbackMediaListTable46aa22_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListTableFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *MediaListTable) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackMediaListTable46aa22_Destroyed
func callbackMediaListTable46aa22_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackMediaListTable46aa22_DisconnectNotify
func callbackMediaListTable46aa22_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewMediaListTableFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *MediaListTable) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackMediaListTable46aa22_ObjectNameChanged
func callbackMediaListTable46aa22_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackMediaListTable46aa22_TimerEvent
func callbackMediaListTable46aa22_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewMediaListTableFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *MediaListTable) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTable46aa22_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type MediaListTabs_ITF interface {
	std_widgets.QTabWidget_ITF
	MediaListTabs_PTR() *MediaListTabs
}

func (ptr *MediaListTabs) MediaListTabs_PTR() *MediaListTabs {
	return ptr
}

func (ptr *MediaListTabs) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTabWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *MediaListTabs) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTabWidget_PTR().SetPointer(p)
	}
}

func PointerFromMediaListTabs(ptr MediaListTabs_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.MediaListTabs_PTR().Pointer()
	}
	return nil
}

func NewMediaListTabsFromPointer(ptr unsafe.Pointer) (n *MediaListTabs) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(MediaListTabs)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *MediaListTabs:
			n = deduced

		case *std_widgets.QTabWidget:
			n = &MediaListTabs{QTabWidget: *deduced}

		default:
			n = new(MediaListTabs)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *MediaListTabs) Init() { this.init() }

//export callbackMediaListTabs46aa22_Constructor
func callbackMediaListTabs46aa22_Constructor(ptr unsafe.Pointer) {
	this := NewMediaListTabsFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

func MediaListTabs_QRegisterMetaType() int {
	return int(int32(C.MediaListTabs46aa22_MediaListTabs46aa22_QRegisterMetaType()))
}

func (ptr *MediaListTabs) QRegisterMetaType() int {
	return int(int32(C.MediaListTabs46aa22_MediaListTabs46aa22_QRegisterMetaType()))
}

func MediaListTabs_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.MediaListTabs46aa22_MediaListTabs46aa22_QRegisterMetaType2(typeNameC)))
}

func (ptr *MediaListTabs) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.MediaListTabs46aa22_MediaListTabs46aa22_QRegisterMetaType2(typeNameC)))
}

func MediaListTabs_QmlRegisterType() int {
	return int(int32(C.MediaListTabs46aa22_MediaListTabs46aa22_QmlRegisterType()))
}

func (ptr *MediaListTabs) QmlRegisterType() int {
	return int(int32(C.MediaListTabs46aa22_MediaListTabs46aa22_QmlRegisterType()))
}

func MediaListTabs_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.MediaListTabs46aa22_MediaListTabs46aa22_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *MediaListTabs) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.MediaListTabs46aa22_MediaListTabs46aa22_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *MediaListTabs) __addActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.MediaListTabs46aa22___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MediaListTabs) __addActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22___addActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *MediaListTabs) __addActions_actions_newList() unsafe.Pointer {
	return C.MediaListTabs46aa22___addActions_actions_newList(ptr.Pointer())
}

func (ptr *MediaListTabs) __insertActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.MediaListTabs46aa22___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MediaListTabs) __insertActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22___insertActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *MediaListTabs) __insertActions_actions_newList() unsafe.Pointer {
	return C.MediaListTabs46aa22___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *MediaListTabs) __actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.MediaListTabs46aa22___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MediaListTabs) __actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22___actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *MediaListTabs) __actions_newList() unsafe.Pointer {
	return C.MediaListTabs46aa22___actions_newList(ptr.Pointer())
}

func (ptr *MediaListTabs) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.MediaListTabs46aa22___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *MediaListTabs) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *MediaListTabs) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.MediaListTabs46aa22___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *MediaListTabs) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.MediaListTabs46aa22___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MediaListTabs) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *MediaListTabs) __findChildren_newList2() unsafe.Pointer {
	return C.MediaListTabs46aa22___findChildren_newList2(ptr.Pointer())
}

func (ptr *MediaListTabs) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.MediaListTabs46aa22___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MediaListTabs) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *MediaListTabs) __findChildren_newList3() unsafe.Pointer {
	return C.MediaListTabs46aa22___findChildren_newList3(ptr.Pointer())
}

func (ptr *MediaListTabs) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.MediaListTabs46aa22___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MediaListTabs) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *MediaListTabs) __findChildren_newList() unsafe.Pointer {
	return C.MediaListTabs46aa22___findChildren_newList(ptr.Pointer())
}

func (ptr *MediaListTabs) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.MediaListTabs46aa22___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MediaListTabs) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *MediaListTabs) __children_newList() unsafe.Pointer {
	return C.MediaListTabs46aa22___children_newList(ptr.Pointer())
}

func NewMediaListTabs(parent std_widgets.QWidget_ITF) *MediaListTabs {
	tmpValue := NewMediaListTabsFromPointer(C.MediaListTabs46aa22_NewMediaListTabs(std_widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackMediaListTabs46aa22_DestroyMediaListTabs
func callbackMediaListTabs46aa22_DestroyMediaListTabs(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~MediaListTabs"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListTabsFromPointer(ptr).DestroyMediaListTabsDefault()
	}
}

func (ptr *MediaListTabs) ConnectDestroyMediaListTabs(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~MediaListTabs"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~MediaListTabs", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~MediaListTabs", unsafe.Pointer(&f))
		}
	}
}

func (ptr *MediaListTabs) DisconnectDestroyMediaListTabs() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~MediaListTabs")
	}
}

func (ptr *MediaListTabs) DestroyMediaListTabs() {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_DestroyMediaListTabs(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *MediaListTabs) DestroyMediaListTabsDefault() {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_DestroyMediaListTabsDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackMediaListTabs46aa22_Event
func callbackMediaListTabs46aa22_Event(ptr unsafe.Pointer, ev unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(ev)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListTabsFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(ev)))))
}

func (ptr *MediaListTabs) EventDefault(ev std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.MediaListTabs46aa22_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(ev))) != 0
	}
	return false
}

//export callbackMediaListTabs46aa22_ChangeEvent
func callbackMediaListTabs46aa22_ChangeEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(ev))
	} else {
		NewMediaListTabsFromPointer(ptr).ChangeEventDefault(std_core.NewQEventFromPointer(ev))
	}
}

func (ptr *MediaListTabs) ChangeEventDefault(ev std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_ChangeEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(ev))
	}
}

//export callbackMediaListTabs46aa22_CurrentChanged
func callbackMediaListTabs46aa22_CurrentChanged(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "currentChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

//export callbackMediaListTabs46aa22_KeyPressEvent
func callbackMediaListTabs46aa22_KeyPressEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		(*(*func(*std_gui.QKeyEvent))(signal))(std_gui.NewQKeyEventFromPointer(e))
	} else {
		NewMediaListTabsFromPointer(ptr).KeyPressEventDefault(std_gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *MediaListTabs) KeyPressEventDefault(e std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_KeyPressEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(e))
	}
}

//export callbackMediaListTabs46aa22_PaintEvent
func callbackMediaListTabs46aa22_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		(*(*func(*std_gui.QPaintEvent))(signal))(std_gui.NewQPaintEventFromPointer(event))
	} else {
		NewMediaListTabsFromPointer(ptr).PaintEventDefault(std_gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *MediaListTabs) PaintEventDefault(event std_gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_PaintEventDefault(ptr.Pointer(), std_gui.PointerFromQPaintEvent(event))
	}
}

//export callbackMediaListTabs46aa22_ResizeEvent
func callbackMediaListTabs46aa22_ResizeEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		(*(*func(*std_gui.QResizeEvent))(signal))(std_gui.NewQResizeEventFromPointer(e))
	} else {
		NewMediaListTabsFromPointer(ptr).ResizeEventDefault(std_gui.NewQResizeEventFromPointer(e))
	}
}

func (ptr *MediaListTabs) ResizeEventDefault(e std_gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_ResizeEventDefault(ptr.Pointer(), std_gui.PointerFromQResizeEvent(e))
	}
}

//export callbackMediaListTabs46aa22_SetCurrentIndex
func callbackMediaListTabs46aa22_SetCurrentIndex(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "setCurrentIndex"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	} else {
		NewMediaListTabsFromPointer(ptr).SetCurrentIndexDefault(int(int32(index)))
	}
}

func (ptr *MediaListTabs) SetCurrentIndexDefault(index int) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_SetCurrentIndexDefault(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackMediaListTabs46aa22_SetCurrentWidget
func callbackMediaListTabs46aa22_SetCurrentWidget(ptr unsafe.Pointer, widget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setCurrentWidget"); signal != nil {
		(*(*func(*std_widgets.QWidget))(signal))(std_widgets.NewQWidgetFromPointer(widget))
	} else {
		NewMediaListTabsFromPointer(ptr).SetCurrentWidgetDefault(std_widgets.NewQWidgetFromPointer(widget))
	}
}

func (ptr *MediaListTabs) SetCurrentWidgetDefault(widget std_widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_SetCurrentWidgetDefault(ptr.Pointer(), std_widgets.PointerFromQWidget(widget))
	}
}

//export callbackMediaListTabs46aa22_ShowEvent
func callbackMediaListTabs46aa22_ShowEvent(ptr unsafe.Pointer, vqs unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		(*(*func(*std_gui.QShowEvent))(signal))(std_gui.NewQShowEventFromPointer(vqs))
	} else {
		NewMediaListTabsFromPointer(ptr).ShowEventDefault(std_gui.NewQShowEventFromPointer(vqs))
	}
}

func (ptr *MediaListTabs) ShowEventDefault(vqs std_gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_ShowEventDefault(ptr.Pointer(), std_gui.PointerFromQShowEvent(vqs))
	}
}

//export callbackMediaListTabs46aa22_TabBarClicked
func callbackMediaListTabs46aa22_TabBarClicked(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "tabBarClicked"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

//export callbackMediaListTabs46aa22_TabBarDoubleClicked
func callbackMediaListTabs46aa22_TabBarDoubleClicked(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "tabBarDoubleClicked"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

//export callbackMediaListTabs46aa22_TabCloseRequested
func callbackMediaListTabs46aa22_TabCloseRequested(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "tabCloseRequested"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

//export callbackMediaListTabs46aa22_TabInserted
func callbackMediaListTabs46aa22_TabInserted(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "tabInserted"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	} else {
		NewMediaListTabsFromPointer(ptr).TabInsertedDefault(int(int32(index)))
	}
}

func (ptr *MediaListTabs) TabInsertedDefault(index int) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_TabInsertedDefault(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackMediaListTabs46aa22_TabRemoved
func callbackMediaListTabs46aa22_TabRemoved(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "tabRemoved"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	} else {
		NewMediaListTabsFromPointer(ptr).TabRemovedDefault(int(int32(index)))
	}
}

func (ptr *MediaListTabs) TabRemovedDefault(index int) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_TabRemovedDefault(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackMediaListTabs46aa22_MinimumSizeHint
func callbackMediaListTabs46aa22_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return std_core.PointerFromQSize((*(*func() *std_core.QSize)(signal))())
	}

	return std_core.PointerFromQSize(NewMediaListTabsFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *MediaListTabs) MinimumSizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.MediaListTabs46aa22_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackMediaListTabs46aa22_SizeHint
func callbackMediaListTabs46aa22_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return std_core.PointerFromQSize((*(*func() *std_core.QSize)(signal))())
	}

	return std_core.PointerFromQSize(NewMediaListTabsFromPointer(ptr).SizeHintDefault())
}

func (ptr *MediaListTabs) SizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.MediaListTabs46aa22_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackMediaListTabs46aa22_HasHeightForWidth
func callbackMediaListTabs46aa22_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListTabsFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *MediaListTabs) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.MediaListTabs46aa22_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackMediaListTabs46aa22_HeightForWidth
func callbackMediaListTabs46aa22_HeightForWidth(ptr unsafe.Pointer, width C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32((*(*func(int) int)(signal))(int(int32(width)))))
	}

	return C.int(int32(NewMediaListTabsFromPointer(ptr).HeightForWidthDefault(int(int32(width)))))
}

func (ptr *MediaListTabs) HeightForWidthDefault(width int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.MediaListTabs46aa22_HeightForWidthDefault(ptr.Pointer(), C.int(int32(width)))))
	}
	return 0
}

//export callbackMediaListTabs46aa22_Close
func callbackMediaListTabs46aa22_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListTabsFromPointer(ptr).CloseDefault())))
}

func (ptr *MediaListTabs) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.MediaListTabs46aa22_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackMediaListTabs46aa22_FocusNextPrevChild
func callbackMediaListTabs46aa22_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(bool) bool)(signal))(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListTabsFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *MediaListTabs) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.MediaListTabs46aa22_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackMediaListTabs46aa22_NativeEvent
func callbackMediaListTabs46aa22_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result *C.long) C.char {
	var resultR int
	if result != nil {
		resultR = int(int32(*result))
		defer func() { *result = C.long(int32(resultR)) }()
	}
	if signal := qt.GetSignal(ptr, "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QByteArray, unsafe.Pointer, *int) bool)(signal))(std_core.NewQByteArrayFromPointer(eventType), message, &resultR))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListTabsFromPointer(ptr).NativeEventDefault(std_core.NewQByteArrayFromPointer(eventType), message, &resultR))))
}

func (ptr *MediaListTabs) NativeEventDefault(eventType std_core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {
	if ptr.Pointer() != nil {
		var resultC C.long
		if result != nil {
			resultC = C.long(int32(*result))
			defer func() { *result = int(int32(resultC)) }()
		}
		return int8(C.MediaListTabs46aa22_NativeEventDefault(ptr.Pointer(), std_core.PointerFromQByteArray(eventType), message, &resultC)) != 0
	}
	return false
}

//export callbackMediaListTabs46aa22_ActionEvent
func callbackMediaListTabs46aa22_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		(*(*func(*std_gui.QActionEvent))(signal))(std_gui.NewQActionEventFromPointer(event))
	} else {
		NewMediaListTabsFromPointer(ptr).ActionEventDefault(std_gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *MediaListTabs) ActionEventDefault(event std_gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_ActionEventDefault(ptr.Pointer(), std_gui.PointerFromQActionEvent(event))
	}
}

//export callbackMediaListTabs46aa22_CloseEvent
func callbackMediaListTabs46aa22_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		(*(*func(*std_gui.QCloseEvent))(signal))(std_gui.NewQCloseEventFromPointer(event))
	} else {
		NewMediaListTabsFromPointer(ptr).CloseEventDefault(std_gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *MediaListTabs) CloseEventDefault(event std_gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_CloseEventDefault(ptr.Pointer(), std_gui.PointerFromQCloseEvent(event))
	}
}

//export callbackMediaListTabs46aa22_ContextMenuEvent
func callbackMediaListTabs46aa22_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		(*(*func(*std_gui.QContextMenuEvent))(signal))(std_gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewMediaListTabsFromPointer(ptr).ContextMenuEventDefault(std_gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *MediaListTabs) ContextMenuEventDefault(event std_gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_ContextMenuEventDefault(ptr.Pointer(), std_gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackMediaListTabs46aa22_CustomContextMenuRequested
func callbackMediaListTabs46aa22_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		(*(*func(*std_core.QPoint))(signal))(std_core.NewQPointFromPointer(pos))
	}

}

//export callbackMediaListTabs46aa22_DragEnterEvent
func callbackMediaListTabs46aa22_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		(*(*func(*std_gui.QDragEnterEvent))(signal))(std_gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewMediaListTabsFromPointer(ptr).DragEnterEventDefault(std_gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *MediaListTabs) DragEnterEventDefault(event std_gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_DragEnterEventDefault(ptr.Pointer(), std_gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackMediaListTabs46aa22_DragLeaveEvent
func callbackMediaListTabs46aa22_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		(*(*func(*std_gui.QDragLeaveEvent))(signal))(std_gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewMediaListTabsFromPointer(ptr).DragLeaveEventDefault(std_gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *MediaListTabs) DragLeaveEventDefault(event std_gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_DragLeaveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackMediaListTabs46aa22_DragMoveEvent
func callbackMediaListTabs46aa22_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		(*(*func(*std_gui.QDragMoveEvent))(signal))(std_gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewMediaListTabsFromPointer(ptr).DragMoveEventDefault(std_gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *MediaListTabs) DragMoveEventDefault(event std_gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_DragMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackMediaListTabs46aa22_DropEvent
func callbackMediaListTabs46aa22_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		(*(*func(*std_gui.QDropEvent))(signal))(std_gui.NewQDropEventFromPointer(event))
	} else {
		NewMediaListTabsFromPointer(ptr).DropEventDefault(std_gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *MediaListTabs) DropEventDefault(event std_gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_DropEventDefault(ptr.Pointer(), std_gui.PointerFromQDropEvent(event))
	}
}

//export callbackMediaListTabs46aa22_EnterEvent
func callbackMediaListTabs46aa22_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewMediaListTabsFromPointer(ptr).EnterEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *MediaListTabs) EnterEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_EnterEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackMediaListTabs46aa22_FocusInEvent
func callbackMediaListTabs46aa22_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		(*(*func(*std_gui.QFocusEvent))(signal))(std_gui.NewQFocusEventFromPointer(event))
	} else {
		NewMediaListTabsFromPointer(ptr).FocusInEventDefault(std_gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *MediaListTabs) FocusInEventDefault(event std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_FocusInEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(event))
	}
}

//export callbackMediaListTabs46aa22_FocusOutEvent
func callbackMediaListTabs46aa22_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		(*(*func(*std_gui.QFocusEvent))(signal))(std_gui.NewQFocusEventFromPointer(event))
	} else {
		NewMediaListTabsFromPointer(ptr).FocusOutEventDefault(std_gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *MediaListTabs) FocusOutEventDefault(event std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_FocusOutEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(event))
	}
}

//export callbackMediaListTabs46aa22_Hide
func callbackMediaListTabs46aa22_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListTabsFromPointer(ptr).HideDefault()
	}
}

func (ptr *MediaListTabs) HideDefault() {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_HideDefault(ptr.Pointer())
	}
}

//export callbackMediaListTabs46aa22_HideEvent
func callbackMediaListTabs46aa22_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		(*(*func(*std_gui.QHideEvent))(signal))(std_gui.NewQHideEventFromPointer(event))
	} else {
		NewMediaListTabsFromPointer(ptr).HideEventDefault(std_gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *MediaListTabs) HideEventDefault(event std_gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_HideEventDefault(ptr.Pointer(), std_gui.PointerFromQHideEvent(event))
	}
}

//export callbackMediaListTabs46aa22_InputMethodEvent
func callbackMediaListTabs46aa22_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		(*(*func(*std_gui.QInputMethodEvent))(signal))(std_gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewMediaListTabsFromPointer(ptr).InputMethodEventDefault(std_gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *MediaListTabs) InputMethodEventDefault(event std_gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_InputMethodEventDefault(ptr.Pointer(), std_gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackMediaListTabs46aa22_KeyReleaseEvent
func callbackMediaListTabs46aa22_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		(*(*func(*std_gui.QKeyEvent))(signal))(std_gui.NewQKeyEventFromPointer(event))
	} else {
		NewMediaListTabsFromPointer(ptr).KeyReleaseEventDefault(std_gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *MediaListTabs) KeyReleaseEventDefault(event std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_KeyReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(event))
	}
}

//export callbackMediaListTabs46aa22_LeaveEvent
func callbackMediaListTabs46aa22_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewMediaListTabsFromPointer(ptr).LeaveEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *MediaListTabs) LeaveEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_LeaveEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackMediaListTabs46aa22_Lower
func callbackMediaListTabs46aa22_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListTabsFromPointer(ptr).LowerDefault()
	}
}

func (ptr *MediaListTabs) LowerDefault() {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_LowerDefault(ptr.Pointer())
	}
}

//export callbackMediaListTabs46aa22_MouseDoubleClickEvent
func callbackMediaListTabs46aa22_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewMediaListTabsFromPointer(ptr).MouseDoubleClickEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *MediaListTabs) MouseDoubleClickEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_MouseDoubleClickEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackMediaListTabs46aa22_MouseMoveEvent
func callbackMediaListTabs46aa22_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewMediaListTabsFromPointer(ptr).MouseMoveEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *MediaListTabs) MouseMoveEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_MouseMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackMediaListTabs46aa22_MousePressEvent
func callbackMediaListTabs46aa22_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewMediaListTabsFromPointer(ptr).MousePressEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *MediaListTabs) MousePressEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_MousePressEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackMediaListTabs46aa22_MouseReleaseEvent
func callbackMediaListTabs46aa22_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewMediaListTabsFromPointer(ptr).MouseReleaseEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *MediaListTabs) MouseReleaseEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_MouseReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackMediaListTabs46aa22_MoveEvent
func callbackMediaListTabs46aa22_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		(*(*func(*std_gui.QMoveEvent))(signal))(std_gui.NewQMoveEventFromPointer(event))
	} else {
		NewMediaListTabsFromPointer(ptr).MoveEventDefault(std_gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *MediaListTabs) MoveEventDefault(event std_gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_MoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMoveEvent(event))
	}
}

//export callbackMediaListTabs46aa22_Raise
func callbackMediaListTabs46aa22_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListTabsFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *MediaListTabs) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_RaiseDefault(ptr.Pointer())
	}
}

//export callbackMediaListTabs46aa22_Repaint
func callbackMediaListTabs46aa22_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListTabsFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *MediaListTabs) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_RepaintDefault(ptr.Pointer())
	}
}

//export callbackMediaListTabs46aa22_SetDisabled
func callbackMediaListTabs46aa22_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		(*(*func(bool))(signal))(int8(disable) != 0)
	} else {
		NewMediaListTabsFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *MediaListTabs) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackMediaListTabs46aa22_SetEnabled
func callbackMediaListTabs46aa22_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewMediaListTabsFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *MediaListTabs) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackMediaListTabs46aa22_SetFocus2
func callbackMediaListTabs46aa22_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListTabsFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *MediaListTabs) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackMediaListTabs46aa22_SetHidden
func callbackMediaListTabs46aa22_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		(*(*func(bool))(signal))(int8(hidden) != 0)
	} else {
		NewMediaListTabsFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *MediaListTabs) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackMediaListTabs46aa22_SetStyleSheet
func callbackMediaListTabs46aa22_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(styleSheet))
	} else {
		NewMediaListTabsFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *MediaListTabs) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.MediaListTabs46aa22_SetStyleSheetDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackMediaListTabs46aa22_SetVisible
func callbackMediaListTabs46aa22_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	} else {
		NewMediaListTabsFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *MediaListTabs) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackMediaListTabs46aa22_SetWindowModified
func callbackMediaListTabs46aa22_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewMediaListTabsFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *MediaListTabs) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackMediaListTabs46aa22_SetWindowTitle
func callbackMediaListTabs46aa22_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(vqs))
	} else {
		NewMediaListTabsFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *MediaListTabs) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.MediaListTabs46aa22_SetWindowTitleDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackMediaListTabs46aa22_Show
func callbackMediaListTabs46aa22_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListTabsFromPointer(ptr).ShowDefault()
	}
}

func (ptr *MediaListTabs) ShowDefault() {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_ShowDefault(ptr.Pointer())
	}
}

//export callbackMediaListTabs46aa22_ShowFullScreen
func callbackMediaListTabs46aa22_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListTabsFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *MediaListTabs) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackMediaListTabs46aa22_ShowMaximized
func callbackMediaListTabs46aa22_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListTabsFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *MediaListTabs) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackMediaListTabs46aa22_ShowMinimized
func callbackMediaListTabs46aa22_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListTabsFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *MediaListTabs) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackMediaListTabs46aa22_ShowNormal
func callbackMediaListTabs46aa22_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListTabsFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *MediaListTabs) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackMediaListTabs46aa22_TabletEvent
func callbackMediaListTabs46aa22_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		(*(*func(*std_gui.QTabletEvent))(signal))(std_gui.NewQTabletEventFromPointer(event))
	} else {
		NewMediaListTabsFromPointer(ptr).TabletEventDefault(std_gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *MediaListTabs) TabletEventDefault(event std_gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_TabletEventDefault(ptr.Pointer(), std_gui.PointerFromQTabletEvent(event))
	}
}

//export callbackMediaListTabs46aa22_Update
func callbackMediaListTabs46aa22_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListTabsFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *MediaListTabs) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_UpdateDefault(ptr.Pointer())
	}
}

//export callbackMediaListTabs46aa22_UpdateMicroFocus
func callbackMediaListTabs46aa22_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListTabsFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *MediaListTabs) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackMediaListTabs46aa22_WheelEvent
func callbackMediaListTabs46aa22_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		(*(*func(*std_gui.QWheelEvent))(signal))(std_gui.NewQWheelEventFromPointer(event))
	} else {
		NewMediaListTabsFromPointer(ptr).WheelEventDefault(std_gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *MediaListTabs) WheelEventDefault(event std_gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_WheelEventDefault(ptr.Pointer(), std_gui.PointerFromQWheelEvent(event))
	}
}

//export callbackMediaListTabs46aa22_WindowIconChanged
func callbackMediaListTabs46aa22_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		(*(*func(*std_gui.QIcon))(signal))(std_gui.NewQIconFromPointer(icon))
	}

}

//export callbackMediaListTabs46aa22_WindowTitleChanged
func callbackMediaListTabs46aa22_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(title))
	}

}

//export callbackMediaListTabs46aa22_PaintEngine
func callbackMediaListTabs46aa22_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return std_gui.PointerFromQPaintEngine((*(*func() *std_gui.QPaintEngine)(signal))())
	}

	return std_gui.PointerFromQPaintEngine(NewMediaListTabsFromPointer(ptr).PaintEngineDefault())
}

func (ptr *MediaListTabs) PaintEngineDefault() *std_gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return std_gui.NewQPaintEngineFromPointer(C.MediaListTabs46aa22_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackMediaListTabs46aa22_InputMethodQuery
func callbackMediaListTabs46aa22_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(std_core.Qt__InputMethodQuery) *std_core.QVariant)(signal))(std_core.Qt__InputMethodQuery(query)))
	}

	return std_core.PointerFromQVariant(NewMediaListTabsFromPointer(ptr).InputMethodQueryDefault(std_core.Qt__InputMethodQuery(query)))
}

func (ptr *MediaListTabs) InputMethodQueryDefault(query std_core.Qt__InputMethodQuery) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.MediaListTabs46aa22_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackMediaListTabs46aa22_Metric
func callbackMediaListTabs46aa22_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32((*(*func(std_gui.QPaintDevice__PaintDeviceMetric) int)(signal))(std_gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewMediaListTabsFromPointer(ptr).MetricDefault(std_gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *MediaListTabs) MetricDefault(m std_gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.MediaListTabs46aa22_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackMediaListTabs46aa22_InitPainter
func callbackMediaListTabs46aa22_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		(*(*func(*std_gui.QPainter))(signal))(std_gui.NewQPainterFromPointer(painter))
	} else {
		NewMediaListTabsFromPointer(ptr).InitPainterDefault(std_gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *MediaListTabs) InitPainterDefault(painter std_gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_InitPainterDefault(ptr.Pointer(), std_gui.PointerFromQPainter(painter))
	}
}

//export callbackMediaListTabs46aa22_EventFilter
func callbackMediaListTabs46aa22_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListTabsFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *MediaListTabs) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.MediaListTabs46aa22_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackMediaListTabs46aa22_ChildEvent
func callbackMediaListTabs46aa22_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewMediaListTabsFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *MediaListTabs) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackMediaListTabs46aa22_ConnectNotify
func callbackMediaListTabs46aa22_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewMediaListTabsFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *MediaListTabs) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackMediaListTabs46aa22_CustomEvent
func callbackMediaListTabs46aa22_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewMediaListTabsFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *MediaListTabs) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackMediaListTabs46aa22_DeleteLater
func callbackMediaListTabs46aa22_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListTabsFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *MediaListTabs) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackMediaListTabs46aa22_Destroyed
func callbackMediaListTabs46aa22_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackMediaListTabs46aa22_DisconnectNotify
func callbackMediaListTabs46aa22_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewMediaListTabsFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *MediaListTabs) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackMediaListTabs46aa22_ObjectNameChanged
func callbackMediaListTabs46aa22_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackMediaListTabs46aa22_TimerEvent
func callbackMediaListTabs46aa22_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewMediaListTabsFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *MediaListTabs) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListTabs46aa22_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type MenuBar_ITF interface {
	std_widgets.QMenuBar_ITF
	MenuBar_PTR() *MenuBar
}

func (ptr *MenuBar) MenuBar_PTR() *MenuBar {
	return ptr
}

func (ptr *MenuBar) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QMenuBar_PTR().Pointer()
	}
	return nil
}

func (ptr *MenuBar) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QMenuBar_PTR().SetPointer(p)
	}
}

func PointerFromMenuBar(ptr MenuBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.MenuBar_PTR().Pointer()
	}
	return nil
}

func NewMenuBarFromPointer(ptr unsafe.Pointer) (n *MenuBar) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(MenuBar)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *MenuBar:
			n = deduced

		case *std_widgets.QMenuBar:
			n = &MenuBar{QMenuBar: *deduced}

		default:
			n = new(MenuBar)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *MenuBar) Init() { this.init() }

//export callbackMenuBar46aa22_Constructor
func callbackMenuBar46aa22_Constructor(ptr unsafe.Pointer) {
	this := NewMenuBarFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

func MenuBar_QRegisterMetaType() int {
	return int(int32(C.MenuBar46aa22_MenuBar46aa22_QRegisterMetaType()))
}

func (ptr *MenuBar) QRegisterMetaType() int {
	return int(int32(C.MenuBar46aa22_MenuBar46aa22_QRegisterMetaType()))
}

func MenuBar_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.MenuBar46aa22_MenuBar46aa22_QRegisterMetaType2(typeNameC)))
}

func (ptr *MenuBar) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.MenuBar46aa22_MenuBar46aa22_QRegisterMetaType2(typeNameC)))
}

func MenuBar_QmlRegisterType() int {
	return int(int32(C.MenuBar46aa22_MenuBar46aa22_QmlRegisterType()))
}

func (ptr *MenuBar) QmlRegisterType() int {
	return int(int32(C.MenuBar46aa22_MenuBar46aa22_QmlRegisterType()))
}

func MenuBar_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.MenuBar46aa22_MenuBar46aa22_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *MenuBar) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.MenuBar46aa22_MenuBar46aa22_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *MenuBar) __addActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.MenuBar46aa22___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MenuBar) __addActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22___addActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *MenuBar) __addActions_actions_newList() unsafe.Pointer {
	return C.MenuBar46aa22___addActions_actions_newList(ptr.Pointer())
}

func (ptr *MenuBar) __insertActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.MenuBar46aa22___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MenuBar) __insertActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22___insertActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *MenuBar) __insertActions_actions_newList() unsafe.Pointer {
	return C.MenuBar46aa22___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *MenuBar) __actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.MenuBar46aa22___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MenuBar) __actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22___actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *MenuBar) __actions_newList() unsafe.Pointer {
	return C.MenuBar46aa22___actions_newList(ptr.Pointer())
}

func (ptr *MenuBar) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.MenuBar46aa22___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *MenuBar) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *MenuBar) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.MenuBar46aa22___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *MenuBar) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.MenuBar46aa22___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MenuBar) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *MenuBar) __findChildren_newList2() unsafe.Pointer {
	return C.MenuBar46aa22___findChildren_newList2(ptr.Pointer())
}

func (ptr *MenuBar) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.MenuBar46aa22___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MenuBar) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *MenuBar) __findChildren_newList3() unsafe.Pointer {
	return C.MenuBar46aa22___findChildren_newList3(ptr.Pointer())
}

func (ptr *MenuBar) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.MenuBar46aa22___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MenuBar) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *MenuBar) __findChildren_newList() unsafe.Pointer {
	return C.MenuBar46aa22___findChildren_newList(ptr.Pointer())
}

func (ptr *MenuBar) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.MenuBar46aa22___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MenuBar) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *MenuBar) __children_newList() unsafe.Pointer {
	return C.MenuBar46aa22___children_newList(ptr.Pointer())
}

func NewMenuBar(parent std_widgets.QWidget_ITF) *MenuBar {
	tmpValue := NewMenuBarFromPointer(C.MenuBar46aa22_NewMenuBar(std_widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackMenuBar46aa22_DestroyMenuBar
func callbackMenuBar46aa22_DestroyMenuBar(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~MenuBar"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMenuBarFromPointer(ptr).DestroyMenuBarDefault()
	}
}

func (ptr *MenuBar) ConnectDestroyMenuBar(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~MenuBar"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~MenuBar", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~MenuBar", unsafe.Pointer(&f))
		}
	}
}

func (ptr *MenuBar) DisconnectDestroyMenuBar() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~MenuBar")
	}
}

func (ptr *MenuBar) DestroyMenuBar() {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_DestroyMenuBar(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *MenuBar) DestroyMenuBarDefault() {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_DestroyMenuBarDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackMenuBar46aa22_Event
func callbackMenuBar46aa22_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMenuBarFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *MenuBar) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.MenuBar46aa22_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackMenuBar46aa22_EventFilter
func callbackMenuBar46aa22_EventFilter(ptr unsafe.Pointer, object unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(object), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMenuBarFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(object), std_core.NewQEventFromPointer(event)))))
}

func (ptr *MenuBar) EventFilterDefault(object std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.MenuBar46aa22_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(object), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackMenuBar46aa22_ActionEvent
func callbackMenuBar46aa22_ActionEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		(*(*func(*std_gui.QActionEvent))(signal))(std_gui.NewQActionEventFromPointer(e))
	} else {
		NewMenuBarFromPointer(ptr).ActionEventDefault(std_gui.NewQActionEventFromPointer(e))
	}
}

func (ptr *MenuBar) ActionEventDefault(e std_gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_ActionEventDefault(ptr.Pointer(), std_gui.PointerFromQActionEvent(e))
	}
}

//export callbackMenuBar46aa22_ChangeEvent
func callbackMenuBar46aa22_ChangeEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(e))
	} else {
		NewMenuBarFromPointer(ptr).ChangeEventDefault(std_core.NewQEventFromPointer(e))
	}
}

func (ptr *MenuBar) ChangeEventDefault(e std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_ChangeEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))
	}
}

//export callbackMenuBar46aa22_FocusInEvent
func callbackMenuBar46aa22_FocusInEvent(ptr unsafe.Pointer, vqf unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		(*(*func(*std_gui.QFocusEvent))(signal))(std_gui.NewQFocusEventFromPointer(vqf))
	} else {
		NewMenuBarFromPointer(ptr).FocusInEventDefault(std_gui.NewQFocusEventFromPointer(vqf))
	}
}

func (ptr *MenuBar) FocusInEventDefault(vqf std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_FocusInEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(vqf))
	}
}

//export callbackMenuBar46aa22_FocusOutEvent
func callbackMenuBar46aa22_FocusOutEvent(ptr unsafe.Pointer, vqf unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		(*(*func(*std_gui.QFocusEvent))(signal))(std_gui.NewQFocusEventFromPointer(vqf))
	} else {
		NewMenuBarFromPointer(ptr).FocusOutEventDefault(std_gui.NewQFocusEventFromPointer(vqf))
	}
}

func (ptr *MenuBar) FocusOutEventDefault(vqf std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_FocusOutEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(vqf))
	}
}

//export callbackMenuBar46aa22_Hovered
func callbackMenuBar46aa22_Hovered(ptr unsafe.Pointer, action unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hovered"); signal != nil {
		(*(*func(*std_widgets.QAction))(signal))(std_widgets.NewQActionFromPointer(action))
	}

}

//export callbackMenuBar46aa22_KeyPressEvent
func callbackMenuBar46aa22_KeyPressEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		(*(*func(*std_gui.QKeyEvent))(signal))(std_gui.NewQKeyEventFromPointer(e))
	} else {
		NewMenuBarFromPointer(ptr).KeyPressEventDefault(std_gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *MenuBar) KeyPressEventDefault(e std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_KeyPressEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(e))
	}
}

//export callbackMenuBar46aa22_LeaveEvent
func callbackMenuBar46aa22_LeaveEvent(ptr unsafe.Pointer, vqe unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(vqe))
	} else {
		NewMenuBarFromPointer(ptr).LeaveEventDefault(std_core.NewQEventFromPointer(vqe))
	}
}

func (ptr *MenuBar) LeaveEventDefault(vqe std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_LeaveEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(vqe))
	}
}

//export callbackMenuBar46aa22_MouseMoveEvent
func callbackMenuBar46aa22_MouseMoveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(e))
	} else {
		NewMenuBarFromPointer(ptr).MouseMoveEventDefault(std_gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *MenuBar) MouseMoveEventDefault(e std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_MouseMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(e))
	}
}

//export callbackMenuBar46aa22_MousePressEvent
func callbackMenuBar46aa22_MousePressEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(e))
	} else {
		NewMenuBarFromPointer(ptr).MousePressEventDefault(std_gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *MenuBar) MousePressEventDefault(e std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_MousePressEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(e))
	}
}

//export callbackMenuBar46aa22_MouseReleaseEvent
func callbackMenuBar46aa22_MouseReleaseEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(e))
	} else {
		NewMenuBarFromPointer(ptr).MouseReleaseEventDefault(std_gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *MenuBar) MouseReleaseEventDefault(e std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_MouseReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(e))
	}
}

//export callbackMenuBar46aa22_PaintEvent
func callbackMenuBar46aa22_PaintEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		(*(*func(*std_gui.QPaintEvent))(signal))(std_gui.NewQPaintEventFromPointer(e))
	} else {
		NewMenuBarFromPointer(ptr).PaintEventDefault(std_gui.NewQPaintEventFromPointer(e))
	}
}

func (ptr *MenuBar) PaintEventDefault(e std_gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_PaintEventDefault(ptr.Pointer(), std_gui.PointerFromQPaintEvent(e))
	}
}

//export callbackMenuBar46aa22_ResizeEvent
func callbackMenuBar46aa22_ResizeEvent(ptr unsafe.Pointer, vqr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		(*(*func(*std_gui.QResizeEvent))(signal))(std_gui.NewQResizeEventFromPointer(vqr))
	} else {
		NewMenuBarFromPointer(ptr).ResizeEventDefault(std_gui.NewQResizeEventFromPointer(vqr))
	}
}

func (ptr *MenuBar) ResizeEventDefault(vqr std_gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_ResizeEventDefault(ptr.Pointer(), std_gui.PointerFromQResizeEvent(vqr))
	}
}

//export callbackMenuBar46aa22_SetVisible
func callbackMenuBar46aa22_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	} else {
		NewMenuBarFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *MenuBar) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackMenuBar46aa22_TimerEvent
func callbackMenuBar46aa22_TimerEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(e))
	} else {
		NewMenuBarFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(e))
	}
}

func (ptr *MenuBar) TimerEventDefault(e std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(e))
	}
}

//export callbackMenuBar46aa22_Triggered
func callbackMenuBar46aa22_Triggered(ptr unsafe.Pointer, action unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "triggered"); signal != nil {
		(*(*func(*std_widgets.QAction))(signal))(std_widgets.NewQActionFromPointer(action))
	}

}

//export callbackMenuBar46aa22_MinimumSizeHint
func callbackMenuBar46aa22_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return std_core.PointerFromQSize((*(*func() *std_core.QSize)(signal))())
	}

	return std_core.PointerFromQSize(NewMenuBarFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *MenuBar) MinimumSizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.MenuBar46aa22_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackMenuBar46aa22_SizeHint
func callbackMenuBar46aa22_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return std_core.PointerFromQSize((*(*func() *std_core.QSize)(signal))())
	}

	return std_core.PointerFromQSize(NewMenuBarFromPointer(ptr).SizeHintDefault())
}

func (ptr *MenuBar) SizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.MenuBar46aa22_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackMenuBar46aa22_HeightForWidth
func callbackMenuBar46aa22_HeightForWidth(ptr unsafe.Pointer, vin C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32((*(*func(int) int)(signal))(int(int32(vin)))))
	}

	return C.int(int32(NewMenuBarFromPointer(ptr).HeightForWidthDefault(int(int32(vin)))))
}

func (ptr *MenuBar) HeightForWidthDefault(vin int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.MenuBar46aa22_HeightForWidthDefault(ptr.Pointer(), C.int(int32(vin)))))
	}
	return 0
}

//export callbackMenuBar46aa22_Close
func callbackMenuBar46aa22_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewMenuBarFromPointer(ptr).CloseDefault())))
}

func (ptr *MenuBar) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.MenuBar46aa22_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackMenuBar46aa22_FocusNextPrevChild
func callbackMenuBar46aa22_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(bool) bool)(signal))(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMenuBarFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *MenuBar) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.MenuBar46aa22_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackMenuBar46aa22_NativeEvent
func callbackMenuBar46aa22_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result *C.long) C.char {
	var resultR int
	if result != nil {
		resultR = int(int32(*result))
		defer func() { *result = C.long(int32(resultR)) }()
	}
	if signal := qt.GetSignal(ptr, "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QByteArray, unsafe.Pointer, *int) bool)(signal))(std_core.NewQByteArrayFromPointer(eventType), message, &resultR))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMenuBarFromPointer(ptr).NativeEventDefault(std_core.NewQByteArrayFromPointer(eventType), message, &resultR))))
}

func (ptr *MenuBar) NativeEventDefault(eventType std_core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {
	if ptr.Pointer() != nil {
		var resultC C.long
		if result != nil {
			resultC = C.long(int32(*result))
			defer func() { *result = int(int32(resultC)) }()
		}
		return int8(C.MenuBar46aa22_NativeEventDefault(ptr.Pointer(), std_core.PointerFromQByteArray(eventType), message, &resultC)) != 0
	}
	return false
}

//export callbackMenuBar46aa22_CloseEvent
func callbackMenuBar46aa22_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		(*(*func(*std_gui.QCloseEvent))(signal))(std_gui.NewQCloseEventFromPointer(event))
	} else {
		NewMenuBarFromPointer(ptr).CloseEventDefault(std_gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *MenuBar) CloseEventDefault(event std_gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_CloseEventDefault(ptr.Pointer(), std_gui.PointerFromQCloseEvent(event))
	}
}

//export callbackMenuBar46aa22_ContextMenuEvent
func callbackMenuBar46aa22_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		(*(*func(*std_gui.QContextMenuEvent))(signal))(std_gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewMenuBarFromPointer(ptr).ContextMenuEventDefault(std_gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *MenuBar) ContextMenuEventDefault(event std_gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_ContextMenuEventDefault(ptr.Pointer(), std_gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackMenuBar46aa22_CustomContextMenuRequested
func callbackMenuBar46aa22_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		(*(*func(*std_core.QPoint))(signal))(std_core.NewQPointFromPointer(pos))
	}

}

//export callbackMenuBar46aa22_DragEnterEvent
func callbackMenuBar46aa22_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		(*(*func(*std_gui.QDragEnterEvent))(signal))(std_gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewMenuBarFromPointer(ptr).DragEnterEventDefault(std_gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *MenuBar) DragEnterEventDefault(event std_gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_DragEnterEventDefault(ptr.Pointer(), std_gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackMenuBar46aa22_DragLeaveEvent
func callbackMenuBar46aa22_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		(*(*func(*std_gui.QDragLeaveEvent))(signal))(std_gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewMenuBarFromPointer(ptr).DragLeaveEventDefault(std_gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *MenuBar) DragLeaveEventDefault(event std_gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_DragLeaveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackMenuBar46aa22_DragMoveEvent
func callbackMenuBar46aa22_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		(*(*func(*std_gui.QDragMoveEvent))(signal))(std_gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewMenuBarFromPointer(ptr).DragMoveEventDefault(std_gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *MenuBar) DragMoveEventDefault(event std_gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_DragMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackMenuBar46aa22_DropEvent
func callbackMenuBar46aa22_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		(*(*func(*std_gui.QDropEvent))(signal))(std_gui.NewQDropEventFromPointer(event))
	} else {
		NewMenuBarFromPointer(ptr).DropEventDefault(std_gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *MenuBar) DropEventDefault(event std_gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_DropEventDefault(ptr.Pointer(), std_gui.PointerFromQDropEvent(event))
	}
}

//export callbackMenuBar46aa22_EnterEvent
func callbackMenuBar46aa22_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewMenuBarFromPointer(ptr).EnterEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *MenuBar) EnterEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_EnterEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackMenuBar46aa22_Hide
func callbackMenuBar46aa22_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMenuBarFromPointer(ptr).HideDefault()
	}
}

func (ptr *MenuBar) HideDefault() {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_HideDefault(ptr.Pointer())
	}
}

//export callbackMenuBar46aa22_HideEvent
func callbackMenuBar46aa22_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		(*(*func(*std_gui.QHideEvent))(signal))(std_gui.NewQHideEventFromPointer(event))
	} else {
		NewMenuBarFromPointer(ptr).HideEventDefault(std_gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *MenuBar) HideEventDefault(event std_gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_HideEventDefault(ptr.Pointer(), std_gui.PointerFromQHideEvent(event))
	}
}

//export callbackMenuBar46aa22_InputMethodEvent
func callbackMenuBar46aa22_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		(*(*func(*std_gui.QInputMethodEvent))(signal))(std_gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewMenuBarFromPointer(ptr).InputMethodEventDefault(std_gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *MenuBar) InputMethodEventDefault(event std_gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_InputMethodEventDefault(ptr.Pointer(), std_gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackMenuBar46aa22_KeyReleaseEvent
func callbackMenuBar46aa22_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		(*(*func(*std_gui.QKeyEvent))(signal))(std_gui.NewQKeyEventFromPointer(event))
	} else {
		NewMenuBarFromPointer(ptr).KeyReleaseEventDefault(std_gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *MenuBar) KeyReleaseEventDefault(event std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_KeyReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(event))
	}
}

//export callbackMenuBar46aa22_Lower
func callbackMenuBar46aa22_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMenuBarFromPointer(ptr).LowerDefault()
	}
}

func (ptr *MenuBar) LowerDefault() {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_LowerDefault(ptr.Pointer())
	}
}

//export callbackMenuBar46aa22_MouseDoubleClickEvent
func callbackMenuBar46aa22_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewMenuBarFromPointer(ptr).MouseDoubleClickEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *MenuBar) MouseDoubleClickEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_MouseDoubleClickEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackMenuBar46aa22_MoveEvent
func callbackMenuBar46aa22_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		(*(*func(*std_gui.QMoveEvent))(signal))(std_gui.NewQMoveEventFromPointer(event))
	} else {
		NewMenuBarFromPointer(ptr).MoveEventDefault(std_gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *MenuBar) MoveEventDefault(event std_gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_MoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMoveEvent(event))
	}
}

//export callbackMenuBar46aa22_Raise
func callbackMenuBar46aa22_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMenuBarFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *MenuBar) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_RaiseDefault(ptr.Pointer())
	}
}

//export callbackMenuBar46aa22_Repaint
func callbackMenuBar46aa22_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMenuBarFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *MenuBar) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_RepaintDefault(ptr.Pointer())
	}
}

//export callbackMenuBar46aa22_SetDisabled
func callbackMenuBar46aa22_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		(*(*func(bool))(signal))(int8(disable) != 0)
	} else {
		NewMenuBarFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *MenuBar) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackMenuBar46aa22_SetEnabled
func callbackMenuBar46aa22_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewMenuBarFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *MenuBar) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackMenuBar46aa22_SetFocus2
func callbackMenuBar46aa22_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMenuBarFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *MenuBar) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackMenuBar46aa22_SetHidden
func callbackMenuBar46aa22_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		(*(*func(bool))(signal))(int8(hidden) != 0)
	} else {
		NewMenuBarFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *MenuBar) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackMenuBar46aa22_SetStyleSheet
func callbackMenuBar46aa22_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(styleSheet))
	} else {
		NewMenuBarFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *MenuBar) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.MenuBar46aa22_SetStyleSheetDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackMenuBar46aa22_SetWindowModified
func callbackMenuBar46aa22_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewMenuBarFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *MenuBar) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackMenuBar46aa22_SetWindowTitle
func callbackMenuBar46aa22_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(vqs))
	} else {
		NewMenuBarFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *MenuBar) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.MenuBar46aa22_SetWindowTitleDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackMenuBar46aa22_Show
func callbackMenuBar46aa22_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMenuBarFromPointer(ptr).ShowDefault()
	}
}

func (ptr *MenuBar) ShowDefault() {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_ShowDefault(ptr.Pointer())
	}
}

//export callbackMenuBar46aa22_ShowEvent
func callbackMenuBar46aa22_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		(*(*func(*std_gui.QShowEvent))(signal))(std_gui.NewQShowEventFromPointer(event))
	} else {
		NewMenuBarFromPointer(ptr).ShowEventDefault(std_gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *MenuBar) ShowEventDefault(event std_gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_ShowEventDefault(ptr.Pointer(), std_gui.PointerFromQShowEvent(event))
	}
}

//export callbackMenuBar46aa22_ShowFullScreen
func callbackMenuBar46aa22_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMenuBarFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *MenuBar) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackMenuBar46aa22_ShowMaximized
func callbackMenuBar46aa22_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMenuBarFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *MenuBar) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackMenuBar46aa22_ShowMinimized
func callbackMenuBar46aa22_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMenuBarFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *MenuBar) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackMenuBar46aa22_ShowNormal
func callbackMenuBar46aa22_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMenuBarFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *MenuBar) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackMenuBar46aa22_TabletEvent
func callbackMenuBar46aa22_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		(*(*func(*std_gui.QTabletEvent))(signal))(std_gui.NewQTabletEventFromPointer(event))
	} else {
		NewMenuBarFromPointer(ptr).TabletEventDefault(std_gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *MenuBar) TabletEventDefault(event std_gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_TabletEventDefault(ptr.Pointer(), std_gui.PointerFromQTabletEvent(event))
	}
}

//export callbackMenuBar46aa22_Update
func callbackMenuBar46aa22_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMenuBarFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *MenuBar) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_UpdateDefault(ptr.Pointer())
	}
}

//export callbackMenuBar46aa22_UpdateMicroFocus
func callbackMenuBar46aa22_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMenuBarFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *MenuBar) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackMenuBar46aa22_WheelEvent
func callbackMenuBar46aa22_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		(*(*func(*std_gui.QWheelEvent))(signal))(std_gui.NewQWheelEventFromPointer(event))
	} else {
		NewMenuBarFromPointer(ptr).WheelEventDefault(std_gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *MenuBar) WheelEventDefault(event std_gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_WheelEventDefault(ptr.Pointer(), std_gui.PointerFromQWheelEvent(event))
	}
}

//export callbackMenuBar46aa22_WindowIconChanged
func callbackMenuBar46aa22_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		(*(*func(*std_gui.QIcon))(signal))(std_gui.NewQIconFromPointer(icon))
	}

}

//export callbackMenuBar46aa22_WindowTitleChanged
func callbackMenuBar46aa22_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(title))
	}

}

//export callbackMenuBar46aa22_PaintEngine
func callbackMenuBar46aa22_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return std_gui.PointerFromQPaintEngine((*(*func() *std_gui.QPaintEngine)(signal))())
	}

	return std_gui.PointerFromQPaintEngine(NewMenuBarFromPointer(ptr).PaintEngineDefault())
}

func (ptr *MenuBar) PaintEngineDefault() *std_gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return std_gui.NewQPaintEngineFromPointer(C.MenuBar46aa22_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackMenuBar46aa22_InputMethodQuery
func callbackMenuBar46aa22_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(std_core.Qt__InputMethodQuery) *std_core.QVariant)(signal))(std_core.Qt__InputMethodQuery(query)))
	}

	return std_core.PointerFromQVariant(NewMenuBarFromPointer(ptr).InputMethodQueryDefault(std_core.Qt__InputMethodQuery(query)))
}

func (ptr *MenuBar) InputMethodQueryDefault(query std_core.Qt__InputMethodQuery) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.MenuBar46aa22_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackMenuBar46aa22_HasHeightForWidth
func callbackMenuBar46aa22_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewMenuBarFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *MenuBar) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.MenuBar46aa22_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackMenuBar46aa22_Metric
func callbackMenuBar46aa22_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32((*(*func(std_gui.QPaintDevice__PaintDeviceMetric) int)(signal))(std_gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewMenuBarFromPointer(ptr).MetricDefault(std_gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *MenuBar) MetricDefault(m std_gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.MenuBar46aa22_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackMenuBar46aa22_InitPainter
func callbackMenuBar46aa22_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		(*(*func(*std_gui.QPainter))(signal))(std_gui.NewQPainterFromPointer(painter))
	} else {
		NewMenuBarFromPointer(ptr).InitPainterDefault(std_gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *MenuBar) InitPainterDefault(painter std_gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_InitPainterDefault(ptr.Pointer(), std_gui.PointerFromQPainter(painter))
	}
}

//export callbackMenuBar46aa22_ChildEvent
func callbackMenuBar46aa22_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewMenuBarFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *MenuBar) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackMenuBar46aa22_ConnectNotify
func callbackMenuBar46aa22_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewMenuBarFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *MenuBar) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackMenuBar46aa22_CustomEvent
func callbackMenuBar46aa22_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewMenuBarFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *MenuBar) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackMenuBar46aa22_DeleteLater
func callbackMenuBar46aa22_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMenuBarFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *MenuBar) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackMenuBar46aa22_Destroyed
func callbackMenuBar46aa22_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackMenuBar46aa22_DisconnectNotify
func callbackMenuBar46aa22_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewMenuBarFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *MenuBar) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.MenuBar46aa22_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackMenuBar46aa22_ObjectNameChanged
func callbackMenuBar46aa22_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

type Navigation_ITF interface {
	std_widgets.QWidget_ITF
	Navigation_PTR() *Navigation
}

func (ptr *Navigation) Navigation_PTR() *Navigation {
	return ptr
}

func (ptr *Navigation) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *Navigation) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromNavigation(ptr Navigation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Navigation_PTR().Pointer()
	}
	return nil
}

func NewNavigationFromPointer(ptr unsafe.Pointer) (n *Navigation) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(Navigation)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *Navigation:
			n = deduced

		case *std_widgets.QWidget:
			n = &Navigation{QWidget: *deduced}

		default:
			n = new(Navigation)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *Navigation) Init() { this.init() }

//export callbackNavigation46aa22_Constructor
func callbackNavigation46aa22_Constructor(ptr unsafe.Pointer) {
	this := NewNavigationFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

func Navigation_QRegisterMetaType() int {
	return int(int32(C.Navigation46aa22_Navigation46aa22_QRegisterMetaType()))
}

func (ptr *Navigation) QRegisterMetaType() int {
	return int(int32(C.Navigation46aa22_Navigation46aa22_QRegisterMetaType()))
}

func Navigation_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Navigation46aa22_Navigation46aa22_QRegisterMetaType2(typeNameC)))
}

func (ptr *Navigation) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Navigation46aa22_Navigation46aa22_QRegisterMetaType2(typeNameC)))
}

func Navigation_QmlRegisterType() int {
	return int(int32(C.Navigation46aa22_Navigation46aa22_QmlRegisterType()))
}

func (ptr *Navigation) QmlRegisterType() int {
	return int(int32(C.Navigation46aa22_Navigation46aa22_QmlRegisterType()))
}

func Navigation_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.Navigation46aa22_Navigation46aa22_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Navigation) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.Navigation46aa22_Navigation46aa22_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Navigation) __addActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.Navigation46aa22___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Navigation) __addActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22___addActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *Navigation) __addActions_actions_newList() unsafe.Pointer {
	return C.Navigation46aa22___addActions_actions_newList(ptr.Pointer())
}

func (ptr *Navigation) __insertActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.Navigation46aa22___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Navigation) __insertActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22___insertActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *Navigation) __insertActions_actions_newList() unsafe.Pointer {
	return C.Navigation46aa22___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *Navigation) __actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.Navigation46aa22___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Navigation) __actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22___actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *Navigation) __actions_newList() unsafe.Pointer {
	return C.Navigation46aa22___actions_newList(ptr.Pointer())
}

func (ptr *Navigation) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.Navigation46aa22___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *Navigation) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *Navigation) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.Navigation46aa22___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *Navigation) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Navigation46aa22___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Navigation) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Navigation) __findChildren_newList2() unsafe.Pointer {
	return C.Navigation46aa22___findChildren_newList2(ptr.Pointer())
}

func (ptr *Navigation) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Navigation46aa22___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Navigation) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Navigation) __findChildren_newList3() unsafe.Pointer {
	return C.Navigation46aa22___findChildren_newList3(ptr.Pointer())
}

func (ptr *Navigation) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Navigation46aa22___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Navigation) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Navigation) __findChildren_newList() unsafe.Pointer {
	return C.Navigation46aa22___findChildren_newList(ptr.Pointer())
}

func (ptr *Navigation) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Navigation46aa22___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Navigation) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Navigation) __children_newList() unsafe.Pointer {
	return C.Navigation46aa22___children_newList(ptr.Pointer())
}

func NewNavigation(parent std_widgets.QWidget_ITF, ff std_core.Qt__WindowType) *Navigation {
	tmpValue := NewNavigationFromPointer(C.Navigation46aa22_NewNavigation(std_widgets.PointerFromQWidget(parent), C.longlong(ff)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackNavigation46aa22_DestroyNavigation
func callbackNavigation46aa22_DestroyNavigation(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~Navigation"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewNavigationFromPointer(ptr).DestroyNavigationDefault()
	}
}

func (ptr *Navigation) ConnectDestroyNavigation(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~Navigation"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~Navigation", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~Navigation", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Navigation) DisconnectDestroyNavigation() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~Navigation")
	}
}

func (ptr *Navigation) DestroyNavigation() {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_DestroyNavigation(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *Navigation) DestroyNavigationDefault() {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_DestroyNavigationDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackNavigation46aa22_Close
func callbackNavigation46aa22_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewNavigationFromPointer(ptr).CloseDefault())))
}

func (ptr *Navigation) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.Navigation46aa22_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackNavigation46aa22_Event
func callbackNavigation46aa22_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewNavigationFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(event)))))
}

func (ptr *Navigation) EventDefault(event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.Navigation46aa22_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackNavigation46aa22_FocusNextPrevChild
func callbackNavigation46aa22_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(bool) bool)(signal))(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewNavigationFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *Navigation) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.Navigation46aa22_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackNavigation46aa22_NativeEvent
func callbackNavigation46aa22_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result *C.long) C.char {
	var resultR int
	if result != nil {
		resultR = int(int32(*result))
		defer func() { *result = C.long(int32(resultR)) }()
	}
	if signal := qt.GetSignal(ptr, "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QByteArray, unsafe.Pointer, *int) bool)(signal))(std_core.NewQByteArrayFromPointer(eventType), message, &resultR))))
	}

	return C.char(int8(qt.GoBoolToInt(NewNavigationFromPointer(ptr).NativeEventDefault(std_core.NewQByteArrayFromPointer(eventType), message, &resultR))))
}

func (ptr *Navigation) NativeEventDefault(eventType std_core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {
	if ptr.Pointer() != nil {
		var resultC C.long
		if result != nil {
			resultC = C.long(int32(*result))
			defer func() { *result = int(int32(resultC)) }()
		}
		return int8(C.Navigation46aa22_NativeEventDefault(ptr.Pointer(), std_core.PointerFromQByteArray(eventType), message, &resultC)) != 0
	}
	return false
}

//export callbackNavigation46aa22_ActionEvent
func callbackNavigation46aa22_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		(*(*func(*std_gui.QActionEvent))(signal))(std_gui.NewQActionEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).ActionEventDefault(std_gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *Navigation) ActionEventDefault(event std_gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_ActionEventDefault(ptr.Pointer(), std_gui.PointerFromQActionEvent(event))
	}
}

//export callbackNavigation46aa22_ChangeEvent
func callbackNavigation46aa22_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).ChangeEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *Navigation) ChangeEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_ChangeEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackNavigation46aa22_CloseEvent
func callbackNavigation46aa22_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		(*(*func(*std_gui.QCloseEvent))(signal))(std_gui.NewQCloseEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).CloseEventDefault(std_gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *Navigation) CloseEventDefault(event std_gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_CloseEventDefault(ptr.Pointer(), std_gui.PointerFromQCloseEvent(event))
	}
}

//export callbackNavigation46aa22_ContextMenuEvent
func callbackNavigation46aa22_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		(*(*func(*std_gui.QContextMenuEvent))(signal))(std_gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).ContextMenuEventDefault(std_gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *Navigation) ContextMenuEventDefault(event std_gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_ContextMenuEventDefault(ptr.Pointer(), std_gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackNavigation46aa22_CustomContextMenuRequested
func callbackNavigation46aa22_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		(*(*func(*std_core.QPoint))(signal))(std_core.NewQPointFromPointer(pos))
	}

}

//export callbackNavigation46aa22_DragEnterEvent
func callbackNavigation46aa22_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		(*(*func(*std_gui.QDragEnterEvent))(signal))(std_gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).DragEnterEventDefault(std_gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *Navigation) DragEnterEventDefault(event std_gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_DragEnterEventDefault(ptr.Pointer(), std_gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackNavigation46aa22_DragLeaveEvent
func callbackNavigation46aa22_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		(*(*func(*std_gui.QDragLeaveEvent))(signal))(std_gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).DragLeaveEventDefault(std_gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *Navigation) DragLeaveEventDefault(event std_gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_DragLeaveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackNavigation46aa22_DragMoveEvent
func callbackNavigation46aa22_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		(*(*func(*std_gui.QDragMoveEvent))(signal))(std_gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).DragMoveEventDefault(std_gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *Navigation) DragMoveEventDefault(event std_gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_DragMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackNavigation46aa22_DropEvent
func callbackNavigation46aa22_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		(*(*func(*std_gui.QDropEvent))(signal))(std_gui.NewQDropEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).DropEventDefault(std_gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *Navigation) DropEventDefault(event std_gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_DropEventDefault(ptr.Pointer(), std_gui.PointerFromQDropEvent(event))
	}
}

//export callbackNavigation46aa22_EnterEvent
func callbackNavigation46aa22_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).EnterEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *Navigation) EnterEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_EnterEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackNavigation46aa22_FocusInEvent
func callbackNavigation46aa22_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		(*(*func(*std_gui.QFocusEvent))(signal))(std_gui.NewQFocusEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).FocusInEventDefault(std_gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *Navigation) FocusInEventDefault(event std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_FocusInEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(event))
	}
}

//export callbackNavigation46aa22_FocusOutEvent
func callbackNavigation46aa22_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		(*(*func(*std_gui.QFocusEvent))(signal))(std_gui.NewQFocusEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).FocusOutEventDefault(std_gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *Navigation) FocusOutEventDefault(event std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_FocusOutEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(event))
	}
}

//export callbackNavigation46aa22_Hide
func callbackNavigation46aa22_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewNavigationFromPointer(ptr).HideDefault()
	}
}

func (ptr *Navigation) HideDefault() {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_HideDefault(ptr.Pointer())
	}
}

//export callbackNavigation46aa22_HideEvent
func callbackNavigation46aa22_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		(*(*func(*std_gui.QHideEvent))(signal))(std_gui.NewQHideEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).HideEventDefault(std_gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *Navigation) HideEventDefault(event std_gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_HideEventDefault(ptr.Pointer(), std_gui.PointerFromQHideEvent(event))
	}
}

//export callbackNavigation46aa22_InputMethodEvent
func callbackNavigation46aa22_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		(*(*func(*std_gui.QInputMethodEvent))(signal))(std_gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).InputMethodEventDefault(std_gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *Navigation) InputMethodEventDefault(event std_gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_InputMethodEventDefault(ptr.Pointer(), std_gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackNavigation46aa22_KeyPressEvent
func callbackNavigation46aa22_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		(*(*func(*std_gui.QKeyEvent))(signal))(std_gui.NewQKeyEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).KeyPressEventDefault(std_gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *Navigation) KeyPressEventDefault(event std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_KeyPressEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(event))
	}
}

//export callbackNavigation46aa22_KeyReleaseEvent
func callbackNavigation46aa22_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		(*(*func(*std_gui.QKeyEvent))(signal))(std_gui.NewQKeyEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).KeyReleaseEventDefault(std_gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *Navigation) KeyReleaseEventDefault(event std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_KeyReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(event))
	}
}

//export callbackNavigation46aa22_LeaveEvent
func callbackNavigation46aa22_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).LeaveEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *Navigation) LeaveEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_LeaveEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackNavigation46aa22_Lower
func callbackNavigation46aa22_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewNavigationFromPointer(ptr).LowerDefault()
	}
}

func (ptr *Navigation) LowerDefault() {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_LowerDefault(ptr.Pointer())
	}
}

//export callbackNavigation46aa22_MouseDoubleClickEvent
func callbackNavigation46aa22_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).MouseDoubleClickEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *Navigation) MouseDoubleClickEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_MouseDoubleClickEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackNavigation46aa22_MouseMoveEvent
func callbackNavigation46aa22_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).MouseMoveEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *Navigation) MouseMoveEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_MouseMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackNavigation46aa22_MousePressEvent
func callbackNavigation46aa22_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).MousePressEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *Navigation) MousePressEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_MousePressEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackNavigation46aa22_MouseReleaseEvent
func callbackNavigation46aa22_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).MouseReleaseEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *Navigation) MouseReleaseEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_MouseReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackNavigation46aa22_MoveEvent
func callbackNavigation46aa22_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		(*(*func(*std_gui.QMoveEvent))(signal))(std_gui.NewQMoveEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).MoveEventDefault(std_gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *Navigation) MoveEventDefault(event std_gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_MoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMoveEvent(event))
	}
}

//export callbackNavigation46aa22_PaintEvent
func callbackNavigation46aa22_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		(*(*func(*std_gui.QPaintEvent))(signal))(std_gui.NewQPaintEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).PaintEventDefault(std_gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *Navigation) PaintEventDefault(event std_gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_PaintEventDefault(ptr.Pointer(), std_gui.PointerFromQPaintEvent(event))
	}
}

//export callbackNavigation46aa22_Raise
func callbackNavigation46aa22_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewNavigationFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *Navigation) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_RaiseDefault(ptr.Pointer())
	}
}

//export callbackNavigation46aa22_Repaint
func callbackNavigation46aa22_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewNavigationFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *Navigation) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_RepaintDefault(ptr.Pointer())
	}
}

//export callbackNavigation46aa22_ResizeEvent
func callbackNavigation46aa22_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		(*(*func(*std_gui.QResizeEvent))(signal))(std_gui.NewQResizeEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).ResizeEventDefault(std_gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *Navigation) ResizeEventDefault(event std_gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_ResizeEventDefault(ptr.Pointer(), std_gui.PointerFromQResizeEvent(event))
	}
}

//export callbackNavigation46aa22_SetDisabled
func callbackNavigation46aa22_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		(*(*func(bool))(signal))(int8(disable) != 0)
	} else {
		NewNavigationFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *Navigation) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackNavigation46aa22_SetEnabled
func callbackNavigation46aa22_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewNavigationFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *Navigation) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackNavigation46aa22_SetFocus2
func callbackNavigation46aa22_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewNavigationFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *Navigation) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackNavigation46aa22_SetHidden
func callbackNavigation46aa22_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		(*(*func(bool))(signal))(int8(hidden) != 0)
	} else {
		NewNavigationFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *Navigation) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackNavigation46aa22_SetStyleSheet
func callbackNavigation46aa22_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(styleSheet))
	} else {
		NewNavigationFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *Navigation) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.Navigation46aa22_SetStyleSheetDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackNavigation46aa22_SetVisible
func callbackNavigation46aa22_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	} else {
		NewNavigationFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *Navigation) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackNavigation46aa22_SetWindowModified
func callbackNavigation46aa22_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewNavigationFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *Navigation) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackNavigation46aa22_SetWindowTitle
func callbackNavigation46aa22_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(vqs))
	} else {
		NewNavigationFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *Navigation) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.Navigation46aa22_SetWindowTitleDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackNavigation46aa22_Show
func callbackNavigation46aa22_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewNavigationFromPointer(ptr).ShowDefault()
	}
}

func (ptr *Navigation) ShowDefault() {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_ShowDefault(ptr.Pointer())
	}
}

//export callbackNavigation46aa22_ShowEvent
func callbackNavigation46aa22_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		(*(*func(*std_gui.QShowEvent))(signal))(std_gui.NewQShowEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).ShowEventDefault(std_gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *Navigation) ShowEventDefault(event std_gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_ShowEventDefault(ptr.Pointer(), std_gui.PointerFromQShowEvent(event))
	}
}

//export callbackNavigation46aa22_ShowFullScreen
func callbackNavigation46aa22_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewNavigationFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *Navigation) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackNavigation46aa22_ShowMaximized
func callbackNavigation46aa22_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewNavigationFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *Navigation) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackNavigation46aa22_ShowMinimized
func callbackNavigation46aa22_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewNavigationFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *Navigation) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackNavigation46aa22_ShowNormal
func callbackNavigation46aa22_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewNavigationFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *Navigation) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackNavigation46aa22_TabletEvent
func callbackNavigation46aa22_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		(*(*func(*std_gui.QTabletEvent))(signal))(std_gui.NewQTabletEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).TabletEventDefault(std_gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *Navigation) TabletEventDefault(event std_gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_TabletEventDefault(ptr.Pointer(), std_gui.PointerFromQTabletEvent(event))
	}
}

//export callbackNavigation46aa22_Update
func callbackNavigation46aa22_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewNavigationFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *Navigation) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_UpdateDefault(ptr.Pointer())
	}
}

//export callbackNavigation46aa22_UpdateMicroFocus
func callbackNavigation46aa22_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewNavigationFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *Navigation) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackNavigation46aa22_WheelEvent
func callbackNavigation46aa22_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		(*(*func(*std_gui.QWheelEvent))(signal))(std_gui.NewQWheelEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).WheelEventDefault(std_gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *Navigation) WheelEventDefault(event std_gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_WheelEventDefault(ptr.Pointer(), std_gui.PointerFromQWheelEvent(event))
	}
}

//export callbackNavigation46aa22_WindowIconChanged
func callbackNavigation46aa22_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		(*(*func(*std_gui.QIcon))(signal))(std_gui.NewQIconFromPointer(icon))
	}

}

//export callbackNavigation46aa22_WindowTitleChanged
func callbackNavigation46aa22_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(title))
	}

}

//export callbackNavigation46aa22_PaintEngine
func callbackNavigation46aa22_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return std_gui.PointerFromQPaintEngine((*(*func() *std_gui.QPaintEngine)(signal))())
	}

	return std_gui.PointerFromQPaintEngine(NewNavigationFromPointer(ptr).PaintEngineDefault())
}

func (ptr *Navigation) PaintEngineDefault() *std_gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return std_gui.NewQPaintEngineFromPointer(C.Navigation46aa22_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackNavigation46aa22_MinimumSizeHint
func callbackNavigation46aa22_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return std_core.PointerFromQSize((*(*func() *std_core.QSize)(signal))())
	}

	return std_core.PointerFromQSize(NewNavigationFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *Navigation) MinimumSizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.Navigation46aa22_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackNavigation46aa22_SizeHint
func callbackNavigation46aa22_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return std_core.PointerFromQSize((*(*func() *std_core.QSize)(signal))())
	}

	return std_core.PointerFromQSize(NewNavigationFromPointer(ptr).SizeHintDefault())
}

func (ptr *Navigation) SizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.Navigation46aa22_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackNavigation46aa22_InputMethodQuery
func callbackNavigation46aa22_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(std_core.Qt__InputMethodQuery) *std_core.QVariant)(signal))(std_core.Qt__InputMethodQuery(query)))
	}

	return std_core.PointerFromQVariant(NewNavigationFromPointer(ptr).InputMethodQueryDefault(std_core.Qt__InputMethodQuery(query)))
}

func (ptr *Navigation) InputMethodQueryDefault(query std_core.Qt__InputMethodQuery) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.Navigation46aa22_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackNavigation46aa22_HasHeightForWidth
func callbackNavigation46aa22_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewNavigationFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *Navigation) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.Navigation46aa22_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackNavigation46aa22_HeightForWidth
func callbackNavigation46aa22_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32((*(*func(int) int)(signal))(int(int32(w)))))
	}

	return C.int(int32(NewNavigationFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *Navigation) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.Navigation46aa22_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackNavigation46aa22_Metric
func callbackNavigation46aa22_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32((*(*func(std_gui.QPaintDevice__PaintDeviceMetric) int)(signal))(std_gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewNavigationFromPointer(ptr).MetricDefault(std_gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *Navigation) MetricDefault(m std_gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.Navigation46aa22_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackNavigation46aa22_InitPainter
func callbackNavigation46aa22_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		(*(*func(*std_gui.QPainter))(signal))(std_gui.NewQPainterFromPointer(painter))
	} else {
		NewNavigationFromPointer(ptr).InitPainterDefault(std_gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *Navigation) InitPainterDefault(painter std_gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_InitPainterDefault(ptr.Pointer(), std_gui.PointerFromQPainter(painter))
	}
}

//export callbackNavigation46aa22_EventFilter
func callbackNavigation46aa22_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewNavigationFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *Navigation) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.Navigation46aa22_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackNavigation46aa22_ChildEvent
func callbackNavigation46aa22_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *Navigation) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackNavigation46aa22_ConnectNotify
func callbackNavigation46aa22_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewNavigationFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Navigation) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackNavigation46aa22_CustomEvent
func callbackNavigation46aa22_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *Navigation) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackNavigation46aa22_DeleteLater
func callbackNavigation46aa22_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewNavigationFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *Navigation) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackNavigation46aa22_Destroyed
func callbackNavigation46aa22_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackNavigation46aa22_DisconnectNotify
func callbackNavigation46aa22_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewNavigationFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Navigation) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackNavigation46aa22_ObjectNameChanged
func callbackNavigation46aa22_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackNavigation46aa22_TimerEvent
func callbackNavigation46aa22_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewNavigationFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *Navigation) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Navigation46aa22_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type SettingsDialog_ITF interface {
	std_widgets.QDialog_ITF
	SettingsDialog_PTR() *SettingsDialog
}

func (ptr *SettingsDialog) SettingsDialog_PTR() *SettingsDialog {
	return ptr
}

func (ptr *SettingsDialog) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDialog_PTR().Pointer()
	}
	return nil
}

func (ptr *SettingsDialog) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDialog_PTR().SetPointer(p)
	}
}

func PointerFromSettingsDialog(ptr SettingsDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SettingsDialog_PTR().Pointer()
	}
	return nil
}

func NewSettingsDialogFromPointer(ptr unsafe.Pointer) (n *SettingsDialog) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(SettingsDialog)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *SettingsDialog:
			n = deduced

		case *std_widgets.QDialog:
			n = &SettingsDialog{QDialog: *deduced}

		default:
			n = new(SettingsDialog)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *SettingsDialog) Init() { this.init() }

//export callbackSettingsDialog46aa22_Constructor
func callbackSettingsDialog46aa22_Constructor(ptr unsafe.Pointer) {
	this := NewSettingsDialogFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

func SettingsDialog_QRegisterMetaType() int {
	return int(int32(C.SettingsDialog46aa22_SettingsDialog46aa22_QRegisterMetaType()))
}

func (ptr *SettingsDialog) QRegisterMetaType() int {
	return int(int32(C.SettingsDialog46aa22_SettingsDialog46aa22_QRegisterMetaType()))
}

func SettingsDialog_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.SettingsDialog46aa22_SettingsDialog46aa22_QRegisterMetaType2(typeNameC)))
}

func (ptr *SettingsDialog) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.SettingsDialog46aa22_SettingsDialog46aa22_QRegisterMetaType2(typeNameC)))
}

func SettingsDialog_QmlRegisterType() int {
	return int(int32(C.SettingsDialog46aa22_SettingsDialog46aa22_QmlRegisterType()))
}

func (ptr *SettingsDialog) QmlRegisterType() int {
	return int(int32(C.SettingsDialog46aa22_SettingsDialog46aa22_QmlRegisterType()))
}

func SettingsDialog_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.SettingsDialog46aa22_SettingsDialog46aa22_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *SettingsDialog) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.SettingsDialog46aa22_SettingsDialog46aa22_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *SettingsDialog) __addActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.SettingsDialog46aa22___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *SettingsDialog) __addActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22___addActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *SettingsDialog) __addActions_actions_newList() unsafe.Pointer {
	return C.SettingsDialog46aa22___addActions_actions_newList(ptr.Pointer())
}

func (ptr *SettingsDialog) __insertActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.SettingsDialog46aa22___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *SettingsDialog) __insertActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22___insertActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *SettingsDialog) __insertActions_actions_newList() unsafe.Pointer {
	return C.SettingsDialog46aa22___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *SettingsDialog) __actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.SettingsDialog46aa22___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *SettingsDialog) __actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22___actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *SettingsDialog) __actions_newList() unsafe.Pointer {
	return C.SettingsDialog46aa22___actions_newList(ptr.Pointer())
}

func (ptr *SettingsDialog) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.SettingsDialog46aa22___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *SettingsDialog) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *SettingsDialog) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.SettingsDialog46aa22___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *SettingsDialog) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.SettingsDialog46aa22___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *SettingsDialog) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *SettingsDialog) __findChildren_newList2() unsafe.Pointer {
	return C.SettingsDialog46aa22___findChildren_newList2(ptr.Pointer())
}

func (ptr *SettingsDialog) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.SettingsDialog46aa22___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *SettingsDialog) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *SettingsDialog) __findChildren_newList3() unsafe.Pointer {
	return C.SettingsDialog46aa22___findChildren_newList3(ptr.Pointer())
}

func (ptr *SettingsDialog) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.SettingsDialog46aa22___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *SettingsDialog) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *SettingsDialog) __findChildren_newList() unsafe.Pointer {
	return C.SettingsDialog46aa22___findChildren_newList(ptr.Pointer())
}

func (ptr *SettingsDialog) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.SettingsDialog46aa22___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *SettingsDialog) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *SettingsDialog) __children_newList() unsafe.Pointer {
	return C.SettingsDialog46aa22___children_newList(ptr.Pointer())
}

func NewSettingsDialog(parent std_widgets.QWidget_ITF, ff std_core.Qt__WindowType) *SettingsDialog {
	tmpValue := NewSettingsDialogFromPointer(C.SettingsDialog46aa22_NewSettingsDialog(std_widgets.PointerFromQWidget(parent), C.longlong(ff)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackSettingsDialog46aa22_DestroySettingsDialog
func callbackSettingsDialog46aa22_DestroySettingsDialog(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~SettingsDialog"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsDialogFromPointer(ptr).DestroySettingsDialogDefault()
	}
}

func (ptr *SettingsDialog) ConnectDestroySettingsDialog(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~SettingsDialog"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~SettingsDialog", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~SettingsDialog", unsafe.Pointer(&f))
		}
	}
}

func (ptr *SettingsDialog) DisconnectDestroySettingsDialog() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~SettingsDialog")
	}
}

func (ptr *SettingsDialog) DestroySettingsDialog() {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_DestroySettingsDialog(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *SettingsDialog) DestroySettingsDialogDefault() {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_DestroySettingsDialogDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackSettingsDialog46aa22_EventFilter
func callbackSettingsDialog46aa22_EventFilter(ptr unsafe.Pointer, o unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(o), std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSettingsDialogFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(o), std_core.NewQEventFromPointer(e)))))
}

func (ptr *SettingsDialog) EventFilterDefault(o std_core.QObject_ITF, e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.SettingsDialog46aa22_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(o), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackSettingsDialog46aa22_Exec
func callbackSettingsDialog46aa22_Exec(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "exec"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewSettingsDialogFromPointer(ptr).ExecDefault()))
}

func (ptr *SettingsDialog) ExecDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.SettingsDialog46aa22_ExecDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackSettingsDialog46aa22_Accept
func callbackSettingsDialog46aa22_Accept(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "accept"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsDialogFromPointer(ptr).AcceptDefault()
	}
}

func (ptr *SettingsDialog) AcceptDefault() {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_AcceptDefault(ptr.Pointer())
	}
}

//export callbackSettingsDialog46aa22_Accepted
func callbackSettingsDialog46aa22_Accepted(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "accepted"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackSettingsDialog46aa22_CloseEvent
func callbackSettingsDialog46aa22_CloseEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		(*(*func(*std_gui.QCloseEvent))(signal))(std_gui.NewQCloseEventFromPointer(e))
	} else {
		NewSettingsDialogFromPointer(ptr).CloseEventDefault(std_gui.NewQCloseEventFromPointer(e))
	}
}

func (ptr *SettingsDialog) CloseEventDefault(e std_gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_CloseEventDefault(ptr.Pointer(), std_gui.PointerFromQCloseEvent(e))
	}
}

//export callbackSettingsDialog46aa22_ContextMenuEvent
func callbackSettingsDialog46aa22_ContextMenuEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		(*(*func(*std_gui.QContextMenuEvent))(signal))(std_gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewSettingsDialogFromPointer(ptr).ContextMenuEventDefault(std_gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *SettingsDialog) ContextMenuEventDefault(e std_gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_ContextMenuEventDefault(ptr.Pointer(), std_gui.PointerFromQContextMenuEvent(e))
	}
}

//export callbackSettingsDialog46aa22_Done
func callbackSettingsDialog46aa22_Done(ptr unsafe.Pointer, r C.int) {
	if signal := qt.GetSignal(ptr, "done"); signal != nil {
		(*(*func(int))(signal))(int(int32(r)))
	} else {
		NewSettingsDialogFromPointer(ptr).DoneDefault(int(int32(r)))
	}
}

func (ptr *SettingsDialog) DoneDefault(r int) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_DoneDefault(ptr.Pointer(), C.int(int32(r)))
	}
}

//export callbackSettingsDialog46aa22_Finished
func callbackSettingsDialog46aa22_Finished(ptr unsafe.Pointer, result C.int) {
	if signal := qt.GetSignal(ptr, "finished"); signal != nil {
		(*(*func(int))(signal))(int(int32(result)))
	}

}

//export callbackSettingsDialog46aa22_KeyPressEvent
func callbackSettingsDialog46aa22_KeyPressEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		(*(*func(*std_gui.QKeyEvent))(signal))(std_gui.NewQKeyEventFromPointer(e))
	} else {
		NewSettingsDialogFromPointer(ptr).KeyPressEventDefault(std_gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *SettingsDialog) KeyPressEventDefault(e std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_KeyPressEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(e))
	}
}

//export callbackSettingsDialog46aa22_Open
func callbackSettingsDialog46aa22_Open(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "open"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsDialogFromPointer(ptr).OpenDefault()
	}
}

func (ptr *SettingsDialog) OpenDefault() {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_OpenDefault(ptr.Pointer())
	}
}

//export callbackSettingsDialog46aa22_Reject
func callbackSettingsDialog46aa22_Reject(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "reject"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsDialogFromPointer(ptr).RejectDefault()
	}
}

func (ptr *SettingsDialog) RejectDefault() {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_RejectDefault(ptr.Pointer())
	}
}

//export callbackSettingsDialog46aa22_Rejected
func callbackSettingsDialog46aa22_Rejected(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "rejected"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackSettingsDialog46aa22_ResizeEvent
func callbackSettingsDialog46aa22_ResizeEvent(ptr unsafe.Pointer, vqr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		(*(*func(*std_gui.QResizeEvent))(signal))(std_gui.NewQResizeEventFromPointer(vqr))
	} else {
		NewSettingsDialogFromPointer(ptr).ResizeEventDefault(std_gui.NewQResizeEventFromPointer(vqr))
	}
}

func (ptr *SettingsDialog) ResizeEventDefault(vqr std_gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_ResizeEventDefault(ptr.Pointer(), std_gui.PointerFromQResizeEvent(vqr))
	}
}

//export callbackSettingsDialog46aa22_SetVisible
func callbackSettingsDialog46aa22_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	} else {
		NewSettingsDialogFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *SettingsDialog) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackSettingsDialog46aa22_ShowEvent
func callbackSettingsDialog46aa22_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		(*(*func(*std_gui.QShowEvent))(signal))(std_gui.NewQShowEventFromPointer(event))
	} else {
		NewSettingsDialogFromPointer(ptr).ShowEventDefault(std_gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *SettingsDialog) ShowEventDefault(event std_gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_ShowEventDefault(ptr.Pointer(), std_gui.PointerFromQShowEvent(event))
	}
}

//export callbackSettingsDialog46aa22_MinimumSizeHint
func callbackSettingsDialog46aa22_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return std_core.PointerFromQSize((*(*func() *std_core.QSize)(signal))())
	}

	return std_core.PointerFromQSize(NewSettingsDialogFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *SettingsDialog) MinimumSizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.SettingsDialog46aa22_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackSettingsDialog46aa22_SizeHint
func callbackSettingsDialog46aa22_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return std_core.PointerFromQSize((*(*func() *std_core.QSize)(signal))())
	}

	return std_core.PointerFromQSize(NewSettingsDialogFromPointer(ptr).SizeHintDefault())
}

func (ptr *SettingsDialog) SizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.SettingsDialog46aa22_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackSettingsDialog46aa22_Close
func callbackSettingsDialog46aa22_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewSettingsDialogFromPointer(ptr).CloseDefault())))
}

func (ptr *SettingsDialog) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.SettingsDialog46aa22_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackSettingsDialog46aa22_Event
func callbackSettingsDialog46aa22_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSettingsDialogFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(event)))))
}

func (ptr *SettingsDialog) EventDefault(event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.SettingsDialog46aa22_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackSettingsDialog46aa22_FocusNextPrevChild
func callbackSettingsDialog46aa22_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(bool) bool)(signal))(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSettingsDialogFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *SettingsDialog) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.SettingsDialog46aa22_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackSettingsDialog46aa22_NativeEvent
func callbackSettingsDialog46aa22_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result *C.long) C.char {
	var resultR int
	if result != nil {
		resultR = int(int32(*result))
		defer func() { *result = C.long(int32(resultR)) }()
	}
	if signal := qt.GetSignal(ptr, "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QByteArray, unsafe.Pointer, *int) bool)(signal))(std_core.NewQByteArrayFromPointer(eventType), message, &resultR))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSettingsDialogFromPointer(ptr).NativeEventDefault(std_core.NewQByteArrayFromPointer(eventType), message, &resultR))))
}

func (ptr *SettingsDialog) NativeEventDefault(eventType std_core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {
	if ptr.Pointer() != nil {
		var resultC C.long
		if result != nil {
			resultC = C.long(int32(*result))
			defer func() { *result = int(int32(resultC)) }()
		}
		return int8(C.SettingsDialog46aa22_NativeEventDefault(ptr.Pointer(), std_core.PointerFromQByteArray(eventType), message, &resultC)) != 0
	}
	return false
}

//export callbackSettingsDialog46aa22_ActionEvent
func callbackSettingsDialog46aa22_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		(*(*func(*std_gui.QActionEvent))(signal))(std_gui.NewQActionEventFromPointer(event))
	} else {
		NewSettingsDialogFromPointer(ptr).ActionEventDefault(std_gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *SettingsDialog) ActionEventDefault(event std_gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_ActionEventDefault(ptr.Pointer(), std_gui.PointerFromQActionEvent(event))
	}
}

//export callbackSettingsDialog46aa22_ChangeEvent
func callbackSettingsDialog46aa22_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewSettingsDialogFromPointer(ptr).ChangeEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *SettingsDialog) ChangeEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_ChangeEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackSettingsDialog46aa22_CustomContextMenuRequested
func callbackSettingsDialog46aa22_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		(*(*func(*std_core.QPoint))(signal))(std_core.NewQPointFromPointer(pos))
	}

}

//export callbackSettingsDialog46aa22_DragEnterEvent
func callbackSettingsDialog46aa22_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		(*(*func(*std_gui.QDragEnterEvent))(signal))(std_gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewSettingsDialogFromPointer(ptr).DragEnterEventDefault(std_gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *SettingsDialog) DragEnterEventDefault(event std_gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_DragEnterEventDefault(ptr.Pointer(), std_gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackSettingsDialog46aa22_DragLeaveEvent
func callbackSettingsDialog46aa22_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		(*(*func(*std_gui.QDragLeaveEvent))(signal))(std_gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewSettingsDialogFromPointer(ptr).DragLeaveEventDefault(std_gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *SettingsDialog) DragLeaveEventDefault(event std_gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_DragLeaveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackSettingsDialog46aa22_DragMoveEvent
func callbackSettingsDialog46aa22_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		(*(*func(*std_gui.QDragMoveEvent))(signal))(std_gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewSettingsDialogFromPointer(ptr).DragMoveEventDefault(std_gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *SettingsDialog) DragMoveEventDefault(event std_gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_DragMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackSettingsDialog46aa22_DropEvent
func callbackSettingsDialog46aa22_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		(*(*func(*std_gui.QDropEvent))(signal))(std_gui.NewQDropEventFromPointer(event))
	} else {
		NewSettingsDialogFromPointer(ptr).DropEventDefault(std_gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *SettingsDialog) DropEventDefault(event std_gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_DropEventDefault(ptr.Pointer(), std_gui.PointerFromQDropEvent(event))
	}
}

//export callbackSettingsDialog46aa22_EnterEvent
func callbackSettingsDialog46aa22_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewSettingsDialogFromPointer(ptr).EnterEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *SettingsDialog) EnterEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_EnterEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackSettingsDialog46aa22_FocusInEvent
func callbackSettingsDialog46aa22_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		(*(*func(*std_gui.QFocusEvent))(signal))(std_gui.NewQFocusEventFromPointer(event))
	} else {
		NewSettingsDialogFromPointer(ptr).FocusInEventDefault(std_gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *SettingsDialog) FocusInEventDefault(event std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_FocusInEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(event))
	}
}

//export callbackSettingsDialog46aa22_FocusOutEvent
func callbackSettingsDialog46aa22_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		(*(*func(*std_gui.QFocusEvent))(signal))(std_gui.NewQFocusEventFromPointer(event))
	} else {
		NewSettingsDialogFromPointer(ptr).FocusOutEventDefault(std_gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *SettingsDialog) FocusOutEventDefault(event std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_FocusOutEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(event))
	}
}

//export callbackSettingsDialog46aa22_Hide
func callbackSettingsDialog46aa22_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsDialogFromPointer(ptr).HideDefault()
	}
}

func (ptr *SettingsDialog) HideDefault() {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_HideDefault(ptr.Pointer())
	}
}

//export callbackSettingsDialog46aa22_HideEvent
func callbackSettingsDialog46aa22_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		(*(*func(*std_gui.QHideEvent))(signal))(std_gui.NewQHideEventFromPointer(event))
	} else {
		NewSettingsDialogFromPointer(ptr).HideEventDefault(std_gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *SettingsDialog) HideEventDefault(event std_gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_HideEventDefault(ptr.Pointer(), std_gui.PointerFromQHideEvent(event))
	}
}

//export callbackSettingsDialog46aa22_InputMethodEvent
func callbackSettingsDialog46aa22_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		(*(*func(*std_gui.QInputMethodEvent))(signal))(std_gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewSettingsDialogFromPointer(ptr).InputMethodEventDefault(std_gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *SettingsDialog) InputMethodEventDefault(event std_gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_InputMethodEventDefault(ptr.Pointer(), std_gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackSettingsDialog46aa22_KeyReleaseEvent
func callbackSettingsDialog46aa22_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		(*(*func(*std_gui.QKeyEvent))(signal))(std_gui.NewQKeyEventFromPointer(event))
	} else {
		NewSettingsDialogFromPointer(ptr).KeyReleaseEventDefault(std_gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *SettingsDialog) KeyReleaseEventDefault(event std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_KeyReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(event))
	}
}

//export callbackSettingsDialog46aa22_LeaveEvent
func callbackSettingsDialog46aa22_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewSettingsDialogFromPointer(ptr).LeaveEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *SettingsDialog) LeaveEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_LeaveEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackSettingsDialog46aa22_Lower
func callbackSettingsDialog46aa22_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsDialogFromPointer(ptr).LowerDefault()
	}
}

func (ptr *SettingsDialog) LowerDefault() {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_LowerDefault(ptr.Pointer())
	}
}

//export callbackSettingsDialog46aa22_MouseDoubleClickEvent
func callbackSettingsDialog46aa22_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewSettingsDialogFromPointer(ptr).MouseDoubleClickEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *SettingsDialog) MouseDoubleClickEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_MouseDoubleClickEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackSettingsDialog46aa22_MouseMoveEvent
func callbackSettingsDialog46aa22_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewSettingsDialogFromPointer(ptr).MouseMoveEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *SettingsDialog) MouseMoveEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_MouseMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackSettingsDialog46aa22_MousePressEvent
func callbackSettingsDialog46aa22_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewSettingsDialogFromPointer(ptr).MousePressEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *SettingsDialog) MousePressEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_MousePressEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackSettingsDialog46aa22_MouseReleaseEvent
func callbackSettingsDialog46aa22_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewSettingsDialogFromPointer(ptr).MouseReleaseEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *SettingsDialog) MouseReleaseEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_MouseReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackSettingsDialog46aa22_MoveEvent
func callbackSettingsDialog46aa22_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		(*(*func(*std_gui.QMoveEvent))(signal))(std_gui.NewQMoveEventFromPointer(event))
	} else {
		NewSettingsDialogFromPointer(ptr).MoveEventDefault(std_gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *SettingsDialog) MoveEventDefault(event std_gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_MoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMoveEvent(event))
	}
}

//export callbackSettingsDialog46aa22_PaintEvent
func callbackSettingsDialog46aa22_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		(*(*func(*std_gui.QPaintEvent))(signal))(std_gui.NewQPaintEventFromPointer(event))
	} else {
		NewSettingsDialogFromPointer(ptr).PaintEventDefault(std_gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *SettingsDialog) PaintEventDefault(event std_gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_PaintEventDefault(ptr.Pointer(), std_gui.PointerFromQPaintEvent(event))
	}
}

//export callbackSettingsDialog46aa22_Raise
func callbackSettingsDialog46aa22_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsDialogFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *SettingsDialog) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_RaiseDefault(ptr.Pointer())
	}
}

//export callbackSettingsDialog46aa22_Repaint
func callbackSettingsDialog46aa22_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsDialogFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *SettingsDialog) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_RepaintDefault(ptr.Pointer())
	}
}

//export callbackSettingsDialog46aa22_SetDisabled
func callbackSettingsDialog46aa22_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		(*(*func(bool))(signal))(int8(disable) != 0)
	} else {
		NewSettingsDialogFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *SettingsDialog) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackSettingsDialog46aa22_SetEnabled
func callbackSettingsDialog46aa22_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewSettingsDialogFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *SettingsDialog) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackSettingsDialog46aa22_SetFocus2
func callbackSettingsDialog46aa22_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsDialogFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *SettingsDialog) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackSettingsDialog46aa22_SetHidden
func callbackSettingsDialog46aa22_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		(*(*func(bool))(signal))(int8(hidden) != 0)
	} else {
		NewSettingsDialogFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *SettingsDialog) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackSettingsDialog46aa22_SetStyleSheet
func callbackSettingsDialog46aa22_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(styleSheet))
	} else {
		NewSettingsDialogFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *SettingsDialog) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.SettingsDialog46aa22_SetStyleSheetDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackSettingsDialog46aa22_SetWindowModified
func callbackSettingsDialog46aa22_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewSettingsDialogFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *SettingsDialog) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackSettingsDialog46aa22_SetWindowTitle
func callbackSettingsDialog46aa22_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(vqs))
	} else {
		NewSettingsDialogFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *SettingsDialog) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.SettingsDialog46aa22_SetWindowTitleDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackSettingsDialog46aa22_Show
func callbackSettingsDialog46aa22_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsDialogFromPointer(ptr).ShowDefault()
	}
}

func (ptr *SettingsDialog) ShowDefault() {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_ShowDefault(ptr.Pointer())
	}
}

//export callbackSettingsDialog46aa22_ShowFullScreen
func callbackSettingsDialog46aa22_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsDialogFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *SettingsDialog) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackSettingsDialog46aa22_ShowMaximized
func callbackSettingsDialog46aa22_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsDialogFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *SettingsDialog) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackSettingsDialog46aa22_ShowMinimized
func callbackSettingsDialog46aa22_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsDialogFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *SettingsDialog) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackSettingsDialog46aa22_ShowNormal
func callbackSettingsDialog46aa22_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsDialogFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *SettingsDialog) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackSettingsDialog46aa22_TabletEvent
func callbackSettingsDialog46aa22_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		(*(*func(*std_gui.QTabletEvent))(signal))(std_gui.NewQTabletEventFromPointer(event))
	} else {
		NewSettingsDialogFromPointer(ptr).TabletEventDefault(std_gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *SettingsDialog) TabletEventDefault(event std_gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_TabletEventDefault(ptr.Pointer(), std_gui.PointerFromQTabletEvent(event))
	}
}

//export callbackSettingsDialog46aa22_Update
func callbackSettingsDialog46aa22_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsDialogFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *SettingsDialog) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_UpdateDefault(ptr.Pointer())
	}
}

//export callbackSettingsDialog46aa22_UpdateMicroFocus
func callbackSettingsDialog46aa22_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsDialogFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *SettingsDialog) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackSettingsDialog46aa22_WheelEvent
func callbackSettingsDialog46aa22_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		(*(*func(*std_gui.QWheelEvent))(signal))(std_gui.NewQWheelEventFromPointer(event))
	} else {
		NewSettingsDialogFromPointer(ptr).WheelEventDefault(std_gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *SettingsDialog) WheelEventDefault(event std_gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_WheelEventDefault(ptr.Pointer(), std_gui.PointerFromQWheelEvent(event))
	}
}

//export callbackSettingsDialog46aa22_WindowIconChanged
func callbackSettingsDialog46aa22_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		(*(*func(*std_gui.QIcon))(signal))(std_gui.NewQIconFromPointer(icon))
	}

}

//export callbackSettingsDialog46aa22_WindowTitleChanged
func callbackSettingsDialog46aa22_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(title))
	}

}

//export callbackSettingsDialog46aa22_PaintEngine
func callbackSettingsDialog46aa22_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return std_gui.PointerFromQPaintEngine((*(*func() *std_gui.QPaintEngine)(signal))())
	}

	return std_gui.PointerFromQPaintEngine(NewSettingsDialogFromPointer(ptr).PaintEngineDefault())
}

func (ptr *SettingsDialog) PaintEngineDefault() *std_gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return std_gui.NewQPaintEngineFromPointer(C.SettingsDialog46aa22_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackSettingsDialog46aa22_InputMethodQuery
func callbackSettingsDialog46aa22_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(std_core.Qt__InputMethodQuery) *std_core.QVariant)(signal))(std_core.Qt__InputMethodQuery(query)))
	}

	return std_core.PointerFromQVariant(NewSettingsDialogFromPointer(ptr).InputMethodQueryDefault(std_core.Qt__InputMethodQuery(query)))
}

func (ptr *SettingsDialog) InputMethodQueryDefault(query std_core.Qt__InputMethodQuery) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.SettingsDialog46aa22_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackSettingsDialog46aa22_HasHeightForWidth
func callbackSettingsDialog46aa22_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewSettingsDialogFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *SettingsDialog) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.SettingsDialog46aa22_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackSettingsDialog46aa22_HeightForWidth
func callbackSettingsDialog46aa22_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32((*(*func(int) int)(signal))(int(int32(w)))))
	}

	return C.int(int32(NewSettingsDialogFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *SettingsDialog) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.SettingsDialog46aa22_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackSettingsDialog46aa22_Metric
func callbackSettingsDialog46aa22_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32((*(*func(std_gui.QPaintDevice__PaintDeviceMetric) int)(signal))(std_gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewSettingsDialogFromPointer(ptr).MetricDefault(std_gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *SettingsDialog) MetricDefault(m std_gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.SettingsDialog46aa22_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackSettingsDialog46aa22_InitPainter
func callbackSettingsDialog46aa22_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		(*(*func(*std_gui.QPainter))(signal))(std_gui.NewQPainterFromPointer(painter))
	} else {
		NewSettingsDialogFromPointer(ptr).InitPainterDefault(std_gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *SettingsDialog) InitPainterDefault(painter std_gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_InitPainterDefault(ptr.Pointer(), std_gui.PointerFromQPainter(painter))
	}
}

//export callbackSettingsDialog46aa22_ChildEvent
func callbackSettingsDialog46aa22_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewSettingsDialogFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *SettingsDialog) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackSettingsDialog46aa22_ConnectNotify
func callbackSettingsDialog46aa22_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewSettingsDialogFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *SettingsDialog) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackSettingsDialog46aa22_CustomEvent
func callbackSettingsDialog46aa22_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewSettingsDialogFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *SettingsDialog) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackSettingsDialog46aa22_DeleteLater
func callbackSettingsDialog46aa22_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsDialogFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *SettingsDialog) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackSettingsDialog46aa22_Destroyed
func callbackSettingsDialog46aa22_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackSettingsDialog46aa22_DisconnectNotify
func callbackSettingsDialog46aa22_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewSettingsDialogFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *SettingsDialog) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackSettingsDialog46aa22_ObjectNameChanged
func callbackSettingsDialog46aa22_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackSettingsDialog46aa22_TimerEvent
func callbackSettingsDialog46aa22_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewSettingsDialogFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *SettingsDialog) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SettingsDialog46aa22_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type CurrentMediaPage_ITF interface {
	std_widgets.QTabWidget_ITF
	CurrentMediaPage_PTR() *CurrentMediaPage
}

func (ptr *CurrentMediaPage) CurrentMediaPage_PTR() *CurrentMediaPage {
	return ptr
}

func (ptr *CurrentMediaPage) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTabWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *CurrentMediaPage) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTabWidget_PTR().SetPointer(p)
	}
}

func PointerFromCurrentMediaPage(ptr CurrentMediaPage_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.CurrentMediaPage_PTR().Pointer()
	}
	return nil
}

func NewCurrentMediaPageFromPointer(ptr unsafe.Pointer) (n *CurrentMediaPage) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(CurrentMediaPage)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *CurrentMediaPage:
			n = deduced

		case *std_widgets.QTabWidget:
			n = &CurrentMediaPage{QTabWidget: *deduced}

		default:
			n = new(CurrentMediaPage)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *CurrentMediaPage) Init() { this.init() }

//export callbackCurrentMediaPage46aa22_Constructor
func callbackCurrentMediaPage46aa22_Constructor(ptr unsafe.Pointer) {
	this := NewCurrentMediaPageFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

func CurrentMediaPage_QRegisterMetaType() int {
	return int(int32(C.CurrentMediaPage46aa22_CurrentMediaPage46aa22_QRegisterMetaType()))
}

func (ptr *CurrentMediaPage) QRegisterMetaType() int {
	return int(int32(C.CurrentMediaPage46aa22_CurrentMediaPage46aa22_QRegisterMetaType()))
}

func CurrentMediaPage_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.CurrentMediaPage46aa22_CurrentMediaPage46aa22_QRegisterMetaType2(typeNameC)))
}

func (ptr *CurrentMediaPage) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.CurrentMediaPage46aa22_CurrentMediaPage46aa22_QRegisterMetaType2(typeNameC)))
}

func CurrentMediaPage_QmlRegisterType() int {
	return int(int32(C.CurrentMediaPage46aa22_CurrentMediaPage46aa22_QmlRegisterType()))
}

func (ptr *CurrentMediaPage) QmlRegisterType() int {
	return int(int32(C.CurrentMediaPage46aa22_CurrentMediaPage46aa22_QmlRegisterType()))
}

func CurrentMediaPage_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.CurrentMediaPage46aa22_CurrentMediaPage46aa22_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *CurrentMediaPage) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.CurrentMediaPage46aa22_CurrentMediaPage46aa22_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *CurrentMediaPage) __addActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.CurrentMediaPage46aa22___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CurrentMediaPage) __addActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22___addActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *CurrentMediaPage) __addActions_actions_newList() unsafe.Pointer {
	return C.CurrentMediaPage46aa22___addActions_actions_newList(ptr.Pointer())
}

func (ptr *CurrentMediaPage) __insertActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.CurrentMediaPage46aa22___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CurrentMediaPage) __insertActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22___insertActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *CurrentMediaPage) __insertActions_actions_newList() unsafe.Pointer {
	return C.CurrentMediaPage46aa22___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *CurrentMediaPage) __actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.CurrentMediaPage46aa22___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CurrentMediaPage) __actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22___actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *CurrentMediaPage) __actions_newList() unsafe.Pointer {
	return C.CurrentMediaPage46aa22___actions_newList(ptr.Pointer())
}

func (ptr *CurrentMediaPage) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.CurrentMediaPage46aa22___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *CurrentMediaPage) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *CurrentMediaPage) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.CurrentMediaPage46aa22___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *CurrentMediaPage) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CurrentMediaPage46aa22___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CurrentMediaPage) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CurrentMediaPage) __findChildren_newList2() unsafe.Pointer {
	return C.CurrentMediaPage46aa22___findChildren_newList2(ptr.Pointer())
}

func (ptr *CurrentMediaPage) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CurrentMediaPage46aa22___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CurrentMediaPage) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CurrentMediaPage) __findChildren_newList3() unsafe.Pointer {
	return C.CurrentMediaPage46aa22___findChildren_newList3(ptr.Pointer())
}

func (ptr *CurrentMediaPage) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CurrentMediaPage46aa22___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CurrentMediaPage) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CurrentMediaPage) __findChildren_newList() unsafe.Pointer {
	return C.CurrentMediaPage46aa22___findChildren_newList(ptr.Pointer())
}

func (ptr *CurrentMediaPage) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CurrentMediaPage46aa22___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CurrentMediaPage) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CurrentMediaPage) __children_newList() unsafe.Pointer {
	return C.CurrentMediaPage46aa22___children_newList(ptr.Pointer())
}

func NewCurrentMediaPage(parent std_widgets.QWidget_ITF) *CurrentMediaPage {
	tmpValue := NewCurrentMediaPageFromPointer(C.CurrentMediaPage46aa22_NewCurrentMediaPage(std_widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackCurrentMediaPage46aa22_DestroyCurrentMediaPage
func callbackCurrentMediaPage46aa22_DestroyCurrentMediaPage(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~CurrentMediaPage"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewCurrentMediaPageFromPointer(ptr).DestroyCurrentMediaPageDefault()
	}
}

func (ptr *CurrentMediaPage) ConnectDestroyCurrentMediaPage(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~CurrentMediaPage"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~CurrentMediaPage", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~CurrentMediaPage", unsafe.Pointer(&f))
		}
	}
}

func (ptr *CurrentMediaPage) DisconnectDestroyCurrentMediaPage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~CurrentMediaPage")
	}
}

func (ptr *CurrentMediaPage) DestroyCurrentMediaPage() {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_DestroyCurrentMediaPage(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *CurrentMediaPage) DestroyCurrentMediaPageDefault() {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_DestroyCurrentMediaPageDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackCurrentMediaPage46aa22_Event
func callbackCurrentMediaPage46aa22_Event(ptr unsafe.Pointer, ev unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(ev)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCurrentMediaPageFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(ev)))))
}

func (ptr *CurrentMediaPage) EventDefault(ev std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CurrentMediaPage46aa22_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(ev))) != 0
	}
	return false
}

//export callbackCurrentMediaPage46aa22_ChangeEvent
func callbackCurrentMediaPage46aa22_ChangeEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(ev))
	} else {
		NewCurrentMediaPageFromPointer(ptr).ChangeEventDefault(std_core.NewQEventFromPointer(ev))
	}
}

func (ptr *CurrentMediaPage) ChangeEventDefault(ev std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_ChangeEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(ev))
	}
}

//export callbackCurrentMediaPage46aa22_CurrentChanged
func callbackCurrentMediaPage46aa22_CurrentChanged(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "currentChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

//export callbackCurrentMediaPage46aa22_KeyPressEvent
func callbackCurrentMediaPage46aa22_KeyPressEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		(*(*func(*std_gui.QKeyEvent))(signal))(std_gui.NewQKeyEventFromPointer(e))
	} else {
		NewCurrentMediaPageFromPointer(ptr).KeyPressEventDefault(std_gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *CurrentMediaPage) KeyPressEventDefault(e std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_KeyPressEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(e))
	}
}

//export callbackCurrentMediaPage46aa22_PaintEvent
func callbackCurrentMediaPage46aa22_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		(*(*func(*std_gui.QPaintEvent))(signal))(std_gui.NewQPaintEventFromPointer(event))
	} else {
		NewCurrentMediaPageFromPointer(ptr).PaintEventDefault(std_gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *CurrentMediaPage) PaintEventDefault(event std_gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_PaintEventDefault(ptr.Pointer(), std_gui.PointerFromQPaintEvent(event))
	}
}

//export callbackCurrentMediaPage46aa22_ResizeEvent
func callbackCurrentMediaPage46aa22_ResizeEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		(*(*func(*std_gui.QResizeEvent))(signal))(std_gui.NewQResizeEventFromPointer(e))
	} else {
		NewCurrentMediaPageFromPointer(ptr).ResizeEventDefault(std_gui.NewQResizeEventFromPointer(e))
	}
}

func (ptr *CurrentMediaPage) ResizeEventDefault(e std_gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_ResizeEventDefault(ptr.Pointer(), std_gui.PointerFromQResizeEvent(e))
	}
}

//export callbackCurrentMediaPage46aa22_SetCurrentIndex
func callbackCurrentMediaPage46aa22_SetCurrentIndex(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "setCurrentIndex"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	} else {
		NewCurrentMediaPageFromPointer(ptr).SetCurrentIndexDefault(int(int32(index)))
	}
}

func (ptr *CurrentMediaPage) SetCurrentIndexDefault(index int) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_SetCurrentIndexDefault(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackCurrentMediaPage46aa22_SetCurrentWidget
func callbackCurrentMediaPage46aa22_SetCurrentWidget(ptr unsafe.Pointer, widget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setCurrentWidget"); signal != nil {
		(*(*func(*std_widgets.QWidget))(signal))(std_widgets.NewQWidgetFromPointer(widget))
	} else {
		NewCurrentMediaPageFromPointer(ptr).SetCurrentWidgetDefault(std_widgets.NewQWidgetFromPointer(widget))
	}
}

func (ptr *CurrentMediaPage) SetCurrentWidgetDefault(widget std_widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_SetCurrentWidgetDefault(ptr.Pointer(), std_widgets.PointerFromQWidget(widget))
	}
}

//export callbackCurrentMediaPage46aa22_ShowEvent
func callbackCurrentMediaPage46aa22_ShowEvent(ptr unsafe.Pointer, vqs unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		(*(*func(*std_gui.QShowEvent))(signal))(std_gui.NewQShowEventFromPointer(vqs))
	} else {
		NewCurrentMediaPageFromPointer(ptr).ShowEventDefault(std_gui.NewQShowEventFromPointer(vqs))
	}
}

func (ptr *CurrentMediaPage) ShowEventDefault(vqs std_gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_ShowEventDefault(ptr.Pointer(), std_gui.PointerFromQShowEvent(vqs))
	}
}

//export callbackCurrentMediaPage46aa22_TabBarClicked
func callbackCurrentMediaPage46aa22_TabBarClicked(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "tabBarClicked"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

//export callbackCurrentMediaPage46aa22_TabBarDoubleClicked
func callbackCurrentMediaPage46aa22_TabBarDoubleClicked(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "tabBarDoubleClicked"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

//export callbackCurrentMediaPage46aa22_TabCloseRequested
func callbackCurrentMediaPage46aa22_TabCloseRequested(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "tabCloseRequested"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

//export callbackCurrentMediaPage46aa22_TabInserted
func callbackCurrentMediaPage46aa22_TabInserted(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "tabInserted"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	} else {
		NewCurrentMediaPageFromPointer(ptr).TabInsertedDefault(int(int32(index)))
	}
}

func (ptr *CurrentMediaPage) TabInsertedDefault(index int) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_TabInsertedDefault(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackCurrentMediaPage46aa22_TabRemoved
func callbackCurrentMediaPage46aa22_TabRemoved(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "tabRemoved"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	} else {
		NewCurrentMediaPageFromPointer(ptr).TabRemovedDefault(int(int32(index)))
	}
}

func (ptr *CurrentMediaPage) TabRemovedDefault(index int) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_TabRemovedDefault(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackCurrentMediaPage46aa22_MinimumSizeHint
func callbackCurrentMediaPage46aa22_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return std_core.PointerFromQSize((*(*func() *std_core.QSize)(signal))())
	}

	return std_core.PointerFromQSize(NewCurrentMediaPageFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *CurrentMediaPage) MinimumSizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.CurrentMediaPage46aa22_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackCurrentMediaPage46aa22_SizeHint
func callbackCurrentMediaPage46aa22_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return std_core.PointerFromQSize((*(*func() *std_core.QSize)(signal))())
	}

	return std_core.PointerFromQSize(NewCurrentMediaPageFromPointer(ptr).SizeHintDefault())
}

func (ptr *CurrentMediaPage) SizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.CurrentMediaPage46aa22_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackCurrentMediaPage46aa22_HasHeightForWidth
func callbackCurrentMediaPage46aa22_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewCurrentMediaPageFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *CurrentMediaPage) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.CurrentMediaPage46aa22_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackCurrentMediaPage46aa22_HeightForWidth
func callbackCurrentMediaPage46aa22_HeightForWidth(ptr unsafe.Pointer, width C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32((*(*func(int) int)(signal))(int(int32(width)))))
	}

	return C.int(int32(NewCurrentMediaPageFromPointer(ptr).HeightForWidthDefault(int(int32(width)))))
}

func (ptr *CurrentMediaPage) HeightForWidthDefault(width int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CurrentMediaPage46aa22_HeightForWidthDefault(ptr.Pointer(), C.int(int32(width)))))
	}
	return 0
}

//export callbackCurrentMediaPage46aa22_Close
func callbackCurrentMediaPage46aa22_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewCurrentMediaPageFromPointer(ptr).CloseDefault())))
}

func (ptr *CurrentMediaPage) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.CurrentMediaPage46aa22_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackCurrentMediaPage46aa22_FocusNextPrevChild
func callbackCurrentMediaPage46aa22_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(bool) bool)(signal))(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCurrentMediaPageFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *CurrentMediaPage) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.CurrentMediaPage46aa22_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackCurrentMediaPage46aa22_NativeEvent
func callbackCurrentMediaPage46aa22_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result *C.long) C.char {
	var resultR int
	if result != nil {
		resultR = int(int32(*result))
		defer func() { *result = C.long(int32(resultR)) }()
	}
	if signal := qt.GetSignal(ptr, "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QByteArray, unsafe.Pointer, *int) bool)(signal))(std_core.NewQByteArrayFromPointer(eventType), message, &resultR))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCurrentMediaPageFromPointer(ptr).NativeEventDefault(std_core.NewQByteArrayFromPointer(eventType), message, &resultR))))
}

func (ptr *CurrentMediaPage) NativeEventDefault(eventType std_core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {
	if ptr.Pointer() != nil {
		var resultC C.long
		if result != nil {
			resultC = C.long(int32(*result))
			defer func() { *result = int(int32(resultC)) }()
		}
		return int8(C.CurrentMediaPage46aa22_NativeEventDefault(ptr.Pointer(), std_core.PointerFromQByteArray(eventType), message, &resultC)) != 0
	}
	return false
}

//export callbackCurrentMediaPage46aa22_ActionEvent
func callbackCurrentMediaPage46aa22_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		(*(*func(*std_gui.QActionEvent))(signal))(std_gui.NewQActionEventFromPointer(event))
	} else {
		NewCurrentMediaPageFromPointer(ptr).ActionEventDefault(std_gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *CurrentMediaPage) ActionEventDefault(event std_gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_ActionEventDefault(ptr.Pointer(), std_gui.PointerFromQActionEvent(event))
	}
}

//export callbackCurrentMediaPage46aa22_CloseEvent
func callbackCurrentMediaPage46aa22_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		(*(*func(*std_gui.QCloseEvent))(signal))(std_gui.NewQCloseEventFromPointer(event))
	} else {
		NewCurrentMediaPageFromPointer(ptr).CloseEventDefault(std_gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *CurrentMediaPage) CloseEventDefault(event std_gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_CloseEventDefault(ptr.Pointer(), std_gui.PointerFromQCloseEvent(event))
	}
}

//export callbackCurrentMediaPage46aa22_ContextMenuEvent
func callbackCurrentMediaPage46aa22_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		(*(*func(*std_gui.QContextMenuEvent))(signal))(std_gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewCurrentMediaPageFromPointer(ptr).ContextMenuEventDefault(std_gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *CurrentMediaPage) ContextMenuEventDefault(event std_gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_ContextMenuEventDefault(ptr.Pointer(), std_gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackCurrentMediaPage46aa22_CustomContextMenuRequested
func callbackCurrentMediaPage46aa22_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		(*(*func(*std_core.QPoint))(signal))(std_core.NewQPointFromPointer(pos))
	}

}

//export callbackCurrentMediaPage46aa22_DragEnterEvent
func callbackCurrentMediaPage46aa22_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		(*(*func(*std_gui.QDragEnterEvent))(signal))(std_gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewCurrentMediaPageFromPointer(ptr).DragEnterEventDefault(std_gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *CurrentMediaPage) DragEnterEventDefault(event std_gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_DragEnterEventDefault(ptr.Pointer(), std_gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackCurrentMediaPage46aa22_DragLeaveEvent
func callbackCurrentMediaPage46aa22_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		(*(*func(*std_gui.QDragLeaveEvent))(signal))(std_gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewCurrentMediaPageFromPointer(ptr).DragLeaveEventDefault(std_gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *CurrentMediaPage) DragLeaveEventDefault(event std_gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_DragLeaveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackCurrentMediaPage46aa22_DragMoveEvent
func callbackCurrentMediaPage46aa22_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		(*(*func(*std_gui.QDragMoveEvent))(signal))(std_gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewCurrentMediaPageFromPointer(ptr).DragMoveEventDefault(std_gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *CurrentMediaPage) DragMoveEventDefault(event std_gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_DragMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackCurrentMediaPage46aa22_DropEvent
func callbackCurrentMediaPage46aa22_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		(*(*func(*std_gui.QDropEvent))(signal))(std_gui.NewQDropEventFromPointer(event))
	} else {
		NewCurrentMediaPageFromPointer(ptr).DropEventDefault(std_gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *CurrentMediaPage) DropEventDefault(event std_gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_DropEventDefault(ptr.Pointer(), std_gui.PointerFromQDropEvent(event))
	}
}

//export callbackCurrentMediaPage46aa22_EnterEvent
func callbackCurrentMediaPage46aa22_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewCurrentMediaPageFromPointer(ptr).EnterEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *CurrentMediaPage) EnterEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_EnterEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackCurrentMediaPage46aa22_FocusInEvent
func callbackCurrentMediaPage46aa22_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		(*(*func(*std_gui.QFocusEvent))(signal))(std_gui.NewQFocusEventFromPointer(event))
	} else {
		NewCurrentMediaPageFromPointer(ptr).FocusInEventDefault(std_gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *CurrentMediaPage) FocusInEventDefault(event std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_FocusInEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(event))
	}
}

//export callbackCurrentMediaPage46aa22_FocusOutEvent
func callbackCurrentMediaPage46aa22_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		(*(*func(*std_gui.QFocusEvent))(signal))(std_gui.NewQFocusEventFromPointer(event))
	} else {
		NewCurrentMediaPageFromPointer(ptr).FocusOutEventDefault(std_gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *CurrentMediaPage) FocusOutEventDefault(event std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_FocusOutEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(event))
	}
}

//export callbackCurrentMediaPage46aa22_Hide
func callbackCurrentMediaPage46aa22_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewCurrentMediaPageFromPointer(ptr).HideDefault()
	}
}

func (ptr *CurrentMediaPage) HideDefault() {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_HideDefault(ptr.Pointer())
	}
}

//export callbackCurrentMediaPage46aa22_HideEvent
func callbackCurrentMediaPage46aa22_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		(*(*func(*std_gui.QHideEvent))(signal))(std_gui.NewQHideEventFromPointer(event))
	} else {
		NewCurrentMediaPageFromPointer(ptr).HideEventDefault(std_gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *CurrentMediaPage) HideEventDefault(event std_gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_HideEventDefault(ptr.Pointer(), std_gui.PointerFromQHideEvent(event))
	}
}

//export callbackCurrentMediaPage46aa22_InputMethodEvent
func callbackCurrentMediaPage46aa22_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		(*(*func(*std_gui.QInputMethodEvent))(signal))(std_gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewCurrentMediaPageFromPointer(ptr).InputMethodEventDefault(std_gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *CurrentMediaPage) InputMethodEventDefault(event std_gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_InputMethodEventDefault(ptr.Pointer(), std_gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackCurrentMediaPage46aa22_KeyReleaseEvent
func callbackCurrentMediaPage46aa22_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		(*(*func(*std_gui.QKeyEvent))(signal))(std_gui.NewQKeyEventFromPointer(event))
	} else {
		NewCurrentMediaPageFromPointer(ptr).KeyReleaseEventDefault(std_gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *CurrentMediaPage) KeyReleaseEventDefault(event std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_KeyReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(event))
	}
}

//export callbackCurrentMediaPage46aa22_LeaveEvent
func callbackCurrentMediaPage46aa22_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewCurrentMediaPageFromPointer(ptr).LeaveEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *CurrentMediaPage) LeaveEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_LeaveEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackCurrentMediaPage46aa22_Lower
func callbackCurrentMediaPage46aa22_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewCurrentMediaPageFromPointer(ptr).LowerDefault()
	}
}

func (ptr *CurrentMediaPage) LowerDefault() {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_LowerDefault(ptr.Pointer())
	}
}

//export callbackCurrentMediaPage46aa22_MouseDoubleClickEvent
func callbackCurrentMediaPage46aa22_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewCurrentMediaPageFromPointer(ptr).MouseDoubleClickEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *CurrentMediaPage) MouseDoubleClickEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_MouseDoubleClickEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackCurrentMediaPage46aa22_MouseMoveEvent
func callbackCurrentMediaPage46aa22_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewCurrentMediaPageFromPointer(ptr).MouseMoveEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *CurrentMediaPage) MouseMoveEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_MouseMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackCurrentMediaPage46aa22_MousePressEvent
func callbackCurrentMediaPage46aa22_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewCurrentMediaPageFromPointer(ptr).MousePressEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *CurrentMediaPage) MousePressEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_MousePressEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackCurrentMediaPage46aa22_MouseReleaseEvent
func callbackCurrentMediaPage46aa22_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewCurrentMediaPageFromPointer(ptr).MouseReleaseEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *CurrentMediaPage) MouseReleaseEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_MouseReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackCurrentMediaPage46aa22_MoveEvent
func callbackCurrentMediaPage46aa22_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		(*(*func(*std_gui.QMoveEvent))(signal))(std_gui.NewQMoveEventFromPointer(event))
	} else {
		NewCurrentMediaPageFromPointer(ptr).MoveEventDefault(std_gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *CurrentMediaPage) MoveEventDefault(event std_gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_MoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMoveEvent(event))
	}
}

//export callbackCurrentMediaPage46aa22_Raise
func callbackCurrentMediaPage46aa22_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewCurrentMediaPageFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *CurrentMediaPage) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_RaiseDefault(ptr.Pointer())
	}
}

//export callbackCurrentMediaPage46aa22_Repaint
func callbackCurrentMediaPage46aa22_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewCurrentMediaPageFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *CurrentMediaPage) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_RepaintDefault(ptr.Pointer())
	}
}

//export callbackCurrentMediaPage46aa22_SetDisabled
func callbackCurrentMediaPage46aa22_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		(*(*func(bool))(signal))(int8(disable) != 0)
	} else {
		NewCurrentMediaPageFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *CurrentMediaPage) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackCurrentMediaPage46aa22_SetEnabled
func callbackCurrentMediaPage46aa22_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewCurrentMediaPageFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *CurrentMediaPage) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackCurrentMediaPage46aa22_SetFocus2
func callbackCurrentMediaPage46aa22_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewCurrentMediaPageFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *CurrentMediaPage) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackCurrentMediaPage46aa22_SetHidden
func callbackCurrentMediaPage46aa22_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		(*(*func(bool))(signal))(int8(hidden) != 0)
	} else {
		NewCurrentMediaPageFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *CurrentMediaPage) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackCurrentMediaPage46aa22_SetStyleSheet
func callbackCurrentMediaPage46aa22_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(styleSheet))
	} else {
		NewCurrentMediaPageFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *CurrentMediaPage) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.CurrentMediaPage46aa22_SetStyleSheetDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackCurrentMediaPage46aa22_SetVisible
func callbackCurrentMediaPage46aa22_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	} else {
		NewCurrentMediaPageFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *CurrentMediaPage) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackCurrentMediaPage46aa22_SetWindowModified
func callbackCurrentMediaPage46aa22_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewCurrentMediaPageFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *CurrentMediaPage) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackCurrentMediaPage46aa22_SetWindowTitle
func callbackCurrentMediaPage46aa22_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(vqs))
	} else {
		NewCurrentMediaPageFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *CurrentMediaPage) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.CurrentMediaPage46aa22_SetWindowTitleDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackCurrentMediaPage46aa22_Show
func callbackCurrentMediaPage46aa22_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewCurrentMediaPageFromPointer(ptr).ShowDefault()
	}
}

func (ptr *CurrentMediaPage) ShowDefault() {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_ShowDefault(ptr.Pointer())
	}
}

//export callbackCurrentMediaPage46aa22_ShowFullScreen
func callbackCurrentMediaPage46aa22_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewCurrentMediaPageFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *CurrentMediaPage) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackCurrentMediaPage46aa22_ShowMaximized
func callbackCurrentMediaPage46aa22_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewCurrentMediaPageFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *CurrentMediaPage) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackCurrentMediaPage46aa22_ShowMinimized
func callbackCurrentMediaPage46aa22_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewCurrentMediaPageFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *CurrentMediaPage) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackCurrentMediaPage46aa22_ShowNormal
func callbackCurrentMediaPage46aa22_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewCurrentMediaPageFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *CurrentMediaPage) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackCurrentMediaPage46aa22_TabletEvent
func callbackCurrentMediaPage46aa22_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		(*(*func(*std_gui.QTabletEvent))(signal))(std_gui.NewQTabletEventFromPointer(event))
	} else {
		NewCurrentMediaPageFromPointer(ptr).TabletEventDefault(std_gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *CurrentMediaPage) TabletEventDefault(event std_gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_TabletEventDefault(ptr.Pointer(), std_gui.PointerFromQTabletEvent(event))
	}
}

//export callbackCurrentMediaPage46aa22_Update
func callbackCurrentMediaPage46aa22_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewCurrentMediaPageFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *CurrentMediaPage) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_UpdateDefault(ptr.Pointer())
	}
}

//export callbackCurrentMediaPage46aa22_UpdateMicroFocus
func callbackCurrentMediaPage46aa22_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewCurrentMediaPageFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *CurrentMediaPage) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackCurrentMediaPage46aa22_WheelEvent
func callbackCurrentMediaPage46aa22_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		(*(*func(*std_gui.QWheelEvent))(signal))(std_gui.NewQWheelEventFromPointer(event))
	} else {
		NewCurrentMediaPageFromPointer(ptr).WheelEventDefault(std_gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *CurrentMediaPage) WheelEventDefault(event std_gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_WheelEventDefault(ptr.Pointer(), std_gui.PointerFromQWheelEvent(event))
	}
}

//export callbackCurrentMediaPage46aa22_WindowIconChanged
func callbackCurrentMediaPage46aa22_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		(*(*func(*std_gui.QIcon))(signal))(std_gui.NewQIconFromPointer(icon))
	}

}

//export callbackCurrentMediaPage46aa22_WindowTitleChanged
func callbackCurrentMediaPage46aa22_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(title))
	}

}

//export callbackCurrentMediaPage46aa22_PaintEngine
func callbackCurrentMediaPage46aa22_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return std_gui.PointerFromQPaintEngine((*(*func() *std_gui.QPaintEngine)(signal))())
	}

	return std_gui.PointerFromQPaintEngine(NewCurrentMediaPageFromPointer(ptr).PaintEngineDefault())
}

func (ptr *CurrentMediaPage) PaintEngineDefault() *std_gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return std_gui.NewQPaintEngineFromPointer(C.CurrentMediaPage46aa22_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackCurrentMediaPage46aa22_InputMethodQuery
func callbackCurrentMediaPage46aa22_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(std_core.Qt__InputMethodQuery) *std_core.QVariant)(signal))(std_core.Qt__InputMethodQuery(query)))
	}

	return std_core.PointerFromQVariant(NewCurrentMediaPageFromPointer(ptr).InputMethodQueryDefault(std_core.Qt__InputMethodQuery(query)))
}

func (ptr *CurrentMediaPage) InputMethodQueryDefault(query std_core.Qt__InputMethodQuery) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.CurrentMediaPage46aa22_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackCurrentMediaPage46aa22_Metric
func callbackCurrentMediaPage46aa22_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32((*(*func(std_gui.QPaintDevice__PaintDeviceMetric) int)(signal))(std_gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewCurrentMediaPageFromPointer(ptr).MetricDefault(std_gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *CurrentMediaPage) MetricDefault(m std_gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CurrentMediaPage46aa22_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackCurrentMediaPage46aa22_InitPainter
func callbackCurrentMediaPage46aa22_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		(*(*func(*std_gui.QPainter))(signal))(std_gui.NewQPainterFromPointer(painter))
	} else {
		NewCurrentMediaPageFromPointer(ptr).InitPainterDefault(std_gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *CurrentMediaPage) InitPainterDefault(painter std_gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_InitPainterDefault(ptr.Pointer(), std_gui.PointerFromQPainter(painter))
	}
}

//export callbackCurrentMediaPage46aa22_EventFilter
func callbackCurrentMediaPage46aa22_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCurrentMediaPageFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *CurrentMediaPage) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CurrentMediaPage46aa22_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackCurrentMediaPage46aa22_ChildEvent
func callbackCurrentMediaPage46aa22_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewCurrentMediaPageFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *CurrentMediaPage) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackCurrentMediaPage46aa22_ConnectNotify
func callbackCurrentMediaPage46aa22_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewCurrentMediaPageFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *CurrentMediaPage) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackCurrentMediaPage46aa22_CustomEvent
func callbackCurrentMediaPage46aa22_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewCurrentMediaPageFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *CurrentMediaPage) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackCurrentMediaPage46aa22_DeleteLater
func callbackCurrentMediaPage46aa22_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewCurrentMediaPageFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *CurrentMediaPage) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackCurrentMediaPage46aa22_Destroyed
func callbackCurrentMediaPage46aa22_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackCurrentMediaPage46aa22_DisconnectNotify
func callbackCurrentMediaPage46aa22_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewCurrentMediaPageFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *CurrentMediaPage) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackCurrentMediaPage46aa22_ObjectNameChanged
func callbackCurrentMediaPage46aa22_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackCurrentMediaPage46aa22_TimerEvent
func callbackCurrentMediaPage46aa22_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewCurrentMediaPageFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *CurrentMediaPage) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CurrentMediaPage46aa22_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type MainWindow_ITF interface {
	std_widgets.QMainWindow_ITF
	MainWindow_PTR() *MainWindow
}

func (ptr *MainWindow) MainWindow_PTR() *MainWindow {
	return ptr
}

func (ptr *MainWindow) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QMainWindow_PTR().Pointer()
	}
	return nil
}

func (ptr *MainWindow) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QMainWindow_PTR().SetPointer(p)
	}
}

func PointerFromMainWindow(ptr MainWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.MainWindow_PTR().Pointer()
	}
	return nil
}

func NewMainWindowFromPointer(ptr unsafe.Pointer) (n *MainWindow) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(MainWindow)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *MainWindow:
			n = deduced

		case *std_widgets.QMainWindow:
			n = &MainWindow{QMainWindow: *deduced}

		default:
			n = new(MainWindow)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *MainWindow) Init() { this.init() }

//export callbackMainWindow46aa22_Constructor
func callbackMainWindow46aa22_Constructor(ptr unsafe.Pointer) {
	this := NewMainWindowFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

func MainWindow_QRegisterMetaType() int {
	return int(int32(C.MainWindow46aa22_MainWindow46aa22_QRegisterMetaType()))
}

func (ptr *MainWindow) QRegisterMetaType() int {
	return int(int32(C.MainWindow46aa22_MainWindow46aa22_QRegisterMetaType()))
}

func MainWindow_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.MainWindow46aa22_MainWindow46aa22_QRegisterMetaType2(typeNameC)))
}

func (ptr *MainWindow) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.MainWindow46aa22_MainWindow46aa22_QRegisterMetaType2(typeNameC)))
}

func MainWindow_QmlRegisterType() int {
	return int(int32(C.MainWindow46aa22_MainWindow46aa22_QmlRegisterType()))
}

func (ptr *MainWindow) QmlRegisterType() int {
	return int(int32(C.MainWindow46aa22_MainWindow46aa22_QmlRegisterType()))
}

func MainWindow_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.MainWindow46aa22_MainWindow46aa22_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *MainWindow) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.MainWindow46aa22_MainWindow46aa22_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *MainWindow) __resizeDocks_docks_atList(i int) *std_widgets.QDockWidget {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQDockWidgetFromPointer(C.MainWindow46aa22___resizeDocks_docks_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MainWindow) __resizeDocks_docks_setList(i std_widgets.QDockWidget_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22___resizeDocks_docks_setList(ptr.Pointer(), std_widgets.PointerFromQDockWidget(i))
	}
}

func (ptr *MainWindow) __resizeDocks_docks_newList() unsafe.Pointer {
	return C.MainWindow46aa22___resizeDocks_docks_newList(ptr.Pointer())
}

func (ptr *MainWindow) __resizeDocks_sizes_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.MainWindow46aa22___resizeDocks_sizes_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *MainWindow) __resizeDocks_sizes_setList(i int) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22___resizeDocks_sizes_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *MainWindow) __resizeDocks_sizes_newList() unsafe.Pointer {
	return C.MainWindow46aa22___resizeDocks_sizes_newList(ptr.Pointer())
}

func (ptr *MainWindow) __tabifiedDockWidgets_atList(i int) *std_widgets.QDockWidget {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQDockWidgetFromPointer(C.MainWindow46aa22___tabifiedDockWidgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MainWindow) __tabifiedDockWidgets_setList(i std_widgets.QDockWidget_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22___tabifiedDockWidgets_setList(ptr.Pointer(), std_widgets.PointerFromQDockWidget(i))
	}
}

func (ptr *MainWindow) __tabifiedDockWidgets_newList() unsafe.Pointer {
	return C.MainWindow46aa22___tabifiedDockWidgets_newList(ptr.Pointer())
}

func (ptr *MainWindow) __addActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.MainWindow46aa22___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MainWindow) __addActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22___addActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *MainWindow) __addActions_actions_newList() unsafe.Pointer {
	return C.MainWindow46aa22___addActions_actions_newList(ptr.Pointer())
}

func (ptr *MainWindow) __insertActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.MainWindow46aa22___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MainWindow) __insertActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22___insertActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *MainWindow) __insertActions_actions_newList() unsafe.Pointer {
	return C.MainWindow46aa22___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *MainWindow) __actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.MainWindow46aa22___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MainWindow) __actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22___actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *MainWindow) __actions_newList() unsafe.Pointer {
	return C.MainWindow46aa22___actions_newList(ptr.Pointer())
}

func (ptr *MainWindow) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.MainWindow46aa22___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *MainWindow) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *MainWindow) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.MainWindow46aa22___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *MainWindow) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.MainWindow46aa22___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MainWindow) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *MainWindow) __findChildren_newList2() unsafe.Pointer {
	return C.MainWindow46aa22___findChildren_newList2(ptr.Pointer())
}

func (ptr *MainWindow) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.MainWindow46aa22___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MainWindow) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *MainWindow) __findChildren_newList3() unsafe.Pointer {
	return C.MainWindow46aa22___findChildren_newList3(ptr.Pointer())
}

func (ptr *MainWindow) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.MainWindow46aa22___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MainWindow) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *MainWindow) __findChildren_newList() unsafe.Pointer {
	return C.MainWindow46aa22___findChildren_newList(ptr.Pointer())
}

func (ptr *MainWindow) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.MainWindow46aa22___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MainWindow) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *MainWindow) __children_newList() unsafe.Pointer {
	return C.MainWindow46aa22___children_newList(ptr.Pointer())
}

func NewMainWindow(parent std_widgets.QWidget_ITF, flags std_core.Qt__WindowType) *MainWindow {
	tmpValue := NewMainWindowFromPointer(C.MainWindow46aa22_NewMainWindow(std_widgets.PointerFromQWidget(parent), C.longlong(flags)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackMainWindow46aa22_DestroyMainWindow
func callbackMainWindow46aa22_DestroyMainWindow(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~MainWindow"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).DestroyMainWindowDefault()
	}
}

func (ptr *MainWindow) ConnectDestroyMainWindow(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~MainWindow"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~MainWindow", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~MainWindow", unsafe.Pointer(&f))
		}
	}
}

func (ptr *MainWindow) DisconnectDestroyMainWindow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~MainWindow")
	}
}

func (ptr *MainWindow) DestroyMainWindow() {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_DestroyMainWindow(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *MainWindow) DestroyMainWindowDefault() {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_DestroyMainWindowDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackMainWindow46aa22_CreatePopupMenu
func callbackMainWindow46aa22_CreatePopupMenu(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createPopupMenu"); signal != nil {
		return std_widgets.PointerFromQMenu((*(*func() *std_widgets.QMenu)(signal))())
	}

	return std_widgets.PointerFromQMenu(NewMainWindowFromPointer(ptr).CreatePopupMenuDefault())
}

func (ptr *MainWindow) CreatePopupMenuDefault() *std_widgets.QMenu {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQMenuFromPointer(C.MainWindow46aa22_CreatePopupMenuDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackMainWindow46aa22_Event
func callbackMainWindow46aa22_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMainWindowFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(event)))))
}

func (ptr *MainWindow) EventDefault(event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.MainWindow46aa22_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackMainWindow46aa22_ContextMenuEvent
func callbackMainWindow46aa22_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		(*(*func(*std_gui.QContextMenuEvent))(signal))(std_gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).ContextMenuEventDefault(std_gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *MainWindow) ContextMenuEventDefault(event std_gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_ContextMenuEventDefault(ptr.Pointer(), std_gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackMainWindow46aa22_IconSizeChanged
func callbackMainWindow46aa22_IconSizeChanged(ptr unsafe.Pointer, iconSize unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "iconSizeChanged"); signal != nil {
		(*(*func(*std_core.QSize))(signal))(std_core.NewQSizeFromPointer(iconSize))
	}

}

//export callbackMainWindow46aa22_SetAnimated
func callbackMainWindow46aa22_SetAnimated(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(ptr, "setAnimated"); signal != nil {
		(*(*func(bool))(signal))(int8(enabled) != 0)
	} else {
		NewMainWindowFromPointer(ptr).SetAnimatedDefault(int8(enabled) != 0)
	}
}

func (ptr *MainWindow) SetAnimatedDefault(enabled bool) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_SetAnimatedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackMainWindow46aa22_SetDockNestingEnabled
func callbackMainWindow46aa22_SetDockNestingEnabled(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(ptr, "setDockNestingEnabled"); signal != nil {
		(*(*func(bool))(signal))(int8(enabled) != 0)
	} else {
		NewMainWindowFromPointer(ptr).SetDockNestingEnabledDefault(int8(enabled) != 0)
	}
}

func (ptr *MainWindow) SetDockNestingEnabledDefault(enabled bool) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_SetDockNestingEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackMainWindow46aa22_SetUnifiedTitleAndToolBarOnMac
func callbackMainWindow46aa22_SetUnifiedTitleAndToolBarOnMac(ptr unsafe.Pointer, set C.char) {
	if signal := qt.GetSignal(ptr, "setUnifiedTitleAndToolBarOnMac"); signal != nil {
		(*(*func(bool))(signal))(int8(set) != 0)
	} else {
		NewMainWindowFromPointer(ptr).SetUnifiedTitleAndToolBarOnMacDefault(int8(set) != 0)
	}
}

func (ptr *MainWindow) SetUnifiedTitleAndToolBarOnMacDefault(set bool) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_SetUnifiedTitleAndToolBarOnMacDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(set))))
	}
}

//export callbackMainWindow46aa22_TabifiedDockWidgetActivated
func callbackMainWindow46aa22_TabifiedDockWidgetActivated(ptr unsafe.Pointer, dockWidget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabifiedDockWidgetActivated"); signal != nil {
		(*(*func(*std_widgets.QDockWidget))(signal))(std_widgets.NewQDockWidgetFromPointer(dockWidget))
	}

}

//export callbackMainWindow46aa22_ToolButtonStyleChanged
func callbackMainWindow46aa22_ToolButtonStyleChanged(ptr unsafe.Pointer, toolButtonStyle C.longlong) {
	if signal := qt.GetSignal(ptr, "toolButtonStyleChanged"); signal != nil {
		(*(*func(std_core.Qt__ToolButtonStyle))(signal))(std_core.Qt__ToolButtonStyle(toolButtonStyle))
	}

}

//export callbackMainWindow46aa22_Close
func callbackMainWindow46aa22_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewMainWindowFromPointer(ptr).CloseDefault())))
}

func (ptr *MainWindow) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.MainWindow46aa22_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackMainWindow46aa22_FocusNextPrevChild
func callbackMainWindow46aa22_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(bool) bool)(signal))(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMainWindowFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *MainWindow) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.MainWindow46aa22_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackMainWindow46aa22_NativeEvent
func callbackMainWindow46aa22_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result *C.long) C.char {
	var resultR int
	if result != nil {
		resultR = int(int32(*result))
		defer func() { *result = C.long(int32(resultR)) }()
	}
	if signal := qt.GetSignal(ptr, "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QByteArray, unsafe.Pointer, *int) bool)(signal))(std_core.NewQByteArrayFromPointer(eventType), message, &resultR))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMainWindowFromPointer(ptr).NativeEventDefault(std_core.NewQByteArrayFromPointer(eventType), message, &resultR))))
}

func (ptr *MainWindow) NativeEventDefault(eventType std_core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {
	if ptr.Pointer() != nil {
		var resultC C.long
		if result != nil {
			resultC = C.long(int32(*result))
			defer func() { *result = int(int32(resultC)) }()
		}
		return int8(C.MainWindow46aa22_NativeEventDefault(ptr.Pointer(), std_core.PointerFromQByteArray(eventType), message, &resultC)) != 0
	}
	return false
}

//export callbackMainWindow46aa22_ActionEvent
func callbackMainWindow46aa22_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		(*(*func(*std_gui.QActionEvent))(signal))(std_gui.NewQActionEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).ActionEventDefault(std_gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *MainWindow) ActionEventDefault(event std_gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_ActionEventDefault(ptr.Pointer(), std_gui.PointerFromQActionEvent(event))
	}
}

//export callbackMainWindow46aa22_ChangeEvent
func callbackMainWindow46aa22_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).ChangeEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *MainWindow) ChangeEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_ChangeEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackMainWindow46aa22_CloseEvent
func callbackMainWindow46aa22_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		(*(*func(*std_gui.QCloseEvent))(signal))(std_gui.NewQCloseEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).CloseEventDefault(std_gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *MainWindow) CloseEventDefault(event std_gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_CloseEventDefault(ptr.Pointer(), std_gui.PointerFromQCloseEvent(event))
	}
}

//export callbackMainWindow46aa22_CustomContextMenuRequested
func callbackMainWindow46aa22_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		(*(*func(*std_core.QPoint))(signal))(std_core.NewQPointFromPointer(pos))
	}

}

//export callbackMainWindow46aa22_DragEnterEvent
func callbackMainWindow46aa22_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		(*(*func(*std_gui.QDragEnterEvent))(signal))(std_gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).DragEnterEventDefault(std_gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *MainWindow) DragEnterEventDefault(event std_gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_DragEnterEventDefault(ptr.Pointer(), std_gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackMainWindow46aa22_DragLeaveEvent
func callbackMainWindow46aa22_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		(*(*func(*std_gui.QDragLeaveEvent))(signal))(std_gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).DragLeaveEventDefault(std_gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *MainWindow) DragLeaveEventDefault(event std_gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_DragLeaveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackMainWindow46aa22_DragMoveEvent
func callbackMainWindow46aa22_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		(*(*func(*std_gui.QDragMoveEvent))(signal))(std_gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).DragMoveEventDefault(std_gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *MainWindow) DragMoveEventDefault(event std_gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_DragMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackMainWindow46aa22_DropEvent
func callbackMainWindow46aa22_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		(*(*func(*std_gui.QDropEvent))(signal))(std_gui.NewQDropEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).DropEventDefault(std_gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *MainWindow) DropEventDefault(event std_gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_DropEventDefault(ptr.Pointer(), std_gui.PointerFromQDropEvent(event))
	}
}

//export callbackMainWindow46aa22_EnterEvent
func callbackMainWindow46aa22_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).EnterEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *MainWindow) EnterEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_EnterEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackMainWindow46aa22_FocusInEvent
func callbackMainWindow46aa22_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		(*(*func(*std_gui.QFocusEvent))(signal))(std_gui.NewQFocusEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).FocusInEventDefault(std_gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *MainWindow) FocusInEventDefault(event std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_FocusInEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(event))
	}
}

//export callbackMainWindow46aa22_FocusOutEvent
func callbackMainWindow46aa22_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		(*(*func(*std_gui.QFocusEvent))(signal))(std_gui.NewQFocusEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).FocusOutEventDefault(std_gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *MainWindow) FocusOutEventDefault(event std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_FocusOutEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(event))
	}
}

//export callbackMainWindow46aa22_Hide
func callbackMainWindow46aa22_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).HideDefault()
	}
}

func (ptr *MainWindow) HideDefault() {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_HideDefault(ptr.Pointer())
	}
}

//export callbackMainWindow46aa22_HideEvent
func callbackMainWindow46aa22_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		(*(*func(*std_gui.QHideEvent))(signal))(std_gui.NewQHideEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).HideEventDefault(std_gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *MainWindow) HideEventDefault(event std_gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_HideEventDefault(ptr.Pointer(), std_gui.PointerFromQHideEvent(event))
	}
}

//export callbackMainWindow46aa22_InputMethodEvent
func callbackMainWindow46aa22_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		(*(*func(*std_gui.QInputMethodEvent))(signal))(std_gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).InputMethodEventDefault(std_gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *MainWindow) InputMethodEventDefault(event std_gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_InputMethodEventDefault(ptr.Pointer(), std_gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackMainWindow46aa22_KeyPressEvent
func callbackMainWindow46aa22_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		(*(*func(*std_gui.QKeyEvent))(signal))(std_gui.NewQKeyEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).KeyPressEventDefault(std_gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *MainWindow) KeyPressEventDefault(event std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_KeyPressEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(event))
	}
}

//export callbackMainWindow46aa22_KeyReleaseEvent
func callbackMainWindow46aa22_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		(*(*func(*std_gui.QKeyEvent))(signal))(std_gui.NewQKeyEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).KeyReleaseEventDefault(std_gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *MainWindow) KeyReleaseEventDefault(event std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_KeyReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(event))
	}
}

//export callbackMainWindow46aa22_LeaveEvent
func callbackMainWindow46aa22_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).LeaveEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *MainWindow) LeaveEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_LeaveEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackMainWindow46aa22_Lower
func callbackMainWindow46aa22_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).LowerDefault()
	}
}

func (ptr *MainWindow) LowerDefault() {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_LowerDefault(ptr.Pointer())
	}
}

//export callbackMainWindow46aa22_MouseDoubleClickEvent
func callbackMainWindow46aa22_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).MouseDoubleClickEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *MainWindow) MouseDoubleClickEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_MouseDoubleClickEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackMainWindow46aa22_MouseMoveEvent
func callbackMainWindow46aa22_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).MouseMoveEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *MainWindow) MouseMoveEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_MouseMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackMainWindow46aa22_MousePressEvent
func callbackMainWindow46aa22_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).MousePressEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *MainWindow) MousePressEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_MousePressEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackMainWindow46aa22_MouseReleaseEvent
func callbackMainWindow46aa22_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).MouseReleaseEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *MainWindow) MouseReleaseEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_MouseReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackMainWindow46aa22_MoveEvent
func callbackMainWindow46aa22_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		(*(*func(*std_gui.QMoveEvent))(signal))(std_gui.NewQMoveEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).MoveEventDefault(std_gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *MainWindow) MoveEventDefault(event std_gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_MoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMoveEvent(event))
	}
}

//export callbackMainWindow46aa22_PaintEvent
func callbackMainWindow46aa22_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		(*(*func(*std_gui.QPaintEvent))(signal))(std_gui.NewQPaintEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).PaintEventDefault(std_gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *MainWindow) PaintEventDefault(event std_gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_PaintEventDefault(ptr.Pointer(), std_gui.PointerFromQPaintEvent(event))
	}
}

//export callbackMainWindow46aa22_Raise
func callbackMainWindow46aa22_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *MainWindow) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_RaiseDefault(ptr.Pointer())
	}
}

//export callbackMainWindow46aa22_Repaint
func callbackMainWindow46aa22_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *MainWindow) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_RepaintDefault(ptr.Pointer())
	}
}

//export callbackMainWindow46aa22_ResizeEvent
func callbackMainWindow46aa22_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		(*(*func(*std_gui.QResizeEvent))(signal))(std_gui.NewQResizeEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).ResizeEventDefault(std_gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *MainWindow) ResizeEventDefault(event std_gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_ResizeEventDefault(ptr.Pointer(), std_gui.PointerFromQResizeEvent(event))
	}
}

//export callbackMainWindow46aa22_SetDisabled
func callbackMainWindow46aa22_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		(*(*func(bool))(signal))(int8(disable) != 0)
	} else {
		NewMainWindowFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *MainWindow) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackMainWindow46aa22_SetEnabled
func callbackMainWindow46aa22_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewMainWindowFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *MainWindow) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackMainWindow46aa22_SetFocus2
func callbackMainWindow46aa22_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *MainWindow) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackMainWindow46aa22_SetHidden
func callbackMainWindow46aa22_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		(*(*func(bool))(signal))(int8(hidden) != 0)
	} else {
		NewMainWindowFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *MainWindow) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackMainWindow46aa22_SetStyleSheet
func callbackMainWindow46aa22_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(styleSheet))
	} else {
		NewMainWindowFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *MainWindow) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.MainWindow46aa22_SetStyleSheetDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackMainWindow46aa22_SetVisible
func callbackMainWindow46aa22_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	} else {
		NewMainWindowFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *MainWindow) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackMainWindow46aa22_SetWindowModified
func callbackMainWindow46aa22_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewMainWindowFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *MainWindow) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackMainWindow46aa22_SetWindowTitle
func callbackMainWindow46aa22_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(vqs))
	} else {
		NewMainWindowFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *MainWindow) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.MainWindow46aa22_SetWindowTitleDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackMainWindow46aa22_Show
func callbackMainWindow46aa22_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).ShowDefault()
	}
}

func (ptr *MainWindow) ShowDefault() {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_ShowDefault(ptr.Pointer())
	}
}

//export callbackMainWindow46aa22_ShowEvent
func callbackMainWindow46aa22_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		(*(*func(*std_gui.QShowEvent))(signal))(std_gui.NewQShowEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).ShowEventDefault(std_gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *MainWindow) ShowEventDefault(event std_gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_ShowEventDefault(ptr.Pointer(), std_gui.PointerFromQShowEvent(event))
	}
}

//export callbackMainWindow46aa22_ShowFullScreen
func callbackMainWindow46aa22_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *MainWindow) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackMainWindow46aa22_ShowMaximized
func callbackMainWindow46aa22_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *MainWindow) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackMainWindow46aa22_ShowMinimized
func callbackMainWindow46aa22_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *MainWindow) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackMainWindow46aa22_ShowNormal
func callbackMainWindow46aa22_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *MainWindow) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackMainWindow46aa22_TabletEvent
func callbackMainWindow46aa22_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		(*(*func(*std_gui.QTabletEvent))(signal))(std_gui.NewQTabletEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).TabletEventDefault(std_gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *MainWindow) TabletEventDefault(event std_gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_TabletEventDefault(ptr.Pointer(), std_gui.PointerFromQTabletEvent(event))
	}
}

//export callbackMainWindow46aa22_Update
func callbackMainWindow46aa22_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *MainWindow) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_UpdateDefault(ptr.Pointer())
	}
}

//export callbackMainWindow46aa22_UpdateMicroFocus
func callbackMainWindow46aa22_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *MainWindow) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackMainWindow46aa22_WheelEvent
func callbackMainWindow46aa22_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		(*(*func(*std_gui.QWheelEvent))(signal))(std_gui.NewQWheelEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).WheelEventDefault(std_gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *MainWindow) WheelEventDefault(event std_gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_WheelEventDefault(ptr.Pointer(), std_gui.PointerFromQWheelEvent(event))
	}
}

//export callbackMainWindow46aa22_WindowIconChanged
func callbackMainWindow46aa22_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		(*(*func(*std_gui.QIcon))(signal))(std_gui.NewQIconFromPointer(icon))
	}

}

//export callbackMainWindow46aa22_WindowTitleChanged
func callbackMainWindow46aa22_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(title))
	}

}

//export callbackMainWindow46aa22_PaintEngine
func callbackMainWindow46aa22_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return std_gui.PointerFromQPaintEngine((*(*func() *std_gui.QPaintEngine)(signal))())
	}

	return std_gui.PointerFromQPaintEngine(NewMainWindowFromPointer(ptr).PaintEngineDefault())
}

func (ptr *MainWindow) PaintEngineDefault() *std_gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return std_gui.NewQPaintEngineFromPointer(C.MainWindow46aa22_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackMainWindow46aa22_MinimumSizeHint
func callbackMainWindow46aa22_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return std_core.PointerFromQSize((*(*func() *std_core.QSize)(signal))())
	}

	return std_core.PointerFromQSize(NewMainWindowFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *MainWindow) MinimumSizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.MainWindow46aa22_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackMainWindow46aa22_SizeHint
func callbackMainWindow46aa22_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return std_core.PointerFromQSize((*(*func() *std_core.QSize)(signal))())
	}

	return std_core.PointerFromQSize(NewMainWindowFromPointer(ptr).SizeHintDefault())
}

func (ptr *MainWindow) SizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.MainWindow46aa22_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackMainWindow46aa22_InputMethodQuery
func callbackMainWindow46aa22_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(std_core.Qt__InputMethodQuery) *std_core.QVariant)(signal))(std_core.Qt__InputMethodQuery(query)))
	}

	return std_core.PointerFromQVariant(NewMainWindowFromPointer(ptr).InputMethodQueryDefault(std_core.Qt__InputMethodQuery(query)))
}

func (ptr *MainWindow) InputMethodQueryDefault(query std_core.Qt__InputMethodQuery) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.MainWindow46aa22_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackMainWindow46aa22_HasHeightForWidth
func callbackMainWindow46aa22_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewMainWindowFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *MainWindow) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.MainWindow46aa22_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackMainWindow46aa22_HeightForWidth
func callbackMainWindow46aa22_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32((*(*func(int) int)(signal))(int(int32(w)))))
	}

	return C.int(int32(NewMainWindowFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *MainWindow) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.MainWindow46aa22_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackMainWindow46aa22_Metric
func callbackMainWindow46aa22_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32((*(*func(std_gui.QPaintDevice__PaintDeviceMetric) int)(signal))(std_gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewMainWindowFromPointer(ptr).MetricDefault(std_gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *MainWindow) MetricDefault(m std_gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.MainWindow46aa22_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackMainWindow46aa22_InitPainter
func callbackMainWindow46aa22_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		(*(*func(*std_gui.QPainter))(signal))(std_gui.NewQPainterFromPointer(painter))
	} else {
		NewMainWindowFromPointer(ptr).InitPainterDefault(std_gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *MainWindow) InitPainterDefault(painter std_gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_InitPainterDefault(ptr.Pointer(), std_gui.PointerFromQPainter(painter))
	}
}

//export callbackMainWindow46aa22_EventFilter
func callbackMainWindow46aa22_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMainWindowFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *MainWindow) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.MainWindow46aa22_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackMainWindow46aa22_ChildEvent
func callbackMainWindow46aa22_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *MainWindow) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackMainWindow46aa22_ConnectNotify
func callbackMainWindow46aa22_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewMainWindowFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *MainWindow) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackMainWindow46aa22_CustomEvent
func callbackMainWindow46aa22_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *MainWindow) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackMainWindow46aa22_DeleteLater
func callbackMainWindow46aa22_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *MainWindow) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackMainWindow46aa22_Destroyed
func callbackMainWindow46aa22_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackMainWindow46aa22_DisconnectNotify
func callbackMainWindow46aa22_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewMainWindowFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *MainWindow) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackMainWindow46aa22_ObjectNameChanged
func callbackMainWindow46aa22_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackMainWindow46aa22_TimerEvent
func callbackMainWindow46aa22_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *MainWindow) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow46aa22_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type MediaListPage_ITF interface {
	std_widgets.QWidget_ITF
	MediaListPage_PTR() *MediaListPage
}

func (ptr *MediaListPage) MediaListPage_PTR() *MediaListPage {
	return ptr
}

func (ptr *MediaListPage) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *MediaListPage) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromMediaListPage(ptr MediaListPage_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.MediaListPage_PTR().Pointer()
	}
	return nil
}

func NewMediaListPageFromPointer(ptr unsafe.Pointer) (n *MediaListPage) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(MediaListPage)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *MediaListPage:
			n = deduced

		case *std_widgets.QWidget:
			n = &MediaListPage{QWidget: *deduced}

		default:
			n = new(MediaListPage)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *MediaListPage) Init() { this.init() }

//export callbackMediaListPage46aa22_Constructor
func callbackMediaListPage46aa22_Constructor(ptr unsafe.Pointer) {
	this := NewMediaListPageFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

func MediaListPage_QRegisterMetaType() int {
	return int(int32(C.MediaListPage46aa22_MediaListPage46aa22_QRegisterMetaType()))
}

func (ptr *MediaListPage) QRegisterMetaType() int {
	return int(int32(C.MediaListPage46aa22_MediaListPage46aa22_QRegisterMetaType()))
}

func MediaListPage_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.MediaListPage46aa22_MediaListPage46aa22_QRegisterMetaType2(typeNameC)))
}

func (ptr *MediaListPage) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.MediaListPage46aa22_MediaListPage46aa22_QRegisterMetaType2(typeNameC)))
}

func MediaListPage_QmlRegisterType() int {
	return int(int32(C.MediaListPage46aa22_MediaListPage46aa22_QmlRegisterType()))
}

func (ptr *MediaListPage) QmlRegisterType() int {
	return int(int32(C.MediaListPage46aa22_MediaListPage46aa22_QmlRegisterType()))
}

func MediaListPage_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.MediaListPage46aa22_MediaListPage46aa22_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *MediaListPage) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.MediaListPage46aa22_MediaListPage46aa22_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *MediaListPage) __addActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.MediaListPage46aa22___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MediaListPage) __addActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22___addActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *MediaListPage) __addActions_actions_newList() unsafe.Pointer {
	return C.MediaListPage46aa22___addActions_actions_newList(ptr.Pointer())
}

func (ptr *MediaListPage) __insertActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.MediaListPage46aa22___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MediaListPage) __insertActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22___insertActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *MediaListPage) __insertActions_actions_newList() unsafe.Pointer {
	return C.MediaListPage46aa22___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *MediaListPage) __actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.MediaListPage46aa22___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MediaListPage) __actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22___actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *MediaListPage) __actions_newList() unsafe.Pointer {
	return C.MediaListPage46aa22___actions_newList(ptr.Pointer())
}

func (ptr *MediaListPage) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.MediaListPage46aa22___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *MediaListPage) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *MediaListPage) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.MediaListPage46aa22___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *MediaListPage) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.MediaListPage46aa22___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MediaListPage) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *MediaListPage) __findChildren_newList2() unsafe.Pointer {
	return C.MediaListPage46aa22___findChildren_newList2(ptr.Pointer())
}

func (ptr *MediaListPage) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.MediaListPage46aa22___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MediaListPage) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *MediaListPage) __findChildren_newList3() unsafe.Pointer {
	return C.MediaListPage46aa22___findChildren_newList3(ptr.Pointer())
}

func (ptr *MediaListPage) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.MediaListPage46aa22___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MediaListPage) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *MediaListPage) __findChildren_newList() unsafe.Pointer {
	return C.MediaListPage46aa22___findChildren_newList(ptr.Pointer())
}

func (ptr *MediaListPage) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.MediaListPage46aa22___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MediaListPage) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *MediaListPage) __children_newList() unsafe.Pointer {
	return C.MediaListPage46aa22___children_newList(ptr.Pointer())
}

func NewMediaListPage(parent std_widgets.QWidget_ITF, ff std_core.Qt__WindowType) *MediaListPage {
	tmpValue := NewMediaListPageFromPointer(C.MediaListPage46aa22_NewMediaListPage(std_widgets.PointerFromQWidget(parent), C.longlong(ff)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackMediaListPage46aa22_DestroyMediaListPage
func callbackMediaListPage46aa22_DestroyMediaListPage(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~MediaListPage"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListPageFromPointer(ptr).DestroyMediaListPageDefault()
	}
}

func (ptr *MediaListPage) ConnectDestroyMediaListPage(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~MediaListPage"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~MediaListPage", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~MediaListPage", unsafe.Pointer(&f))
		}
	}
}

func (ptr *MediaListPage) DisconnectDestroyMediaListPage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~MediaListPage")
	}
}

func (ptr *MediaListPage) DestroyMediaListPage() {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_DestroyMediaListPage(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *MediaListPage) DestroyMediaListPageDefault() {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_DestroyMediaListPageDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackMediaListPage46aa22_Close
func callbackMediaListPage46aa22_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListPageFromPointer(ptr).CloseDefault())))
}

func (ptr *MediaListPage) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.MediaListPage46aa22_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackMediaListPage46aa22_Event
func callbackMediaListPage46aa22_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListPageFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(event)))))
}

func (ptr *MediaListPage) EventDefault(event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.MediaListPage46aa22_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackMediaListPage46aa22_FocusNextPrevChild
func callbackMediaListPage46aa22_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(bool) bool)(signal))(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListPageFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *MediaListPage) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.MediaListPage46aa22_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackMediaListPage46aa22_NativeEvent
func callbackMediaListPage46aa22_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result *C.long) C.char {
	var resultR int
	if result != nil {
		resultR = int(int32(*result))
		defer func() { *result = C.long(int32(resultR)) }()
	}
	if signal := qt.GetSignal(ptr, "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QByteArray, unsafe.Pointer, *int) bool)(signal))(std_core.NewQByteArrayFromPointer(eventType), message, &resultR))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListPageFromPointer(ptr).NativeEventDefault(std_core.NewQByteArrayFromPointer(eventType), message, &resultR))))
}

func (ptr *MediaListPage) NativeEventDefault(eventType std_core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {
	if ptr.Pointer() != nil {
		var resultC C.long
		if result != nil {
			resultC = C.long(int32(*result))
			defer func() { *result = int(int32(resultC)) }()
		}
		return int8(C.MediaListPage46aa22_NativeEventDefault(ptr.Pointer(), std_core.PointerFromQByteArray(eventType), message, &resultC)) != 0
	}
	return false
}

//export callbackMediaListPage46aa22_ActionEvent
func callbackMediaListPage46aa22_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		(*(*func(*std_gui.QActionEvent))(signal))(std_gui.NewQActionEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).ActionEventDefault(std_gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *MediaListPage) ActionEventDefault(event std_gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_ActionEventDefault(ptr.Pointer(), std_gui.PointerFromQActionEvent(event))
	}
}

//export callbackMediaListPage46aa22_ChangeEvent
func callbackMediaListPage46aa22_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).ChangeEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *MediaListPage) ChangeEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_ChangeEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackMediaListPage46aa22_CloseEvent
func callbackMediaListPage46aa22_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		(*(*func(*std_gui.QCloseEvent))(signal))(std_gui.NewQCloseEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).CloseEventDefault(std_gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *MediaListPage) CloseEventDefault(event std_gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_CloseEventDefault(ptr.Pointer(), std_gui.PointerFromQCloseEvent(event))
	}
}

//export callbackMediaListPage46aa22_ContextMenuEvent
func callbackMediaListPage46aa22_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		(*(*func(*std_gui.QContextMenuEvent))(signal))(std_gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).ContextMenuEventDefault(std_gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *MediaListPage) ContextMenuEventDefault(event std_gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_ContextMenuEventDefault(ptr.Pointer(), std_gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackMediaListPage46aa22_CustomContextMenuRequested
func callbackMediaListPage46aa22_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		(*(*func(*std_core.QPoint))(signal))(std_core.NewQPointFromPointer(pos))
	}

}

//export callbackMediaListPage46aa22_DragEnterEvent
func callbackMediaListPage46aa22_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		(*(*func(*std_gui.QDragEnterEvent))(signal))(std_gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).DragEnterEventDefault(std_gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *MediaListPage) DragEnterEventDefault(event std_gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_DragEnterEventDefault(ptr.Pointer(), std_gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackMediaListPage46aa22_DragLeaveEvent
func callbackMediaListPage46aa22_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		(*(*func(*std_gui.QDragLeaveEvent))(signal))(std_gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).DragLeaveEventDefault(std_gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *MediaListPage) DragLeaveEventDefault(event std_gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_DragLeaveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackMediaListPage46aa22_DragMoveEvent
func callbackMediaListPage46aa22_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		(*(*func(*std_gui.QDragMoveEvent))(signal))(std_gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).DragMoveEventDefault(std_gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *MediaListPage) DragMoveEventDefault(event std_gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_DragMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackMediaListPage46aa22_DropEvent
func callbackMediaListPage46aa22_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		(*(*func(*std_gui.QDropEvent))(signal))(std_gui.NewQDropEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).DropEventDefault(std_gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *MediaListPage) DropEventDefault(event std_gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_DropEventDefault(ptr.Pointer(), std_gui.PointerFromQDropEvent(event))
	}
}

//export callbackMediaListPage46aa22_EnterEvent
func callbackMediaListPage46aa22_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).EnterEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *MediaListPage) EnterEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_EnterEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackMediaListPage46aa22_FocusInEvent
func callbackMediaListPage46aa22_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		(*(*func(*std_gui.QFocusEvent))(signal))(std_gui.NewQFocusEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).FocusInEventDefault(std_gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *MediaListPage) FocusInEventDefault(event std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_FocusInEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(event))
	}
}

//export callbackMediaListPage46aa22_FocusOutEvent
func callbackMediaListPage46aa22_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		(*(*func(*std_gui.QFocusEvent))(signal))(std_gui.NewQFocusEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).FocusOutEventDefault(std_gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *MediaListPage) FocusOutEventDefault(event std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_FocusOutEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(event))
	}
}

//export callbackMediaListPage46aa22_Hide
func callbackMediaListPage46aa22_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListPageFromPointer(ptr).HideDefault()
	}
}

func (ptr *MediaListPage) HideDefault() {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_HideDefault(ptr.Pointer())
	}
}

//export callbackMediaListPage46aa22_HideEvent
func callbackMediaListPage46aa22_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		(*(*func(*std_gui.QHideEvent))(signal))(std_gui.NewQHideEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).HideEventDefault(std_gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *MediaListPage) HideEventDefault(event std_gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_HideEventDefault(ptr.Pointer(), std_gui.PointerFromQHideEvent(event))
	}
}

//export callbackMediaListPage46aa22_InputMethodEvent
func callbackMediaListPage46aa22_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		(*(*func(*std_gui.QInputMethodEvent))(signal))(std_gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).InputMethodEventDefault(std_gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *MediaListPage) InputMethodEventDefault(event std_gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_InputMethodEventDefault(ptr.Pointer(), std_gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackMediaListPage46aa22_KeyPressEvent
func callbackMediaListPage46aa22_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		(*(*func(*std_gui.QKeyEvent))(signal))(std_gui.NewQKeyEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).KeyPressEventDefault(std_gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *MediaListPage) KeyPressEventDefault(event std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_KeyPressEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(event))
	}
}

//export callbackMediaListPage46aa22_KeyReleaseEvent
func callbackMediaListPage46aa22_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		(*(*func(*std_gui.QKeyEvent))(signal))(std_gui.NewQKeyEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).KeyReleaseEventDefault(std_gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *MediaListPage) KeyReleaseEventDefault(event std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_KeyReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(event))
	}
}

//export callbackMediaListPage46aa22_LeaveEvent
func callbackMediaListPage46aa22_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).LeaveEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *MediaListPage) LeaveEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_LeaveEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackMediaListPage46aa22_Lower
func callbackMediaListPage46aa22_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListPageFromPointer(ptr).LowerDefault()
	}
}

func (ptr *MediaListPage) LowerDefault() {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_LowerDefault(ptr.Pointer())
	}
}

//export callbackMediaListPage46aa22_MouseDoubleClickEvent
func callbackMediaListPage46aa22_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).MouseDoubleClickEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *MediaListPage) MouseDoubleClickEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_MouseDoubleClickEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackMediaListPage46aa22_MouseMoveEvent
func callbackMediaListPage46aa22_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).MouseMoveEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *MediaListPage) MouseMoveEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_MouseMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackMediaListPage46aa22_MousePressEvent
func callbackMediaListPage46aa22_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).MousePressEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *MediaListPage) MousePressEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_MousePressEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackMediaListPage46aa22_MouseReleaseEvent
func callbackMediaListPage46aa22_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).MouseReleaseEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *MediaListPage) MouseReleaseEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_MouseReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackMediaListPage46aa22_MoveEvent
func callbackMediaListPage46aa22_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		(*(*func(*std_gui.QMoveEvent))(signal))(std_gui.NewQMoveEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).MoveEventDefault(std_gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *MediaListPage) MoveEventDefault(event std_gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_MoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMoveEvent(event))
	}
}

//export callbackMediaListPage46aa22_PaintEvent
func callbackMediaListPage46aa22_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		(*(*func(*std_gui.QPaintEvent))(signal))(std_gui.NewQPaintEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).PaintEventDefault(std_gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *MediaListPage) PaintEventDefault(event std_gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_PaintEventDefault(ptr.Pointer(), std_gui.PointerFromQPaintEvent(event))
	}
}

//export callbackMediaListPage46aa22_Raise
func callbackMediaListPage46aa22_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListPageFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *MediaListPage) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_RaiseDefault(ptr.Pointer())
	}
}

//export callbackMediaListPage46aa22_Repaint
func callbackMediaListPage46aa22_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListPageFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *MediaListPage) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_RepaintDefault(ptr.Pointer())
	}
}

//export callbackMediaListPage46aa22_ResizeEvent
func callbackMediaListPage46aa22_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		(*(*func(*std_gui.QResizeEvent))(signal))(std_gui.NewQResizeEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).ResizeEventDefault(std_gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *MediaListPage) ResizeEventDefault(event std_gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_ResizeEventDefault(ptr.Pointer(), std_gui.PointerFromQResizeEvent(event))
	}
}

//export callbackMediaListPage46aa22_SetDisabled
func callbackMediaListPage46aa22_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		(*(*func(bool))(signal))(int8(disable) != 0)
	} else {
		NewMediaListPageFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *MediaListPage) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackMediaListPage46aa22_SetEnabled
func callbackMediaListPage46aa22_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewMediaListPageFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *MediaListPage) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackMediaListPage46aa22_SetFocus2
func callbackMediaListPage46aa22_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListPageFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *MediaListPage) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackMediaListPage46aa22_SetHidden
func callbackMediaListPage46aa22_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		(*(*func(bool))(signal))(int8(hidden) != 0)
	} else {
		NewMediaListPageFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *MediaListPage) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackMediaListPage46aa22_SetStyleSheet
func callbackMediaListPage46aa22_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(styleSheet))
	} else {
		NewMediaListPageFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *MediaListPage) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.MediaListPage46aa22_SetStyleSheetDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackMediaListPage46aa22_SetVisible
func callbackMediaListPage46aa22_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	} else {
		NewMediaListPageFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *MediaListPage) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackMediaListPage46aa22_SetWindowModified
func callbackMediaListPage46aa22_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewMediaListPageFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *MediaListPage) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackMediaListPage46aa22_SetWindowTitle
func callbackMediaListPage46aa22_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(vqs))
	} else {
		NewMediaListPageFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *MediaListPage) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.MediaListPage46aa22_SetWindowTitleDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackMediaListPage46aa22_Show
func callbackMediaListPage46aa22_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListPageFromPointer(ptr).ShowDefault()
	}
}

func (ptr *MediaListPage) ShowDefault() {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_ShowDefault(ptr.Pointer())
	}
}

//export callbackMediaListPage46aa22_ShowEvent
func callbackMediaListPage46aa22_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		(*(*func(*std_gui.QShowEvent))(signal))(std_gui.NewQShowEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).ShowEventDefault(std_gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *MediaListPage) ShowEventDefault(event std_gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_ShowEventDefault(ptr.Pointer(), std_gui.PointerFromQShowEvent(event))
	}
}

//export callbackMediaListPage46aa22_ShowFullScreen
func callbackMediaListPage46aa22_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListPageFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *MediaListPage) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackMediaListPage46aa22_ShowMaximized
func callbackMediaListPage46aa22_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListPageFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *MediaListPage) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackMediaListPage46aa22_ShowMinimized
func callbackMediaListPage46aa22_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListPageFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *MediaListPage) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackMediaListPage46aa22_ShowNormal
func callbackMediaListPage46aa22_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListPageFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *MediaListPage) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackMediaListPage46aa22_TabletEvent
func callbackMediaListPage46aa22_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		(*(*func(*std_gui.QTabletEvent))(signal))(std_gui.NewQTabletEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).TabletEventDefault(std_gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *MediaListPage) TabletEventDefault(event std_gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_TabletEventDefault(ptr.Pointer(), std_gui.PointerFromQTabletEvent(event))
	}
}

//export callbackMediaListPage46aa22_Update
func callbackMediaListPage46aa22_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListPageFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *MediaListPage) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_UpdateDefault(ptr.Pointer())
	}
}

//export callbackMediaListPage46aa22_UpdateMicroFocus
func callbackMediaListPage46aa22_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListPageFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *MediaListPage) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackMediaListPage46aa22_WheelEvent
func callbackMediaListPage46aa22_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		(*(*func(*std_gui.QWheelEvent))(signal))(std_gui.NewQWheelEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).WheelEventDefault(std_gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *MediaListPage) WheelEventDefault(event std_gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_WheelEventDefault(ptr.Pointer(), std_gui.PointerFromQWheelEvent(event))
	}
}

//export callbackMediaListPage46aa22_WindowIconChanged
func callbackMediaListPage46aa22_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		(*(*func(*std_gui.QIcon))(signal))(std_gui.NewQIconFromPointer(icon))
	}

}

//export callbackMediaListPage46aa22_WindowTitleChanged
func callbackMediaListPage46aa22_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(title))
	}

}

//export callbackMediaListPage46aa22_PaintEngine
func callbackMediaListPage46aa22_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return std_gui.PointerFromQPaintEngine((*(*func() *std_gui.QPaintEngine)(signal))())
	}

	return std_gui.PointerFromQPaintEngine(NewMediaListPageFromPointer(ptr).PaintEngineDefault())
}

func (ptr *MediaListPage) PaintEngineDefault() *std_gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return std_gui.NewQPaintEngineFromPointer(C.MediaListPage46aa22_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackMediaListPage46aa22_MinimumSizeHint
func callbackMediaListPage46aa22_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return std_core.PointerFromQSize((*(*func() *std_core.QSize)(signal))())
	}

	return std_core.PointerFromQSize(NewMediaListPageFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *MediaListPage) MinimumSizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.MediaListPage46aa22_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackMediaListPage46aa22_SizeHint
func callbackMediaListPage46aa22_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return std_core.PointerFromQSize((*(*func() *std_core.QSize)(signal))())
	}

	return std_core.PointerFromQSize(NewMediaListPageFromPointer(ptr).SizeHintDefault())
}

func (ptr *MediaListPage) SizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.MediaListPage46aa22_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackMediaListPage46aa22_InputMethodQuery
func callbackMediaListPage46aa22_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(std_core.Qt__InputMethodQuery) *std_core.QVariant)(signal))(std_core.Qt__InputMethodQuery(query)))
	}

	return std_core.PointerFromQVariant(NewMediaListPageFromPointer(ptr).InputMethodQueryDefault(std_core.Qt__InputMethodQuery(query)))
}

func (ptr *MediaListPage) InputMethodQueryDefault(query std_core.Qt__InputMethodQuery) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.MediaListPage46aa22_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackMediaListPage46aa22_HasHeightForWidth
func callbackMediaListPage46aa22_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListPageFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *MediaListPage) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.MediaListPage46aa22_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackMediaListPage46aa22_HeightForWidth
func callbackMediaListPage46aa22_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32((*(*func(int) int)(signal))(int(int32(w)))))
	}

	return C.int(int32(NewMediaListPageFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *MediaListPage) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.MediaListPage46aa22_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackMediaListPage46aa22_Metric
func callbackMediaListPage46aa22_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32((*(*func(std_gui.QPaintDevice__PaintDeviceMetric) int)(signal))(std_gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewMediaListPageFromPointer(ptr).MetricDefault(std_gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *MediaListPage) MetricDefault(m std_gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.MediaListPage46aa22_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackMediaListPage46aa22_InitPainter
func callbackMediaListPage46aa22_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		(*(*func(*std_gui.QPainter))(signal))(std_gui.NewQPainterFromPointer(painter))
	} else {
		NewMediaListPageFromPointer(ptr).InitPainterDefault(std_gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *MediaListPage) InitPainterDefault(painter std_gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_InitPainterDefault(ptr.Pointer(), std_gui.PointerFromQPainter(painter))
	}
}

//export callbackMediaListPage46aa22_EventFilter
func callbackMediaListPage46aa22_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMediaListPageFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *MediaListPage) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.MediaListPage46aa22_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackMediaListPage46aa22_ChildEvent
func callbackMediaListPage46aa22_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *MediaListPage) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackMediaListPage46aa22_ConnectNotify
func callbackMediaListPage46aa22_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewMediaListPageFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *MediaListPage) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackMediaListPage46aa22_CustomEvent
func callbackMediaListPage46aa22_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *MediaListPage) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackMediaListPage46aa22_DeleteLater
func callbackMediaListPage46aa22_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMediaListPageFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *MediaListPage) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackMediaListPage46aa22_Destroyed
func callbackMediaListPage46aa22_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackMediaListPage46aa22_DisconnectNotify
func callbackMediaListPage46aa22_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewMediaListPageFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *MediaListPage) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackMediaListPage46aa22_ObjectNameChanged
func callbackMediaListPage46aa22_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackMediaListPage46aa22_TimerEvent
func callbackMediaListPage46aa22_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewMediaListPageFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *MediaListPage) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MediaListPage46aa22_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
