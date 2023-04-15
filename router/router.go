package router

import (
	"net/http"

	"github.com/hy-reza/mygram-api/controller"
	"github.com/hy-reza/mygram-api/middleware"
	"github.com/hy-reza/mygram-api/repository"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func NewRouter(userRepository repository.UsersRepository, photoRepository repository.PhotosRepository, commentRepository *repository.CommentRepositoryImpl, mediaRepository repository.MediaRepository, authenticationController *controller.AuthenticationController, usersController *controller.UserController, photosController *controller.PhotoController, commentController *controller.CommentController, mediaController *controller.MediaController) *gin.Engine {
	service := gin.Default()

	service.StaticFile("/docs_file.yaml", "./docs/docs_file.yaml")

	url := ginSwagger.URL("http://localhost:8888/docs_file.yaml")
	service.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api/v1")
	authenticationRouter := router.Group("/authentication")
	authenticationRouter.POST("/register", authenticationController.Register)
	authenticationRouter.POST("/login", authenticationController.Login)

	usersRouter := router.Group("/users")
	usersRouter.GET("/list", middleware.DeserializeUser(userRepository), usersController.GetUsers)

	photosRouter := router.Group("/photos")
	photosRouter.POST("/", middleware.DeserializeUser(userRepository), photosController.CreatePhoto)
	photosRouter.GET("/", middleware.DeserializeUser(userRepository), photosController.GetPhotos)
	photosRouter.GET("/:id", middleware.DeserializeUser(userRepository), photosController.GetPhoto)
	photosRouter.PUT("/:id", middleware.DeserializeUser(userRepository), middleware.ProtecPhoto(photoRepository), photosController.UpdatePhoto)
	photosRouter.DELETE("/:id", middleware.DeserializeUser(userRepository), middleware.ProtecPhoto(photoRepository), photosController.DeletePhoto)

	commentsRouter := router.Group("/comments")
	commentsRouter.POST("/", middleware.DeserializeUser(userRepository), commentController.CreateComment)
	commentsRouter.GET("/", middleware.DeserializeUser(userRepository), commentController.GetAllComments)
	commentsRouter.GET("/:id", middleware.DeserializeUser(userRepository), commentController.GetComment)
	commentsRouter.PUT("/:id", middleware.DeserializeUser(userRepository), middleware.ProtecComment(*commentRepository), commentController.UpdateComment)
	commentsRouter.DELETE("/:id", middleware.DeserializeUser(userRepository), middleware.ProtecComment(*commentRepository), commentController.DeleteComment)

	mediaRouter := router.Group("/media")
	mediaRouter.POST("/", middleware.DeserializeUser(userRepository), mediaController.CreateMedia)
	mediaRouter.GET("/", middleware.DeserializeUser(userRepository), mediaController.GetMedias)
	mediaRouter.GET("/:id", middleware.DeserializeUser(userRepository), mediaController.GetMedia)
	mediaRouter.PUT("/:id", middleware.DeserializeUser(userRepository), middleware.ProtecMedia(mediaRepository), mediaController.UpdateMedia)
	mediaRouter.DELETE("/:id", middleware.DeserializeUser(userRepository), middleware.ProtecMedia(mediaRepository), mediaController.DeleteMedia)
	return service
}
