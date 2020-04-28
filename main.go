package main

import (
	"github.com/tevity/go-getting-started/models"
	"fmt"
)

func main() {
	u := models.User {
		ID: 2,
		FirstName: "Trisha",
		LastName: "McMillan",
	}
	fmt.Println(u)
}
