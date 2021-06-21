package history

import (
	"testing"
)

func TestHistoryLoad(t *testing.T) {

	// this should succeed
	h, err := Load("../testdata/history/history.txt")
	if err != nil {
		t.Errorf("got error when attempting to load history struct: %s", err.Error())
	}

	if len(h.Entries) != 0 {
		t.Errorf("history entries count is %d instead of 0", len(h.Entries))
	}

	// this should fail gracefully
	_, fail := Load("../testdata/foobar/history.txt")
	if fail == nil {
		t.Errorf("expected to fail history load with an invalid path error. got success")
	}
}

func TestHistoryAccess(t *testing.T) {
	h, _ := Load("../testdata/history/history.txt")

	var lookup []HistoryEntry

	lookup = h.Find("foobar")
	if len(lookup) != 0 {
		t.Errorf("Expected no found history items for lookup foobar. Found %d", len(lookup))
	}

	err := h.Append("foobar fizz buzz bin bazz")
	if err != nil {
		t.Errorf("got error while appending value to history: %s", err.Error())
	}

	lookup = h.Find("foobar")
	if len(lookup) != 1 {
		t.Errorf("expected post-append foobar lookup to equal 1. actual equal is %d", len(lookup))
	}

}
