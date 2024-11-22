package utility

import (
	"math/rand"

	"github.com/go-faker/faker/v4"
)

func RandomID(maxnum int) int {
	min := 10
	max := maxnum
	var result = rand.Intn(max-min) + min
	return result
}

func RandomUsername() string {
	return faker.Username()
}

func RandomFullname() string {
	return faker.FirstName() + " " + faker.LastName()
}

func RandomUUID() string {
	return faker.UUIDDigit()
}
