package models

type Checkout struct {
	UserID      int     `json:"user_id"`
	ItemID      int     `json:"item_id"`
	VoucherCode string  `json:"voucher_code"`
	TotalPrice  float64 `json:"total_price"`
	EarnedPoint float64 `json:"earned_point"`
}

type CheckoutResponse struct {
	Message     string  `json:"message"`
	TotalPrice  float64 `json:"total_price"`
	EarnedPoint float64 `json:"earned_point"`
}
