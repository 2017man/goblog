package article

import (
	"fmt"
	"goblog/pkg/model"
	"goblog/pkg/types"
)

// Get 通过Id获取文章
func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToInt(idstr)
	fmt.Println(id, 1)

	if err := model.DB.First(&article, id).Error; err != nil {
		return article, err
	}

	return article, nil
}

// GetAll 获取所有文章
func GetAll() ([]Article, error) {
	var articles []Article
	if err := model.DB.Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}
