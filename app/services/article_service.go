package services

import (
	"blog/app/models"
	"blog/vendors/model"
	"blog/vendors/pagination"
	"net/http"
)

type ArticleService struct {
}

// Get 通过 ID 获取文章
func (ArticleService) GetBySlug(slug string) (models.Article, error) {
	var article models.Article
	err := model.DB.Preload("Category").First(&article, map[string]interface{}{"slug": slug}).Error
	return article, err
}

// GetAll 获取全部文章
func (ArticleService) GetAll(r *http.Request, perPage int, where map[string]interface{}) ([]models.Article, pagination.PagerData, error) {
	// 1. 初始化分页实例
	db := model.DB.Preload("Tags").Model(models.Article{}).Where(where).Order("is_top desc, `order`")
	_pager := pagination.New(r, db, perPage, "")

	// 2. 获取视图数据
	PagerData := _pager.Paging()

	// 3. 获取数据
	var articles []models.Article
	err := _pager.Results(&articles)
	return articles, PagerData, err
}

//Hots 获取最近热门文章
func (ArticleService) Hots(limit int) []models.Article {
	var articles []models.Article
	model.DB.Order("views desc").Limit(limit).Find(&articles)

	return articles
}
