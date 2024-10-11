package main

import (
	"fmt"
	"helper"
)

func main() {
	h := helper.NewHelper()
	for _, s := range []string{"pelun", "pelun", "pehrr"} {
		h.AddString(s)
	}
	fmt.Println(h.MakeNewRequest("pe"))
	fmt.Println(h.ExpandRequest("h"))
	for _, s := range []string{"pehem", "pehem", "pehem"} {
		h.AddString(s)
	}
	fmt.Println(h.MakeNewRequest("pe"))
}