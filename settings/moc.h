

#pragma once

#ifndef GO_MOC_c4adcb_H
#define GO_MOC_c4adcb_H

#include <stdint.h>

#ifdef __cplusplus
class Settingsc4adcb;
void Settingsc4adcb_Settingsc4adcb_QRegisterMetaTypes();
extern "C" {
#endif

struct Moc_PackedString { char* data; long long len; };
struct Moc_PackedList { void* data; long long len; };
int Settingsc4adcb_Settingsc4adcb_QRegisterMetaType();
int Settingsc4adcb_Settingsc4adcb_QRegisterMetaType2(char* typeName);
int Settingsc4adcb_Settingsc4adcb_QmlRegisterType();
int Settingsc4adcb_Settingsc4adcb_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* Settingsc4adcb___dynamicPropertyNames_atList(void* ptr, int i);
void Settingsc4adcb___dynamicPropertyNames_setList(void* ptr, void* i);
void* Settingsc4adcb___dynamicPropertyNames_newList(void* ptr);
void* Settingsc4adcb___findChildren_atList2(void* ptr, int i);
void Settingsc4adcb___findChildren_setList2(void* ptr, void* i);
void* Settingsc4adcb___findChildren_newList2(void* ptr);
void* Settingsc4adcb___findChildren_atList3(void* ptr, int i);
void Settingsc4adcb___findChildren_setList3(void* ptr, void* i);
void* Settingsc4adcb___findChildren_newList3(void* ptr);
void* Settingsc4adcb___findChildren_atList(void* ptr, int i);
void Settingsc4adcb___findChildren_setList(void* ptr, void* i);
void* Settingsc4adcb___findChildren_newList(void* ptr);
void* Settingsc4adcb___children_atList(void* ptr, int i);
void Settingsc4adcb___children_setList(void* ptr, void* i);
void* Settingsc4adcb___children_newList(void* ptr);
void* Settingsc4adcb_NewSettings5(void* parent);
void* Settingsc4adcb_NewSettings3(long long format, long long scope, struct Moc_PackedString organization, struct Moc_PackedString application, void* parent);
void* Settingsc4adcb_NewSettings2(long long scope, struct Moc_PackedString organization, struct Moc_PackedString application, void* parent);
void* Settingsc4adcb_NewSettings4(struct Moc_PackedString fileName, long long format, void* parent);
void* Settingsc4adcb_NewSettings(struct Moc_PackedString organization, struct Moc_PackedString application, void* parent);
void Settingsc4adcb_DestroySettings(void* ptr);
void Settingsc4adcb_DestroySettingsDefault(void* ptr);
char Settingsc4adcb_EventDefault(void* ptr, void* event);
;
char Settingsc4adcb_EventFilterDefault(void* ptr, void* watched, void* event);
void Settingsc4adcb_ChildEventDefault(void* ptr, void* event);
void Settingsc4adcb_ConnectNotifyDefault(void* ptr, void* sign);
void Settingsc4adcb_CustomEventDefault(void* ptr, void* event);
void Settingsc4adcb_DeleteLaterDefault(void* ptr);
void Settingsc4adcb_DisconnectNotifyDefault(void* ptr, void* sign);
void Settingsc4adcb_TimerEventDefault(void* ptr, void* event);

#ifdef __cplusplus
}
#endif

#endif