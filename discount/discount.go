package discount

import "github.com/uptrace/bun"

type CouponEntity struct {
	bun.BaseModel `bun:"table:coupon,alias:u"`
	ID            int    `bun:",pk,autoincrement"`
	ProductName   string `bun:"product_name"`
	Description   string `bun:"description"`
	Amount        int    `bun:"amount"`
}

type Coupon struct {
	ID          int    `json:"id"`
	ProductName string `json:"product_name"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}

type CreateCouponDTO struct {
	ProductName string `json:"product_name"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}
