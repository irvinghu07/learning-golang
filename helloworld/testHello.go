package main

import "testing"

func test(t *testing.T){
	if Fuck() != 31{
		t.errorf("Fuck error")
	}
}