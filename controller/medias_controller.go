package controller

import (
	"net/http"

	"github.com/hy-reza/mygram-api/data/request"
	"github.com/hy-reza/mygram-api/data/response"
	"github.com/hy-reza/mygram-api/helper"
	"github.com/hy-reza/mygram-api/model"
	"github.com/hy-reza/mygram-api/repository"

	"github.com/gin-gonic/gin"
)

type MediaController struct {
	mediaRepository repository.MediaRepository
}

func NewMediaController(repository repository.MediaRepository) *MediaController {
	return &MediaController{mediaRepository: repository}
}

func (controller *MediaController) GetMedias(ctx *gin.Context) {
	media, err := controller.mediaRepository.FindAll()
	if err != nil {
		helper.ErrorResponse(ctx, err, "media not found")
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all media data!",
		Data:    media,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *MediaController) CreateMedia(ctx *gin.Context) {
	createMediaRequest := request.CreateMediaRequest{}
	userId := ctx.Value("userId")

	ctx.ShouldBindJSON(&createMediaRequest)
	media := model.Media{
		Name:             createMediaRequest.Name,
		Social_media_url: createMediaRequest.Social_media_url,
		UserID:           userId.(string),
	}
	err := validate.Struct(createMediaRequest)
	if err != nil {
		helper.ErrorBinding(ctx, err, http.StatusBadRequest, "Create Media Failed!")
		return
	}

	newMedia, err := controller.mediaRepository.Save(media)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Create Media Failed!",
			Errors:  []string{err.Error()},
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully created media!",
		Data:    newMedia,
	})

}

func (controller *MediaController) UpdateMedia(ctx *gin.Context) {
	updateMediaRequest := request.UpdateMediaRequest{}
	mediaId := ctx.Param("id")

	ctx.ShouldBindJSON(&updateMediaRequest)
	media, err := controller.mediaRepository.FindById(mediaId)
	if err != nil {
		helper.ErrorResponse(ctx, err, "Media not found")
		return
	}

	if updateMediaRequest.Name != nil {
		media.Name = *updateMediaRequest.Name
	}
	if updateMediaRequest.Social_media_url != nil {
		media.Social_media_url = *updateMediaRequest.Social_media_url
	}

	err = validate.Struct(&updateMediaRequest)
	if err != nil {
		helper.ErrorBinding(ctx, err, http.StatusBadRequest, "Update Media Failed!")
		return
	}

	updatedMedia, err := controller.mediaRepository.Update(media)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Update Failed!",
			Errors:  []string{err.Error()},
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully updated media!",
		Data:    updatedMedia,
	})
}

func (controller *MediaController) DeleteMedia(ctx *gin.Context) {
	mediaId := ctx.Param("id")

	err := controller.mediaRepository.Delete(mediaId)
	if err != nil {
		helper.ErrorResponse(ctx, err, "Media not found")
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully deleted media!",
	})
}

func (controller *MediaController) GetMedia(ctx *gin.Context) {
	id := ctx.Param("id")
	media, err := controller.mediaRepository.FindById(id)
	if err != nil {
		helper.ErrorResponse(ctx, err, "media not found")
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully fetch media data!",
		Data:    media,
	})
}
