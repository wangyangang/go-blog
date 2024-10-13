package dao

import (
	"go-blog/models"
	"log"
)

func GetCategoryNameById(cId int) string {
	row := DB.QueryRow("SELECT name FROM blog_category WHERE cid=?", cId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var categoryName string
	_ = row.Scan(&categoryName)
	return categoryName
}
func GetAllCategory() ([]models.Category, error) {
	rows, err := DB.Query("SELECT * FROM blog_category")
	if err != nil {
		log.Println("GetAllCategory 查询出错:", err)
		return nil, err
	}
	var categories []models.Category
	for rows.Next() {
		var obj models.Category
		err = rows.Scan(&obj.Cid, &obj.Name, &obj.CreateAt, &obj.UpdateAt)
		if err != nil {
			log.Println("GetAllCategory取值出错:", err)
			return nil, err
		}
		categories = append(categories, obj)
	}
	return categories, nil
}
