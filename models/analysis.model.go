package models

type Order struct {
	OrderID         int    `gorm:"primaryKey"`
	ProductID       string `gorm:"unique"`
	CustomerID      string `gorm:"unique"`
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

type Revenue struct {
	DateFrom string `json:"dateFrom"`
	DateTo   string `json:"dateTo"`
	Product  string `json:"product"`
	Category string `json:"category"`
	Region   string `json:"region"`
}

type TopProduct struct {
	ProductName   string
	TotalQuantity int
}
