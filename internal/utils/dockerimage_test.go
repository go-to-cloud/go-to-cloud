package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDockerImageTagBuild(t *testing.T) {
	testCases := []struct {
		buildId  uint
		expected string
	}{
		{1, "00001"},
		{10, "00010"},
		{123, "00123"},
		{12345, "12345"},
	}

	for _, tc := range testCases {
		result := DockerImageTagBuild(tc.buildId)
		assert.Equal(t, tc.expected, result)
	}
}
