package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/labstack/echo/v4"
)

type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hel412lo, Wo123rld!")
}

func getCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("your cat name is %s\nand his type is %s\n", catName, catType))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": catName,
			"type": catType,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "you need to lets us know if you want json or string data",
	})

}


func addCat(c echo.Context) error {
	fmt.Println("Hello ji hun")

	if c.Request().Body != nil {
		return c.String(http.StatusOK, "we got your cat!")
	}

	cat := Cat{}

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &cat)

	if err != nil {
		log.Printf("Failed unmarshaling in addCats: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("this is your cat: %#v", cat)
	return c.String(http.StatusOK, "we got your cat!")

}


func main() {
	fmt.Println("Hello ji hun")

	e := echo.New()

	e.GET("/", hello)
	e.GET("/cats/:data", getCats)

	e.GET("/cats", func (c echo.Context) error {
		return c.String(http.StatusOK, "Hel412lo, Wo123rld!")
	})

	e.POST("/cats", addCat)

	e.Logger.Fatal(e.Start(":7777"))
}
