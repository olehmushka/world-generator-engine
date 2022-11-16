package tools

import (
	"fmt"
	"strings"
)

func PreparePath(path, requiredSuffix string) string {
	path = removeLastSlash(path)
	if strings.HasSuffix(path, requiredSuffix) {
		return path + "/"
	}

	return fmt.Sprintf("%s/%s/", path, requiredSuffix)
}

func removeLastSlash(path string) string {
	if string(path[len(path)-1]) == "/" {
		return removeLastSlash(path[:len(path)-1])
	}

	return path
}
