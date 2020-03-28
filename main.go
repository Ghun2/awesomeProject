package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Dog struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Hamster struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
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


func addDog(c echo.Context) error {
	dog := Dog{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&dog)
	if err != nil {
		log.Printf("Failed processing addDog request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("this is your dog: %#v", dog)
	return c.String(http.StatusOK, "we got your dog!")
}

func addHamster(c echo.Context) error {
	hamster := Hamster{}

	err := c.Bind(&hamster)
	if err != nil {
		log.Printf("Failed processing addHamster request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("this is your hamster: %#v", hamster)
	return c.String(http.StatusOK, "we got your hamster!")
}


func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "ji hun you are on the secret admin main page!")
}

func mainCookie(c echo.Context) error {
	return c.String(http.StatusOK, "you are on the secret cookie page!")
}

func mainJwt(c echo.Context) error {
	user := c.Get("user")
	token := user.(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)

	log.Print("User Name: ", claims["name"], " User ID: ", claims["jti"])

	return c.String(http.StatusOK, "you are on the top secret jwt page!")
}

func login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	// check username and password against DB after hashing the password
	if username == "yelim" && password == "0526" {
		// this is the same
		//cookie := new(http.Cookie)
		cookie := &http.Cookie{}

		cookie.Name = "sessionID"
		cookie.Value = "some_string"
		cookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(cookie)

		// create jwt token
		token, err := createJwtToken()
		if err != nil {
			log.Print("Error Creating JWT token", err)
			return c.String(http.StatusInternalServerError, "someting went wrong")
		}

		// this is the same
		//cookie := new(http.Cookie)
		jwtCookie := &http.Cookie{}

		jwtCookie.Name = "JWTCookie"
		jwtCookie.Value = token
		jwtCookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(jwtCookie)


		return c.JSON(http.StatusOK, map[string] string {
			"message": "You were logged in!",
			"token": token,
		})
	}
	return c.String(http.StatusUnauthorized, "You username or password were wrong")
}

func createJwtToken() (string, error) {
	claims := JwtClaims{
		"yelim",
		jwt.StandardClaims{
			Id:        "main_user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := rawToken.SignedString([]byte("mySecret"))
	if err != nil {
		return "", err
	}

	return token, nil
}

///////////////////////////////////// middlewares ////////////////////////////////////

////////// 커스텀 서버 헤더 ///////////
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "BlueBot/1.0")
		c.Response().Header().Set("notReallyHeader", "thisHaveNoMearning")
		return next(c)
	}
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")
		if err != nil {
			if strings.Contains(err.Error(), "named cookie not present") {
				return c.String(http.StatusUnauthorized, "you dont have any cookie")
			}
			log.Print(err)
			return err
		}
		if cookie.Value == "some_string" {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "you dont have the right cookie, cookie")
	}
}

func main() {
	fmt.Println("Hello ji hun")
	fmt.Println(time.Now())

	e := echo.New()
	e.Use(ServerHeader)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:"[${time_rfc3339}]  ${status}  ${method}  ${host}${path} ${latency_human}" + "\n",
	}))

	adminGroup := e.Group("/admin")
	anmalGroup := e.Group("/animal")

	cookieGroup := e.Group("/cookie")
	jwtGroup := e.Group("/jwt")

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "static",
		HTML5: true,
	}))

	adminGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// check in the DB
		if username == "hun" && password == "3133" {
			return true, nil;
		}
		return false, nil;
	}))

	jwtGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey: []byte("mySecret"),
		TokenLookup: "cookie:JWTCookie",
	}))

	cookieGroup.Use(checkCookie)

	cookieGroup.GET("/main", mainCookie)
	adminGroup.GET("/main", mainAdmin)
	jwtGroup.GET("/main", mainJwt)

	e.GET("/hello", hello)
	e.GET("/login", login)

	anmalGroup.GET("/cats/:data", getCats)

	anmalGroup.GET("/cats", func (c echo.Context) error {
		return c.String(http.StatusOK, "Hel412lo, Wo123rld!")
	})

	anmalGroup.POST("/cats", addCat)
	anmalGroup.POST("/dogs", addDog)
	anmalGroup.POST("/hamsters", addHamster)

	e.Logger.Fatal(e.Start(":7777"))
}
