package hashtable_test

import (
	"testing"

	"github.com/victorprb/data-structures/hashtable"
)

func TestHashTable(t *testing.T) {
	ht := hashtable.New(7)

	ht.Add("fruit", "orange")
	if got := ht.Get("fruit"); got != "orange" {
		t.Errorf("got %v, want %v", got, "orange")
	}

	ht.Delete("fruit")
	if got := ht.Get("fruit"); got != nil {
		t.Errorf("got %v, want %v", got, nil)
	}

	t.Run("handle hash collisions", func(t *testing.T) {
		ht := hashtable.New(1)

		name := "john"
		age := 42

		ht.Add("name", name)
		ht.Add("age", age)

		if got := ht.Get("name"); got != name {
			t.Errorf("got %v, want %v", got, name)
		}

		if got := ht.Get("age"); got != age {
			t.Errorf("got %v, want %v", got, age)
		}
	})
}
