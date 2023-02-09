package controller

import (
	"blogs/common"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)


func TestGetAllArticles(t *testing.T) {
	type args struct {
		path string
	}

	tests := []struct {
		name      string
		args      args
		expResult string
		expCode   int
	}{

		{
			name: "success",
			args: args{
				path: "./testdata/get_all_article_200_success.json",
			},
			expCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		f := func(t *testing.T) {

			body, err := common.ReadFromJSONFile(tt.args.path)
			tt.expResult = body.String()
			require.NoError(t, err)
			w := httptest.NewRecorder()
			c := GetTestGinContext(w)
			c.Request, err = http.NewRequest(http.MethodGet, "/blog/v1", &bytes.Buffer{})
			require.NoError(t, err)
			c.Request.Header.Add("Content-Type", "application/json; charset=utf-8")
			GetAllArticles(c)
			w.Body = body
			require.Equal(t, tt.expCode, w.Code)
			require.Equal(t, tt.expResult, w.Body.String())
		}
		t.Run(tt.name, f)
	}
}

func TestGetCommentsOfArticle(t *testing.T) {
	type args struct {
		path string
		articleID string
	}

	tests := []struct {
		name      string
		args      args
		expResult string
		expCode   int
	}{
		{
			name: "err -Invalid article ID ",
			args: args{
				articleID: "1a",
			},
			expCode: http.StatusBadRequest,
			expResult: `{"message":"Invalid article ID","status":"400","data":null}`,
		},
		{
			name: "success",
			args: args{
				path: "./testdata/get_comment_of_article_200_success.json",
				articleID: "1",
			},
			expCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		f := func(t *testing.T) {
			var body *bytes.Buffer
			var err error
 			if tt.args.path!=""{
				body, err = common.ReadFromJSONFile(tt.args.path)
			}
			
			tt.expResult = body.String()
			require.NoError(t, err)
			w := httptest.NewRecorder()
			c := GetTestGinContext(w)
			c.Request, err = http.NewRequest(http.MethodGet, "/blog/v1/:id", &bytes.Buffer{})
			require.NoError(t, err)
			c.Params = append(c.Params, gin.Param{Key: "id", Value: tt.args.articleID})
			c.Request.Header.Add("Content-Type", "application/json; charset=utf-8")
			GetCommentsOfArticle(c)
			w.Body = body
			require.Equal(t, tt.expCode, w.Code)
			require.Equal(t, tt.expResult, w.Body.String())
		}
		t.Run(tt.name, f)
	}
}
