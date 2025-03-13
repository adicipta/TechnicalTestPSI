package controllers

import (
	"net/http"
	"soal1/models"
	"soal1/services"

	"github.com/gin-gonic/gin"
)

var users = map[int]models.User{
	1: {ID: 1, Name: "John Doe", Point: 0},
}

var items = map[int]models.Item{
	1: {ID: 1, Name: "Bootcamp Coding", Price: 5000000},
}

var vouchers = map[string]models.Voucher{
	"Discount50": {Code: "Discount50", Value: 50},
}

func CheckoutHandler(c *gin.Context) {
	var req models.Checkout
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, userExists := users[req.UserID]
	if !userExists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	item, itemExists := items[req.ItemID]
	if !itemExists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	var voucher models.Voucher
	voucherFound := false
	for _, v := range vouchers {
		if v.Code == req.VoucherCode {
			voucher = v
			voucherFound = true
			break
		}
	}

	if req.VoucherCode != "" && !voucherFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "Voucher not found"})
		return
	}

	totalPrice, earnedPoint := services.CalculatePriceAndPoints(item.Price, voucher, voucherFound)

	c.JSON(http.StatusOK, models.CheckoutResponse{
		Message:     "Checkout Success",
		TotalPrice:  totalPrice,
		EarnedPoint: earnedPoint,
	})

}
