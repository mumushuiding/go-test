package test

import (
	"fmt"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	now := time.Now()
	fmt.Println("now:", now)
	fmt.Println("昨天:", now.Add(time.Hour*-24))
}
