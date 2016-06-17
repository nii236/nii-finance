package main

import (
	"fmt"
	"testing"
	"time"
)

func TestPublish(t *testing.T) {
	fmt.Println(time.Now().UnixNano())
	input := `{"price": 469.0, "timestamp": "1464340140", "amount": 1.80891143, "type": 0, "id": 11226014}`
	publish(input)

}
