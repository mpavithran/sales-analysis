package repositories

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/mpavithran/sales-analysis/models"
	"gorm.io/gorm"
)

type AnalysisRepository struct {
	DB *gorm.DB
}

func NewAnalysisRepository(db *gorm.DB) *AnalysisRepository {
	return &AnalysisRepository{DB: db}
}

func (r *AnalysisRepository) loadCSVData(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Unable to open file", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read()

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		orderID, _ := strconv.Atoi(record[0])
		quantitySold, _ := strconv.Atoi(record[7])
		unitPrice, _ := strconv.ParseFloat(record[8], 64)
		discount, _ := strconv.ParseFloat(record[9], 64)
		shippingCost, _ := strconv.ParseFloat(record[10], 64)

		order := models.Order{
			OrderID:         orderID,
			ProductID:       record[1],
			CustomerID:      record[2],
			ProductName:     record[3],
			Category:        record[4],
			Region:          record[5],
			DateOfSale:      record[6],
			QuantitySold:    quantitySold,
			UnitPrice:       unitPrice,
			Discount:        discount,
			ShippingCost:    shippingCost,
			PaymentMethod:   record[11],
			CustomerName:    record[12],
			CustomerEmail:   record[13],
			CustomerAddress: record[14],
		}
		r.DB.Create(order)
		if err != nil {
			log.Println("Failed to insert record", err)
		}
	}
}

func (r *AnalysisRepository) UploadCSV(filePath string) error {
	r.loadCSVData(filePath)
	return nil
}

func (r *AnalysisRepository) GetRevenue(revenue models.Revenue) (float64, error) {
	var revenueData float64
	query := r.DB.Table("orders").
		Select("COALESCE(SUM(quantity_sold * unit_price * (1 - discount)), 0)").
		Where("date_of_sale BETWEEN ? AND ?", revenue.DateFrom, revenue.DateTo)

	if revenue.Product != "" {
		query = query.Where("product_name = ?", revenue.Product)
	}

	if revenue.Category != "" {
		query = query.Where("category = ?", revenue.Category)
	}

	if revenue.Region != "" {
		query = query.Where("region = ?", revenue.Region)
	}

	if err := query.Scan(&revenueData).Error; err != nil {
		return 0, err
	}

	return revenueData, nil

}

func (r *AnalysisRepository) TopProducts(topProduct models.TopProduct) ([]models.TopProductData, error) {

	var products []models.TopProductData

	query := r.DB.Table("orders").
		Select("product_name, SUM(quantity_sold) as total_quantity").
		Where("date_of_sale BETWEEN ? AND ?", topProduct.DateFrom, topProduct.DateTo).
		Group("product_name").
		Order("total_quantity DESC").
		Limit(topProduct.Limit)

	if topProduct.Category != "" {
		query = query.Where("category = ?", topProduct.Category)
	}
	if topProduct.Region != "" {
		query = query.Where("region = ?", topProduct.Region)
	}

	if err := query.Scan(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
