package main

import (
	"encoding/json"
	"fmt"
	"git.imooc.com/wendell1000/resk/services"
)

func main() {

	d, e := json.Marshal(&acservices.AccountTransferDTO{})
	fmt.Println(e)
	fmt.Println(string(d))
}
