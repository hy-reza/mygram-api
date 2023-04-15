package request

type CreateCommentRequest struct {
	PhotoId string `json:"photo_id" validate:"required"`
	Message string `json:"message" binding:"required" validate:"required"`
}
type UpdateCommentRequest struct {
	Message string `json:"message" binding:"required" validate:"required"`
}
