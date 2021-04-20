package handlers

import (
	"crypto/sha1"
	"encoding/hex"
)

func HashKey(key string) (string, error) {

	h := sha1.New()
	h.Write([]byte(key))
	keyhash := hex.EncodeToString(h.Sum(nil))

	return keyhash, nil

}
