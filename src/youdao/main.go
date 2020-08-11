package main

import (
	"alfred/youdao/logic"
	"encoding/json"
	"fmt"
)

func main() {
	res, _ := json.Marshal(logic.Excute())
	fmt.Print(string(res))
}
