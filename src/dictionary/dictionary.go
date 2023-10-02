package dictionary

type DictionaryErr string
type Dictionary map[string]string

const (
	ErrKeyExists   = DictionaryErr("Key already exists.")
	ErrKeyNotFound = DictionaryErr("Key could not be found.")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Update(word, definition string) (err error) {
	_, searchErr := d.Search(word)

	switch searchErr {
	case ErrKeyNotFound:
		err = ErrKeyNotFound
	case nil:
		d[word] = definition
	}

	return
}

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

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
