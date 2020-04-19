package proc

import (
	"errors"
	"unsafe"

	"golang.org/x/sys/windows"
)

type PID uint32

var ErrProcessNotFound = errors.New("process not found")

func GetProcId(procName string) (PID, error) {

	hSnap, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return 0, err
	}

	var procEntry windows.ProcessEntry32
	var exeFile string

	procEntry.Size = uint32(unsafe.Sizeof(procEntry))

	err = windows.Process32First(hSnap, &procEntry)
	if err != nil {
		return 0, err
	}
	for {

		exeFile = windows.UTF16ToString(procEntry.ExeFile[:])
		if exeFile == procName {
			return PID(procEntry.ProcessID), nil
		}

		err = windows.Process32Next(hSnap, &procEntry)
		if err != nil {
			if err == windows.Errno(18) {
				return 0, ErrProcessNotFound
			}
			return 0, err
		}
	}

	//return 0, ErrProcessNotFound

}
