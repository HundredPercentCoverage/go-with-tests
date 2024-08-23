package dictionary

type Dictionary map[string]string

// Creating a custom error - first declare type
type DictionaryErr string

// Then declare actual errors - const to make them immutable and reusable(?)
const (
	ErrNotFound     = DictionaryErr("term not found")
	ErrWordExists   = DictionaryErr("term already exists")
	ErrWordNoExists = DictionaryErr("term does not exist")
)

// Then implement error interface by implementing Error()
func (e DictionaryErr) Error() string {
	return string(e)
}

// Maps are already (sort of) reference types, so no
// pointer receiver.
// They are pointers to runtime types.
// https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it
func (dict Dictionary) Search(term string) (string, error) {
	result, ok := dict[term]

	if !ok {
		return "", ErrNotFound
	}

	return result, nil
}

func (dict Dictionary) Add(term, definition string) error {
	_, err := dict.Search(term)

	// Using a switch is good in case some unknown error is returned by Search()
	switch err {
	case ErrNotFound:
		dict[term] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (dict Dictionary) Update(term, definition string) error {
	_, err := dict.Search(term)

	switch err {
	case ErrNotFound:
		// Could reuse ErrNotFound here
		// but better to use specific errors
		return ErrWordNoExists
	case nil:
		dict[term] = definition
	default:
		return err
	}

	return nil
}

// Don't need to return an error - deleting anything
// that doesn't exist has no effect
func (dict Dictionary) Delete(term string) {
	delete(dict, term)
}
