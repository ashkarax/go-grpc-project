package models

type Product struct {
	Id                int64 `gorm:"primaryKey"`
	Name              string
	Stock             int64
	Price             int64
	StockDecreaseLogs StockDecreaseLog `gorm:"foreignKey:ProductRefer"`
}
