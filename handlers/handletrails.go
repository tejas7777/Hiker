package handlers

import (
	models "hiker/models"
	payloads "hiker/payloads"
)

func GetAllTrails() (*payloads.Trails, error) {

	s, err := models.GetAllTrails()

	if err != nil {

		return nil, err
	}

	trails := payloads.Trails{Trails: s}

	if err != nil {
		return nil, err
	}

	return &trails, err

}
