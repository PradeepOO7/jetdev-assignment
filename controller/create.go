package controller

import (
	"blogs/common"
	"blogs/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateArticle(ctx *gin.Context) {
	var article model.Article

	if err := ctx.ShouldBind(&article); err != nil {
		log.Printf("Request Body is invalid %s \n", err)
		ctx.JSON(http.StatusBadRequest, common.Respone{Message: "Invalid RequestBody", Status: "400"})
		return
	}
	err := article.CreateArticle()
	if err != nil {
		log.Println("Error Inserting into DB ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, common.Respone{Message: "StatusInternalServerError", Status: "500"})
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

func AddCommentOnArticle(ctx *gin.Context) {
	var comment *model.Comment
	if err := ctx.Bind(&comment); err != nil {
		log.Printf("Request Body is invalid %s \n", err)
		ctx.JSON(http.StatusBadRequest, common.Respone{Message: "Invalid RequestBody", Status: "400"})
		return
	}
	articleID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println("Error Parsing articleID ", err)
		ctx.JSON(http.StatusBadRequest, common.Respone{Message: "invalid article id", Status: "400"})
		return
	}
	var article model.Article
	article.ID = uint(articleID)
	err = comment.CommentOnArticle(article)
	if err != nil {
		log.Println("Error Inserting into DB ", err)
		ctx.JSON(http.StatusInternalServerError, common.Respone{Message: "StatusInternalServerError", Status: "500"})
		return
	}

	ctx.JSON(http.StatusOK, common.Respone{Message: "success", Status: "200"})
}

func AddCommentOnComment(ctx *gin.Context) {
	var comment model.Comment
	if err := ctx.Bind(&comment); err != nil {
		log.Printf("Request Body is invalid %s \n", err)
		ctx.JSON(http.StatusBadRequest, common.Respone{Message: "Invalid RequestBody", Status: "400"})
		return
	}
	commentID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println("Error Parsing commentID ", err)
		ctx.JSON(http.StatusBadRequest, common.Respone{Message: "invalid comment id", Status: "400"})
		return
	}
	err = comment.CommentOnComment(commentID)
	if err != nil {
		log.Println("Error Inserting into DB ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, common.Respone{Message: "StatusInternalServerError", Status: "500"})
	}

	ctx.JSON(http.StatusOK, common.Respone{Message: "success", Status: "200"})
}
