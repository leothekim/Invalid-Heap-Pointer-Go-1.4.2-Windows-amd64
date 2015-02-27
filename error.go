package main

import (
	"net"
	"strings"
	"syscall"
)

const connectionResetErrno = 0x68

func IsConnectionResetErr(err error) bool {
	if err == nil {
		return false
	}
	operr, ok := err.(*net.OpError)
	if !ok {
		return false
	}

	errno, ok := operr.Err.(syscall.Errno)
	return (ok && (errno == connectionResetErrno)) || strings.Contains(err.Error(), "connection reset by peer")
}
