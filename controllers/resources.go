package controllers

import (
	"house-payment/models"
)

type (
	// Models for JSON resources
	//For Post - /user/register
	//UserResource struct {
	//	Data models.User `json:"data"`
	//}

	//For Post - /user/login
	//LoginResource struct {
	//	Data LoginModel `json:"data"`
	//}
	//Response for authorized user Post - /user/login
	//AuthUserResource struct {
	//	Data AuthUserModel `json:"data"`
	//}
	// For Get - /cities
	CityResource struct {
		Data []models.City `json:"data"`
	}
)
