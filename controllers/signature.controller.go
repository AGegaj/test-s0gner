package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	db "api/db/sqlc"
	"api/utils"
)

type SignatureController struct {
	db *db.Queries
}

func NewSignatureController(db *db.Queries) *SignatureController {
	return &SignatureController{db}
}

func (sc *SignatureController) SignTest(ctx *gin.Context) {
	var request *db.Signature

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	currentTime := time.Now()
	signedTest, err := utils.SignTestCompletion(request.UserID, currentTime)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return

	}

	args := &db.CreateSignatureParams{
		UserID:    request.UserID,
		Signature: signedTest,
		Answers:   request.Answers,
		Questions: request.Questions,
		Timestamp: currentTime,
	}

	signature, err := sc.db.CreateSignature(ctx, *args)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Test signed successfully", "data": signature})
}

func (sc *SignatureController) VerifyTest(ctx *gin.Context) {
	userId := ctx.Param("userId")
	signatureValue := ctx.Param("signature")

	verified, err := utils.VerifyTestCompletion(signatureValue, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return

	}

	if verified == false {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Test verification failed"})
		return

	}

	args := &db.GetSignatureByUserIdAndSignatureParams{
		UserID:    userId,
		Signature: signatureValue,
	}
	signature, err := sc.db.GetSignatureByUserIdAndSignature(ctx, *args)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Test verified successfully", "data": signature})
}
