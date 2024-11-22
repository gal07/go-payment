package middleware

import (
	"github.com/gin-gonic/gin"
)

type middleware struct {
	// useCaseStudent students.IStudentUseCase
}

func NewMiddleWare() middleware {
	return middleware{
		// useCaseStudent: useCaseStudent,
	}
}

func MiddlewareAuth(c *gin.Context) {

	// Verifiy Token
	// util.VerifyToken(c)

}
