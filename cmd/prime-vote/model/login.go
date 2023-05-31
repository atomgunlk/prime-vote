package model

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"` // Raw password string
}

type LoginReqponse struct {
	ResponseStatus ResponseStatus `json:"responseStatus"`
	Token          string         `json:"token"`
}
