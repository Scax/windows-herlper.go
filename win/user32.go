package win

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

var user32 = windows.NewLazySystemDLL("user32.dll")

var (
	pEnumWindows              = user32.NewProc("EnumWindows")
	pFindWindowW              = user32.NewProc("FindWindowW")
	pGetWindowThreadProcessID = user32.NewProc("GetWindowThreadProcessId")
)

func FindWindow(class, title string) (uintptr, error) {

	var pClass uintptr
	var pTitle uintptr
	if class != "" {
		upClass, err := windows.UTF16PtrFromString(class)
		if err != nil {
			return 0, err
		}
		pClass = uintptr(unsafe.Pointer(upClass))

	}
	if title != "" {
		upTitle, err := windows.UTF16PtrFromString(title)
		if err != nil {
			return 0, err
		}
		pTitle = uintptr(unsafe.Pointer(upTitle))
	}
	r0, _, err := pFindWindowW.Call(pClass, pTitle)

	if err != windows.Errno(0) {
		return 0, err
	}
	return r0, nil

}

func GetWindowThreadProcessID(windowHwnd uintptr) (uintptr, uintptr, error) {
	var pid uintptr
	r0, _, err := pGetWindowThreadProcessID.Call(windowHwnd, uintptr(unsafe.Pointer(&pid)))

	if err != windows.Errno(0) {
		return 0, 0, err
	}
	return r0, pid, nil

}

func EnumWindows(callback func(hwnd uintptr, repass interface{}) bool, repass interface{}) (bool, error) {

	goCallback := func(hwnd uintptr, uRepass uintptr) uintptr {
		retBool := callback(hwnd, *(*interface{})(unsafe.Pointer(uRepass)))

		if retBool {
			return 1
		}
		return 0

	}

	wCallback := windows.NewCallback(goCallback)

	r0, _, err := pEnumWindows.Call(wCallback, uintptr(unsafe.Pointer(&repass)))

	if err != windows.Errno(0) {
		return false, err
	}
	return r0 != 0, nil
}
