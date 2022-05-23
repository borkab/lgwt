package main

type Dictionary map[string]string

type DictionaryErr string

const (
	ErrNotFound   = DictionaryErr("couldn't find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because already exists")
)

func (e DictionaryErr) Error() string {
	return string(e)
} //We made the errors constant; this required us to create our own
//DictionaryErr type which implements the error interface.

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
