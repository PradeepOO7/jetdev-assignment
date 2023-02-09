package controller

import "github.com/gin-gonic/gin"

func AddRoutes(router *gin.Engine) {

	grp := router.Group("/blog/v1")
	grp.GET("", GetAllArticles)
	grp.POST("",CreateArticle)
	grp.GET("/:id/comment",GetCommentsOfArticle)
	grp.PUT("/:id/article",AddCommentOnArticle)
	grp.PUT("/:id/comment",AddCommentOnComment)
}
