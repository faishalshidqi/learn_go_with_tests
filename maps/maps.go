package maps

import (
	"errors"
)

type Dictionary map[string]string
type DictMethods interface {
	Search(key string) (string, error)
	Add(key, value string)
}

const (
	notFoundKeyError   = "could not find the word you were looking for"
	wordExistsErr      = "cannot add word because it already exists"
	wordNotExistsError = "cannot perform operation because word does not exist"
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	if v, ok := d[key]; ok {
		return v, nil
	} else {
		return "", errors.New(notFoundKeyError)
	}
}

func (d Dictionary) Add(key, value string) error {
	/*
		_, err := d.Search(key)
		switch err {
		case fmt.Errorf(notFoundKeyError):
			d[key] = value
		case nil:
			return errors.New(wordExistsErr)
		default:
			return err
		}
	*/
	if _, ok := d[key]; ok {
		return errors.New(wordExistsErr)
	}
	d[key] = value
	return nil
}

func (d Dictionary) Update(key, value string) error {
	if _, ok := d[key]; ok {
		d[key] = value
		return nil
	}
	return errors.New(wordNotExistsError)
}

func (d Dictionary) Delete(key string) error {
	if _, ok := d[key]; ok {
		delete(d, key)
		return nil
	}
	return errors.New(wordNotExistsError)
}
