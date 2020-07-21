package benchmark

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

type Post struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

func newBody() io.Reader {
	return strings.NewReader(`{
	"id":1000, 
	"title":"this is a blog about json unmarshal vs decoder", 
	"description":"It really depends on what your input is. If you look at the implementation of the Decode method of json.Decoder, it buffers the entire JSON value in memory before unmarshalling it into a Go value. So in most cases it won't be any more memory efficient (although this could easily change in a future version of the language)." ,
	"created_at": 1595303417,
    "updated_at":1595303417 
}`)
}

func BenchmarkJsonUnmarshal(b *testing.B) {
	var post Post
	for i := 0; i < b.N; i++ {
		body, _ := ioutil.ReadAll(newBody())
		_ = json.Unmarshal(body, &post)
	}
}

func BenchmarkJsonDecoder(b *testing.B) {
	var post Post
	for i := 0; i < b.N; i++ {
		_ = json.NewDecoder(newBody()).Decode(&post)
	}
}
