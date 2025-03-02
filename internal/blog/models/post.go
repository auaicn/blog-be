package models

type Post struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// type Post struct {
// 	ID    uint   `gorm:"primaryKey"`
// 	Name  string `gorm:"size:100"`
// 	Email string `gorm:"size:100;unique"`
// }
