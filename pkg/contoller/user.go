package contoller

import (
	"encoding/json"
	"net/http"

	usermodel "github.com/api-assignment/pkg/model/userModel"
	"github.com/api-assignment/pkg/utils/logger"
)

func GetUserData(w http.ResponseWriter, r *http.Request) {
	log := logger.InitializeAuditLogger()
	ctx := r.Context()
	userId, _ := ctx.Value("userId").(uint64)
	userData, err := usermodel.FindUserByID(userId)
	if err != nil {
		log.Errorf("unable to find user with ID %v ", userId, err)
		return
	}
	json.NewEncoder(w).Encode(userData)
}

func DeActivateUser(w http.ResponseWriter, r *http.Request) {
	log := logger.InitializeAuditLogger()
	ctx := r.Context()
	userId, _ := ctx.Value("userId").(uint64)
	// var userData usermodel.UserStatus
	userData, err := usermodel.FindUserByID(userId)
	if err != nil {
		log.Errorf("unable to find user with ID %v ", userId, err)
		return
	}
	err = userData.Disable()
	if err != nil {
		http.Error(w, "user has already been disabled", http.StatusBadRequest)
		log.Error(err)
		return
	}
	err = userData.Save()
	if err != nil {
		http.Error(w, "unable to update user status", http.StatusBadRequest)
		log.Error("unable to update user status ", err)
		return
	}
	w.Write([]byte("user has been disabled and will reflect in 5 mins"))
}
