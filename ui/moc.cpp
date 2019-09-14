

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QAbstractItemModel>
#include <QAbstractTableModel>
#include <QAction>
#include <QActionEvent>
#include <QByteArray>
#include <QChildEvent>
#include <QCloseEvent>
#include <QContextMenuEvent>
#include <QDialog>
#include <QDockWidget>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEvent>
#include <QFocusEvent>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QHash>
#include <QHideEvent>
#include <QIcon>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QLayout>
#include <QMainWindow>
#include <QMap>
#include <QMenu>
#include <QMenuBar>
#include <QMetaMethod>
#include <QMimeData>
#include <QModelIndex>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPaintEngine>
#include <QPaintEvent>
#include <QPainter>
#include <QPdfWriter>
#include <QPersistentModelIndex>
#include <QPoint>
#include <QResizeEvent>
#include <QSettings>
#include <QShowEvent>
#include <QSize>
#include <QString>
#include <QTabWidget>
#include <QTabletEvent>
#include <QTimerEvent>
#include <QVariant>
#include <QVector>
#include <QWheelEvent>
#include <QWidget>
#include <QWindow>

#ifdef QT_QML_LIB
	#include <QQmlEngine>
#endif

class Settingsc4adcb: public QSettings{
public:
	Settingsc4adcb(QObject *parent) : QSettings(parent) {};
	Settingsc4adcb(QSettings::Format format, QSettings::Scope scope, const QString &organization, const QString &application, QObject *parent) : QSettings(format, scope, organization, application, parent) {};
	Settingsc4adcb(QSettings::Scope scope, const QString &organization, const QString &application, QObject *parent) : QSettings(scope, organization, application, parent) {};
	Settingsc4adcb(const QString &fileName, QSettings::Format format, QObject *parent) : QSettings(fileName, format, parent) {};
	Settingsc4adcb(const QString &organization, const QString &application, QObject *parent) : QSettings(organization, application, parent) {};

};

class MediaListTable46aa22: public QAbstractTableModel
{
Q_OBJECT
public:
	MediaListTable46aa22(QObject *parent = Q_NULLPTR) : QAbstractTableModel(parent) {qRegisterMetaType<quintptr>("quintptr");MediaListTable46aa22_MediaListTable46aa22_QRegisterMetaType();MediaListTable46aa22_MediaListTable46aa22_QRegisterMetaTypes();callbackMediaListTable46aa22_Constructor(this);};
	 ~MediaListTable46aa22() { callbackMediaListTable46aa22_DestroyMediaListTable(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackMediaListTable46aa22_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackMediaListTable46aa22_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackMediaListTable46aa22_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackMediaListTable46aa22_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackMediaListTable46aa22_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackMediaListTable46aa22_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackMediaListTable46aa22_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackMediaListTable46aa22_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackMediaListTable46aa22_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackMediaListTable46aa22_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackMediaListTable46aa22_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackMediaListTable46aa22_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackMediaListTable46aa22_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	bool submit() { return callbackMediaListTable46aa22_Submit(this) != 0; };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackMediaListTable46aa22_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackMediaListTable46aa22_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackMediaListTable46aa22_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackMediaListTable46aa22_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackMediaListTable46aa22_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackMediaListTable46aa22_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackMediaListTable46aa22_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = new QVector<int>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackMediaListTable46aa22_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackMediaListTable46aa22_HeaderDataChanged(this, orientation, first, last); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackMediaListTable46aa22_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackMediaListTable46aa22_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_ModelAboutToBeReset() { callbackMediaListTable46aa22_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackMediaListTable46aa22_ModelReset(this); };
	void resetInternalData() { callbackMediaListTable46aa22_ResetInternalData(this); };
	void revert() { callbackMediaListTable46aa22_Revert(this); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackMediaListTable46aa22_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackMediaListTable46aa22_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackMediaListTable46aa22_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackMediaListTable46aa22_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackMediaListTable46aa22_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackMediaListTable46aa22_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void sort(int column, Qt::SortOrder order) { callbackMediaListTable46aa22_Sort(this, column, order); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackMediaListTable46aa22_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackMediaListTable46aa22_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackMediaListTable46aa22_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackMediaListTable46aa22_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackMediaListTable46aa22_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackMediaListTable46aa22_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackMediaListTable46aa22_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackMediaListTable46aa22_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackMediaListTable46aa22_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackMediaListTable46aa22_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackMediaListTable46aa22_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackMediaListTable46aa22_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackMediaListTable46aa22_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackMediaListTable46aa22_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	bool hasChildren(const QModelIndex & parent) const { return callbackMediaListTable46aa22_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackMediaListTable46aa22_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	int rowCount(const QModelIndex & parent) const { return callbackMediaListTable46aa22_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	bool event(QEvent * e) { return callbackMediaListTable46aa22_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackMediaListTable46aa22_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackMediaListTable46aa22_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackMediaListTable46aa22_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackMediaListTable46aa22_CustomEvent(this, event); };
	void deleteLater() { callbackMediaListTable46aa22_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackMediaListTable46aa22_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackMediaListTable46aa22_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackMediaListTable46aa22_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackMediaListTable46aa22_TimerEvent(this, event); };
signals:
public slots:
private:
};

Q_DECLARE_METATYPE(MediaListTable46aa22*)


void MediaListTable46aa22_MediaListTable46aa22_QRegisterMetaTypes() {
}

class MediaListTabs46aa22: public QTabWidget
{
Q_OBJECT
public:
	MediaListTabs46aa22(QWidget *parent = Q_NULLPTR) : QTabWidget(parent) {qRegisterMetaType<quintptr>("quintptr");MediaListTabs46aa22_MediaListTabs46aa22_QRegisterMetaType();MediaListTabs46aa22_MediaListTabs46aa22_QRegisterMetaTypes();callbackMediaListTabs46aa22_Constructor(this);};
	 ~MediaListTabs46aa22() { callbackMediaListTabs46aa22_DestroyMediaListTabs(this); };
	bool event(QEvent * ev) { return callbackMediaListTabs46aa22_Event(this, ev) != 0; };
	void changeEvent(QEvent * ev) { callbackMediaListTabs46aa22_ChangeEvent(this, ev); };
	void Signal_CurrentChanged(int index) { callbackMediaListTabs46aa22_CurrentChanged(this, index); };
	void keyPressEvent(QKeyEvent * e) { callbackMediaListTabs46aa22_KeyPressEvent(this, e); };
	void paintEvent(QPaintEvent * event) { callbackMediaListTabs46aa22_PaintEvent(this, event); };
	void resizeEvent(QResizeEvent * e) { callbackMediaListTabs46aa22_ResizeEvent(this, e); };
	void setCurrentIndex(int index) { callbackMediaListTabs46aa22_SetCurrentIndex(this, index); };
	void setCurrentWidget(QWidget * widget) { callbackMediaListTabs46aa22_SetCurrentWidget(this, widget); };
	void showEvent(QShowEvent * vqs) { callbackMediaListTabs46aa22_ShowEvent(this, vqs); };
	void Signal_TabBarClicked(int index) { callbackMediaListTabs46aa22_TabBarClicked(this, index); };
	void Signal_TabBarDoubleClicked(int index) { callbackMediaListTabs46aa22_TabBarDoubleClicked(this, index); };
	void Signal_TabCloseRequested(int index) { callbackMediaListTabs46aa22_TabCloseRequested(this, index); };
	void tabInserted(int index) { callbackMediaListTabs46aa22_TabInserted(this, index); };
	void tabRemoved(int index) { callbackMediaListTabs46aa22_TabRemoved(this, index); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackMediaListTabs46aa22_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackMediaListTabs46aa22_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	bool hasHeightForWidth() const { return callbackMediaListTabs46aa22_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int width) const { return callbackMediaListTabs46aa22_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), width); };
	bool close() { return callbackMediaListTabs46aa22_Close(this) != 0; };
	bool focusNextPrevChild(bool next) { return callbackMediaListTabs46aa22_FocusNextPrevChild(this, next) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackMediaListTabs46aa22_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, result) != 0; };
	void actionEvent(QActionEvent * event) { callbackMediaListTabs46aa22_ActionEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackMediaListTabs46aa22_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackMediaListTabs46aa22_ContextMenuEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackMediaListTabs46aa22_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackMediaListTabs46aa22_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackMediaListTabs46aa22_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackMediaListTabs46aa22_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackMediaListTabs46aa22_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackMediaListTabs46aa22_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackMediaListTabs46aa22_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackMediaListTabs46aa22_FocusOutEvent(this, event); };
	void hide() { callbackMediaListTabs46aa22_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackMediaListTabs46aa22_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackMediaListTabs46aa22_InputMethodEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackMediaListTabs46aa22_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackMediaListTabs46aa22_LeaveEvent(this, event); };
	void lower() { callbackMediaListTabs46aa22_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackMediaListTabs46aa22_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackMediaListTabs46aa22_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackMediaListTabs46aa22_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackMediaListTabs46aa22_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackMediaListTabs46aa22_MoveEvent(this, event); };
	void raise() { callbackMediaListTabs46aa22_Raise(this); };
	void repaint() { callbackMediaListTabs46aa22_Repaint(this); };
	void setDisabled(bool disable) { callbackMediaListTabs46aa22_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackMediaListTabs46aa22_SetEnabled(this, vbo); };
	void setFocus() { callbackMediaListTabs46aa22_SetFocus2(this); };
	void setHidden(bool hidden) { callbackMediaListTabs46aa22_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); Moc_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackMediaListTabs46aa22_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackMediaListTabs46aa22_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackMediaListTabs46aa22_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); Moc_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackMediaListTabs46aa22_SetWindowTitle(this, vqsPacked); };
	void show() { callbackMediaListTabs46aa22_Show(this); };
	void showFullScreen() { callbackMediaListTabs46aa22_ShowFullScreen(this); };
	void showMaximized() { callbackMediaListTabs46aa22_ShowMaximized(this); };
	void showMinimized() { callbackMediaListTabs46aa22_ShowMinimized(this); };
	void showNormal() { callbackMediaListTabs46aa22_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackMediaListTabs46aa22_TabletEvent(this, event); };
	void update() { callbackMediaListTabs46aa22_Update(this); };
	void updateMicroFocus() { callbackMediaListTabs46aa22_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackMediaListTabs46aa22_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackMediaListTabs46aa22_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); Moc_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackMediaListTabs46aa22_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackMediaListTabs46aa22_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackMediaListTabs46aa22_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackMediaListTabs46aa22_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void initPainter(QPainter * painter) const { callbackMediaListTabs46aa22_InitPainter(const_cast<void*>(static_cast<const void*>(this)), painter); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackMediaListTabs46aa22_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackMediaListTabs46aa22_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackMediaListTabs46aa22_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackMediaListTabs46aa22_CustomEvent(this, event); };
	void deleteLater() { callbackMediaListTabs46aa22_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackMediaListTabs46aa22_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackMediaListTabs46aa22_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackMediaListTabs46aa22_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackMediaListTabs46aa22_TimerEvent(this, event); };
signals:
public slots:
private:
};

Q_DECLARE_METATYPE(MediaListTabs46aa22*)


void MediaListTabs46aa22_MediaListTabs46aa22_QRegisterMetaTypes() {
}

class MenuBar46aa22: public QMenuBar
{
Q_OBJECT
public:
	MenuBar46aa22(QWidget *parent = Q_NULLPTR) : QMenuBar(parent) {qRegisterMetaType<quintptr>("quintptr");MenuBar46aa22_MenuBar46aa22_QRegisterMetaType();MenuBar46aa22_MenuBar46aa22_QRegisterMetaTypes();callbackMenuBar46aa22_Constructor(this);};
	 ~MenuBar46aa22() { callbackMenuBar46aa22_DestroyMenuBar(this); };
	bool event(QEvent * e) { return callbackMenuBar46aa22_Event(this, e) != 0; };
	bool eventFilter(QObject * object, QEvent * event) { return callbackMenuBar46aa22_EventFilter(this, object, event) != 0; };
	void actionEvent(QActionEvent * e) { callbackMenuBar46aa22_ActionEvent(this, e); };
	void changeEvent(QEvent * e) { callbackMenuBar46aa22_ChangeEvent(this, e); };
	void focusInEvent(QFocusEvent * vqf) { callbackMenuBar46aa22_FocusInEvent(this, vqf); };
	void focusOutEvent(QFocusEvent * vqf) { callbackMenuBar46aa22_FocusOutEvent(this, vqf); };
	void Signal_Hovered(QAction * action) { callbackMenuBar46aa22_Hovered(this, action); };
	void keyPressEvent(QKeyEvent * e) { callbackMenuBar46aa22_KeyPressEvent(this, e); };
	void leaveEvent(QEvent * vqe) { callbackMenuBar46aa22_LeaveEvent(this, vqe); };
	void mouseMoveEvent(QMouseEvent * e) { callbackMenuBar46aa22_MouseMoveEvent(this, e); };
	void mousePressEvent(QMouseEvent * e) { callbackMenuBar46aa22_MousePressEvent(this, e); };
	void mouseReleaseEvent(QMouseEvent * e) { callbackMenuBar46aa22_MouseReleaseEvent(this, e); };
	void paintEvent(QPaintEvent * e) { callbackMenuBar46aa22_PaintEvent(this, e); };
	void resizeEvent(QResizeEvent * vqr) { callbackMenuBar46aa22_ResizeEvent(this, vqr); };
	void setVisible(bool visible) { callbackMenuBar46aa22_SetVisible(this, visible); };
	void timerEvent(QTimerEvent * e) { callbackMenuBar46aa22_TimerEvent(this, e); };
	void Signal_Triggered(QAction * action) { callbackMenuBar46aa22_Triggered(this, action); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackMenuBar46aa22_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackMenuBar46aa22_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	int heightForWidth(int vin) const { return callbackMenuBar46aa22_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), vin); };
	bool close() { return callbackMenuBar46aa22_Close(this) != 0; };
	bool focusNextPrevChild(bool next) { return callbackMenuBar46aa22_FocusNextPrevChild(this, next) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackMenuBar46aa22_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, result) != 0; };
	void closeEvent(QCloseEvent * event) { callbackMenuBar46aa22_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackMenuBar46aa22_ContextMenuEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackMenuBar46aa22_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackMenuBar46aa22_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackMenuBar46aa22_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackMenuBar46aa22_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackMenuBar46aa22_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackMenuBar46aa22_EnterEvent(this, event); };
	void hide() { callbackMenuBar46aa22_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackMenuBar46aa22_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackMenuBar46aa22_InputMethodEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackMenuBar46aa22_KeyReleaseEvent(this, event); };
	void lower() { callbackMenuBar46aa22_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackMenuBar46aa22_MouseDoubleClickEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackMenuBar46aa22_MoveEvent(this, event); };
	void raise() { callbackMenuBar46aa22_Raise(this); };
	void repaint() { callbackMenuBar46aa22_Repaint(this); };
	void setDisabled(bool disable) { callbackMenuBar46aa22_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackMenuBar46aa22_SetEnabled(this, vbo); };
	void setFocus() { callbackMenuBar46aa22_SetFocus2(this); };
	void setHidden(bool hidden) { callbackMenuBar46aa22_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); Moc_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackMenuBar46aa22_SetStyleSheet(this, styleSheetPacked); };
	void setWindowModified(bool vbo) { callbackMenuBar46aa22_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); Moc_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackMenuBar46aa22_SetWindowTitle(this, vqsPacked); };
	void show() { callbackMenuBar46aa22_Show(this); };
	void showEvent(QShowEvent * event) { callbackMenuBar46aa22_ShowEvent(this, event); };
	void showFullScreen() { callbackMenuBar46aa22_ShowFullScreen(this); };
	void showMaximized() { callbackMenuBar46aa22_ShowMaximized(this); };
	void showMinimized() { callbackMenuBar46aa22_ShowMinimized(this); };
	void showNormal() { callbackMenuBar46aa22_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackMenuBar46aa22_TabletEvent(this, event); };
	void update() { callbackMenuBar46aa22_Update(this); };
	void updateMicroFocus() { callbackMenuBar46aa22_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackMenuBar46aa22_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackMenuBar46aa22_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); Moc_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackMenuBar46aa22_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackMenuBar46aa22_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackMenuBar46aa22_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackMenuBar46aa22_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackMenuBar46aa22_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void initPainter(QPainter * painter) const { callbackMenuBar46aa22_InitPainter(const_cast<void*>(static_cast<const void*>(this)), painter); };
	void childEvent(QChildEvent * event) { callbackMenuBar46aa22_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackMenuBar46aa22_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackMenuBar46aa22_CustomEvent(this, event); };
	void deleteLater() { callbackMenuBar46aa22_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackMenuBar46aa22_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackMenuBar46aa22_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackMenuBar46aa22_ObjectNameChanged(this, objectNamePacked); };
signals:
public slots:
private:
};

Q_DECLARE_METATYPE(MenuBar46aa22*)


void MenuBar46aa22_MenuBar46aa22_QRegisterMetaTypes() {
}

class Navigation46aa22: public QWidget
{
Q_OBJECT
public:
	Navigation46aa22(QWidget *parent = Q_NULLPTR, Qt::WindowFlags ff = Qt::WindowFlags()) : QWidget(parent, ff) {qRegisterMetaType<quintptr>("quintptr");Navigation46aa22_Navigation46aa22_QRegisterMetaType();Navigation46aa22_Navigation46aa22_QRegisterMetaTypes();callbackNavigation46aa22_Constructor(this);};
	 ~Navigation46aa22() { callbackNavigation46aa22_DestroyNavigation(this); };
	bool close() { return callbackNavigation46aa22_Close(this) != 0; };
	bool event(QEvent * event) { return callbackNavigation46aa22_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackNavigation46aa22_FocusNextPrevChild(this, next) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackNavigation46aa22_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, result) != 0; };
	void actionEvent(QActionEvent * event) { callbackNavigation46aa22_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackNavigation46aa22_ChangeEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackNavigation46aa22_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackNavigation46aa22_ContextMenuEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackNavigation46aa22_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackNavigation46aa22_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackNavigation46aa22_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackNavigation46aa22_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackNavigation46aa22_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackNavigation46aa22_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackNavigation46aa22_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackNavigation46aa22_FocusOutEvent(this, event); };
	void hide() { callbackNavigation46aa22_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackNavigation46aa22_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackNavigation46aa22_InputMethodEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackNavigation46aa22_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackNavigation46aa22_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackNavigation46aa22_LeaveEvent(this, event); };
	void lower() { callbackNavigation46aa22_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackNavigation46aa22_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackNavigation46aa22_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackNavigation46aa22_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackNavigation46aa22_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackNavigation46aa22_MoveEvent(this, event); };
	void paintEvent(QPaintEvent * event) { callbackNavigation46aa22_PaintEvent(this, event); };
	void raise() { callbackNavigation46aa22_Raise(this); };
	void repaint() { callbackNavigation46aa22_Repaint(this); };
	void resizeEvent(QResizeEvent * event) { callbackNavigation46aa22_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackNavigation46aa22_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackNavigation46aa22_SetEnabled(this, vbo); };
	void setFocus() { callbackNavigation46aa22_SetFocus2(this); };
	void setHidden(bool hidden) { callbackNavigation46aa22_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); Moc_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackNavigation46aa22_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackNavigation46aa22_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackNavigation46aa22_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); Moc_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackNavigation46aa22_SetWindowTitle(this, vqsPacked); };
	void show() { callbackNavigation46aa22_Show(this); };
	void showEvent(QShowEvent * event) { callbackNavigation46aa22_ShowEvent(this, event); };
	void showFullScreen() { callbackNavigation46aa22_ShowFullScreen(this); };
	void showMaximized() { callbackNavigation46aa22_ShowMaximized(this); };
	void showMinimized() { callbackNavigation46aa22_ShowMinimized(this); };
	void showNormal() { callbackNavigation46aa22_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackNavigation46aa22_TabletEvent(this, event); };
	void update() { callbackNavigation46aa22_Update(this); };
	void updateMicroFocus() { callbackNavigation46aa22_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackNavigation46aa22_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackNavigation46aa22_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); Moc_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackNavigation46aa22_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackNavigation46aa22_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackNavigation46aa22_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackNavigation46aa22_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackNavigation46aa22_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackNavigation46aa22_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackNavigation46aa22_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackNavigation46aa22_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void initPainter(QPainter * painter) const { callbackNavigation46aa22_InitPainter(const_cast<void*>(static_cast<const void*>(this)), painter); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackNavigation46aa22_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackNavigation46aa22_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackNavigation46aa22_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackNavigation46aa22_CustomEvent(this, event); };
	void deleteLater() { callbackNavigation46aa22_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackNavigation46aa22_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackNavigation46aa22_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackNavigation46aa22_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackNavigation46aa22_TimerEvent(this, event); };
signals:
public slots:
private:
};

Q_DECLARE_METATYPE(Navigation46aa22*)


void Navigation46aa22_Navigation46aa22_QRegisterMetaTypes() {
}

class SettingsDialog46aa22: public QDialog
{
Q_OBJECT
public:
	SettingsDialog46aa22(QWidget *parent = Q_NULLPTR, Qt::WindowFlags ff = Qt::WindowFlags()) : QDialog(parent, ff) {qRegisterMetaType<quintptr>("quintptr");SettingsDialog46aa22_SettingsDialog46aa22_QRegisterMetaType();SettingsDialog46aa22_SettingsDialog46aa22_QRegisterMetaTypes();callbackSettingsDialog46aa22_Constructor(this);};
	 ~SettingsDialog46aa22() { callbackSettingsDialog46aa22_DestroySettingsDialog(this); };
	bool eventFilter(QObject * o, QEvent * e) { return callbackSettingsDialog46aa22_EventFilter(this, o, e) != 0; };
	int exec() { return callbackSettingsDialog46aa22_Exec(this); };
	void accept() { callbackSettingsDialog46aa22_Accept(this); };
	void Signal_Accepted() { callbackSettingsDialog46aa22_Accepted(this); };
	void closeEvent(QCloseEvent * e) { callbackSettingsDialog46aa22_CloseEvent(this, e); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackSettingsDialog46aa22_ContextMenuEvent(this, e); };
	void done(int r) { callbackSettingsDialog46aa22_Done(this, r); };
	void Signal_Finished(int result) { callbackSettingsDialog46aa22_Finished(this, result); };
	void keyPressEvent(QKeyEvent * e) { callbackSettingsDialog46aa22_KeyPressEvent(this, e); };
	void open() { callbackSettingsDialog46aa22_Open(this); };
	void reject() { callbackSettingsDialog46aa22_Reject(this); };
	void Signal_Rejected() { callbackSettingsDialog46aa22_Rejected(this); };
	void resizeEvent(QResizeEvent * vqr) { callbackSettingsDialog46aa22_ResizeEvent(this, vqr); };
	void setVisible(bool visible) { callbackSettingsDialog46aa22_SetVisible(this, visible); };
	void showEvent(QShowEvent * event) { callbackSettingsDialog46aa22_ShowEvent(this, event); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackSettingsDialog46aa22_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackSettingsDialog46aa22_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	bool close() { return callbackSettingsDialog46aa22_Close(this) != 0; };
	bool event(QEvent * event) { return callbackSettingsDialog46aa22_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackSettingsDialog46aa22_FocusNextPrevChild(this, next) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackSettingsDialog46aa22_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, result) != 0; };
	void actionEvent(QActionEvent * event) { callbackSettingsDialog46aa22_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackSettingsDialog46aa22_ChangeEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackSettingsDialog46aa22_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackSettingsDialog46aa22_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackSettingsDialog46aa22_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackSettingsDialog46aa22_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackSettingsDialog46aa22_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackSettingsDialog46aa22_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackSettingsDialog46aa22_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackSettingsDialog46aa22_FocusOutEvent(this, event); };
	void hide() { callbackSettingsDialog46aa22_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackSettingsDialog46aa22_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackSettingsDialog46aa22_InputMethodEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackSettingsDialog46aa22_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackSettingsDialog46aa22_LeaveEvent(this, event); };
	void lower() { callbackSettingsDialog46aa22_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackSettingsDialog46aa22_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackSettingsDialog46aa22_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackSettingsDialog46aa22_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackSettingsDialog46aa22_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackSettingsDialog46aa22_MoveEvent(this, event); };
	void paintEvent(QPaintEvent * event) { callbackSettingsDialog46aa22_PaintEvent(this, event); };
	void raise() { callbackSettingsDialog46aa22_Raise(this); };
	void repaint() { callbackSettingsDialog46aa22_Repaint(this); };
	void setDisabled(bool disable) { callbackSettingsDialog46aa22_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackSettingsDialog46aa22_SetEnabled(this, vbo); };
	void setFocus() { callbackSettingsDialog46aa22_SetFocus2(this); };
	void setHidden(bool hidden) { callbackSettingsDialog46aa22_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); Moc_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackSettingsDialog46aa22_SetStyleSheet(this, styleSheetPacked); };
	void setWindowModified(bool vbo) { callbackSettingsDialog46aa22_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); Moc_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackSettingsDialog46aa22_SetWindowTitle(this, vqsPacked); };
	void show() { callbackSettingsDialog46aa22_Show(this); };
	void showFullScreen() { callbackSettingsDialog46aa22_ShowFullScreen(this); };
	void showMaximized() { callbackSettingsDialog46aa22_ShowMaximized(this); };
	void showMinimized() { callbackSettingsDialog46aa22_ShowMinimized(this); };
	void showNormal() { callbackSettingsDialog46aa22_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackSettingsDialog46aa22_TabletEvent(this, event); };
	void update() { callbackSettingsDialog46aa22_Update(this); };
	void updateMicroFocus() { callbackSettingsDialog46aa22_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackSettingsDialog46aa22_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackSettingsDialog46aa22_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); Moc_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackSettingsDialog46aa22_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackSettingsDialog46aa22_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackSettingsDialog46aa22_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackSettingsDialog46aa22_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackSettingsDialog46aa22_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackSettingsDialog46aa22_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void initPainter(QPainter * painter) const { callbackSettingsDialog46aa22_InitPainter(const_cast<void*>(static_cast<const void*>(this)), painter); };
	void childEvent(QChildEvent * event) { callbackSettingsDialog46aa22_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackSettingsDialog46aa22_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackSettingsDialog46aa22_CustomEvent(this, event); };
	void deleteLater() { callbackSettingsDialog46aa22_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackSettingsDialog46aa22_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackSettingsDialog46aa22_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackSettingsDialog46aa22_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackSettingsDialog46aa22_TimerEvent(this, event); };
signals:
public slots:
private:
};

Q_DECLARE_METATYPE(SettingsDialog46aa22*)


void SettingsDialog46aa22_SettingsDialog46aa22_QRegisterMetaTypes() {
}

class CurrentMediaPage46aa22: public QTabWidget
{
Q_OBJECT
public:
	CurrentMediaPage46aa22(QWidget *parent = Q_NULLPTR) : QTabWidget(parent) {qRegisterMetaType<quintptr>("quintptr");CurrentMediaPage46aa22_CurrentMediaPage46aa22_QRegisterMetaType();CurrentMediaPage46aa22_CurrentMediaPage46aa22_QRegisterMetaTypes();callbackCurrentMediaPage46aa22_Constructor(this);};
	 ~CurrentMediaPage46aa22() { callbackCurrentMediaPage46aa22_DestroyCurrentMediaPage(this); };
	bool event(QEvent * ev) { return callbackCurrentMediaPage46aa22_Event(this, ev) != 0; };
	void changeEvent(QEvent * ev) { callbackCurrentMediaPage46aa22_ChangeEvent(this, ev); };
	void Signal_CurrentChanged(int index) { callbackCurrentMediaPage46aa22_CurrentChanged(this, index); };
	void keyPressEvent(QKeyEvent * e) { callbackCurrentMediaPage46aa22_KeyPressEvent(this, e); };
	void paintEvent(QPaintEvent * event) { callbackCurrentMediaPage46aa22_PaintEvent(this, event); };
	void resizeEvent(QResizeEvent * e) { callbackCurrentMediaPage46aa22_ResizeEvent(this, e); };
	void setCurrentIndex(int index) { callbackCurrentMediaPage46aa22_SetCurrentIndex(this, index); };
	void setCurrentWidget(QWidget * widget) { callbackCurrentMediaPage46aa22_SetCurrentWidget(this, widget); };
	void showEvent(QShowEvent * vqs) { callbackCurrentMediaPage46aa22_ShowEvent(this, vqs); };
	void Signal_TabBarClicked(int index) { callbackCurrentMediaPage46aa22_TabBarClicked(this, index); };
	void Signal_TabBarDoubleClicked(int index) { callbackCurrentMediaPage46aa22_TabBarDoubleClicked(this, index); };
	void Signal_TabCloseRequested(int index) { callbackCurrentMediaPage46aa22_TabCloseRequested(this, index); };
	void tabInserted(int index) { callbackCurrentMediaPage46aa22_TabInserted(this, index); };
	void tabRemoved(int index) { callbackCurrentMediaPage46aa22_TabRemoved(this, index); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackCurrentMediaPage46aa22_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackCurrentMediaPage46aa22_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	bool hasHeightForWidth() const { return callbackCurrentMediaPage46aa22_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int width) const { return callbackCurrentMediaPage46aa22_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), width); };
	bool close() { return callbackCurrentMediaPage46aa22_Close(this) != 0; };
	bool focusNextPrevChild(bool next) { return callbackCurrentMediaPage46aa22_FocusNextPrevChild(this, next) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackCurrentMediaPage46aa22_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, result) != 0; };
	void actionEvent(QActionEvent * event) { callbackCurrentMediaPage46aa22_ActionEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackCurrentMediaPage46aa22_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackCurrentMediaPage46aa22_ContextMenuEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackCurrentMediaPage46aa22_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackCurrentMediaPage46aa22_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackCurrentMediaPage46aa22_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackCurrentMediaPage46aa22_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackCurrentMediaPage46aa22_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackCurrentMediaPage46aa22_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackCurrentMediaPage46aa22_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackCurrentMediaPage46aa22_FocusOutEvent(this, event); };
	void hide() { callbackCurrentMediaPage46aa22_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackCurrentMediaPage46aa22_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackCurrentMediaPage46aa22_InputMethodEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackCurrentMediaPage46aa22_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackCurrentMediaPage46aa22_LeaveEvent(this, event); };
	void lower() { callbackCurrentMediaPage46aa22_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackCurrentMediaPage46aa22_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackCurrentMediaPage46aa22_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackCurrentMediaPage46aa22_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackCurrentMediaPage46aa22_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackCurrentMediaPage46aa22_MoveEvent(this, event); };
	void raise() { callbackCurrentMediaPage46aa22_Raise(this); };
	void repaint() { callbackCurrentMediaPage46aa22_Repaint(this); };
	void setDisabled(bool disable) { callbackCurrentMediaPage46aa22_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackCurrentMediaPage46aa22_SetEnabled(this, vbo); };
	void setFocus() { callbackCurrentMediaPage46aa22_SetFocus2(this); };
	void setHidden(bool hidden) { callbackCurrentMediaPage46aa22_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); Moc_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackCurrentMediaPage46aa22_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackCurrentMediaPage46aa22_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackCurrentMediaPage46aa22_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); Moc_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackCurrentMediaPage46aa22_SetWindowTitle(this, vqsPacked); };
	void show() { callbackCurrentMediaPage46aa22_Show(this); };
	void showFullScreen() { callbackCurrentMediaPage46aa22_ShowFullScreen(this); };
	void showMaximized() { callbackCurrentMediaPage46aa22_ShowMaximized(this); };
	void showMinimized() { callbackCurrentMediaPage46aa22_ShowMinimized(this); };
	void showNormal() { callbackCurrentMediaPage46aa22_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackCurrentMediaPage46aa22_TabletEvent(this, event); };
	void update() { callbackCurrentMediaPage46aa22_Update(this); };
	void updateMicroFocus() { callbackCurrentMediaPage46aa22_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackCurrentMediaPage46aa22_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackCurrentMediaPage46aa22_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); Moc_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackCurrentMediaPage46aa22_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackCurrentMediaPage46aa22_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackCurrentMediaPage46aa22_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackCurrentMediaPage46aa22_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void initPainter(QPainter * painter) const { callbackCurrentMediaPage46aa22_InitPainter(const_cast<void*>(static_cast<const void*>(this)), painter); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackCurrentMediaPage46aa22_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackCurrentMediaPage46aa22_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackCurrentMediaPage46aa22_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackCurrentMediaPage46aa22_CustomEvent(this, event); };
	void deleteLater() { callbackCurrentMediaPage46aa22_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackCurrentMediaPage46aa22_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackCurrentMediaPage46aa22_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackCurrentMediaPage46aa22_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackCurrentMediaPage46aa22_TimerEvent(this, event); };
signals:
public slots:
private:
};

Q_DECLARE_METATYPE(CurrentMediaPage46aa22*)


void CurrentMediaPage46aa22_CurrentMediaPage46aa22_QRegisterMetaTypes() {
}

class MainWindow46aa22: public QMainWindow
{
Q_OBJECT
public:
	MainWindow46aa22(QWidget *parent = Q_NULLPTR, Qt::WindowFlags flags = Qt::WindowFlags()) : QMainWindow(parent, flags) {qRegisterMetaType<quintptr>("quintptr");MainWindow46aa22_MainWindow46aa22_QRegisterMetaType();MainWindow46aa22_MainWindow46aa22_QRegisterMetaTypes();callbackMainWindow46aa22_Constructor(this);};
	 ~MainWindow46aa22() { callbackMainWindow46aa22_DestroyMainWindow(this); };
	QMenu * createPopupMenu() { return static_cast<QMenu*>(callbackMainWindow46aa22_CreatePopupMenu(this)); };
	bool event(QEvent * event) { return callbackMainWindow46aa22_Event(this, event) != 0; };
	void contextMenuEvent(QContextMenuEvent * event) { callbackMainWindow46aa22_ContextMenuEvent(this, event); };
	void Signal_IconSizeChanged(const QSize & iconSize) { callbackMainWindow46aa22_IconSizeChanged(this, const_cast<QSize*>(&iconSize)); };
	void setAnimated(bool enabled) { callbackMainWindow46aa22_SetAnimated(this, enabled); };
	void setDockNestingEnabled(bool enabled) { callbackMainWindow46aa22_SetDockNestingEnabled(this, enabled); };
	void setUnifiedTitleAndToolBarOnMac(bool set) { callbackMainWindow46aa22_SetUnifiedTitleAndToolBarOnMac(this, set); };
	void Signal_TabifiedDockWidgetActivated(QDockWidget * dockWidget) { callbackMainWindow46aa22_TabifiedDockWidgetActivated(this, dockWidget); };
	void Signal_ToolButtonStyleChanged(Qt::ToolButtonStyle toolButtonStyle) { callbackMainWindow46aa22_ToolButtonStyleChanged(this, toolButtonStyle); };
	bool close() { return callbackMainWindow46aa22_Close(this) != 0; };
	bool focusNextPrevChild(bool next) { return callbackMainWindow46aa22_FocusNextPrevChild(this, next) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackMainWindow46aa22_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, result) != 0; };
	void actionEvent(QActionEvent * event) { callbackMainWindow46aa22_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackMainWindow46aa22_ChangeEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackMainWindow46aa22_CloseEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackMainWindow46aa22_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackMainWindow46aa22_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackMainWindow46aa22_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackMainWindow46aa22_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackMainWindow46aa22_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackMainWindow46aa22_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackMainWindow46aa22_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackMainWindow46aa22_FocusOutEvent(this, event); };
	void hide() { callbackMainWindow46aa22_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackMainWindow46aa22_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackMainWindow46aa22_InputMethodEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackMainWindow46aa22_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackMainWindow46aa22_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackMainWindow46aa22_LeaveEvent(this, event); };
	void lower() { callbackMainWindow46aa22_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackMainWindow46aa22_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackMainWindow46aa22_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackMainWindow46aa22_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackMainWindow46aa22_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackMainWindow46aa22_MoveEvent(this, event); };
	void paintEvent(QPaintEvent * event) { callbackMainWindow46aa22_PaintEvent(this, event); };
	void raise() { callbackMainWindow46aa22_Raise(this); };
	void repaint() { callbackMainWindow46aa22_Repaint(this); };
	void resizeEvent(QResizeEvent * event) { callbackMainWindow46aa22_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackMainWindow46aa22_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackMainWindow46aa22_SetEnabled(this, vbo); };
	void setFocus() { callbackMainWindow46aa22_SetFocus2(this); };
	void setHidden(bool hidden) { callbackMainWindow46aa22_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); Moc_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackMainWindow46aa22_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackMainWindow46aa22_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackMainWindow46aa22_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); Moc_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackMainWindow46aa22_SetWindowTitle(this, vqsPacked); };
	void show() { callbackMainWindow46aa22_Show(this); };
	void showEvent(QShowEvent * event) { callbackMainWindow46aa22_ShowEvent(this, event); };
	void showFullScreen() { callbackMainWindow46aa22_ShowFullScreen(this); };
	void showMaximized() { callbackMainWindow46aa22_ShowMaximized(this); };
	void showMinimized() { callbackMainWindow46aa22_ShowMinimized(this); };
	void showNormal() { callbackMainWindow46aa22_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackMainWindow46aa22_TabletEvent(this, event); };
	void update() { callbackMainWindow46aa22_Update(this); };
	void updateMicroFocus() { callbackMainWindow46aa22_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackMainWindow46aa22_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackMainWindow46aa22_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); Moc_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackMainWindow46aa22_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackMainWindow46aa22_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackMainWindow46aa22_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackMainWindow46aa22_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackMainWindow46aa22_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackMainWindow46aa22_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackMainWindow46aa22_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackMainWindow46aa22_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void initPainter(QPainter * painter) const { callbackMainWindow46aa22_InitPainter(const_cast<void*>(static_cast<const void*>(this)), painter); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackMainWindow46aa22_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackMainWindow46aa22_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackMainWindow46aa22_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackMainWindow46aa22_CustomEvent(this, event); };
	void deleteLater() { callbackMainWindow46aa22_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackMainWindow46aa22_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackMainWindow46aa22_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackMainWindow46aa22_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackMainWindow46aa22_TimerEvent(this, event); };
signals:
public slots:
private:
};

Q_DECLARE_METATYPE(MainWindow46aa22*)


void MainWindow46aa22_MainWindow46aa22_QRegisterMetaTypes() {
}

class MediaListPage46aa22: public QWidget
{
Q_OBJECT
public:
	MediaListPage46aa22(QWidget *parent = Q_NULLPTR, Qt::WindowFlags ff = Qt::WindowFlags()) : QWidget(parent, ff) {qRegisterMetaType<quintptr>("quintptr");MediaListPage46aa22_MediaListPage46aa22_QRegisterMetaType();MediaListPage46aa22_MediaListPage46aa22_QRegisterMetaTypes();callbackMediaListPage46aa22_Constructor(this);};
	 ~MediaListPage46aa22() { callbackMediaListPage46aa22_DestroyMediaListPage(this); };
	bool close() { return callbackMediaListPage46aa22_Close(this) != 0; };
	bool event(QEvent * event) { return callbackMediaListPage46aa22_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackMediaListPage46aa22_FocusNextPrevChild(this, next) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackMediaListPage46aa22_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, result) != 0; };
	void actionEvent(QActionEvent * event) { callbackMediaListPage46aa22_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackMediaListPage46aa22_ChangeEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackMediaListPage46aa22_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackMediaListPage46aa22_ContextMenuEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackMediaListPage46aa22_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackMediaListPage46aa22_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackMediaListPage46aa22_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackMediaListPage46aa22_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackMediaListPage46aa22_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackMediaListPage46aa22_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackMediaListPage46aa22_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackMediaListPage46aa22_FocusOutEvent(this, event); };
	void hide() { callbackMediaListPage46aa22_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackMediaListPage46aa22_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackMediaListPage46aa22_InputMethodEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackMediaListPage46aa22_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackMediaListPage46aa22_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackMediaListPage46aa22_LeaveEvent(this, event); };
	void lower() { callbackMediaListPage46aa22_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackMediaListPage46aa22_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackMediaListPage46aa22_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackMediaListPage46aa22_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackMediaListPage46aa22_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackMediaListPage46aa22_MoveEvent(this, event); };
	void paintEvent(QPaintEvent * event) { callbackMediaListPage46aa22_PaintEvent(this, event); };
	void raise() { callbackMediaListPage46aa22_Raise(this); };
	void repaint() { callbackMediaListPage46aa22_Repaint(this); };
	void resizeEvent(QResizeEvent * event) { callbackMediaListPage46aa22_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackMediaListPage46aa22_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackMediaListPage46aa22_SetEnabled(this, vbo); };
	void setFocus() { callbackMediaListPage46aa22_SetFocus2(this); };
	void setHidden(bool hidden) { callbackMediaListPage46aa22_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); Moc_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackMediaListPage46aa22_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackMediaListPage46aa22_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackMediaListPage46aa22_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); Moc_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackMediaListPage46aa22_SetWindowTitle(this, vqsPacked); };
	void show() { callbackMediaListPage46aa22_Show(this); };
	void showEvent(QShowEvent * event) { callbackMediaListPage46aa22_ShowEvent(this, event); };
	void showFullScreen() { callbackMediaListPage46aa22_ShowFullScreen(this); };
	void showMaximized() { callbackMediaListPage46aa22_ShowMaximized(this); };
	void showMinimized() { callbackMediaListPage46aa22_ShowMinimized(this); };
	void showNormal() { callbackMediaListPage46aa22_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackMediaListPage46aa22_TabletEvent(this, event); };
	void update() { callbackMediaListPage46aa22_Update(this); };
	void updateMicroFocus() { callbackMediaListPage46aa22_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackMediaListPage46aa22_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackMediaListPage46aa22_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); Moc_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackMediaListPage46aa22_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackMediaListPage46aa22_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackMediaListPage46aa22_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackMediaListPage46aa22_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackMediaListPage46aa22_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackMediaListPage46aa22_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackMediaListPage46aa22_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackMediaListPage46aa22_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void initPainter(QPainter * painter) const { callbackMediaListPage46aa22_InitPainter(const_cast<void*>(static_cast<const void*>(this)), painter); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackMediaListPage46aa22_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackMediaListPage46aa22_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackMediaListPage46aa22_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackMediaListPage46aa22_CustomEvent(this, event); };
	void deleteLater() { callbackMediaListPage46aa22_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackMediaListPage46aa22_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackMediaListPage46aa22_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackMediaListPage46aa22_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackMediaListPage46aa22_TimerEvent(this, event); };
signals:
public slots:
private:
};

Q_DECLARE_METATYPE(MediaListPage46aa22*)


void MediaListPage46aa22_MediaListPage46aa22_QRegisterMetaTypes() {
}

int MainWindow46aa22_MainWindow46aa22_QRegisterMetaType()
{
	return qRegisterMetaType<MainWindow46aa22*>();
}

int MainWindow46aa22_MainWindow46aa22_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<MainWindow46aa22*>(const_cast<const char*>(typeName));
}

int MainWindow46aa22_MainWindow46aa22_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<MainWindow46aa22>();
#else
	return 0;
#endif
}

int MainWindow46aa22_MainWindow46aa22_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<MainWindow46aa22>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* MainWindow46aa22___resizeDocks_docks_atList(void* ptr, int i)
{
	return ({QDockWidget * tmp = static_cast<QList<QDockWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QDockWidget *>*>(ptr)->size()-1) { static_cast<QList<QDockWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MainWindow46aa22___resizeDocks_docks_setList(void* ptr, void* i)
{
	static_cast<QList<QDockWidget *>*>(ptr)->append(static_cast<QDockWidget*>(i));
}

void* MainWindow46aa22___resizeDocks_docks_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QDockWidget *>();
}

int MainWindow46aa22___resizeDocks_sizes_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MainWindow46aa22___resizeDocks_sizes_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* MainWindow46aa22___resizeDocks_sizes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* MainWindow46aa22___tabifiedDockWidgets_atList(void* ptr, int i)
{
	return ({QDockWidget * tmp = static_cast<QList<QDockWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QDockWidget *>*>(ptr)->size()-1) { static_cast<QList<QDockWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MainWindow46aa22___tabifiedDockWidgets_setList(void* ptr, void* i)
{
	static_cast<QList<QDockWidget *>*>(ptr)->append(static_cast<QDockWidget*>(i));
}

void* MainWindow46aa22___tabifiedDockWidgets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QDockWidget *>();
}

void* MainWindow46aa22___addActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MainWindow46aa22___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* MainWindow46aa22___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* MainWindow46aa22___insertActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MainWindow46aa22___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* MainWindow46aa22___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* MainWindow46aa22___actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MainWindow46aa22___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* MainWindow46aa22___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* MainWindow46aa22___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void MainWindow46aa22___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* MainWindow46aa22___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* MainWindow46aa22___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MainWindow46aa22___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MainWindow46aa22___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* MainWindow46aa22___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MainWindow46aa22___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MainWindow46aa22___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* MainWindow46aa22___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MainWindow46aa22___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MainWindow46aa22___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* MainWindow46aa22___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MainWindow46aa22___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MainWindow46aa22___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* MainWindow46aa22_NewMainWindow(void* parent, long long flags)
{
		return new MainWindow46aa22(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void MainWindow46aa22_DestroyMainWindow(void* ptr)
{
	static_cast<MainWindow46aa22*>(ptr)->~MainWindow46aa22();
}

void MainWindow46aa22_DestroyMainWindowDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* MainWindow46aa22_CreatePopupMenuDefault(void* ptr)
{
	return static_cast<MainWindow46aa22*>(ptr)->QMainWindow::createPopupMenu();
}

char MainWindow46aa22_EventDefault(void* ptr, void* event)
{
	return static_cast<MainWindow46aa22*>(ptr)->QMainWindow::event(static_cast<QEvent*>(event));
}

void MainWindow46aa22_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void MainWindow46aa22_SetAnimatedDefault(void* ptr, char enabled)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::setAnimated(enabled != 0);
}

void MainWindow46aa22_SetDockNestingEnabledDefault(void* ptr, char enabled)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::setDockNestingEnabled(enabled != 0);
}

void MainWindow46aa22_SetUnifiedTitleAndToolBarOnMacDefault(void* ptr, char set)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::setUnifiedTitleAndToolBarOnMac(set != 0);
}



char MainWindow46aa22_CloseDefault(void* ptr)
{
	return static_cast<MainWindow46aa22*>(ptr)->QMainWindow::close();
}

char MainWindow46aa22_FocusNextPrevChildDefault(void* ptr, char next)
{
	return static_cast<MainWindow46aa22*>(ptr)->QMainWindow::focusNextPrevChild(next != 0);
}

char MainWindow46aa22_NativeEventDefault(void* ptr, void* eventType, void* message, long* result)
{
	return static_cast<MainWindow46aa22*>(ptr)->QMainWindow::nativeEvent(*static_cast<QByteArray*>(eventType), message, result);
}

void MainWindow46aa22_ActionEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::actionEvent(static_cast<QActionEvent*>(event));
}

void MainWindow46aa22_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::changeEvent(static_cast<QEvent*>(event));
}

void MainWindow46aa22_CloseEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::closeEvent(static_cast<QCloseEvent*>(event));
}

void MainWindow46aa22_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void MainWindow46aa22_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void MainWindow46aa22_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void MainWindow46aa22_DropEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::dropEvent(static_cast<QDropEvent*>(event));
}

void MainWindow46aa22_EnterEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::enterEvent(static_cast<QEvent*>(event));
}

void MainWindow46aa22_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::focusInEvent(static_cast<QFocusEvent*>(event));
}

void MainWindow46aa22_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void MainWindow46aa22_HideDefault(void* ptr)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::hide();
}

void MainWindow46aa22_HideEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::hideEvent(static_cast<QHideEvent*>(event));
}

void MainWindow46aa22_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void MainWindow46aa22_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void MainWindow46aa22_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void MainWindow46aa22_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::leaveEvent(static_cast<QEvent*>(event));
}

void MainWindow46aa22_LowerDefault(void* ptr)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::lower();
}

void MainWindow46aa22_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void MainWindow46aa22_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void MainWindow46aa22_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void MainWindow46aa22_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void MainWindow46aa22_MoveEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::moveEvent(static_cast<QMoveEvent*>(event));
}

void MainWindow46aa22_PaintEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::paintEvent(static_cast<QPaintEvent*>(event));
}

void MainWindow46aa22_RaiseDefault(void* ptr)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::raise();
}

void MainWindow46aa22_RepaintDefault(void* ptr)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::repaint();
}

void MainWindow46aa22_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::resizeEvent(static_cast<QResizeEvent*>(event));
}

void MainWindow46aa22_SetDisabledDefault(void* ptr, char disable)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::setDisabled(disable != 0);
}

void MainWindow46aa22_SetEnabledDefault(void* ptr, char vbo)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::setEnabled(vbo != 0);
}

void MainWindow46aa22_SetFocus2Default(void* ptr)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::setFocus();
}

void MainWindow46aa22_SetHiddenDefault(void* ptr, char hidden)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::setHidden(hidden != 0);
}

void MainWindow46aa22_SetStyleSheetDefault(void* ptr, struct Moc_PackedString styleSheet)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void MainWindow46aa22_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::setVisible(visible != 0);
}

void MainWindow46aa22_SetWindowModifiedDefault(void* ptr, char vbo)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::setWindowModified(vbo != 0);
}

void MainWindow46aa22_SetWindowTitleDefault(void* ptr, struct Moc_PackedString vqs)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void MainWindow46aa22_ShowDefault(void* ptr)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::show();
}

void MainWindow46aa22_ShowEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::showEvent(static_cast<QShowEvent*>(event));
}

void MainWindow46aa22_ShowFullScreenDefault(void* ptr)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::showFullScreen();
}

void MainWindow46aa22_ShowMaximizedDefault(void* ptr)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::showMaximized();
}

void MainWindow46aa22_ShowMinimizedDefault(void* ptr)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::showMinimized();
}

void MainWindow46aa22_ShowNormalDefault(void* ptr)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::showNormal();
}

void MainWindow46aa22_TabletEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::tabletEvent(static_cast<QTabletEvent*>(event));
}

void MainWindow46aa22_UpdateDefault(void* ptr)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::update();
}

void MainWindow46aa22_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::updateMicroFocus();
}

void MainWindow46aa22_WheelEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* MainWindow46aa22_PaintEngineDefault(void* ptr)
{
	return static_cast<MainWindow46aa22*>(ptr)->QMainWindow::paintEngine();
}

void* MainWindow46aa22_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<MainWindow46aa22*>(ptr)->QMainWindow::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* MainWindow46aa22_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<MainWindow46aa22*>(ptr)->QMainWindow::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* MainWindow46aa22_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<MainWindow46aa22*>(ptr)->QMainWindow::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char MainWindow46aa22_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<MainWindow46aa22*>(ptr)->QMainWindow::hasHeightForWidth();
}

int MainWindow46aa22_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<MainWindow46aa22*>(ptr)->QMainWindow::heightForWidth(w);
}

int MainWindow46aa22_MetricDefault(void* ptr, long long m)
{
	return static_cast<MainWindow46aa22*>(ptr)->QMainWindow::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void MainWindow46aa22_InitPainterDefault(void* ptr, void* painter)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::initPainter(static_cast<QPainter*>(painter));
}

char MainWindow46aa22_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<MainWindow46aa22*>(ptr)->QMainWindow::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void MainWindow46aa22_ChildEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::childEvent(static_cast<QChildEvent*>(event));
}

void MainWindow46aa22_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void MainWindow46aa22_CustomEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::customEvent(static_cast<QEvent*>(event));
}

void MainWindow46aa22_DeleteLaterDefault(void* ptr)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::deleteLater();
}

void MainWindow46aa22_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void MainWindow46aa22_TimerEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow46aa22*>(ptr)->QMainWindow::timerEvent(static_cast<QTimerEvent*>(event));
}

int MediaListPage46aa22_MediaListPage46aa22_QRegisterMetaType()
{
	return qRegisterMetaType<MediaListPage46aa22*>();
}

int MediaListPage46aa22_MediaListPage46aa22_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<MediaListPage46aa22*>(const_cast<const char*>(typeName));
}

int MediaListPage46aa22_MediaListPage46aa22_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<MediaListPage46aa22>();
#else
	return 0;
#endif
}

int MediaListPage46aa22_MediaListPage46aa22_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<MediaListPage46aa22>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* MediaListPage46aa22___addActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MediaListPage46aa22___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* MediaListPage46aa22___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* MediaListPage46aa22___insertActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MediaListPage46aa22___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* MediaListPage46aa22___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* MediaListPage46aa22___actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MediaListPage46aa22___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* MediaListPage46aa22___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* MediaListPage46aa22___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void MediaListPage46aa22___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* MediaListPage46aa22___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* MediaListPage46aa22___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MediaListPage46aa22___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MediaListPage46aa22___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* MediaListPage46aa22___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MediaListPage46aa22___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MediaListPage46aa22___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* MediaListPage46aa22___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MediaListPage46aa22___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MediaListPage46aa22___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* MediaListPage46aa22___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MediaListPage46aa22___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MediaListPage46aa22___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* MediaListPage46aa22_NewMediaListPage(void* parent, long long ff)
{
		return new MediaListPage46aa22(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(ff));
}

void MediaListPage46aa22_DestroyMediaListPage(void* ptr)
{
	static_cast<MediaListPage46aa22*>(ptr)->~MediaListPage46aa22();
}

void MediaListPage46aa22_DestroyMediaListPageDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char MediaListPage46aa22_CloseDefault(void* ptr)
{
	return static_cast<MediaListPage46aa22*>(ptr)->QWidget::close();
}

char MediaListPage46aa22_EventDefault(void* ptr, void* event)
{
	return static_cast<MediaListPage46aa22*>(ptr)->QWidget::event(static_cast<QEvent*>(event));
}

char MediaListPage46aa22_FocusNextPrevChildDefault(void* ptr, char next)
{
	return static_cast<MediaListPage46aa22*>(ptr)->QWidget::focusNextPrevChild(next != 0);
}

char MediaListPage46aa22_NativeEventDefault(void* ptr, void* eventType, void* message, long* result)
{
	return static_cast<MediaListPage46aa22*>(ptr)->QWidget::nativeEvent(*static_cast<QByteArray*>(eventType), message, result);
}

void MediaListPage46aa22_ActionEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void MediaListPage46aa22_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::changeEvent(static_cast<QEvent*>(event));
}

void MediaListPage46aa22_CloseEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void MediaListPage46aa22_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void MediaListPage46aa22_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void MediaListPage46aa22_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void MediaListPage46aa22_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void MediaListPage46aa22_DropEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::dropEvent(static_cast<QDropEvent*>(event));
}

void MediaListPage46aa22_EnterEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::enterEvent(static_cast<QEvent*>(event));
}

void MediaListPage46aa22_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

void MediaListPage46aa22_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void MediaListPage46aa22_HideDefault(void* ptr)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::hide();
}

void MediaListPage46aa22_HideEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::hideEvent(static_cast<QHideEvent*>(event));
}

void MediaListPage46aa22_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void MediaListPage46aa22_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void MediaListPage46aa22_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void MediaListPage46aa22_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::leaveEvent(static_cast<QEvent*>(event));
}

void MediaListPage46aa22_LowerDefault(void* ptr)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::lower();
}

void MediaListPage46aa22_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void MediaListPage46aa22_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void MediaListPage46aa22_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void MediaListPage46aa22_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void MediaListPage46aa22_MoveEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void MediaListPage46aa22_PaintEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::paintEvent(static_cast<QPaintEvent*>(event));
}

void MediaListPage46aa22_RaiseDefault(void* ptr)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::raise();
}

void MediaListPage46aa22_RepaintDefault(void* ptr)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::repaint();
}

void MediaListPage46aa22_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::resizeEvent(static_cast<QResizeEvent*>(event));
}

void MediaListPage46aa22_SetDisabledDefault(void* ptr, char disable)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::setDisabled(disable != 0);
}

void MediaListPage46aa22_SetEnabledDefault(void* ptr, char vbo)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::setEnabled(vbo != 0);
}

void MediaListPage46aa22_SetFocus2Default(void* ptr)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::setFocus();
}

void MediaListPage46aa22_SetHiddenDefault(void* ptr, char hidden)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::setHidden(hidden != 0);
}

void MediaListPage46aa22_SetStyleSheetDefault(void* ptr, struct Moc_PackedString styleSheet)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void MediaListPage46aa22_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::setVisible(visible != 0);
}

void MediaListPage46aa22_SetWindowModifiedDefault(void* ptr, char vbo)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::setWindowModified(vbo != 0);
}

void MediaListPage46aa22_SetWindowTitleDefault(void* ptr, struct Moc_PackedString vqs)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void MediaListPage46aa22_ShowDefault(void* ptr)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::show();
}

void MediaListPage46aa22_ShowEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::showEvent(static_cast<QShowEvent*>(event));
}

void MediaListPage46aa22_ShowFullScreenDefault(void* ptr)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::showFullScreen();
}

void MediaListPage46aa22_ShowMaximizedDefault(void* ptr)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::showMaximized();
}

void MediaListPage46aa22_ShowMinimizedDefault(void* ptr)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::showMinimized();
}

void MediaListPage46aa22_ShowNormalDefault(void* ptr)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::showNormal();
}

void MediaListPage46aa22_TabletEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void MediaListPage46aa22_UpdateDefault(void* ptr)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::update();
}

void MediaListPage46aa22_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::updateMicroFocus();
}

void MediaListPage46aa22_WheelEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* MediaListPage46aa22_PaintEngineDefault(void* ptr)
{
	return static_cast<MediaListPage46aa22*>(ptr)->QWidget::paintEngine();
}

void* MediaListPage46aa22_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<MediaListPage46aa22*>(ptr)->QWidget::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* MediaListPage46aa22_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<MediaListPage46aa22*>(ptr)->QWidget::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* MediaListPage46aa22_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<MediaListPage46aa22*>(ptr)->QWidget::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char MediaListPage46aa22_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<MediaListPage46aa22*>(ptr)->QWidget::hasHeightForWidth();
}



int MediaListPage46aa22_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<MediaListPage46aa22*>(ptr)->QWidget::heightForWidth(w);
}

int MediaListPage46aa22_MetricDefault(void* ptr, long long m)
{
	return static_cast<MediaListPage46aa22*>(ptr)->QWidget::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void MediaListPage46aa22_InitPainterDefault(void* ptr, void* painter)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::initPainter(static_cast<QPainter*>(painter));
}

char MediaListPage46aa22_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<MediaListPage46aa22*>(ptr)->QWidget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void MediaListPage46aa22_ChildEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::childEvent(static_cast<QChildEvent*>(event));
}

void MediaListPage46aa22_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void MediaListPage46aa22_CustomEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::customEvent(static_cast<QEvent*>(event));
}

void MediaListPage46aa22_DeleteLaterDefault(void* ptr)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::deleteLater();
}

void MediaListPage46aa22_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void MediaListPage46aa22_TimerEventDefault(void* ptr, void* event)
{
	static_cast<MediaListPage46aa22*>(ptr)->QWidget::timerEvent(static_cast<QTimerEvent*>(event));
}

int MediaListTable46aa22_MediaListTable46aa22_QRegisterMetaType()
{
	return qRegisterMetaType<MediaListTable46aa22*>();
}

int MediaListTable46aa22_MediaListTable46aa22_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<MediaListTable46aa22*>(const_cast<const char*>(typeName));
}

int MediaListTable46aa22_MediaListTable46aa22_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<MediaListTable46aa22>();
#else
	return 0;
#endif
}

int MediaListTable46aa22_MediaListTable46aa22_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<MediaListTable46aa22>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int MediaListTable46aa22_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MediaListTable46aa22_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* MediaListTable46aa22_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int MediaListTable46aa22_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MediaListTable46aa22_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* MediaListTable46aa22_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int MediaListTable46aa22_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MediaListTable46aa22_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* MediaListTable46aa22_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* MediaListTable46aa22___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void MediaListTable46aa22___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* MediaListTable46aa22___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList MediaListTable46aa22___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* MediaListTable46aa22___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void MediaListTable46aa22___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* MediaListTable46aa22___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* MediaListTable46aa22___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void MediaListTable46aa22___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* MediaListTable46aa22___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int MediaListTable46aa22___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void MediaListTable46aa22___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* MediaListTable46aa22___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* MediaListTable46aa22___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void MediaListTable46aa22___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* MediaListTable46aa22___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* MediaListTable46aa22___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void MediaListTable46aa22___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* MediaListTable46aa22___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* MediaListTable46aa22___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void MediaListTable46aa22___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* MediaListTable46aa22___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList MediaListTable46aa22___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* MediaListTable46aa22___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void MediaListTable46aa22___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* MediaListTable46aa22___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList MediaListTable46aa22___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* MediaListTable46aa22___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void MediaListTable46aa22___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* MediaListTable46aa22___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* MediaListTable46aa22___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void MediaListTable46aa22___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* MediaListTable46aa22___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* MediaListTable46aa22___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void MediaListTable46aa22___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* MediaListTable46aa22___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int MediaListTable46aa22_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MediaListTable46aa22_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* MediaListTable46aa22_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int MediaListTable46aa22_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MediaListTable46aa22_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* MediaListTable46aa22_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* MediaListTable46aa22___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void MediaListTable46aa22___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* MediaListTable46aa22___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* MediaListTable46aa22___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MediaListTable46aa22___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MediaListTable46aa22___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* MediaListTable46aa22___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MediaListTable46aa22___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MediaListTable46aa22___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* MediaListTable46aa22___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MediaListTable46aa22___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MediaListTable46aa22___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* MediaListTable46aa22___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MediaListTable46aa22___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MediaListTable46aa22___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* MediaListTable46aa22_NewMediaListTable(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MediaListTable46aa22(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MediaListTable46aa22(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MediaListTable46aa22(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MediaListTable46aa22(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MediaListTable46aa22(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MediaListTable46aa22(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MediaListTable46aa22(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MediaListTable46aa22(static_cast<QWindow*>(parent));
	} else {
		return new MediaListTable46aa22(static_cast<QObject*>(parent));
	}
}

void MediaListTable46aa22_DestroyMediaListTable(void* ptr)
{
	static_cast<MediaListTable46aa22*>(ptr)->~MediaListTable46aa22();
}

void MediaListTable46aa22_DestroyMediaListTableDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char MediaListTable46aa22_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void* MediaListTable46aa22_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* MediaListTable46aa22_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

long long MediaListTable46aa22_FlagsDefault(void* ptr, void* index)
{
	return static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::flags(*static_cast<QModelIndex*>(index));
}



char MediaListTable46aa22_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char MediaListTable46aa22_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char MediaListTable46aa22_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char MediaListTable46aa22_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char MediaListTable46aa22_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char MediaListTable46aa22_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

char MediaListTable46aa22_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char MediaListTable46aa22_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char MediaListTable46aa22_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

char MediaListTable46aa22_SubmitDefault(void* ptr)
{
	return static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::submit();
}

void MediaListTable46aa22_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

void MediaListTable46aa22_ResetInternalDataDefault(void* ptr)
{
	static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::resetInternalData();
}

void MediaListTable46aa22_RevertDefault(void* ptr)
{
	static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::revert();
}

void MediaListTable46aa22_SortDefault(void* ptr, int column, long long order)
{
	static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::sort(column, static_cast<Qt::SortOrder>(order));
}

struct Moc_PackedList MediaListTable46aa22_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList MediaListTable46aa22_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* MediaListTable46aa22_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void* MediaListTable46aa22_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::buddy(*static_cast<QModelIndex*>(index)));
}

void* MediaListTable46aa22_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::parent(*static_cast<QModelIndex*>(index)));
}

struct Moc_PackedList MediaListTable46aa22_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* MediaListTable46aa22_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct Moc_PackedString MediaListTable46aa22_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray td75236 = static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::mimeTypes().join("¡¦!").toUtf8(); Moc_PackedString { const_cast<char*>(td75236.prepend("WHITESPACE").constData()+10), td75236.size()-10 }; });
}

void* MediaListTable46aa22_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void* MediaListTable46aa22_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

long long MediaListTable46aa22_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::supportedDragActions();
}

long long MediaListTable46aa22_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::supportedDropActions();
}

char MediaListTable46aa22_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char MediaListTable46aa22_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

char MediaListTable46aa22_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

int MediaListTable46aa22_ColumnCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

int MediaListTable46aa22_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char MediaListTable46aa22_EventDefault(void* ptr, void* e)
{
	return static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::event(static_cast<QEvent*>(e));
}

char MediaListTable46aa22_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void MediaListTable46aa22_ChildEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::childEvent(static_cast<QChildEvent*>(event));
}

void MediaListTable46aa22_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void MediaListTable46aa22_CustomEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::customEvent(static_cast<QEvent*>(event));
}

void MediaListTable46aa22_DeleteLaterDefault(void* ptr)
{
	static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::deleteLater();
}

void MediaListTable46aa22_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void MediaListTable46aa22_TimerEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTable46aa22*>(ptr)->QAbstractTableModel::timerEvent(static_cast<QTimerEvent*>(event));
}

int MediaListTabs46aa22_MediaListTabs46aa22_QRegisterMetaType()
{
	return qRegisterMetaType<MediaListTabs46aa22*>();
}

int MediaListTabs46aa22_MediaListTabs46aa22_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<MediaListTabs46aa22*>(const_cast<const char*>(typeName));
}

int MediaListTabs46aa22_MediaListTabs46aa22_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<MediaListTabs46aa22>();
#else
	return 0;
#endif
}

int MediaListTabs46aa22_MediaListTabs46aa22_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<MediaListTabs46aa22>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* MediaListTabs46aa22___addActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MediaListTabs46aa22___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* MediaListTabs46aa22___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* MediaListTabs46aa22___insertActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MediaListTabs46aa22___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* MediaListTabs46aa22___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* MediaListTabs46aa22___actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MediaListTabs46aa22___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* MediaListTabs46aa22___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* MediaListTabs46aa22___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void MediaListTabs46aa22___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* MediaListTabs46aa22___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* MediaListTabs46aa22___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MediaListTabs46aa22___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MediaListTabs46aa22___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* MediaListTabs46aa22___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MediaListTabs46aa22___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MediaListTabs46aa22___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* MediaListTabs46aa22___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MediaListTabs46aa22___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MediaListTabs46aa22___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* MediaListTabs46aa22___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MediaListTabs46aa22___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MediaListTabs46aa22___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* MediaListTabs46aa22_NewMediaListTabs(void* parent)
{
		return new MediaListTabs46aa22(static_cast<QWidget*>(parent));
}

void MediaListTabs46aa22_DestroyMediaListTabs(void* ptr)
{
	static_cast<MediaListTabs46aa22*>(ptr)->~MediaListTabs46aa22();
}

void MediaListTabs46aa22_DestroyMediaListTabsDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char MediaListTabs46aa22_EventDefault(void* ptr, void* ev)
{
	return static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::event(static_cast<QEvent*>(ev));
}

void MediaListTabs46aa22_ChangeEventDefault(void* ptr, void* ev)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::changeEvent(static_cast<QEvent*>(ev));
}

void MediaListTabs46aa22_KeyPressEventDefault(void* ptr, void* e)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::keyPressEvent(static_cast<QKeyEvent*>(e));
}

void MediaListTabs46aa22_PaintEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::paintEvent(static_cast<QPaintEvent*>(event));
}

void MediaListTabs46aa22_ResizeEventDefault(void* ptr, void* e)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::resizeEvent(static_cast<QResizeEvent*>(e));
}

void MediaListTabs46aa22_SetCurrentIndexDefault(void* ptr, int index)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::setCurrentIndex(index);
}

void MediaListTabs46aa22_SetCurrentWidgetDefault(void* ptr, void* widget)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::setCurrentWidget(static_cast<QWidget*>(widget));
}

void MediaListTabs46aa22_ShowEventDefault(void* ptr, void* vqs)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::showEvent(static_cast<QShowEvent*>(vqs));
}

void MediaListTabs46aa22_TabInsertedDefault(void* ptr, int index)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::tabInserted(index);
}

void MediaListTabs46aa22_TabRemovedDefault(void* ptr, int index)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::tabRemoved(index);
}

void* MediaListTabs46aa22_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* MediaListTabs46aa22_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

char MediaListTabs46aa22_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::hasHeightForWidth();
}



int MediaListTabs46aa22_HeightForWidthDefault(void* ptr, int width)
{
	return static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::heightForWidth(width);
}

char MediaListTabs46aa22_CloseDefault(void* ptr)
{
	return static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::close();
}

char MediaListTabs46aa22_FocusNextPrevChildDefault(void* ptr, char next)
{
	return static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::focusNextPrevChild(next != 0);
}

char MediaListTabs46aa22_NativeEventDefault(void* ptr, void* eventType, void* message, long* result)
{
	return static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::nativeEvent(*static_cast<QByteArray*>(eventType), message, result);
}

void MediaListTabs46aa22_ActionEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void MediaListTabs46aa22_CloseEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void MediaListTabs46aa22_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void MediaListTabs46aa22_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void MediaListTabs46aa22_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void MediaListTabs46aa22_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void MediaListTabs46aa22_DropEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::dropEvent(static_cast<QDropEvent*>(event));
}

void MediaListTabs46aa22_EnterEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::enterEvent(static_cast<QEvent*>(event));
}

void MediaListTabs46aa22_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

void MediaListTabs46aa22_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void MediaListTabs46aa22_HideDefault(void* ptr)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::hide();
}

void MediaListTabs46aa22_HideEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::hideEvent(static_cast<QHideEvent*>(event));
}

void MediaListTabs46aa22_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void MediaListTabs46aa22_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void MediaListTabs46aa22_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::leaveEvent(static_cast<QEvent*>(event));
}

void MediaListTabs46aa22_LowerDefault(void* ptr)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::lower();
}

void MediaListTabs46aa22_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void MediaListTabs46aa22_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void MediaListTabs46aa22_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void MediaListTabs46aa22_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void MediaListTabs46aa22_MoveEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void MediaListTabs46aa22_RaiseDefault(void* ptr)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::raise();
}

void MediaListTabs46aa22_RepaintDefault(void* ptr)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::repaint();
}

void MediaListTabs46aa22_SetDisabledDefault(void* ptr, char disable)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::setDisabled(disable != 0);
}

void MediaListTabs46aa22_SetEnabledDefault(void* ptr, char vbo)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::setEnabled(vbo != 0);
}

void MediaListTabs46aa22_SetFocus2Default(void* ptr)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::setFocus();
}

void MediaListTabs46aa22_SetHiddenDefault(void* ptr, char hidden)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::setHidden(hidden != 0);
}

void MediaListTabs46aa22_SetStyleSheetDefault(void* ptr, struct Moc_PackedString styleSheet)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void MediaListTabs46aa22_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::setVisible(visible != 0);
}

void MediaListTabs46aa22_SetWindowModifiedDefault(void* ptr, char vbo)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::setWindowModified(vbo != 0);
}

void MediaListTabs46aa22_SetWindowTitleDefault(void* ptr, struct Moc_PackedString vqs)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void MediaListTabs46aa22_ShowDefault(void* ptr)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::show();
}

void MediaListTabs46aa22_ShowFullScreenDefault(void* ptr)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::showFullScreen();
}

void MediaListTabs46aa22_ShowMaximizedDefault(void* ptr)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::showMaximized();
}

void MediaListTabs46aa22_ShowMinimizedDefault(void* ptr)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::showMinimized();
}

void MediaListTabs46aa22_ShowNormalDefault(void* ptr)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::showNormal();
}

void MediaListTabs46aa22_TabletEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void MediaListTabs46aa22_UpdateDefault(void* ptr)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::update();
}

void MediaListTabs46aa22_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::updateMicroFocus();
}

void MediaListTabs46aa22_WheelEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* MediaListTabs46aa22_PaintEngineDefault(void* ptr)
{
	return static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::paintEngine();
}

void* MediaListTabs46aa22_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

int MediaListTabs46aa22_MetricDefault(void* ptr, long long m)
{
	return static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void MediaListTabs46aa22_InitPainterDefault(void* ptr, void* painter)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::initPainter(static_cast<QPainter*>(painter));
}

char MediaListTabs46aa22_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void MediaListTabs46aa22_ChildEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::childEvent(static_cast<QChildEvent*>(event));
}

void MediaListTabs46aa22_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void MediaListTabs46aa22_CustomEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::customEvent(static_cast<QEvent*>(event));
}

void MediaListTabs46aa22_DeleteLaterDefault(void* ptr)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::deleteLater();
}

void MediaListTabs46aa22_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void MediaListTabs46aa22_TimerEventDefault(void* ptr, void* event)
{
	static_cast<MediaListTabs46aa22*>(ptr)->QTabWidget::timerEvent(static_cast<QTimerEvent*>(event));
}

int MenuBar46aa22_MenuBar46aa22_QRegisterMetaType()
{
	return qRegisterMetaType<MenuBar46aa22*>();
}

int MenuBar46aa22_MenuBar46aa22_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<MenuBar46aa22*>(const_cast<const char*>(typeName));
}

int MenuBar46aa22_MenuBar46aa22_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<MenuBar46aa22>();
#else
	return 0;
#endif
}

int MenuBar46aa22_MenuBar46aa22_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<MenuBar46aa22>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* MenuBar46aa22___addActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MenuBar46aa22___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* MenuBar46aa22___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* MenuBar46aa22___insertActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MenuBar46aa22___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* MenuBar46aa22___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* MenuBar46aa22___actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MenuBar46aa22___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* MenuBar46aa22___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* MenuBar46aa22___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void MenuBar46aa22___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* MenuBar46aa22___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* MenuBar46aa22___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MenuBar46aa22___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MenuBar46aa22___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* MenuBar46aa22___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MenuBar46aa22___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MenuBar46aa22___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* MenuBar46aa22___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MenuBar46aa22___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MenuBar46aa22___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* MenuBar46aa22___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MenuBar46aa22___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MenuBar46aa22___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* MenuBar46aa22_NewMenuBar(void* parent)
{
		return new MenuBar46aa22(static_cast<QWidget*>(parent));
}

void MenuBar46aa22_DestroyMenuBar(void* ptr)
{
	static_cast<MenuBar46aa22*>(ptr)->~MenuBar46aa22();
}

void MenuBar46aa22_DestroyMenuBarDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char MenuBar46aa22_EventDefault(void* ptr, void* e)
{
	return static_cast<MenuBar46aa22*>(ptr)->QMenuBar::event(static_cast<QEvent*>(e));
}

char MenuBar46aa22_EventFilterDefault(void* ptr, void* object, void* event)
{
	return static_cast<MenuBar46aa22*>(ptr)->QMenuBar::eventFilter(static_cast<QObject*>(object), static_cast<QEvent*>(event));
}

void MenuBar46aa22_ActionEventDefault(void* ptr, void* e)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::actionEvent(static_cast<QActionEvent*>(e));
}

void MenuBar46aa22_ChangeEventDefault(void* ptr, void* e)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::changeEvent(static_cast<QEvent*>(e));
}

void MenuBar46aa22_FocusInEventDefault(void* ptr, void* vqf)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::focusInEvent(static_cast<QFocusEvent*>(vqf));
}

void MenuBar46aa22_FocusOutEventDefault(void* ptr, void* vqf)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::focusOutEvent(static_cast<QFocusEvent*>(vqf));
}

void MenuBar46aa22_KeyPressEventDefault(void* ptr, void* e)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::keyPressEvent(static_cast<QKeyEvent*>(e));
}

void MenuBar46aa22_LeaveEventDefault(void* ptr, void* vqe)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::leaveEvent(static_cast<QEvent*>(vqe));
}

void MenuBar46aa22_MouseMoveEventDefault(void* ptr, void* e)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::mouseMoveEvent(static_cast<QMouseEvent*>(e));
}

void MenuBar46aa22_MousePressEventDefault(void* ptr, void* e)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::mousePressEvent(static_cast<QMouseEvent*>(e));
}

void MenuBar46aa22_MouseReleaseEventDefault(void* ptr, void* e)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::mouseReleaseEvent(static_cast<QMouseEvent*>(e));
}

void MenuBar46aa22_PaintEventDefault(void* ptr, void* e)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::paintEvent(static_cast<QPaintEvent*>(e));
}

void MenuBar46aa22_ResizeEventDefault(void* ptr, void* vqr)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::resizeEvent(static_cast<QResizeEvent*>(vqr));
}

void MenuBar46aa22_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::setVisible(visible != 0);
}

void MenuBar46aa22_TimerEventDefault(void* ptr, void* e)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::timerEvent(static_cast<QTimerEvent*>(e));
}

void* MenuBar46aa22_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<MenuBar46aa22*>(ptr)->QMenuBar::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* MenuBar46aa22_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<MenuBar46aa22*>(ptr)->QMenuBar::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}



int MenuBar46aa22_HeightForWidthDefault(void* ptr, int vin)
{
	return static_cast<MenuBar46aa22*>(ptr)->QMenuBar::heightForWidth(vin);
}

char MenuBar46aa22_CloseDefault(void* ptr)
{
	return static_cast<MenuBar46aa22*>(ptr)->QMenuBar::close();
}

char MenuBar46aa22_FocusNextPrevChildDefault(void* ptr, char next)
{
	return static_cast<MenuBar46aa22*>(ptr)->QMenuBar::focusNextPrevChild(next != 0);
}

char MenuBar46aa22_NativeEventDefault(void* ptr, void* eventType, void* message, long* result)
{
	return static_cast<MenuBar46aa22*>(ptr)->QMenuBar::nativeEvent(*static_cast<QByteArray*>(eventType), message, result);
}

void MenuBar46aa22_CloseEventDefault(void* ptr, void* event)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::closeEvent(static_cast<QCloseEvent*>(event));
}

void MenuBar46aa22_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void MenuBar46aa22_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void MenuBar46aa22_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void MenuBar46aa22_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void MenuBar46aa22_DropEventDefault(void* ptr, void* event)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::dropEvent(static_cast<QDropEvent*>(event));
}

void MenuBar46aa22_EnterEventDefault(void* ptr, void* event)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::enterEvent(static_cast<QEvent*>(event));
}

void MenuBar46aa22_HideDefault(void* ptr)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::hide();
}

void MenuBar46aa22_HideEventDefault(void* ptr, void* event)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::hideEvent(static_cast<QHideEvent*>(event));
}

void MenuBar46aa22_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void MenuBar46aa22_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void MenuBar46aa22_LowerDefault(void* ptr)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::lower();
}

void MenuBar46aa22_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void MenuBar46aa22_MoveEventDefault(void* ptr, void* event)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::moveEvent(static_cast<QMoveEvent*>(event));
}

void MenuBar46aa22_RaiseDefault(void* ptr)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::raise();
}

void MenuBar46aa22_RepaintDefault(void* ptr)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::repaint();
}

void MenuBar46aa22_SetDisabledDefault(void* ptr, char disable)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::setDisabled(disable != 0);
}

void MenuBar46aa22_SetEnabledDefault(void* ptr, char vbo)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::setEnabled(vbo != 0);
}

void MenuBar46aa22_SetFocus2Default(void* ptr)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::setFocus();
}

void MenuBar46aa22_SetHiddenDefault(void* ptr, char hidden)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::setHidden(hidden != 0);
}

void MenuBar46aa22_SetStyleSheetDefault(void* ptr, struct Moc_PackedString styleSheet)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void MenuBar46aa22_SetWindowModifiedDefault(void* ptr, char vbo)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::setWindowModified(vbo != 0);
}

void MenuBar46aa22_SetWindowTitleDefault(void* ptr, struct Moc_PackedString vqs)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void MenuBar46aa22_ShowDefault(void* ptr)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::show();
}

void MenuBar46aa22_ShowEventDefault(void* ptr, void* event)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::showEvent(static_cast<QShowEvent*>(event));
}

void MenuBar46aa22_ShowFullScreenDefault(void* ptr)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::showFullScreen();
}

void MenuBar46aa22_ShowMaximizedDefault(void* ptr)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::showMaximized();
}

void MenuBar46aa22_ShowMinimizedDefault(void* ptr)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::showMinimized();
}

void MenuBar46aa22_ShowNormalDefault(void* ptr)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::showNormal();
}

void MenuBar46aa22_TabletEventDefault(void* ptr, void* event)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::tabletEvent(static_cast<QTabletEvent*>(event));
}

void MenuBar46aa22_UpdateDefault(void* ptr)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::update();
}

void MenuBar46aa22_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::updateMicroFocus();
}

void MenuBar46aa22_WheelEventDefault(void* ptr, void* event)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* MenuBar46aa22_PaintEngineDefault(void* ptr)
{
	return static_cast<MenuBar46aa22*>(ptr)->QMenuBar::paintEngine();
}

void* MenuBar46aa22_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<MenuBar46aa22*>(ptr)->QMenuBar::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char MenuBar46aa22_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<MenuBar46aa22*>(ptr)->QMenuBar::hasHeightForWidth();
}

int MenuBar46aa22_MetricDefault(void* ptr, long long m)
{
	return static_cast<MenuBar46aa22*>(ptr)->QMenuBar::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void MenuBar46aa22_InitPainterDefault(void* ptr, void* painter)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::initPainter(static_cast<QPainter*>(painter));
}

void MenuBar46aa22_ChildEventDefault(void* ptr, void* event)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::childEvent(static_cast<QChildEvent*>(event));
}

void MenuBar46aa22_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void MenuBar46aa22_CustomEventDefault(void* ptr, void* event)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::customEvent(static_cast<QEvent*>(event));
}

void MenuBar46aa22_DeleteLaterDefault(void* ptr)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::deleteLater();
}

void MenuBar46aa22_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<MenuBar46aa22*>(ptr)->QMenuBar::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int Navigation46aa22_Navigation46aa22_QRegisterMetaType()
{
	return qRegisterMetaType<Navigation46aa22*>();
}

int Navigation46aa22_Navigation46aa22_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<Navigation46aa22*>(const_cast<const char*>(typeName));
}

int Navigation46aa22_Navigation46aa22_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Navigation46aa22>();
#else
	return 0;
#endif
}

int Navigation46aa22_Navigation46aa22_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Navigation46aa22>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* Navigation46aa22___addActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Navigation46aa22___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* Navigation46aa22___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* Navigation46aa22___insertActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Navigation46aa22___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* Navigation46aa22___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* Navigation46aa22___actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Navigation46aa22___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* Navigation46aa22___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* Navigation46aa22___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void Navigation46aa22___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* Navigation46aa22___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* Navigation46aa22___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Navigation46aa22___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Navigation46aa22___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Navigation46aa22___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Navigation46aa22___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Navigation46aa22___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Navigation46aa22___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Navigation46aa22___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Navigation46aa22___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Navigation46aa22___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Navigation46aa22___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Navigation46aa22___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* Navigation46aa22_NewNavigation(void* parent, long long ff)
{
		return new Navigation46aa22(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(ff));
}

void Navigation46aa22_DestroyNavigation(void* ptr)
{
	static_cast<Navigation46aa22*>(ptr)->~Navigation46aa22();
}

void Navigation46aa22_DestroyNavigationDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char Navigation46aa22_CloseDefault(void* ptr)
{
	return static_cast<Navigation46aa22*>(ptr)->QWidget::close();
}

char Navigation46aa22_EventDefault(void* ptr, void* event)
{
	return static_cast<Navigation46aa22*>(ptr)->QWidget::event(static_cast<QEvent*>(event));
}

char Navigation46aa22_FocusNextPrevChildDefault(void* ptr, char next)
{
	return static_cast<Navigation46aa22*>(ptr)->QWidget::focusNextPrevChild(next != 0);
}

char Navigation46aa22_NativeEventDefault(void* ptr, void* eventType, void* message, long* result)
{
	return static_cast<Navigation46aa22*>(ptr)->QWidget::nativeEvent(*static_cast<QByteArray*>(eventType), message, result);
}

void Navigation46aa22_ActionEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void Navigation46aa22_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::changeEvent(static_cast<QEvent*>(event));
}

void Navigation46aa22_CloseEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void Navigation46aa22_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void Navigation46aa22_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void Navigation46aa22_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void Navigation46aa22_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void Navigation46aa22_DropEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::dropEvent(static_cast<QDropEvent*>(event));
}

void Navigation46aa22_EnterEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::enterEvent(static_cast<QEvent*>(event));
}

void Navigation46aa22_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

void Navigation46aa22_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void Navigation46aa22_HideDefault(void* ptr)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::hide();
}

void Navigation46aa22_HideEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::hideEvent(static_cast<QHideEvent*>(event));
}

void Navigation46aa22_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void Navigation46aa22_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void Navigation46aa22_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void Navigation46aa22_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::leaveEvent(static_cast<QEvent*>(event));
}

void Navigation46aa22_LowerDefault(void* ptr)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::lower();
}

void Navigation46aa22_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void Navigation46aa22_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void Navigation46aa22_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void Navigation46aa22_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void Navigation46aa22_MoveEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void Navigation46aa22_PaintEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::paintEvent(static_cast<QPaintEvent*>(event));
}

void Navigation46aa22_RaiseDefault(void* ptr)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::raise();
}

void Navigation46aa22_RepaintDefault(void* ptr)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::repaint();
}

void Navigation46aa22_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::resizeEvent(static_cast<QResizeEvent*>(event));
}

void Navigation46aa22_SetDisabledDefault(void* ptr, char disable)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::setDisabled(disable != 0);
}

void Navigation46aa22_SetEnabledDefault(void* ptr, char vbo)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::setEnabled(vbo != 0);
}

void Navigation46aa22_SetFocus2Default(void* ptr)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::setFocus();
}

void Navigation46aa22_SetHiddenDefault(void* ptr, char hidden)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::setHidden(hidden != 0);
}

void Navigation46aa22_SetStyleSheetDefault(void* ptr, struct Moc_PackedString styleSheet)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void Navigation46aa22_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::setVisible(visible != 0);
}

void Navigation46aa22_SetWindowModifiedDefault(void* ptr, char vbo)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::setWindowModified(vbo != 0);
}

void Navigation46aa22_SetWindowTitleDefault(void* ptr, struct Moc_PackedString vqs)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void Navigation46aa22_ShowDefault(void* ptr)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::show();
}

void Navigation46aa22_ShowEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::showEvent(static_cast<QShowEvent*>(event));
}

void Navigation46aa22_ShowFullScreenDefault(void* ptr)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::showFullScreen();
}

void Navigation46aa22_ShowMaximizedDefault(void* ptr)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::showMaximized();
}

void Navigation46aa22_ShowMinimizedDefault(void* ptr)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::showMinimized();
}

void Navigation46aa22_ShowNormalDefault(void* ptr)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::showNormal();
}

void Navigation46aa22_TabletEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void Navigation46aa22_UpdateDefault(void* ptr)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::update();
}

void Navigation46aa22_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::updateMicroFocus();
}

void Navigation46aa22_WheelEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* Navigation46aa22_PaintEngineDefault(void* ptr)
{
	return static_cast<Navigation46aa22*>(ptr)->QWidget::paintEngine();
}

void* Navigation46aa22_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<Navigation46aa22*>(ptr)->QWidget::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* Navigation46aa22_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<Navigation46aa22*>(ptr)->QWidget::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* Navigation46aa22_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<Navigation46aa22*>(ptr)->QWidget::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char Navigation46aa22_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<Navigation46aa22*>(ptr)->QWidget::hasHeightForWidth();
}



int Navigation46aa22_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<Navigation46aa22*>(ptr)->QWidget::heightForWidth(w);
}

int Navigation46aa22_MetricDefault(void* ptr, long long m)
{
	return static_cast<Navigation46aa22*>(ptr)->QWidget::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void Navigation46aa22_InitPainterDefault(void* ptr, void* painter)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::initPainter(static_cast<QPainter*>(painter));
}

char Navigation46aa22_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<Navigation46aa22*>(ptr)->QWidget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void Navigation46aa22_ChildEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::childEvent(static_cast<QChildEvent*>(event));
}

void Navigation46aa22_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void Navigation46aa22_CustomEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::customEvent(static_cast<QEvent*>(event));
}

void Navigation46aa22_DeleteLaterDefault(void* ptr)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::deleteLater();
}

void Navigation46aa22_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void Navigation46aa22_TimerEventDefault(void* ptr, void* event)
{
	static_cast<Navigation46aa22*>(ptr)->QWidget::timerEvent(static_cast<QTimerEvent*>(event));
}

int SettingsDialog46aa22_SettingsDialog46aa22_QRegisterMetaType()
{
	return qRegisterMetaType<SettingsDialog46aa22*>();
}

int SettingsDialog46aa22_SettingsDialog46aa22_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<SettingsDialog46aa22*>(const_cast<const char*>(typeName));
}

int SettingsDialog46aa22_SettingsDialog46aa22_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<SettingsDialog46aa22>();
#else
	return 0;
#endif
}

int SettingsDialog46aa22_SettingsDialog46aa22_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<SettingsDialog46aa22>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* SettingsDialog46aa22___addActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void SettingsDialog46aa22___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* SettingsDialog46aa22___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* SettingsDialog46aa22___insertActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void SettingsDialog46aa22___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* SettingsDialog46aa22___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* SettingsDialog46aa22___actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void SettingsDialog46aa22___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* SettingsDialog46aa22___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* SettingsDialog46aa22___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void SettingsDialog46aa22___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* SettingsDialog46aa22___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* SettingsDialog46aa22___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void SettingsDialog46aa22___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* SettingsDialog46aa22___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* SettingsDialog46aa22___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void SettingsDialog46aa22___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* SettingsDialog46aa22___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* SettingsDialog46aa22___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void SettingsDialog46aa22___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* SettingsDialog46aa22___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* SettingsDialog46aa22___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void SettingsDialog46aa22___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* SettingsDialog46aa22___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* SettingsDialog46aa22_NewSettingsDialog(void* parent, long long ff)
{
		return new SettingsDialog46aa22(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(ff));
}

void SettingsDialog46aa22_DestroySettingsDialog(void* ptr)
{
	static_cast<SettingsDialog46aa22*>(ptr)->~SettingsDialog46aa22();
}

void SettingsDialog46aa22_DestroySettingsDialogDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char SettingsDialog46aa22_EventFilterDefault(void* ptr, void* o, void* e)
{
	return static_cast<SettingsDialog46aa22*>(ptr)->QDialog::eventFilter(static_cast<QObject*>(o), static_cast<QEvent*>(e));
}

int SettingsDialog46aa22_ExecDefault(void* ptr)
{
	return static_cast<SettingsDialog46aa22*>(ptr)->QDialog::exec();
}

void SettingsDialog46aa22_AcceptDefault(void* ptr)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::accept();
}

void SettingsDialog46aa22_CloseEventDefault(void* ptr, void* e)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::closeEvent(static_cast<QCloseEvent*>(e));
}

void SettingsDialog46aa22_ContextMenuEventDefault(void* ptr, void* e)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::contextMenuEvent(static_cast<QContextMenuEvent*>(e));
}

void SettingsDialog46aa22_DoneDefault(void* ptr, int r)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::done(r);
}

void SettingsDialog46aa22_KeyPressEventDefault(void* ptr, void* e)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::keyPressEvent(static_cast<QKeyEvent*>(e));
}

void SettingsDialog46aa22_OpenDefault(void* ptr)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::open();
}

void SettingsDialog46aa22_RejectDefault(void* ptr)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::reject();
}

void SettingsDialog46aa22_ResizeEventDefault(void* ptr, void* vqr)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::resizeEvent(static_cast<QResizeEvent*>(vqr));
}

void SettingsDialog46aa22_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::setVisible(visible != 0);
}

void SettingsDialog46aa22_ShowEventDefault(void* ptr, void* event)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::showEvent(static_cast<QShowEvent*>(event));
}

void* SettingsDialog46aa22_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<SettingsDialog46aa22*>(ptr)->QDialog::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* SettingsDialog46aa22_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<SettingsDialog46aa22*>(ptr)->QDialog::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}



char SettingsDialog46aa22_CloseDefault(void* ptr)
{
	return static_cast<SettingsDialog46aa22*>(ptr)->QDialog::close();
}

char SettingsDialog46aa22_EventDefault(void* ptr, void* event)
{
	return static_cast<SettingsDialog46aa22*>(ptr)->QDialog::event(static_cast<QEvent*>(event));
}

char SettingsDialog46aa22_FocusNextPrevChildDefault(void* ptr, char next)
{
	return static_cast<SettingsDialog46aa22*>(ptr)->QDialog::focusNextPrevChild(next != 0);
}

char SettingsDialog46aa22_NativeEventDefault(void* ptr, void* eventType, void* message, long* result)
{
	return static_cast<SettingsDialog46aa22*>(ptr)->QDialog::nativeEvent(*static_cast<QByteArray*>(eventType), message, result);
}

void SettingsDialog46aa22_ActionEventDefault(void* ptr, void* event)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::actionEvent(static_cast<QActionEvent*>(event));
}

void SettingsDialog46aa22_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::changeEvent(static_cast<QEvent*>(event));
}

void SettingsDialog46aa22_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void SettingsDialog46aa22_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void SettingsDialog46aa22_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void SettingsDialog46aa22_DropEventDefault(void* ptr, void* event)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::dropEvent(static_cast<QDropEvent*>(event));
}

void SettingsDialog46aa22_EnterEventDefault(void* ptr, void* event)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::enterEvent(static_cast<QEvent*>(event));
}

void SettingsDialog46aa22_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::focusInEvent(static_cast<QFocusEvent*>(event));
}

void SettingsDialog46aa22_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void SettingsDialog46aa22_HideDefault(void* ptr)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::hide();
}

void SettingsDialog46aa22_HideEventDefault(void* ptr, void* event)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::hideEvent(static_cast<QHideEvent*>(event));
}

void SettingsDialog46aa22_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void SettingsDialog46aa22_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void SettingsDialog46aa22_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::leaveEvent(static_cast<QEvent*>(event));
}

void SettingsDialog46aa22_LowerDefault(void* ptr)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::lower();
}

void SettingsDialog46aa22_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void SettingsDialog46aa22_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void SettingsDialog46aa22_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void SettingsDialog46aa22_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void SettingsDialog46aa22_MoveEventDefault(void* ptr, void* event)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::moveEvent(static_cast<QMoveEvent*>(event));
}

void SettingsDialog46aa22_PaintEventDefault(void* ptr, void* event)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::paintEvent(static_cast<QPaintEvent*>(event));
}

void SettingsDialog46aa22_RaiseDefault(void* ptr)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::raise();
}

void SettingsDialog46aa22_RepaintDefault(void* ptr)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::repaint();
}

void SettingsDialog46aa22_SetDisabledDefault(void* ptr, char disable)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::setDisabled(disable != 0);
}

void SettingsDialog46aa22_SetEnabledDefault(void* ptr, char vbo)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::setEnabled(vbo != 0);
}

void SettingsDialog46aa22_SetFocus2Default(void* ptr)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::setFocus();
}

void SettingsDialog46aa22_SetHiddenDefault(void* ptr, char hidden)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::setHidden(hidden != 0);
}

void SettingsDialog46aa22_SetStyleSheetDefault(void* ptr, struct Moc_PackedString styleSheet)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void SettingsDialog46aa22_SetWindowModifiedDefault(void* ptr, char vbo)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::setWindowModified(vbo != 0);
}

void SettingsDialog46aa22_SetWindowTitleDefault(void* ptr, struct Moc_PackedString vqs)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void SettingsDialog46aa22_ShowDefault(void* ptr)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::show();
}

void SettingsDialog46aa22_ShowFullScreenDefault(void* ptr)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::showFullScreen();
}

void SettingsDialog46aa22_ShowMaximizedDefault(void* ptr)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::showMaximized();
}

void SettingsDialog46aa22_ShowMinimizedDefault(void* ptr)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::showMinimized();
}

void SettingsDialog46aa22_ShowNormalDefault(void* ptr)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::showNormal();
}

void SettingsDialog46aa22_TabletEventDefault(void* ptr, void* event)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::tabletEvent(static_cast<QTabletEvent*>(event));
}

void SettingsDialog46aa22_UpdateDefault(void* ptr)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::update();
}

void SettingsDialog46aa22_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::updateMicroFocus();
}

void SettingsDialog46aa22_WheelEventDefault(void* ptr, void* event)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* SettingsDialog46aa22_PaintEngineDefault(void* ptr)
{
	return static_cast<SettingsDialog46aa22*>(ptr)->QDialog::paintEngine();
}

void* SettingsDialog46aa22_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<SettingsDialog46aa22*>(ptr)->QDialog::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char SettingsDialog46aa22_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<SettingsDialog46aa22*>(ptr)->QDialog::hasHeightForWidth();
}

int SettingsDialog46aa22_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<SettingsDialog46aa22*>(ptr)->QDialog::heightForWidth(w);
}

int SettingsDialog46aa22_MetricDefault(void* ptr, long long m)
{
	return static_cast<SettingsDialog46aa22*>(ptr)->QDialog::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void SettingsDialog46aa22_InitPainterDefault(void* ptr, void* painter)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::initPainter(static_cast<QPainter*>(painter));
}

void SettingsDialog46aa22_ChildEventDefault(void* ptr, void* event)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::childEvent(static_cast<QChildEvent*>(event));
}

void SettingsDialog46aa22_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void SettingsDialog46aa22_CustomEventDefault(void* ptr, void* event)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::customEvent(static_cast<QEvent*>(event));
}

void SettingsDialog46aa22_DeleteLaterDefault(void* ptr)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::deleteLater();
}

void SettingsDialog46aa22_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void SettingsDialog46aa22_TimerEventDefault(void* ptr, void* event)
{
	static_cast<SettingsDialog46aa22*>(ptr)->QDialog::timerEvent(static_cast<QTimerEvent*>(event));
}

int CurrentMediaPage46aa22_CurrentMediaPage46aa22_QRegisterMetaType()
{
	return qRegisterMetaType<CurrentMediaPage46aa22*>();
}

int CurrentMediaPage46aa22_CurrentMediaPage46aa22_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<CurrentMediaPage46aa22*>(const_cast<const char*>(typeName));
}

int CurrentMediaPage46aa22_CurrentMediaPage46aa22_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<CurrentMediaPage46aa22>();
#else
	return 0;
#endif
}

int CurrentMediaPage46aa22_CurrentMediaPage46aa22_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<CurrentMediaPage46aa22>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* CurrentMediaPage46aa22___addActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CurrentMediaPage46aa22___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* CurrentMediaPage46aa22___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* CurrentMediaPage46aa22___insertActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CurrentMediaPage46aa22___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* CurrentMediaPage46aa22___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* CurrentMediaPage46aa22___actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CurrentMediaPage46aa22___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* CurrentMediaPage46aa22___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* CurrentMediaPage46aa22___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CurrentMediaPage46aa22___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* CurrentMediaPage46aa22___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* CurrentMediaPage46aa22___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CurrentMediaPage46aa22___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CurrentMediaPage46aa22___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CurrentMediaPage46aa22___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CurrentMediaPage46aa22___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CurrentMediaPage46aa22___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CurrentMediaPage46aa22___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CurrentMediaPage46aa22___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CurrentMediaPage46aa22___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CurrentMediaPage46aa22___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CurrentMediaPage46aa22___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CurrentMediaPage46aa22___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* CurrentMediaPage46aa22_NewCurrentMediaPage(void* parent)
{
		return new CurrentMediaPage46aa22(static_cast<QWidget*>(parent));
}

void CurrentMediaPage46aa22_DestroyCurrentMediaPage(void* ptr)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->~CurrentMediaPage46aa22();
}

void CurrentMediaPage46aa22_DestroyCurrentMediaPageDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char CurrentMediaPage46aa22_EventDefault(void* ptr, void* ev)
{
	return static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::event(static_cast<QEvent*>(ev));
}

void CurrentMediaPage46aa22_ChangeEventDefault(void* ptr, void* ev)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::changeEvent(static_cast<QEvent*>(ev));
}

void CurrentMediaPage46aa22_KeyPressEventDefault(void* ptr, void* e)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::keyPressEvent(static_cast<QKeyEvent*>(e));
}

void CurrentMediaPage46aa22_PaintEventDefault(void* ptr, void* event)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::paintEvent(static_cast<QPaintEvent*>(event));
}

void CurrentMediaPage46aa22_ResizeEventDefault(void* ptr, void* e)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::resizeEvent(static_cast<QResizeEvent*>(e));
}

void CurrentMediaPage46aa22_SetCurrentIndexDefault(void* ptr, int index)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::setCurrentIndex(index);
}

void CurrentMediaPage46aa22_SetCurrentWidgetDefault(void* ptr, void* widget)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::setCurrentWidget(static_cast<QWidget*>(widget));
}

void CurrentMediaPage46aa22_ShowEventDefault(void* ptr, void* vqs)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::showEvent(static_cast<QShowEvent*>(vqs));
}

void CurrentMediaPage46aa22_TabInsertedDefault(void* ptr, int index)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::tabInserted(index);
}

void CurrentMediaPage46aa22_TabRemovedDefault(void* ptr, int index)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::tabRemoved(index);
}

void* CurrentMediaPage46aa22_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* CurrentMediaPage46aa22_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

char CurrentMediaPage46aa22_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::hasHeightForWidth();
}



int CurrentMediaPage46aa22_HeightForWidthDefault(void* ptr, int width)
{
	return static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::heightForWidth(width);
}

char CurrentMediaPage46aa22_CloseDefault(void* ptr)
{
	return static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::close();
}

char CurrentMediaPage46aa22_FocusNextPrevChildDefault(void* ptr, char next)
{
	return static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::focusNextPrevChild(next != 0);
}

char CurrentMediaPage46aa22_NativeEventDefault(void* ptr, void* eventType, void* message, long* result)
{
	return static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::nativeEvent(*static_cast<QByteArray*>(eventType), message, result);
}

void CurrentMediaPage46aa22_ActionEventDefault(void* ptr, void* event)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void CurrentMediaPage46aa22_CloseEventDefault(void* ptr, void* event)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void CurrentMediaPage46aa22_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void CurrentMediaPage46aa22_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void CurrentMediaPage46aa22_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void CurrentMediaPage46aa22_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void CurrentMediaPage46aa22_DropEventDefault(void* ptr, void* event)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::dropEvent(static_cast<QDropEvent*>(event));
}

void CurrentMediaPage46aa22_EnterEventDefault(void* ptr, void* event)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::enterEvent(static_cast<QEvent*>(event));
}

void CurrentMediaPage46aa22_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

void CurrentMediaPage46aa22_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void CurrentMediaPage46aa22_HideDefault(void* ptr)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::hide();
}

void CurrentMediaPage46aa22_HideEventDefault(void* ptr, void* event)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::hideEvent(static_cast<QHideEvent*>(event));
}

void CurrentMediaPage46aa22_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void CurrentMediaPage46aa22_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void CurrentMediaPage46aa22_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::leaveEvent(static_cast<QEvent*>(event));
}

void CurrentMediaPage46aa22_LowerDefault(void* ptr)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::lower();
}

void CurrentMediaPage46aa22_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void CurrentMediaPage46aa22_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void CurrentMediaPage46aa22_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void CurrentMediaPage46aa22_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void CurrentMediaPage46aa22_MoveEventDefault(void* ptr, void* event)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void CurrentMediaPage46aa22_RaiseDefault(void* ptr)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::raise();
}

void CurrentMediaPage46aa22_RepaintDefault(void* ptr)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::repaint();
}

void CurrentMediaPage46aa22_SetDisabledDefault(void* ptr, char disable)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::setDisabled(disable != 0);
}

void CurrentMediaPage46aa22_SetEnabledDefault(void* ptr, char vbo)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::setEnabled(vbo != 0);
}

void CurrentMediaPage46aa22_SetFocus2Default(void* ptr)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::setFocus();
}

void CurrentMediaPage46aa22_SetHiddenDefault(void* ptr, char hidden)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::setHidden(hidden != 0);
}

void CurrentMediaPage46aa22_SetStyleSheetDefault(void* ptr, struct Moc_PackedString styleSheet)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void CurrentMediaPage46aa22_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::setVisible(visible != 0);
}

void CurrentMediaPage46aa22_SetWindowModifiedDefault(void* ptr, char vbo)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::setWindowModified(vbo != 0);
}

void CurrentMediaPage46aa22_SetWindowTitleDefault(void* ptr, struct Moc_PackedString vqs)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void CurrentMediaPage46aa22_ShowDefault(void* ptr)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::show();
}

void CurrentMediaPage46aa22_ShowFullScreenDefault(void* ptr)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::showFullScreen();
}

void CurrentMediaPage46aa22_ShowMaximizedDefault(void* ptr)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::showMaximized();
}

void CurrentMediaPage46aa22_ShowMinimizedDefault(void* ptr)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::showMinimized();
}

void CurrentMediaPage46aa22_ShowNormalDefault(void* ptr)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::showNormal();
}

void CurrentMediaPage46aa22_TabletEventDefault(void* ptr, void* event)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void CurrentMediaPage46aa22_UpdateDefault(void* ptr)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::update();
}

void CurrentMediaPage46aa22_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::updateMicroFocus();
}

void CurrentMediaPage46aa22_WheelEventDefault(void* ptr, void* event)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* CurrentMediaPage46aa22_PaintEngineDefault(void* ptr)
{
	return static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::paintEngine();
}

void* CurrentMediaPage46aa22_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

int CurrentMediaPage46aa22_MetricDefault(void* ptr, long long m)
{
	return static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void CurrentMediaPage46aa22_InitPainterDefault(void* ptr, void* painter)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::initPainter(static_cast<QPainter*>(painter));
}

char CurrentMediaPage46aa22_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void CurrentMediaPage46aa22_ChildEventDefault(void* ptr, void* event)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::childEvent(static_cast<QChildEvent*>(event));
}

void CurrentMediaPage46aa22_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void CurrentMediaPage46aa22_CustomEventDefault(void* ptr, void* event)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::customEvent(static_cast<QEvent*>(event));
}

void CurrentMediaPage46aa22_DeleteLaterDefault(void* ptr)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::deleteLater();
}

void CurrentMediaPage46aa22_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void CurrentMediaPage46aa22_TimerEventDefault(void* ptr, void* event)
{
	static_cast<CurrentMediaPage46aa22*>(ptr)->QTabWidget::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
