package mydoc_test

import (
    "testing"
)

func TestSuccess(t *testing.T) {
	a := false
	if(a) {
		t.Log("TestXYZ is called");	
	}
	
}

func TestXYZFailed(t *testing.T) {
	b := false
	if(b) {
		t.Errorf("TestXYZ is called");	
	}
}