package main

import (
	"net/http"

	"github.com/tevity/go-getting-started/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
