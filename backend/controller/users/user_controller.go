package users

import (
	"github.com/gin-gonic/gin"
	"github.com/ris967/go-react/backend/domain/users"
	"github.com/ris967/go-react/backend/utils/errors"
)

func Register(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.NewBadRequestError("invalid json body")
		c.JSON(err.Status, err)
		return
	}

	
}