package tests

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type memthodUrl struct {
	method   string
	url      string
	expected int
}

func TestAllPages(t *testing.T) {
	baseUrl := "http://localhost:3000"
	tests := []memthodUrl{
		{"GET", "/", 200},                   //首页
		{"GET", "/about", 200},              //关于页页
		{"GET", "/articles", 200},           //列表页
		{"GET", "/articles/2", 200},         //详情页
		{"GET", "/articles/create", 200},    //创建页
		{"GET", "/articles/2/edit", 200},    //编辑页
		{"POST", "/articles/2", 200},        //编辑
		{"POST", "/articles/1/delete", 404}, //删除
		{"POST", "/articles", 200},          //创建
	}

	var (
		resp *http.Response
		err  error
	)
	for _, test := range tests {
		t.Logf("当前请求 URL: %v \n", test.url)
		switch {
		case test.method == "POST":
			data := make(map[string][]string)
			resp, err = http.PostForm(baseUrl+test.url, data)
		default:
			resp, err = http.Get(baseUrl + test.url)

		}
		assert.NoError(t, err, "请求 "+test.url+" 时报错")
		assert.Equal(t, test.expected, resp.StatusCode, test.url+" 应返回状态码 "+strconv.Itoa(test.expected))
	}

}
