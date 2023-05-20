package scm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListBranch(t *testing.T) {
	if testing.Short() {
		t.Skip("skipped due to ci is seperated from DB")
	}

	branches, err := ListBranches(7, 22)
	assert.NotEmpty(t, branches)
	assert.NoError(t, err)
}
