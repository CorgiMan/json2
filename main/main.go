package main

import (
	"encoding/json"
	"fmt"

	"github.com/CorgiMan/json2"
)

type A struct {
	a int
	B int
}

func main() {
	// bts, _ := json.Marshal(A{1, 2})
	// fmt.Println(string(bts))
	bts, _ := json.Marshal(`{A:1, B:2}`)
	fmt.Println(string(bts))

	bts2, _ := json2.Marshal(A{1, 2})
	fmt.Println(string(bts2))
	var v interface{}
	json.Unmarshal(bts, &v)
	fmt.Printf("%+v\n", v)
}
