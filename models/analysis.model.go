package models

type Order struct {
	OrderID         int `gorm:"primaryKey"`
	ProductID       string
	CustomerID      string
	ProductName     string
	Category        string
	Region          string
	DateOfSale      string
	QuantitySold    int
	UnitPrice       float64
	Discount        float64
	ShippingCost    float64
	PaymentMethod   string
	CustomerName    string
	CustomerEmail   string
	CustomerAddress string
}

type TopProduct struct {
	ProductName   string
	TotalQuantity int
}
