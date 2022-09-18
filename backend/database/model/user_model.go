package model

import (
	"fmt"

	"github.com/iamhabbeboy/elastic-app/ent"
)

type UserModel struct{}

type UserParams struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Company string `json:"company"`
	Address string `json:"address"`
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (*UserModel) Create(params UserParams) (*ent.User, error) {
	record, err := dbConn.User.
		Create().
		SetName(params.Name).
		SetPhone(params.Phone).
		SetCompany(params.Company).
		SetAddress(params.Address).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	return record, nil
}
