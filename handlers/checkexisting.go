package handlers

import (
	model "hiker/models"
)

func CheckUsernameExist(username string) (string, error) {

	key, err := HashKey(username + salt)

	if err != nil {
		return "error", err
	}

	khash, err := HashKey(key)

	if err != nil {
		return "error", err
	}

	err = model.GetHash(khash, "username")

	if err != nil {
		return "", err
	}

	return "Set", err
}

func Authenticate(key string) (bool, error) {
	khash, err := HashKey(key)
	if err != nil {
		return false, err
	}
	result, err := model.CheckKey(khash)

	if err != nil {
		return false, err
	}

	return result, nil

}
