package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func Test_customJson(t *testing.T) {
	// prepare a MyUser instance
	u := &MyUser{1, "Ken", time.Now()}

	// marshal
	bytes, _ := json.Marshal(u)

	// display
	fmt.Println(string(bytes))

	// unmarshal
	u1 := &MyUser{}
	json.Unmarshal(bytes, u1)

	// display
	fmt.Printf("origian instance = \n%+v\n", u)
	fmt.Printf("unmarshal(marshal()) instance = \n%+v\n", u1)
}
