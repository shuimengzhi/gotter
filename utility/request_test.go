package utility

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"1-1", args{url: "https://www.baidu.com"}},
		{"1-2", args{url: "https://www.4399.com"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := Get(tt.args.url)
			if body == "" {
				t.Error(tt.args.url + "请求错误")
			}
			fmt.Printf("%s \n", body)
		})
	}
}

func TestPost(t *testing.T) {
	type args struct {
		url  string
		info interface{}
	}
	tests := []struct {
		name string
		args args
		//want []byte
	}{
		// TODO: Add test cases.
		{"1-1", args{url: "http://127.0.0.1:3000/api/v1/ping"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := Post(tt.args.url, tt.args.info); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Post() = %v, want %v", got, tt.want)
			//}
			got := Post(tt.args.url, tt.args.info)

			response, _ := JsonDecode(got)
			VarDump(response)
			//for k,v:=range response{
			//	fmt.Println(k)
			//	fmt.Println(v)
			//}
			fmt.Println(response["msg"])

		})
	}
}
