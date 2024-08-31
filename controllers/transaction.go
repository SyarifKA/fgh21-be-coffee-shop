package controllers

import (
	"RGT/konis/dtos"
	"RGT/konis/lib"
	"RGT/konis/models"
	"RGT/konis/repository"
	"fmt"
	"math/rand"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context){
	var formTransaction dtos.FormTransaction
	err := c.Bind(&formTransaction)
	userId := c.GetInt("UserId")
	if err != nil {
		lib.HandlerBadReq(c,"Invalide")
		return
	}
	noOrder := rand.Intn(90000)+10000
	fmt.Println(formTransaction)
	fmt.Println(noOrder)
	data, err := repository.CreateTransaction(models.Transaction{
		NoOrder: noOrder,
		AddFullName: formTransaction.FullName,
		AddEmail: formTransaction.Email,
		AddAddress: formTransaction.Address,
		Payment: formTransaction.Payment,
		UserId: userId,
		OrderTypeId: formTransaction.OrderType,
		TransactionStatusId: formTransaction.TransactionStatus,
	})

	if err != nil {
		
		lib.HandlerBadReq(c,"Invalid Data")
		return
	}

	lib.HandlerOK(c,"transaction success", data,nil)
}