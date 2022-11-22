package main

import "testing"

func TestSearch(t *testing.T){
	dictionary := Dictionary{"test": "this is just a test"}


	t.Run("known word", func(t *testing.T) {
		got,_ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T){

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		meaning := "this is just a test"
		dictionary.Add(word, meaning)

		assertDefinition(t, dictionary, word, meaning)
	})

	t.Run("known word", func(t *testing.T) {
		
		word := "test"
		meaning := "this is just a test"
		dictionary := Dictionary{word: meaning}
		err := dictionary.Add(word, meaning)

		assertError(t, err, ErrWordExist)
		assertDefinition(t, dictionary, word, meaning)
	})
}

func TestUpdate(t *testing.T){
	
	t.Run("update exist word", func(t *testing.T){
		word := "test"
		meaning := "this is a new meaning of test"
		dictionary := Dictionary{word: "this is just a test"}
		dictionary.Update(word, meaning)

		assertDefinition(t, dictionary, word, meaning)
	})

	t.Run("update unknown word", func(t *testing.T) {
		word := "test"
		meaning := "this is a new meaning of test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, meaning)

		assertError(t, err, ErrNotFound)
	})

}

func TestDelete(t *testing.T){

	t.Run("delete exist word", func(t *testing.T){
		word := "test"
		dictionary := Dictionary{word:"this is just a test"}
		dictionary.Delete(word)
	
		_, err := dictionary.Search(word)
		
		assertError(t, err, ErrNotFound)
	})

	t.Run("delete unknown word", func(t *testing.T){
		word := "test"
		dictionary := Dictionary{}
		err := dictionary.Delete(word)

		assertError(t, err, ErrDeleteWordNotFound)
	})
	
}

func assertDefinition (t testing.TB, d Dictionary, word, meaning string) {
	t.Helper()
	got, _ := d.Search(word)

	if got != meaning {
		t.Errorf("got %q, want %q", got, meaning)
	} 
}

func assertStrings (t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q | want %q", got, want)
	}
}

func assertError (t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q, want %q", got, want)
	}
}