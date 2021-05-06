package mydict

import (
	"errors"
)

type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found...")
	errWordExists = errors.New("Word Exists...")
	errCantUpdate = errors.New("Can't update non-existing word")
	errCantDelete = errors.New("Can't delete non-existing word")
)

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

func (d Dictionary) Update(word string, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
	// if err != nil {
	// 	return errNotFound
	// }
	// fmt.Println(word, "got successfully updated to ", def)
	// d[word] = def
	// return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		return errCantDelete
	}
	delete(d, word)
	return nil
}
