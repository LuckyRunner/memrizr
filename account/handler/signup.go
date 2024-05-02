package handler

import (
	"github.com/gin-gonic/gin"
)

// signupReq is not exported, hence the lowercase name
// it is used for validation and marshalling
type signupReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

func (h *Handler) Signup(c *gin.Context) {
	// define a variable to which we'll bind incoming
	// json body, {email, password}
	var req signupReq

	// Bind incoming json body to struct and check for validation errors
	if ok := bindData(c, &req); !ok {
		return
	}
}
