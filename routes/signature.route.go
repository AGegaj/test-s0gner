package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

type SignatureRoutes struct {
	signatureController controllers.SignatureController
}

func NewSignatureRoutes(signatureController controllers.SignatureController) SignatureRoutes {
	return SignatureRoutes{signatureController}
}

func (sr *SignatureRoutes) SignatureRoute(rg *gin.RouterGroup) {
	router := rg.Group("/signature")
	router.POST("/sign", sr.signatureController.SignTest)
	router.GET("/verify/:userId/:signature", sr.signatureController.VerifyTest)
}
