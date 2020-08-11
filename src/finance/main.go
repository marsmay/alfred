package main

import (
	"alfred/finance/logic"
	"encoding/json"
	"fmt"
)

func main() {
	res, _ := json.Marshal(logic.Excute())
	fmt.Print(string(res))
}
