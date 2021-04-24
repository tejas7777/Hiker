package api

import (
	"encoding/json"
	"fmt"
	handler "hiker/handlers"

	"github.com/labstack/echo/v4"
)

func Trails(c echo.Context) error {

	key := c.QueryParam("key")

	if key == "" {
		return c.NoContent(400)
	}

	//key authentication part

	res, err := handler.Authenticate(key)

	if err != nil {
		fmt.Errorf(err.Error())
		return c.NoContent(500)
	}

	if res == false {
		return c.NoContent(403)
	}

	//key is authenticated now, we can proceed
	payload, err := handler.GetAllTrails()

	if err != nil {
		panic(err)
		return c.NoContent(502)
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(200)
	return json.NewEncoder(c.Response()).Encode(payload)

}
