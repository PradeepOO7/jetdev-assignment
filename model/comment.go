package model

import "time"

type Comment struct {
	ID         uint      `json:"id"`
	Nickname   string    `json:"nickname" binding:"required"`
	Content    string    `json:"content" binding:"required"`
	CreateDate time.Time `json:"createDate" binding:"required"`
	Comments   []Comment `json:"comments,omitempty" gorm:"many2many:comment_comments"`
}

func (comment Comment) CommentOnArticle(article Article) error {
	err := SqlOrm.Model(&article).Association("Comments").Append(&comment)
	if err != nil {
		return err
	}
	return nil
}

func (comment Comment) CommentOnComment(commentID int) error {

	err := SqlOrm.Model(&Comment{ID: uint(commentID)}).Association("Comments").Append(&comment)
	if err != nil {
		return err
	}
	return nil
}
