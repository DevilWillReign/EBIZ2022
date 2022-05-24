package models

import "apprit/store/api/v1/database/dtos"

type Auth struct {
	ID       uint          `json:"id"`
	Authtype dtos.AuthType `json:"authtype" validate:"required"`
	UserID   uint          `json:"userid" validate:"required"`
}

func (a *Auth) Equals(o Auth) bool {
	return a.Authtype == o.Authtype && a.UserID == o.UserID
}
