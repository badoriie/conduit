//go:build linux && 386

// iSH's SIGURG emulation conflicts with Go's async goroutine preemption,
// causing a "fatal: bad g in signal handler" crash. Disable it at build time.
//go:debug asyncpreemptoff=1

package main
