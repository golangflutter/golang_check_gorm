package model

// type OrderItem struct {
// 	Id        string `gorm:"type:uuid;primary_key"`
// 	OrderId   string `gorm:"type:uuid"`
// 	Order     Order  `gorm:"foreignKey:OrderID"`
// 	ProductId string `gorm:"type:uuid"`
// 	Product   Product
// }

// package model

// type OrderItem struct {
// 	ID        string `gorm:"type:uuid;primary_key"`
// 	OrderID   string `gorm:"type:uuid"`
// 	Order     Order // Define the relationship
// 	ProductID string `gorm:"type:uuid"`
// 	Product   Product // Define the relationship
// }
type OrderItem struct {
    ID        string `gorm:"type:uuid;primary_key"`
    OrderID   string `gorm:"type:uuid"`
    Order     Order `gorm:"foreignKey:OrderID"` // Define the relationship
    ProductID string `gorm:"type:uuid"`
    Product   Product `gorm:"foreignKey:ProductID"` // Define the relationship
}