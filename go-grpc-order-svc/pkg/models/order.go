package models

type Order struct {
	Id        int64 ` gorm:"primaryKey"`
	Price     int64
	ProductId int64
	UserId    int64
}
