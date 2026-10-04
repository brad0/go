// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import (
	"internal/abi"
	"unsafe"
)

//go:cgo_import_dynamic libc_getexecpath getexecpath "libc.so"

func libc_getexecpath_trampoline()

// Getexecpath copies the absolute, canonicalized pathname of the
// running executable into buf as a NUL-terminated string.
func Getexecpath(buf []byte) error {
	_, _, errno := syscall_syscall(abi.FuncPCABI0(libc_getexecpath_trampoline),
		uintptr(unsafe.Pointer(unsafe.SliceData(buf))), uintptr(len(buf)), 0)
	if errno != 0 {
		return errno
	}
	return nil
}
