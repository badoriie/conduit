//go:build linux && 386

package cmd

import "runtime"

func init() {
	// iSH (x86 emulator on iOS) reports the host iPhone's CPU count via
	// sched_getaffinity, causing Go to spawn 6-8 OS threads all running inside
	// iSH's single emulation loop. Forcing GOMAXPROCS=1 eliminates that overhead.
	runtime.GOMAXPROCS(1)
}
