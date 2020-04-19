package win

import (
	"os/exec"
	"testing"
	"time"
)

func TestFindWindow(t *testing.T) {
	cmd := exec.Command("notepad.exe")
	defer func() {
		if cmd.Process != nil {
			cmd.Process.Kill()
		}
	}()
	cmd.Start()
	<-time.NewTimer(1 * time.Second).C

	hwnd, err := FindWindow("", "Untitled - Notepad")
	if err != nil {
		t.Fatal(err)
	}
	if hwnd == 0 {
		t.Log("window handle is 0")
		t.Failed()
	}
}
func TestGetWindowThreadProcessID(t *testing.T) {

	EnumWindows(func(hwnd uintptr, repass interface{}) bool {

		threadID, processID, err := GetWindowThreadProcessID(hwnd)
		t.Logf("ThreadID: %d, PID: %d, error:%v", threadID, processID, err)

		return false
	}, nil)

}
func TestEnumWindows(t *testing.T) {
	wantRepass := "test"
	EnumWindows(func(hwnd uintptr, repass interface{}) bool {

		t.Logf("goCallback, hwnd: %d, repass: %v", hwnd, repass)
		if repass != wantRepass {
			t.Logf("wrong repass want: %v != got: %v", wantRepass, repass)
			t.Failed()
			return false
		}
		return true
	}, wantRepass)

}
