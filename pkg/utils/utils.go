package utils

import (
	"path/filepath"
	"strings"
)

func FileNameWithoutExtension(name string) string {
	return strings.TrimSuffix(filepath.Base(name), filepath.Ext(name))
}
