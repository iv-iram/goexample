package hello_world

import (	
	"testing"
)

func TestHello_world(t *testing.T){
	got := Hello_world()
	expected :="hello world"

	if got != expected {
		t.Errorf("got %q, wanted %q", got, expected)
	}
}

