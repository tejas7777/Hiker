package test

import (
	"fmt"
	models "hiker/models"
)

func TestConnection() {
	_, err := models.GetMongoClient()

	defer models.Disconnect()

	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("Connected!")
	}
}
