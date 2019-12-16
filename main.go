package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	user := User{
		Store: map[string]*Point{},
	}

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.Use(&rest.CorsMiddleware{
		OriginValidator: func(origin string, request *rest.Request) bool {
			return true
		},
	})
	router, err := rest.MakeRouter(
		rest.Post("/user/point", user.PostPoint),
	)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Server started.")
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

type Point struct {
}

type User struct {
	sync.RWMutex
	Store map[string]*Point
}

func (u *User) PostPoint(w rest.ResponseWriter, r *rest.Request) {
	rest.NotFound(w, r)
}
