package utils

import (
	"strconv"
	"strings"
)

func DockerImageTagBuild(buildId uint) string {
	exceptedLen := 5
	buildIdStr := strconv.Itoa(int(buildId))
	if len(buildIdStr) >= exceptedLen {
		return buildIdStr
	}
	padding := strings.Repeat("0", exceptedLen-len(buildIdStr))
	return padding + buildIdStr
}
