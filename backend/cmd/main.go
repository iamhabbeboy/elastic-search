package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/iamhabbeboy/elastic-app/database/model"
)

func main() {
	u := new(model.UserParams)
	addr := gofakeit.Address()

	fulAddr := fmt.Sprintf("%s, %s.%s", addr.Street, addr.City, addr.Country)
	u.Name = gofakeit.Name()
	u.Company = gofakeit.Company()
	u.Address = fulAddr
	u.Phone = gofakeit.Phone()

	_, err := model.NewUserModel().Create(*u)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(umd.ID)
}
