package model

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID         uint      `json:"id,omitempty"`
	Nickname   string    `json:"nickname" binding:"required"`
	Title      string    `json:"title" binding:"required"`
	Content    string    `json:"content" binding:"required"`
	CreateDate time.Time `json:"createDate" binding:"required"`

	Comments []Comment `json:"comments,omitempty" gorm:"many2many:article_comments"`
}

var SqlOrm *gorm.DB

func (article *Article) GetAllComments() ([]Comment, error) {
	var comments []Comment
	err := SqlOrm.Model(article).Association("Comments").Find(&comments)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (article *Article) CreateArticle() error {
	err := SqlOrm.Create(&article).Error
	if err != nil {
		return err
	}
	return nil
}

func (article *Article) GetAllArticles(page int) ([]Article, error) {
	var articles []Article
	offset := (page - 1) * 20
	err := SqlOrm.Model(article).Offset(offset).Limit(20).Find(&articles).Error

	if err != nil {
		return nil, err
	}

	return articles, nil
}
