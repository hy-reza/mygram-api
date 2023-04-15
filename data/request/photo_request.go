package request

type CreatePhotoRequest struct {
	Title     string `json:"title" validate:"required"`
	Caption   string `json:"caption"`
	Photo_url string `json:"photo_url" validate:"required"`
	User_id   string `json:"user_id"`
}

type UpdatePhotoRequest struct {
	Title     string `json:"title" validate:"required"`
	Caption   string `json:"caption"`
	Photo_url string `json:"photo_url" validate:"required"`
}
