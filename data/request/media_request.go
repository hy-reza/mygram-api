package request

type CreateMediaRequest struct {
	Name             string `json:"name" form:"name" validate:"required"`
	Social_media_url string `json:"social_media_url" form:"social_media_url" validate:"required"`
}

type UpdateMediaRequest struct {
	Name             *string `json:"name,omitempty" form:"name,omitempty"`
	Social_media_url *string `json:"social_media_url,omitempty" form:"social_media_url,omitempty" validate:"omitempty"`
}
