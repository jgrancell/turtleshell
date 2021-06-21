package variables

import "testing"

func TestVariables(t *testing.T) {
	v := Load()

	val, ok := v.Get("FOOBAR")
	// We should have failures here
	if ok || val != "" {
		t.Errorf("FOOBAR should be unset, found %s", val)
	}

	v.Set("FOOBAR", "fizzbuzz")
	val, ok = v.Get("FOOBAR")
	// We should have successes here
	if !ok || val != "fizzbuzz" {
		t.Errorf("Expecting FOOBAR to equal fizzbuzz, got %s", val)
	}

	v.Set("FOOBAR", "slartibartfast")
	val, ok = v.Get("FOOBAR")
	// We should have successes here
	if !ok || val != "slartibartfast" {
		t.Errorf("Expecting FOOBAR to equal slartibartfast, got %s", val)
	}
}
