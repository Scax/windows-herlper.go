package window

import "github.com/scax/windows-helper/win"

// GetWindowHandleFromPID return the first window for the pid
func GetWindowHandleFromPID(pid uint32) (uintptr, error) {

	wHwndCh := make(chan uintptr)
	errCh := make(chan error)

	go func() {

		_, err := win.EnumWindows(func(hwnd uintptr, repass interface{}) bool {

			_, pidToMatch, err := win.GetWindowThreadProcessID(hwnd)
			if err != nil {
				errCh <- err
				return false
			}

			if pid == uint32(pidToMatch) {
				wHwndCh <- hwnd
				return false
			}
			return true
		}, nil)

		defer close(errCh)
		defer close(wHwndCh)
		if err != nil {
			errCh <- err
		}
	}()

	select {
	case hwnd := <-wHwndCh:
		return hwnd, nil
	case err := <-errCh:
		return 0, err
	}
}
