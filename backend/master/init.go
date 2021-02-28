package master

import (
	"database/sql"
	"github.com/gorilla/mux"
	"testMekar/master/controllers"
	"testMekar/master/repositories"
	"testMekar/master/usecases"
)

func Init(r *mux.Router, db *sql.DB)  {
	userRepo := repositories.InitUserRepoImpl(db)
	userUsecase := usecases.InitUserUsecase(userRepo)
	controllers.UserController(r, userUsecase)
}