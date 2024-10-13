package dao

import "go-blog/models"

func GetPostPage(page int, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("SELECT * FROM blog_post limit ?, ?", page, pageSize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var obj models.Post
		err = rows.Scan(&obj.Pid, &obj.Title, &obj.Content, &obj.Markdown, &obj.CategoryId, &obj.UserId,
			&obj.ViewCount, &obj.Type, &obj.Slug, &obj.CreateAt, &obj.UpdateAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, obj)
	}

	return posts, nil
}

func GetAllPostCount() (count int) {
	rows := DB.QueryRow("SELECT count(1) FROM blog_post")
	_ = rows.Scan(&count)
	return
}
