package settings

/*
#cgo CFLAGS: -fno-keep-inline-dllexport -O2 -Wall -W -Wextra -DUNICODE -D_UNICODE -DWIN32 -DMINGW_HAS_SECURE_API=1 -DQT_NEEDS_QMAIN -DQT_NO_DEBUG -DQT_WIDGETS_LIB -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -fno-keep-inline-dllexport -O2 -Wall -W -Wextra -fexceptions -mthreads -DUNICODE -D_UNICODE -DWIN32 -DMINGW_HAS_SECURE_API=1 -DQT_NEEDS_QMAIN -DQT_NO_DEBUG -DQT_WIDGETS_LIB -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -I../../hikari -I. -I../../../../Qt/Qt5.12.0/5.12.0/mingw64/include -I../../../../Qt/Qt5.12.0/5.12.0/mingw64/include/QtWidgets -I../../../../Qt/Qt5.12.0/5.12.0/mingw64/include/QtGui -I../../../../Qt/Qt5.12.0/5.12.0/mingw64/include/QtANGLE -I../../../../Qt/Qt5.12.0/5.12.0/mingw64/include/QtCore -Irelease -I/include -I../../../../Qt/Qt5.12.0/5.12.0/mingw64/mkspecs/win32-g++
#cgo LDFLAGS:        -Wl,-subsystem,windows -mthreads
#cgo LDFLAGS:        -LC:/Qt/Qt5.12.0/5.12.0/mingw64/lib -lQt5Widgets -lQt5Gui -lQt5Core  -lmingw32 -lqtmain -LC:/openssl/lib -LC:/Utils/my_sql/mysql-5.6.11-winx64/lib -LC:/Utils/postgresql/pgsql/lib -lshell32 -LC:/Qt/Qt5.12.0/5.12.0/mingw64/lib -lQt5Core
#cgo LDFLAGS: -Wl,--allow-multiple-definition
#cgo CFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
#cgo CXXFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
*/
import "C"
