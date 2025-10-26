package models

type InvoiceItem struct {
    ID        uint    `gorm:"primaryKey" json:"id"`
    InvoiceID uint    `json:"invoice_id"`
    ProductID uint    `json:"product_id"`
    Product   Product `gorm:"foreignKey:ProductID" json:"product"`
    Quantity  int     `json:"quantity"`
    Price     int64   `json:"price"`
    Subtotal  int64   `json:"subtotal"`
}
