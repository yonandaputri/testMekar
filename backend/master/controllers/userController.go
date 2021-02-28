package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"testMekar/master/models"
	"testMekar/master/usecases"
	"testMekar/tools"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserUsecase usecases.UserUsecase
}

func UserController(r *mux.Router, service usecases.UserUsecase) {
	userHandler := UserHandler{service}
	r.HandleFunc("/users", userHandler.ListUser).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", userHandler.GetUserById).Methods(http.MethodGet)
	r.HandleFunc("/user", userHandler.PostUser).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}", userHandler.PutUser).Methods(http.MethodPut)
	r.HandleFunc("/user/{id}", userHandler.DeleteUser).Methods(http.MethodDelete)

	r.HandleFunc("/job", userHandler.ListJob).Methods(http.MethodGet)
	r.HandleFunc("/education", userHandler.ListEducation).Methods(http.MethodGet)
}

func (uh UserHandler) ListUser(w http.ResponseWriter, r *http.Request) {
	users, err := uh.UserUsecase.GetUsers()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}

	byteOfUsers, err := json.Marshal(users)
	if err != nil {
		w.Write([]byte("Something When Wrong!!!"))
	}

	if users != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(byteOfUsers))
	} else {
		msg := fmt.Sprintf("List Data Users Not Found")
		byteOfMsg, _ := json.Marshal(msg)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(byteOfMsg))
	}
	fmt.Println("Endpoint hit: Get Users")
}

func (uh UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	key := "id"
	userId, _ := strconv.Atoi(tools.GetPathVar(key, r))
	user, err := uh.UserUsecase.GetUserById(userId)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}

	byteOfUser, err := json.Marshal(user)
	if err != nil {
		w.Write([]byte("Something When Wrong!!!"))
	}

	if user != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(byteOfUser))
	} else {
		msg := fmt.Sprintf("Data User Not Found")
		byteOfMsg, _ := json.Marshal(msg)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(byteOfMsg))
	}
	fmt.Println("Endpoint hit: Get User By Id")
}

func (uh UserHandler) PostUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	response, err := uh.UserUsecase.PostUser(user)
	if err != nil {
		w.Write([]byte("Cannot Add Data"))
	} else {
		data, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(data))
	}
	fmt.Println("Endpoint hit: Post User")
}

func (uh UserHandler) PutUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	key := "id"
	idUser, _ := strconv.Atoi(tools.GetPathVar(key, r))
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Write([]byte("Something Wrong"))
	}

	response, err := uh.UserUsecase.PutUser(idUser, user)
	if err != nil {
		w.Write([]byte("Cannot Update Data"))
	} else {
		data, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(data))
	}
	fmt.Println("Endpoint hit: Put User")
}

func (uh UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	key := "id"
	idUser, _ := strconv.Atoi(tools.GetPathVar(key, r))
	err := uh.UserUsecase.DeleteUser(idUser)
	if err != nil {
		w.Write([]byte("Cannot Delete Data"))
	} else {
		w.Write([]byte("Success Delete Data"))
	}
	fmt.Println("Endpoint hit: Delete User")
}

func (uh UserHandler) ListJob(w http.ResponseWriter, r *http.Request) {
	jobs, err := uh.UserUsecase.GetJob()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfJobs, err := json.Marshal(jobs)
	if err != nil {
		w.Write([]byte("Something When Wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(byteOfJobs))
	fmt.Println("Endpoint hit: Get List Job")
}

func (uh UserHandler) ListEducation(w http.ResponseWriter, r *http.Request) {
	education, err := uh.UserUsecase.GetEducation()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfEducations, err := json.Marshal(education)
	if err != nil {
		w.Write([]byte("Something When Wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(byteOfEducations))
	fmt.Println("Endpoint hit: Get List Education")
}
