package controllers

import (
	"net/http"
	"encoding/json"
	"house-payment/data"
	"house-payment/models"
	"house-payment/common"
)

// Register add a new User document
// Handler for HTTP Post - "/users/register"
func Register(w http.ResponseWriter, r *http.Request) {
	/*
	Example data for HTTP Post - "user/register"
	============================================
		{
		"data":{
			"FirstName":"Rando",
			"LastName":"Mais",
			"Email": "my.naim@mail.com",
			"Password":"MyPassword"
		}
	}
	*/
	var dataResource UserResource
	//Decode incoming user json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user := &dataResource.Data
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("users")
	repo := &data.UserRepository{C: col}
	// Insert User document
	repo.CreateUser(user)
	// Clean-up the hashpassword to eliminate it from response JSON
	user.HashPassword = nil
	j, err := json.Marshal(UserResource{Data: *user})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

// Login authenticates the HTTP request with username and apssword
// Handler for HTTP Post - "/users/login"
func Login(w http.ResponseWriter, r *http.Request) {
	/*
	Example data for HTTP Post - "user/login"
	=============================================
	{
		"data":{
			"email": "rando.mais@mail.com",
			"password":"RandoMais123"
		}
	}
	*/
	var dataResource LoginResource
	var token string
	//Decode the incoming Login json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid login credentials",
			http.StatusUnauthorized,
		)
		return
	}
	loginModel := dataResource.Data
	loginUser := models.User{
		Email:    loginModel.Email,
		Password: loginModel.Password,
	}
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("users")
	repo := &data.UserRepository{C: col}
	// Authenticate the login user
	user, err := repo.Login(loginUser)
	//Generate JWT token
	token, err = common.GenerateJWT(user.Email, "member")
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Eror while generating the access token",
			http.StatusInternalServerError, //500
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// Clean-up the hashpassword to eliminate it from response JSON
	user.HashPassword = nil
	authUser := AuthUserModel{
		User:  user,
		Token: token,
	}
	j, err := json.Marshal(AuthUserResource{Data: authUser})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			http.StatusInternalServerError, //500
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
