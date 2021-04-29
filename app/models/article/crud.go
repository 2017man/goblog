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

	model.DB.First(&article, id)
	fmt.Println(article)

	return article, nil
}
