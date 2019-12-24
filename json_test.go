package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/mumushuiding/util"
)

type something struct {
	types string
}
type Host struct {
	types string
	Name  string
}

func TestSliceToJson(t *testing.T) {
	a := []int{4, 5, 6}
	json, _ := util.ToJSONStr(a)
	fmt.Println(json)
}
func TestStructToJson(t *testing.T) {
	m := Host{Name: "Sky", types: "192.168.23.92"}

	b, err := json.Marshal(m)
	if err != nil {

		fmt.Println("Umarshal failed:", err)
		return
	}
	fmt.Println("json:", string(b))
}
