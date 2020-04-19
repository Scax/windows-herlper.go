package window

import (
	"fmt"
	"os/exec"
	"testing"
	"time"

	"github.com/scax/windows-helper/win"
)

func TestGetWindowHandleFromPID(t *testing.T) {

	cmd := exec.Command("notepad.exe")
	defer func() {
		if cmd.Process != nil {
			cmd.Process.Kill()
		}
	}()
	cmd.Start()
	<-time.NewTimer(1 * time.Second).C

	pid := cmd.Process.Pid
	t.Log("PID: ", pid)

	hwnd, err := GetWindowHandleFromPID(uint32(pid))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("GetWindowHandleFromPID:", hwnd)
	if err != nil {
		t.Fatal(err)
	}
	fHwnd, err := win.FindWindow("", "Untitled - Notepad")
	t.Log("FindWind:", fHwnd)
	if hwnd != fHwnd {
		t.Logf("%d != %d", hwnd, fHwnd)
		t.Failed()
	}

}
