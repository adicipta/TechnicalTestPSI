package services

import "soal1/models"

func CalculatePriceAndPoints(price float64, voucher models.Voucher, voucherFound bool) (float64, float64) {
	discount := price * (voucher.Value / 100)
	totalPrice := price - discount
	earnedPoint := 0.0
	if discount > 0 {
		earnedPoint = discount * 0.02
	}
	return totalPrice, earnedPoint
}
