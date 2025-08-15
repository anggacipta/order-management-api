package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"` // "admin" atau "customer"
	Orders   []Order
}

type Product struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	OrderItems  []OrderItem
}

type Order struct {
	ID         uint        `gorm:"primaryKey" json:"id"`
	UserID     uint        `json:"user_id"`
	User       User        `json:"user"`
	OrderItems []OrderItem `json:"order_items"`
}

type OrderItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Product   Product `json:"product"`
	Quantity  int     `json:"quantity"`
	Price     int     `json:"price"`
}
