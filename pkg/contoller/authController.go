package contoller

import (
	"encoding/json"
	"net/http"

	usermodel "github.com/api-assignment/pkg/model/userModel"
	"github.com/api-assignment/pkg/utils/validation"
)

type userSignup struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var user userSignup
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if err := validation.Validator.Struct(user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	userData := usermodel.CreateUser(user.Email)
	userData.SetPassword(user.Password)
	if err := userData.Save(); err != nil {
		http.Error(w, "email already exist", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(userData)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user userSignup
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if err := validation.Validator.Struct(user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	userData, err := usermodel.FindUserByEmail(user.Email)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	err = userData.ValidatePassword(user.Password)
	if err != nil {
		http.Error(w, "invalid email or password", http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(userData)
}
