// +build linux netbsd openbsd solaris freebsd dragonfly

package fdlimit

import (
	"github.com/pkg/errors"
	"golang.org/x/sys/unix"
)

func RaiseMaxFDs(max uint64) (uint64, error) {
	var rLimit unix.Rlimit
	if err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rLimit); err != nil {
		return 0, errors.Wrap(err, "error getting rlimit")
	}

	// Try to update the rLimit to the max allowance
	if max > rLimit.Max {
		max = rLimit.Max
	}
	if rLimit.Cur > max {
		return rLimit.Cur, nil
	}
	rLimit.Cur = max

	if err := unix.Setrlimit(unix.RLIMIT_NOFILE, &rLimit); err != nil {
		return 0, errors.Wrap(err, "error setting rlimit")
	}
	if err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rLimit); err != nil {
		return 0, errors.Wrap(err, "error getting rlimit")
	}
	return rLimit.Cur, nil
}

func CurrentFDs() (uint64, error) {
	var rLimit unix.Rlimit
	if err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rLimit); err != nil {
		return 0, errors.Wrap(err, "error getting rlimit")
	}
	return rLimit.Cur, nil
}

func MaxFDs() (uint64, error) {
	var rLimit unix.Rlimit
	if err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rLimit); err != nil {
		return 0, errors.Wrap(err, "error getting rlimit")
	}
	return rLimit.Max, nil
}
