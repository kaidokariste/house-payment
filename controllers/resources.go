package controllers

import (
	"house-payment/models"
)

// Models for JSON resources
type (
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

	// For Post/Put - /cities
	// For Get - /cities/id
	CityResource struct {
		Data models.City `json:"data"`
	}

	// For Get - /cities
	CitiesResource struct {
		Data []models.City `json:"data"`
	}


)
