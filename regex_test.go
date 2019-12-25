package test

import (
	"fmt"
	"strings"
	"testing"
)

// 从url匹配domain
func TestGetDomainFromURL(t *testing.T) {
	url := "http://localhost.com/#/dashboard"
	a1 := strings.Split(url, "//")[1]
	a2 := strings.Split(a1, "/")[0]
	a3 := strings.Split(a2, ":")[0]
	fmt.Println(a3)
}
