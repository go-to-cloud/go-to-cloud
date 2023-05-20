package repositories

import "testing"

func TestHistory(t *testing.T) {
	if testing.Short() {
		t.Skip("skipped due to ci is seperated from DB")
	}

	Deployed(1, 1)
}
