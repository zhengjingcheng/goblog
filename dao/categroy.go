package dao

import (
	"log"
	"zjc_blog/models"
)

func GetCategoryNameById(cId int) string{
	row := DB.QueryRow("select name  from blog_category where cid=?",cId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var categoryName string
	_ = row.Scan(&categoryName)
	return categoryName
}

func GetAllCategory() ([]models.Category,error){
	rows ,err := DB.Query("select * from blog_category")//从目录表中读取出所有的信息
	if err !=  nil {
		log.Println("GetAllCategory 查询出错:",err)
		return nil, err
	}
	var categorys []models.Category
	for rows.Next() {//如果有该行的数据
		var category models.Category
		err = rows.Scan(&category.Cid,&category.Name,&category.CreateAt,&category.UpdateAt)//讲row种的数据读入到catagory中
		if err != nil {
			log.Println("GetAllCategory 取值出错:",err)
			return nil, err
		}
		categorys = append(categorys,category)
	}
	return categorys,nil
}