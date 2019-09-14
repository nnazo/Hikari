

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QLayout>
#include <QMetaMethod>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QSettings>
#include <QString>
#include <QTimerEvent>
#include <QWidget>
#include <QWindow>

#ifdef QT_QML_LIB
	#include <QQmlEngine>
#endif


class Settingsc4adcb: public QSettings
{
Q_OBJECT
public:
	Settingsc4adcb(QObject *parent = Q_NULLPTR) : QSettings(parent) {qRegisterMetaType<quintptr>("quintptr");Settingsc4adcb_Settingsc4adcb_QRegisterMetaType();Settingsc4adcb_Settingsc4adcb_QRegisterMetaTypes();callbackSettingsc4adcb_Constructor(this);};
	Settingsc4adcb(QSettings::Format format, QSettings::Scope scope, const QString &organization, const QString &application = QString(), QObject *parent = Q_NULLPTR) : QSettings(format, scope, organization, application, parent) {qRegisterMetaType<quintptr>("quintptr");Settingsc4adcb_Settingsc4adcb_QRegisterMetaType();Settingsc4adcb_Settingsc4adcb_QRegisterMetaTypes();callbackSettingsc4adcb_Constructor(this);};
	Settingsc4adcb(QSettings::Scope scope, const QString &organization, const QString &application = QString(), QObject *parent = Q_NULLPTR) : QSettings(scope, organization, application, parent) {qRegisterMetaType<quintptr>("quintptr");Settingsc4adcb_Settingsc4adcb_QRegisterMetaType();Settingsc4adcb_Settingsc4adcb_QRegisterMetaTypes();callbackSettingsc4adcb_Constructor(this);};
	Settingsc4adcb(const QString &fileName, QSettings::Format format, QObject *parent = Q_NULLPTR) : QSettings(fileName, format, parent) {qRegisterMetaType<quintptr>("quintptr");Settingsc4adcb_Settingsc4adcb_QRegisterMetaType();Settingsc4adcb_Settingsc4adcb_QRegisterMetaTypes();callbackSettingsc4adcb_Constructor(this);};
	Settingsc4adcb(const QString &organization, const QString &application = QString(), QObject *parent = Q_NULLPTR) : QSettings(organization, application, parent) {qRegisterMetaType<quintptr>("quintptr");Settingsc4adcb_Settingsc4adcb_QRegisterMetaType();Settingsc4adcb_Settingsc4adcb_QRegisterMetaTypes();callbackSettingsc4adcb_Constructor(this);};
	 ~Settingsc4adcb() { callbackSettingsc4adcb_DestroySettings(this); };
	bool event(QEvent * event) { return callbackSettingsc4adcb_Event(this, event) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackSettingsc4adcb_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackSettingsc4adcb_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackSettingsc4adcb_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackSettingsc4adcb_CustomEvent(this, event); };
	void deleteLater() { callbackSettingsc4adcb_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackSettingsc4adcb_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackSettingsc4adcb_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackSettingsc4adcb_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackSettingsc4adcb_TimerEvent(this, event); };
signals:
public slots:
private:
};

Q_DECLARE_METATYPE(Settingsc4adcb*)


void Settingsc4adcb_Settingsc4adcb_QRegisterMetaTypes() {
}

int Settingsc4adcb_Settingsc4adcb_QRegisterMetaType()
{
	return qRegisterMetaType<Settingsc4adcb*>();
}

int Settingsc4adcb_Settingsc4adcb_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<Settingsc4adcb*>(const_cast<const char*>(typeName));
}

int Settingsc4adcb_Settingsc4adcb_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Settingsc4adcb>();
#else
	return 0;
#endif
}

int Settingsc4adcb_Settingsc4adcb_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Settingsc4adcb>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* Settingsc4adcb___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void Settingsc4adcb___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* Settingsc4adcb___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* Settingsc4adcb___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Settingsc4adcb___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Settingsc4adcb___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Settingsc4adcb___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Settingsc4adcb___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Settingsc4adcb___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Settingsc4adcb___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Settingsc4adcb___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Settingsc4adcb___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Settingsc4adcb___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Settingsc4adcb___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Settingsc4adcb___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* Settingsc4adcb_NewSettings5(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(static_cast<QWindow*>(parent));
	} else {
		return new Settingsc4adcb(static_cast<QObject*>(parent));
	}
}

void* Settingsc4adcb_NewSettings3(long long format, long long scope, struct Moc_PackedString organization, struct Moc_PackedString application, void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(static_cast<QSettings::Format>(format), static_cast<QSettings::Scope>(scope), QString::fromUtf8(organization.data, organization.len), QString::fromUtf8(application.data, application.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(static_cast<QSettings::Format>(format), static_cast<QSettings::Scope>(scope), QString::fromUtf8(organization.data, organization.len), QString::fromUtf8(application.data, application.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(static_cast<QSettings::Format>(format), static_cast<QSettings::Scope>(scope), QString::fromUtf8(organization.data, organization.len), QString::fromUtf8(application.data, application.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(static_cast<QSettings::Format>(format), static_cast<QSettings::Scope>(scope), QString::fromUtf8(organization.data, organization.len), QString::fromUtf8(application.data, application.len), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(static_cast<QSettings::Format>(format), static_cast<QSettings::Scope>(scope), QString::fromUtf8(organization.data, organization.len), QString::fromUtf8(application.data, application.len), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(static_cast<QSettings::Format>(format), static_cast<QSettings::Scope>(scope), QString::fromUtf8(organization.data, organization.len), QString::fromUtf8(application.data, application.len), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(static_cast<QSettings::Format>(format), static_cast<QSettings::Scope>(scope), QString::fromUtf8(organization.data, organization.len), QString::fromUtf8(application.data, application.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(static_cast<QSettings::Format>(format), static_cast<QSettings::Scope>(scope), QString::fromUtf8(organization.data, organization.len), QString::fromUtf8(application.data, application.len), static_cast<QWindow*>(parent));
	} else {
		return new Settingsc4adcb(static_cast<QSettings::Format>(format), static_cast<QSettings::Scope>(scope), QString::fromUtf8(organization.data, organization.len), QString::fromUtf8(application.data, application.len), static_cast<QObject*>(parent));
	}
}

void* Settingsc4adcb_NewSettings2(long long scope, struct Moc_PackedString organization, struct Moc_PackedString application, void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(static_cast<QSettings::Scope>(scope), QString::fromUtf8(organization.data, organization.len), QString::fromUtf8(application.data, application.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(static_cast<QSettings::Scope>(scope), QString::fromUtf8(organization.data, organization.len), QString::fromUtf8(application.data, application.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(static_cast<QSettings::Scope>(scope), QString::fromUtf8(organization.data, organization.len), QString::fromUtf8(application.data, application.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(static_cast<QSettings::Scope>(scope), QString::fromUtf8(organization.data, organization.len), QString::fromUtf8(application.data, application.len), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(static_cast<QSettings::Scope>(scope), QString::fromUtf8(organization.data, organization.len), QString::fromUtf8(application.data, application.len), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(static_cast<QSettings::Scope>(scope), QString::fromUtf8(organization.data, organization.len), QString::fromUtf8(application.data, application.len), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(static_cast<QSettings::Scope>(scope), QString::fromUtf8(organization.data, organization.len), QString::fromUtf8(application.data, application.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(static_cast<QSettings::Scope>(scope), QString::fromUtf8(organization.data, organization.len), QString::fromUtf8(application.data, application.len), static_cast<QWindow*>(parent));
	} else {
		return new Settingsc4adcb(static_cast<QSettings::Scope>(scope), QString::fromUtf8(organization.data, organization.len), QString::fromUtf8(application.data, application.len), static_cast<QObject*>(parent));
	}
}

void* Settingsc4adcb_NewSettings4(struct Moc_PackedString fileName, long long format, void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(QString::fromUtf8(fileName.data, fileName.len), static_cast<QSettings::Format>(format), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(QString::fromUtf8(fileName.data, fileName.len), static_cast<QSettings::Format>(format), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(QString::fromUtf8(fileName.data, fileName.len), static_cast<QSettings::Format>(format), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(QString::fromUtf8(fileName.data, fileName.len), static_cast<QSettings::Format>(format), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(QString::fromUtf8(fileName.data, fileName.len), static_cast<QSettings::Format>(format), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(QString::fromUtf8(fileName.data, fileName.len), static_cast<QSettings::Format>(format), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(QString::fromUtf8(fileName.data, fileName.len), static_cast<QSettings::Format>(format), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(QString::fromUtf8(fileName.data, fileName.len), static_cast<QSettings::Format>(format), static_cast<QWindow*>(parent));
	} else {
		return new Settingsc4adcb(QString::fromUtf8(fileName.data, fileName.len), static_cast<QSettings::Format>(format), static_cast<QObject*>(parent));
	}
}

void* Settingsc4adcb_NewSettings(struct Moc_PackedString organization, struct Moc_PackedString application, void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(QString::fromUtf8(organization.data, organization.len), QString::fromUtf8(application.data, application.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(QString::fromUtf8(organization.data, organization.len), QString::fromUtf8(application.data, application.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(QString::fromUtf8(organization.data, organization.len), QString::fromUtf8(application.data, application.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(QString::fromUtf8(organization.data, organization.len), QString::fromUtf8(application.data, application.len), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(QString::fromUtf8(organization.data, organization.len), QString::fromUtf8(application.data, application.len), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(QString::fromUtf8(organization.data, organization.len), QString::fromUtf8(application.data, application.len), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(QString::fromUtf8(organization.data, organization.len), QString::fromUtf8(application.data, application.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new Settingsc4adcb(QString::fromUtf8(organization.data, organization.len), QString::fromUtf8(application.data, application.len), static_cast<QWindow*>(parent));
	} else {
		return new Settingsc4adcb(QString::fromUtf8(organization.data, organization.len), QString::fromUtf8(application.data, application.len), static_cast<QObject*>(parent));
	}
}

void Settingsc4adcb_DestroySettings(void* ptr)
{
	static_cast<Settingsc4adcb*>(ptr)->~Settingsc4adcb();
}

void Settingsc4adcb_DestroySettingsDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char Settingsc4adcb_EventDefault(void* ptr, void* event)
{
	return static_cast<Settingsc4adcb*>(ptr)->QSettings::event(static_cast<QEvent*>(event));
}



char Settingsc4adcb_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<Settingsc4adcb*>(ptr)->QSettings::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void Settingsc4adcb_ChildEventDefault(void* ptr, void* event)
{
	static_cast<Settingsc4adcb*>(ptr)->QSettings::childEvent(static_cast<QChildEvent*>(event));
}

void Settingsc4adcb_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Settingsc4adcb*>(ptr)->QSettings::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void Settingsc4adcb_CustomEventDefault(void* ptr, void* event)
{
	static_cast<Settingsc4adcb*>(ptr)->QSettings::customEvent(static_cast<QEvent*>(event));
}

void Settingsc4adcb_DeleteLaterDefault(void* ptr)
{
	static_cast<Settingsc4adcb*>(ptr)->QSettings::deleteLater();
}

void Settingsc4adcb_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Settingsc4adcb*>(ptr)->QSettings::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void Settingsc4adcb_TimerEventDefault(void* ptr, void* event)
{
	static_cast<Settingsc4adcb*>(ptr)->QSettings::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
