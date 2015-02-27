// +build windows

package main

import "github.com/winlabs/gowin32"

func SysDepTerm() {
	gowin32.UninitializeCOM()
}

