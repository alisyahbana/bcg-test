package main

import (
	"fmt"
	"github.com/alisyahbana/bcg-test/pkg/common/app"
	"github.com/alisyahbana/bcg-test/pkg/common/log"
	"github.com/alisyahbana/bcg-test/pkg/handler"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
	"net/http"
)

func main() {
	router := httprouter.New()
	SetRoute(router)

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(router)

	fmt.Println(fmt.Sprintf("Starting BCG test API HTTP Server on %d", app.GetConfig().Port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", app.GetConfig().Port), n)
	if err != nil {
		log.Error(err.Error())
	}
}

func SetRoute(router *httprouter.Router) {
	router.GET("/", handler.InfoHandler)
	router.POST("/v1/purchase", handler.PurchaseHandler)
}
