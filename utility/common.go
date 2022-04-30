package utility

import (
	"encoding/json"
	"fmt"
)

// VarDump 打印数据
func VarDump(expression ...interface{}) {
	fmt.Println(fmt.Sprintf("%#v", expression))
}
func JsonDecode(data string) (map[string]interface{}, error) {
	var dat map[string]interface{}
	err := json.Unmarshal([]byte(data), &dat)
	return dat, err
}
func JsonEncode(data interface{}) (string, error) {
	jsons, err := json.Marshal(data)
	return string(jsons), err
}
