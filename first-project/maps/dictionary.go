package maps

import "errors"

func Search(dictionary map[string]string, query string) (string, error) {
	value, ok := dictionary[query]
	if !ok {
		return "", errors.New("Key was not found")
	}

	return value, nil
}
