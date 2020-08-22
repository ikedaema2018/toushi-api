// main.go
package main

import (
        "github.com/aws/aws-lambda-go/lambda"
				"net/http"
				"fmt"
				"io/ioutil"
)

func hello() (string, error) {
	url := "http://google.co.jp"
	func (){
		res, _ := http.Get(url)
		defer res.Body.Close()

		byteArray, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(byteArray))
	}()
	return "Hello ƛ!", nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(hello)
}