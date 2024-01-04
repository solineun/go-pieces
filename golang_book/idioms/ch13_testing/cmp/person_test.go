package cmp

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestCreatePerson(t *testing.T) {
	comparer := cmp.Comparer(func (x, y Person) bool {
		return x.Name == y.Name && x.Age == y.Age
	})
	expected := Person{
		Name: "D",
		Age: 3,
		DateAdded: time.Now(),
	}
	result := CreatePerson("D", 3)
	if diff := cmp.Diff(expected, result, comparer); diff != "" {
		t.Error(diff)
	}
}