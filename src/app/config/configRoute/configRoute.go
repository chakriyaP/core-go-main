package configRoute

import "github.com/gorilla/mux"

import userController "app/controllers/userController"
import jwtController "app/controllers/jwtController"

func Route() *mux.Router {

	mux := mux.NewRouter()

	
	mux.HandleFunc("/api/v1/user/selectAll", userController.SelectAll).Methods("GET", "OPTIONS")
	mux.HandleFunc("/api/v1/user/selectBySet", userController.SelectBySet).Methods("POST", "OPTIONS")
	mux.HandleFunc("/api/v1/user/selectByID/{_id}", userController.SelectByID).Methods("GET", "OPTIONS")
	mux.HandleFunc("/api/v1/user/selectByRedis/{_id}", userController.SelectByRedis).Methods("GET", "OPTIONS")
	mux.HandleFunc("/api/v1/user/insert", userController.Insert).Methods("POST", "OPTIONS")
	mux.HandleFunc("/api/v1/user/update/{_id}", userController.Update).Methods("PUT", "OPTIONS")
	mux.HandleFunc("/api/v1/user/delete/{_id}", userController.Delete).Methods("DELETE", "OPTIONS")


	mux.HandleFunc("/api/v1/jwt/create", jwtController.Create).Methods("GET", "OPTIONS")
	mux.HandleFunc("/api/v1/jwt/validate", jwtController.Validate).Methods("GET", "OPTIONS")

	
	return mux

}

