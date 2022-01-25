// Copyright 2017 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

//go:build windows && !arm && !arm64 && !linux
// +build windows,!arm,!arm64,!linux

// borrowed from https://github.com/tamird/cockroach/blob/12859da4c3068a61efb8dc157761b54ef4620402/pkg/util/log/stderr_redirect_windows.go

package main

import (
	"syscall"
)

var (
	kernel32         = syscall.MustLoadDLL("kernel32.dll")
	procSetStdHandle = kernel32.MustFindProc("SetStdHandle")
)

func dupFD(fd uintptr) (uintptr, error) {
	// Cribbed from https://github.com/golang/go/blob/go1.8/src/syscall/exec_windows.go#L303.
	p, err := syscall.GetCurrentProcess()
	if err != nil {
		return 0, err
	}
	var h syscall.Handle
	return uintptr(h), syscall.DuplicateHandle(p, syscall.Handle(fd), p, &h, 0, true, syscall.DUPLICATE_SAME_ACCESS)
}

func dupFD2(oldfd uintptr, newfd uintptr) error {
	// Cribbed from https://groups.google.com/d/msg/golang-nuts/fG8hEAs7ZXs/tahEOuCEPn0J.
	r0, _, e1 := syscall.Syscall(procSetStdHandle.Addr(), 2, oldfd, newfd, 0)
	if r0 == 0 {
		if e1 != 0 {
			return error(e1)
		}
		return syscall.EINVAL
	}
	return nil
}
