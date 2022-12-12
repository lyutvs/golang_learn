package mydict

import (
	"errors"
)

type Dictionary map[string]string

var (
	errorNotFound   = errors.New("Not Found")
	errorCantUpdate = errors.New("Can't Update non-existing word")
	errorWordExists = errors.New("That word allredy exists")
)

func (dictionary Dictionary) Search(word string) (string, error) {
	value, exists := dictionary[word]

	if exists {
		return value, nil
	}
	return "", errorNotFound
}

func (dictionary Dictionary) Add(word, def string) error {
	_, err := dictionary.Search(word)

	switch err {
	case errorNotFound:
		dictionary[word] = def
	case nil:
		return errorWordExists
	}

	return nil
	//if err == errorNotFound {
	//	dictionary[word] = def
	//} else if err == nil {
	//	return errorWordExists
	//}
}

func (dictionary Dictionary) Update(word, definition string) error {
	_, err := dictionary.Search(word)

	switch err {
	case nil:
		dictionary[word] = definition
	case errorNotFound:
		return errorCantUpdate
	}
	return nil
}

func (dictionary Dictionary) Delete(word string) {
	delete(dictionary, word)

}
