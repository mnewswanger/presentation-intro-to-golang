package main

import "testing"

func TestAlwaysSucceeds(t *testing.T) {
	if false {
		t.Errorf("True is not %s!", "true")
	}
}

func TestAlwaysFails(t *testing.T) {
	if true {
		t.Errorf("True is not %s!", "false")
	}
}
