package main

import (
	"fmt"

	"github.com/tevity/go-getting-started/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Trisha",
		LastName:  "McMillan",
	}
	fmt.Println(u)
}
