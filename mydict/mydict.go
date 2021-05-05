package mydict

import "errors"

type Dictionary map[string]string

var errNotFound = errors.New("Not Found...")
var errWordExists = errors.New("Word Exists...")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	// fmt.Println(i, ok)
	if exists {
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
	// if err == errNotFound {
	// 	d[word] = def
	// } else if err == nil {
	// 	return errWordExists
	// }
	// return nil
}
