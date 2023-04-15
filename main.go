package main

import (
	"log"
	"net/http"
	"time"

	"github.com/hy-reza/mygram-api/config"
	"github.com/hy-reza/mygram-api/controller"
	"github.com/hy-reza/mygram-api/helper"
	"github.com/hy-reza/mygram-api/model"
	"github.com/hy-reza/mygram-api/repository"
	"github.com/hy-reza/mygram-api/router"
	"github.com/hy-reza/mygram-api/service"
)

func main() {

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	//Database
	db := config.ConnectionDB(&loadConfig)

	db.Table("users").AutoMigrate(&model.User{})
	db.Table("photos").AutoMigrate(&model.Photo{})
	db.Table("comments").AutoMigrate(&model.Comment{})
	db.Table("media").AutoMigrate(&model.Media{})

	//Init Repository
	userRepository := repository.NewUsersRepositoryImpl(db)
	photoRepository := repository.NewPhotosRepositoryImpl(db)
	commentRepository := repository.NewCommentRepository(db)
	mediaRepository := repository.NewMediaRepositoryImpl(db)

	//Init Service
	authenticationService := service.NewAuthenticationServiceImpl(userRepository)

	//Init controller
	authenticationController := controller.NewAuthenticationController(authenticationService)
	usersController := controller.NewUsersController(userRepository)
	photoController := controller.NewPhotosController(photoRepository)
	commentsController := controller.NewCommentController(commentRepository)
	mediasController := controller.NewMediaController(mediaRepository)

	var _ = commentsController

	routes := router.NewRouter(userRepository, photoRepository, commentRepository, mediaRepository, authenticationController, usersController, photoController, commentsController, mediasController)

	server := &http.Server{
		Addr:           ":" + loadConfig.ServerPort,
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server_err := server.ListenAndServe()
	helper.ErrorPanic(server_err)
}
