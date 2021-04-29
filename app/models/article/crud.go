package article

import (
	"goblog/pkg/model"
	"goblog/pkg/types"
)

// Get 通过Id获取文章
func Get(idstring string) (Article, error) {
	var article Article
	id := types.StringToInt(idstring)
	if err := model.DB.First(&article, id).Error; err != nil {
		return article, err
	}
	return article, nil
}
