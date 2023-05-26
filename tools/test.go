package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func getNames(filename string) string {
	url := "localhost:43221"
	jsonStr := []byte(`{"path":"test.png"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Cotent-Type", "applicaiton/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.String()

}
func main() {
	fmt.Print("working")

	fmt.Print(getNames("test.png"))
	fmt.Print("...")
}
