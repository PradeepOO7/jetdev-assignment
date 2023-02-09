package controller

import (
	"blogs/config"
	"blogs/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func init() {
	config.InitDB()
}
func TestCreateArticle(t *testing.T) {
	type args struct {
		article *model.Article
	}

	tests := []struct {
		name      string
		args      args
		expResult string
		expCode   int
	}{
		{
			name: "err - parse request body",
			args: args{
				article: &model.Article{},
			},
			expCode:   http.StatusBadRequest,
			expResult: `{"message":"Invalid RequestBody","status":"400","data":null}`,
		},

		{
			name: "success",
			args: args{
				article: &model.Article{Nickname: "test", Content: "Testing", Title: "In Test Function", CreateDate: time.Now()},
			},
			expCode:   http.StatusCreated,
			expResult: `null`,
		},
	}
	for _, tt := range tests {
		f := func(t *testing.T) {

			body, err := json.Marshal(tt.args.article)
			require.NoError(t, err)
			w := httptest.NewRecorder()
			c := GetTestGinContext(w)
			c.Request, err = http.NewRequest(http.MethodPost, "/blog/v1", bytes.NewReader(body))
			require.NoError(t, err)
			c.Request.Header.Add("Content-Type", "application/json; charset=utf-8")
			CreateArticle(c)
			require.Equal(t, tt.expCode, w.Code)
			require.Equal(t, tt.expResult, w.Body.String())
		}
		t.Run(tt.name, f)
	}
}

func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func TestAddCommentOnArticle(t *testing.T) {
	type args struct {
		articleID string
		comment   *model.Comment
	}

	tests := []struct {
		name      string
		args      args
		expResult string
		expCode   int
	}{
		{
			name: "err - parse request body",
			args: args{
				articleID: "1",
				comment:   &model.Comment{},
			},
			expCode:   http.StatusBadRequest,
			expResult: `{"message":"Invalid RequestBody","status":"400","data":null}`,
		},
		{
			name: "err - invalid articleID",
			args: args{
				articleID: "1a",
				comment:   &model.Comment{Nickname: "test", Content: "Testing", CreateDate: time.Now()},
			},
			expCode:   http.StatusBadRequest,
			expResult: `{"message":"invalid article id","status":"400","data":null}`,
		},
		{
			name: "success",
			args: args{
				articleID: "1",
				comment:   &model.Comment{Nickname: "test", Content: "Testing", CreateDate: time.Now()},
			},
			expCode:   http.StatusOK,
			expResult: `{"message":"success","status":"200","data":null}`,
		},
	}
	for _, tt := range tests {
		f := func(t *testing.T) {

			body, err := json.Marshal(tt.args.comment)
			require.NoError(t, err)
			w := httptest.NewRecorder()
			c := GetTestGinContext(w)
			c.Request, err = http.NewRequest(http.MethodPut, "/blog/v1/:id/article", bytes.NewReader(body))
			require.NoError(t, err)
			c.Params = append(c.Params, gin.Param{Key: "id", Value: tt.args.articleID})
			c.Request.Header.Add("Content-Type", "application/json; charset=utf-8")
			AddCommentOnArticle(c)
			require.Equal(t, tt.expCode, w.Code)
			require.Equal(t, tt.expResult, w.Body.String())
		}
		t.Run(tt.name, f)
	}
}

func TestAddCommentOnComment(t *testing.T) {
	type args struct {
		commentID string
		comment   *model.Comment
	}

	tests := []struct {
		name      string
		args      args
		expResult string
		expCode   int
	}{
		{
			name: "err - parse request body",
			args: args{
				commentID: "1",
				comment:   &model.Comment{},
			},
			expCode:   http.StatusBadRequest,
			expResult: `{"message":"Invalid RequestBody","status":"400","data":null}`,
		},
		{
			name: "err - invalid commentID",
			args: args{
				commentID: "1a",
				comment:   &model.Comment{Nickname: "test", Content: "Testing", CreateDate: time.Now()},
			},
			expCode:   http.StatusBadRequest,
			expResult: `{"message":"invalid comment id","status":"400","data":null}`,
		},
		{
			name: "success",
			args: args{
				commentID: "1",
				comment:   &model.Comment{Nickname: "test", Content: "Testing", CreateDate: time.Now()},
			},
			expCode:   http.StatusOK,
			expResult: `{"message":"success","status":"200","data":null}`,
		},
	}
	for _, tt := range tests {
		f := func(t *testing.T) {

			body, err := json.Marshal(tt.args.comment)
			require.NoError(t, err)
			w := httptest.NewRecorder()
			c := GetTestGinContext(w)
			c.Request, err = http.NewRequest(http.MethodPut, "/blog/v1/:id/comment", bytes.NewReader(body))
			require.NoError(t, err)
			c.Params = append(c.Params, gin.Param{Key: "id", Value: tt.args.commentID})
			c.Request.Header.Add("Content-Type", "application/json; charset=utf-8")
			AddCommentOnComment(c)
			require.Equal(t, tt.expCode, w.Code)
			require.Equal(t, tt.expResult, w.Body.String())
		}
		t.Run(tt.name, f)
	}
}
