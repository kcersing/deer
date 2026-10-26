//go:build windows

package mtl

import "os"

// getSnapshotSignal 在 Windows 下退化为 os.Interrupt（Ctrl+C）。
// Windows 不真正支持 SIGUSR1，且能可靠接收的信号有限。
func getSnapshotSignal() os.Signal {
	return os.Interrupt
}
