package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gupta29470/golang-sql-crud-with-orm/databases"
	"github.com/gupta29470/golang-sql-crud-with-orm/models"
)

func CreateUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var user models.User

	decodeError := json.NewDecoder(request.Body).Decode(&user)
	if decodeError != nil {
		errorResponse := models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Request is in incorrect format",
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(errorResponse)
		return
	}

	if user.FirstName == "" || user.LastName == "" || user.Email == "" {
		errorResponse := models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "first_name, last_name and email should not be empty",
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(errorResponse)
		return
	}

	databases.DB().Create(&user)
	json.NewEncoder(writer).Encode(user)
}

func GetUsers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var users []models.User
	databases.DB().Find(&users)
	json.NewEncoder(writer).Encode(users)
}

func GetUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	id := strings.Split(request.URL.Path, "/")[2]
	if id == "" {
		errorResponse := models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "User id is missing",
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(errorResponse)
		return
	}
	var user models.User
	var count int64
	databases.DB().Model(&models.User{}).Where("id=?", id).Count(&count)
	if count <= 0 {
		errorResponse := models.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "User not found",
		}
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(errorResponse)
		return
	}

	databases.DB().First(&user, id)
	json.NewEncoder(writer).Encode(user)
}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var user models.User
	decodeError := json.NewDecoder(request.Body).Decode(&user)
	if decodeError != nil {
		errorResponse := models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Request is in incorrect format",
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(errorResponse)
		return
	}

	var fetchedUser models.User
	databases.DB().First(&fetchedUser, user.ID)

	if user.FirstName != "" {
		fetchedUser.FirstName = user.FirstName
	}

	if user.LastName != "" {
		fetchedUser.LastName = user.LastName
	}

	if user.Email != "" {
		fetchedUser.Email = user.Email
	}

	databases.DB().Save(&fetchedUser)
	json.NewEncoder(writer).Encode(fetchedUser)
}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	id := strings.Split(request.URL.Path, "/")[3]
	if id == "" {
		errorResponse := models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "User id is missing",
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(errorResponse)
		return
	}

	intId, _ := strconv.Atoi(id)
	var count int64
	databases.DB().Model(&models.User{}).Where("id=?", intId).Count(&count)
	if count <= 0 {
		errorResponse := models.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "User not found",
		}
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(errorResponse)
		return
	}

	user := models.User{ID: uint(intId)}
	databases.DB().Delete(&user, intId)
	writer.WriteHeader(http.StatusNoContent)
}
