package dictionary

import "errors"

type Dictionary map[string]string

var ErrKeyNotFound = errors.New("Key could not be found.")

func (d Dictionary) Search(word string) (result string, err error) {
	result, found := d[word]

	if !found {
		err = ErrKeyNotFound
	}

	return
}