package font

import (
	"math"
	"os"
)

func round(f float64) int {
	if f < 0 {
		return -int(math.Floor(-f + 0.5))
	}
	return int(math.Floor(f + 0.5))
}

// fileExist returns true if the specified normal file exists
func fileExist(filename string) (ok bool) {
	info, err := os.Stat(filename)
	if err == nil {
		if ^os.ModePerm&info.Mode() == 0 {
			ok = true
		}
	}
	return ok
}

// fileSize returns the size of the specified file; ok will be false
// if the file does not exist or is not an ordinary file
func fileSize(filename string) (size int64, ok bool) {
	info, err := os.Stat(filename)
	ok = err == nil
	if ok {
		size = info.Size()
	}
	return
}
