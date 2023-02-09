package controller

import (
	"blogs/common"
	"blogs/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllArticles(ctx *gin.Context) {
	var article model.Article
	var err error
	//default page to fetch first 20 article
	page := 1
	if ctx.Query("page") != "" {
		page, _ = strconv.Atoi(ctx.Query("page"))
	}
	articles, err := article.GetAllArticles(page)
	if err != nil {
		log.Println("Error Inserting into DB ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, common.Respone{Message: "StatusInternalServerError", Status: "500"})
	}

	ctx.JSON(http.StatusOK, common.Respone{Message: "success", Status: "200", Data: articles})
}

func GetCommentsOfArticle(ctx *gin.Context) {
	articleID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, common.Respone{Message: "Invalid Article Id", Status: "400"})
	}
	var article model.Article
	article.ID = uint(articleID)
	comments, err := article.GetAllComments()
	if err != nil {
		log.Println("Error Fetching from DB ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, common.Respone{Message: "StatusInternalServerError", Status: "500"})
	}

	ctx.JSON(http.StatusOK, common.Respone{Message: "success", Status: "200", Data: comments})
}
