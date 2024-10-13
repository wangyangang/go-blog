package service

import (
	"go-blog/config"
	"go-blog/dao"
	"go-blog/models"
	"html/template"
)

func GetAllIndexInfo(page int, pageSize int) (*models.HomeData, error) {
	categories, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	posts, err := dao.GetPostPage(page, pageSize)
	var postMores []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)
		content := []rune(post.Content)
		if len(content) > 200 {
			content = content[:200]
		}
		postMore := models.PostMore{
			post.Pid,
			post.Title,
			post.Slug,
			template.HTML(content),
			post.CategoryId,
			categoryName,
			post.UserId,
			userName,
			post.ViewCount,
			post.Type,
			post.CreateAt.Format("2006-01-02 15:04:05"),
			post.UpdateAt.Format("2006-01-02 15:04:05"),
		}
		postMores = append(postMores, postMore)
	}
	postCount := dao.GetAllPostCount()
	pageCont := postCount / pageSize
	if postCount%pageSize > 0 {
		pageCont++
	}
	var pages []int
	for i := 1; i <= pageCont; i++ {
		pages = append(pages, i)
	}

	data := &models.HomeData{config.Cfg.Viewer, categories, postMores, postCount,
		page, pages, page != pageCont}
	return data, nil
}
