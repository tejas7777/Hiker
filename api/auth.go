package api

import (
	"encoding/json"
	"fmt"
	handler "hiker/handlers"
	payload "hiker/payloads"

	"github.com/labstack/echo/v4"
)

func Hello() {
	fmt.Println("Hello From Auth")
}

func Auth(c echo.Context) error {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.String(502, "Invalid payload")
	} else {
		//json_map has the JSON Payload decoded into a map

		if _, ok := json_map["username"]; ok {
			if _, ok := json_map["password"]; ok {

				username := json_map["username"].(string)
				password := json_map["password"].(string)

				reply, err := handler.CheckUsernameExist(username)
				fmt.Println("REPLY IS " + reply)

				if reply == "Set" {
					return c.String(409, "Already Registered")
				}

				if reply == "error" && err != nil {
					return c.NoContent(502)
				}

				key, err := handler.GetAPIKey(username, password)

				if err != nil {
					return c.NoContent(502)
				}

				resp := &payload.AuthPayload{Key: key}
				c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c.Response().WriteHeader(200)
				return json.NewEncoder(c.Response()).Encode(resp)

			} else {
				return c.String(400, "Missing password")
			}
		} else {
			if _, ok := json_map["password"]; ok {
				return c.String(400, "Missing username")
			} else {
				return c.String(400, "Missing username and password")
			}
		}

	}
}
