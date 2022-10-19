package main

import (
	"flag"
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"rpc-example/service"
	"time"
)

func main() {
	// Read the inline input
	var token string
	var address string
	var port string
	var format string
	flag.StringVar(&token, "t", "", "Need correct token to fetch the data")
	flag.StringVar(&address, "a", "localhost", "Server address")
	flag.StringVar(&port, "p", "2333", "Server port")
	flag.StringVar(&format, "f", "analogue", "The display mode of time")
	flag.Parse()

	// Connect to server
	client, err := jsonrpc.Dial("tcp", address+":"+port)
	if err != nil {
		log.Fatal("Server connection failure!")
	}
	cnt := 0

	// Verify the token and log in
	authRequest := service.AuthServiceRequest{Token: token}
	var authResponse service.AuthServiceResponse
	err = client.Call("AuthService.LogIn", authRequest, &authResponse)
	if err != nil {
		log.Fatal("AuthService.LogIn error")
	} else if authResponse.Status == "fail" {
		log.Fatal(authResponse.Message)
	} else {
		fmt.Println("Token validation success")
	}

	// Keep getting time from server
	for {
		cnt += 1
		request := service.TimeServiceRequest{}
		var response service.TimeServiceResponse
		err = client.Call("TimeService.GetTime", request, &response)
		if err != nil {
			log.Fatal("TimeService.GetTime error", err)
			break
		}

		// Print thr response time
		if format == "analogue" {
			fmt.Printf("Request %d: %v\n", cnt, response.CurrentTime)
			time.Sleep(time.Duration(1) * time.Minute)
		} else if format == "digital" {
			fmt.Printf("Request %d: %v\n", cnt, response.CurrentTime.Unix())
			time.Sleep(time.Duration(1) * time.Second)
		}
	}

	// Log out when error occurred
	defer client.Call("AuthService.LogOut", authRequest, &authResponse)
	defer client.Close()
}
