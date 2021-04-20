package handlers

import (
	model "hiker/models"
)

var salt = "LUXEMBORUGISH"

func GetAPIKey(username string, password string) (string, error) {
	key, err := HashKey(username + salt)

	if err != nil {
		return "", err
	}

	phash, err := HashKey(password)

	if err != nil {
		return "", err
	}

	khash, err := HashKey(key)

	if err != nil {
		return "", err
	}

	err = model.SetHash(khash, "username", username)
	if err != nil {
		return "", err
	}

	err = model.SetHash(khash, "password", phash)
	if err != nil {
		return "", err
	}

	return key, nil

}
