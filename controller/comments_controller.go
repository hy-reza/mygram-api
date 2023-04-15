package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hy-reza/mygram-api/data/request"
	"github.com/hy-reza/mygram-api/data/response"
	"github.com/hy-reza/mygram-api/helper"
	"github.com/hy-reza/mygram-api/model"
	"github.com/hy-reza/mygram-api/repository"
)

type CommentController struct {
	commentRepository repository.CommentRepository
}

func NewCommentController(repo repository.CommentRepository) *CommentController {
	return &CommentController{commentRepository: repo}
}

func (controller *CommentController) GetAllComments(ctx *gin.Context) {
	comments, err := controller.commentRepository.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Failed to fetch comments!",
			Errors:  []string{err.Error()},
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully fetched comments!",
		Data:    comments,
	})
}

func (controller *CommentController) GetComment(ctx *gin.Context) {
	id := ctx.Param("id")
	comment, err := controller.commentRepository.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.Response{
			Code:    http.StatusNotFound,
			Status:  "Not Found",
			Message: "Comment not found!",
			Errors:  []string{err.Error()},
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully fetched comment!",
		Data:    comment,
	})
}

func (controller *CommentController) CreateComment(ctx *gin.Context) {
	createCommentRequest := request.CreateCommentRequest{}
	ctx.ShouldBindJSON(&createCommentRequest)
	err := validate.Struct(createCommentRequest)
	if err != nil {
		helper.ErrorBinding(ctx, err, http.StatusBadRequest, "Comment Failed!")
		return
	}
	userId := ctx.Value("userId")

	comment := model.Comment{
		UserID:  userId.(string),
		PhotoID: createCommentRequest.PhotoId,
		Message: createCommentRequest.Message,
	}

	newComment, err := controller.commentRepository.Save(comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Failed to create comment!",
			Errors:  []string{err.Error()},
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully created comment!",
		Data:    newComment,
	})
}

func (controller *CommentController) UpdateComment(ctx *gin.Context) {
	id := ctx.Param("id")
	comment, err := controller.commentRepository.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.Response{
			Code:    http.StatusNotFound,
			Status:  "Not Found",
			Message: "Comment not found!",
			Errors:  []string{err.Error()},
		})
		return
	}

	updateCommentRequest := request.UpdateCommentRequest{}
	ctx.ShouldBindJSON(&updateCommentRequest)

	comment.Message = updateCommentRequest.Message

	updatedComment, err := controller.commentRepository.Update(comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Failed to update comment!",
			Errors:  []string{err.Error()},
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully updated comment!",
		Data:    updatedComment,
	})
}

func (controller *CommentController) DeleteComment(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := controller.commentRepository.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.Response{
			Code:    http.StatusNotFound,
			Status:  "Not Found",
			Message: "Comment not found!",
			Errors:  []string{err.Error()},
		})
		return
	}

	err = controller.commentRepository.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Failed to delete comment!",
			Errors:  []string{err.Error()},
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully deleted comment!",
	})
}
