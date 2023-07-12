package _map

import (
	"fmt"
	"testing"
)

func TestDictionary_Search(t *testing.T) {
	var dictionary = Dictionary{"test": "test"}

	t.Run("found the word", func(t *testing.T) {
		word, err := dictionary.Search("test")
		if err != nil {
			t.Errorf("not found the word")
		}
		fmt.Println(word)
	})

	t.Run("not found the word", func(t *testing.T) {
		_, err := dictionary.Search("not found test")
		if err != ErrNotFound {
			t.Errorf("the word ")
		}
	})
}
