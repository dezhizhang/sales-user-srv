package test

import (
	"fmt"
	"testing"
)

type User struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Age       uint8  `validate:"gte=0,lte=130"`
}

var validate *validator.Validate

func TestValidator(t *testing.T) {
	fmt.Println("hello")
}
