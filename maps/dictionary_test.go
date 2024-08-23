package dictionary

import "testing"

func TestSearch(t *testing.T) {
	// Maps are never initialised with null e.g. var m map[string]string
	// The can be initialised with make(map[string]string), or per below
	dict := Dictionary{"test": "this is just a test"}

	t.Run("word exists", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("word no exist", func(t *testing.T) {
		_, err := dict.Search("potato")

		if err == nil {
			// Use fatal to stop further tests
			// Because the later call of err.Error() would give a null pointer
			// and cause a panic
			t.Fatal("expected an error but didn't get one")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}

		term := "new"
		def := "new def"

		dict.Add(term, def)

		assertDefinition(t, dict, term, def)
	})

	t.Run("existing word", func(t *testing.T) {
		term := "word"
		def := "definition"
		dict := Dictionary{term: def}

		err := dict.Add(term, "new def")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, term, def)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("word exists", func(t *testing.T) {
		term := "word"
		def := "definition"
		dict := Dictionary{term: def}

		err := dict.Update(term, "new definition")

		assertError(t, err, nil)
		assertDefinition(t, dict, term, "new definition")
	})

	t.Run("word does not exist", func(t *testing.T) {
		term := "word"
		def := "definition"
		dict := Dictionary{}

		err := dict.Update(term, def)

		assertError(t, err, ErrWordNoExists)
	})
}

func TestDelete(t *testing.T) {
	term := "term"
	definition := "definition"
	dict := Dictionary{term: definition}

	dict.Delete(term)

	_, err := dict.Search(term)

	assertError(t, err, ErrNotFound)
}

func assertDefinition(t testing.TB, dict Dictionary, term, want string) {
	t.Helper()

	got, err := dict.Search(term)

	if err != nil {
		t.Fatal("expected no error but got one", err)
	}

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q, wanted %q", got, want)
	}
}
