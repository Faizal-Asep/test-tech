package model

var Notif chan Notification

type Notification struct {
	Command string      `json:"command"`
	Data    interface{} `json:"data"`
}

type HttpResponse struct {
	Message string `json:"message"`
}
type UserInfo struct {
	Sub           string `json:"sub"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}
