package ordermodel

import "time"

type OrderStatus string

const (
	PENDING  OrderStatus = "pending"
	FINISHED OrderStatus = "finished"
	CANCELED OrderStatus = "canceled"
)

type OrderAdjustmentType string

const (
	OrderAdjustmentFixed      OrderAdjustmentType = "fixed"
	OrderAdjustmentPercentage OrderAdjustmentType = "percentage"
)

type OrderModel struct {
	Id            string
	TenantId      string
	UserId        string
	UserAddressId string
	Status        OrderStatus
	TotalValue    float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}

type OrderItemsModel struct {
	Id             string
	TenantId       string
	OrderId        string
	ComboId        *string
	ProductId      *string
	DrinkId        *string
	TypeOfDiscount OrderAdjustmentType
	DiscountValue  float64
	TypeOfAddition OrderAdjustmentType
	AdditionValue  float64
	UnitValue      float64
	TotalValue     float64
	Amount         float64
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}
