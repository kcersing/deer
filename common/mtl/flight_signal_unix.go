//go:build !windows

package mtl

import (
	"os"
	"syscall"
)

// getSnapshotSignal 返回用于触发 trace 快照的信号。
// Unix 下用 SIGUSR1，不干扰 Ctrl+C(SIGINT) 的退出语义。
func getSnapshotSignal() os.Signal {
	return syscall.SIGUSR1
}
