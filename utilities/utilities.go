package utilities

import "path/filepath"

func GetWriteLocation(filename string, dir string, vol string) string {
	return filepath.Join(vol, dir, filename)
}
