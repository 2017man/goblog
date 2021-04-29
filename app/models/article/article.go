package article

import "gorm.io/gorm"

// Article 文章模型
type Article struct {
	gorm.Model

	ID    int
	Title string
	Body  string
}
