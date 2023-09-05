package mymap

import (
	"reflect"
	"strings"
	"testing"
)

func TestWordCount(t *testing.T) {
	got := WordCount(strings.Fields("one two two three three three"))
	want := map[string]int {
		"one" : 1,
		"two" : 2,
		"three" : 3,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}