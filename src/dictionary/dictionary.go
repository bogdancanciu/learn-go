package dictionary

import "errors"

type Dictionary map[string]string

var (
	ErrKeyExists   = errors.New("Key already exists.")
	ErrKeyNotFound = errors.New("Key could not be found.")
)

func (d Dictionary) Add(word, definition string) (err error) {
	_, searchErr := d.Search(word)

	switch searchErr {
	case ErrKeyNotFound:
		d[word] = definition
	case nil:
		err = ErrKeyExists
	}
	
	return
}

func (d Dictionary) Search(word string) (result string, err error) {
	result, found := d[word]

	if !found {
		err = ErrKeyNotFound
	}

	return
}