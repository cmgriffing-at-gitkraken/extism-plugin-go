package main

import (
	"encoding/json"
	"fmt"

	"github.com/extism/go-pdk"
)

type Payload struct {
	Name string `json:"name"`
}

//go:export greet
func Greet() int32 {
	name := pdk.InputString()
	pdk.OutputString("Hello, " + name)

	payload := Payload{
		Name: name,
	}

	payloadJSON, err := json.Marshal(payload)

	if err != nil {
		fmt.Println("Could not marshal payload into JSON")
		return 1
	}

	req := pdk.NewHTTPRequest(pdk.MethodPost, "https://localhost/echo")
	req.SetBody([]byte(payloadJSON))
	req.SetHeader("Content-Type", "application/json")

	req.Send()

	// TODO: See how FS access works, get help from discord
	// d1 := []byte("hello\ngo\n")
	// fsErr := os.WriteFile("/tmp/dat1", d1, 0644)

	// if fsErr != nil {
	// 	fmt.Println("Could not write to filesystem")
	// 	// pdk.OutputString("Could not write to filesystem")
	// 	pdk.OutputString(fsErr.Error())
	// 	// return 1
	// }

	return 0
}

func main() {}
