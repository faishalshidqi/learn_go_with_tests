package maps

import (
	"testing"
)

const (
	ErrNotFound      = "could not find the word you were looking for"
	ErrWordExists    = "cannot add word because it already exists"
	ErrWordNotExists = "cannot perform operation because word does not exist"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is a test"}
	t.Run("should find the key", func(t *testing.T) {
		got, err := dict.Search("test")
		want := "this is a test"
		assertNoError(t, err, "")
		assertEqual(t, got, want)
	})
	t.Run("should return an error when key not found", func(t *testing.T) {
		_, err := dict.Search("nope")
		want := ErrNotFound
		assertError(t, err, want)
		assertEqual(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("should add a new entry", func(t *testing.T) {
		dict := Dictionary{}
		key, value := "test", "this is a test"
		err := dict.Add(key, value)

		want := "this is a test"
		assertNoError(t, err, "should find added word")
		assertMap(t, dict, key, want)
	})
	t.Run("should return error when key already exists", func(t *testing.T) {
		key, value := "test", "this is a test"
		dict := Dictionary{}
		err := dict.Add(key, value)
		newVal := "another test"
		err = dict.Add(key, newVal)
		assertError(t, err, ErrWordExists)
		assertMap(t, dict, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("should update a word", func(t *testing.T) {
		key := "test"
		value := "this is a test"
		dict := Dictionary{}
		err := dict.Add(key, value)
		assertNoError(t, err, "")
		newVal := "update test"
		err = dict.Update(key, newVal)
		assertNoError(t, err, "")
		assertMap(t, dict, key, newVal)
	})
	t.Run("should return an error when key not found", func(t *testing.T) {
		dict := Dictionary{}
		key := "test"
		newVal := "update test"
		err := dict.Update(key, newVal)
		assertError(t, err, ErrWordNotExists)
	})

}

func TestDelete(t *testing.T) {
	t.Run("should delete a word", func(t *testing.T) {
		key := "test"
		value := "this is a test"
		dict := Dictionary{}
		err := dict.Add(key, value)
		assertNoError(t, err, "")
		dict.Delete(key)
		_, err = dict.Search(key)
		assertError(t, err, ErrNotFound)
	})
	t.Run("should return an error when key not found", func(t *testing.T) {
		key := "test"
		dict := Dictionary{}
		err := dict.Delete(key)
		assertError(t, err, ErrWordNotExists)
	})
}

func assertNoError(t *testing.T, err error, msg string) {
	t.Helper()
	if err != nil {
		if msg == "" {
			msg = "got an error but didn't want one"
		}
		t.Fatalf("%v, %v", msg, err)
	}
}

func assertError(t *testing.T, err error, want string) {
	t.Helper()
	if err == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if err.Error() != want {
		t.Errorf("got \"%s\", want \"%s\"", err, want)
	}
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	t.Helper()
	if a != b {
		t.Errorf("got '%v' want '%v'", a, b)
	}
}

func assertMap(t *testing.T, dictionary Dictionary, key, value string) {
	t.Helper()
	got, err := dictionary.Search(key)
	assertNoError(t, err, "")
	assertEqual(t, got, value)
}
