// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"internal/bytealg"
	"internal/syscall/unix"
)

func executable() (string, error) {
	var buf [1024]byte // PATH_MAX; the kernel caps the path at this length
	if err := unix.Getexecpath(buf[:]); err != nil {
		return "", NewSyscallError("getexecpath", err)
	}
	return string(buf[:bytealg.IndexByte(buf[:], 0)]), nil
}
