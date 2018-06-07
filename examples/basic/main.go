package main

import (
	"net/http"

	"github.com/go-mego/bcrypt"
	"github.com/go-mego/mego"
)

func main() {
	e := mego.Default()
	e.GET("/", bcrypt.New(), func(c *mego.Context, cpt *bcrypt.Crypt) {
		c.String(http.StatusOK, cpt.Hash("myPassword"))
	})
	e.Run()
}
