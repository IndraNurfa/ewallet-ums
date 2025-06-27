package interfaces

import (
	"context"
	"ewallet-ums/internal/models"

	"github.com/gin-gonic/gin"
)

type IRegisterHandler interface {
	Register(c *gin.Context)
}
type IRegisterService interface {
	Register(ctx context.Context, request models.User) (interface{}, error)
}

type IRegisterRepository interface {
	InsertNewUser(ctx context.Context, user *models.User) error
}
