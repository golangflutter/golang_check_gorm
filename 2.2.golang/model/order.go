package model
type Order struct {
    Id        int `gorm:"type:uuid;primary_key"`
    OrderItem []OrderItem `gorm:"foreignKey:OrderId"` // Added the foreign key constraint to the OrderItem field
    IsPaid    bool
    Phone     string
    Address   string
    CreatedAt string `gorm:"type:varchar(19);default:to_char(now(), 'YYYY-MM-DD HH24:MI:SS')"`
    UpdatedAt string `gorm:"type:varchar(19);default:to_char(now(), 'YYYY-MM-DD HH24:MI:SS')"`
}

// type Order struct {
// 	Id        int `gorm:"type:uuid;primary_key"`
//     OrderItem []OrderItem 
//     IsPaid    bool
//     Phone     string
//     Address   string
//     CreatedAt string `gorm:"type:varchar(19);default:to_char(now(), 'YYYY-MM-DD HH24:MI:SS')"`
//     UpdatedAt string `gorm:"type:varchar(19);default:to_char(now(), 'YYYY-MM-DD HH24:MI:SS')"`
// }



// type Order struct {

// 	OrderItems []OrderItem `gorm:"foreignKey:OrderId"`
// 	IsPaid    bool
// 	Phone     string
// 	Address   string
// 	CreatedAt string `gorm:"type:varchar(19);default:to_char(now(), 'YYYY-MM-DD HH24:MI:SS')"`
// 	UpdatedAt   string `gorm:"type:varchar(19);default:to_char(now(), 'YYYY-MM-DD HH24:MI:SS')"`
// }