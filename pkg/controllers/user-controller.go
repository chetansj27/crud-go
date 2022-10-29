package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/chetansj27/crud-go/pkg/models"
	"github.com/chetansj27/crud-go/pkg/utils"
	"github.com/gorilla/mux"
)

func GetUsers(writer http.ResponseWriter, request *http.Request) {
	allUsers := models.GetAllUsers()
	res, _ := json.Marshal(allUsers)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}

func GetUserById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId := vars["userId"]
	id, err := strconv.ParseUint(userId, 0, 0)
	if err != nil {
		fmt.Println("userid not integer")
	}
	userDetails, _ := models.GetUserById(id)
	res, _ := json.Marshal(userDetails)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}

func AddUser(writer http.ResponseWriter, request *http.Request) {
	addUser := &models.User{}
	utils.ParseBody(request, addUser)
	user := models.AddUser(addUser)
	res, _ := json.Marshal(user)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId := vars["userId"]
	id, err := strconv.ParseUint(userId, 0, 0)
	if err != nil {
		fmt.Println("userid not integer")
	}
	user := models.DeleteUserById(id)
	res, _ := json.Marshal(user)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)

}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	var updateUser = &models.User{}
	utils.ParseBody(request, updateUser)
	vars := mux.Vars(request)
	userId := vars["userId"]
	id, err := strconv.ParseUint(userId, 0, 0)
	if err != nil {
		fmt.Println("userid not integer")
	}
	userDetails, db := models.GetUserById(id)
	if updateUser.Name != "" {
		userDetails.Name = updateUser.Name
	}
	if updateUser.Email != "" {
		userDetails.Email = updateUser.Email
	}
	if updateUser.Address != "" {
		userDetails.Address = updateUser.Address
	}
	db.Save(&userDetails)
	res, _ := json.Marshal(userDetails)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}
