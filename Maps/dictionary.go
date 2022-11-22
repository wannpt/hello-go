package main

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound           = DictionaryErr("could not find the word you looking for")
	ErrWordExist          = DictionaryErr("word already exists")
	ErrDeleteWordNotFound = DictionaryErr("could not find the word you want to delete")
)

// error interface for error constant
func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, meaning string) error {

	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = meaning
	case nil:
		return ErrWordExist
	default:
		return err
	}

	return nil
}
func (d Dictionary) Update(word, meaning string) error {

	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = meaning
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		delete(d, word)
	case ErrNotFound:
		return ErrDeleteWordNotFound
	default:
		return err
	}

	return nil
}