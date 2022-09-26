// You can edit this code!
// Click here and start typing.
package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {
	sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}
	current := 0
	var conn net.Conn = nil
	for {
		var err error
		if conn == nil {
			conn, err = net.Dial("tcp", "localhost:8080")
			if err != nil {
				panic(err)
			}
			fmt.Printf("Access: %d\n", current)
		}
		request, err := http.NewRequest("POST", "http://localhost:8080", strings.NewReader(sendMessages[current]))
		if err != nil {
			panic(err)
		}
		request.Header.Set("Accept-Encoding", "gzip")
		err = request.Write(conn)
		if err != nil {
			panic(err)
		}
		response, err := http.ReadResponse(bufio.NewReader(conn), request)
		if err != nil {
			fmt.Println("Retly")
			conn = nil
			continue
		}
		dump, err := httputil.DumpResponse(response, false)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		defer response.Body.Close()

		if response.Header.Get("Content-Encoding") == "gzip" {
			reader, err := gzip.NewReader(response.Body)
			if err != nil {
				panic(err)
			}
			io.Copy(os.Stdout, reader)
			reader.Close()
		}
		if response.Header.Get("Content-Encoding") != "gzip" {
			io.Copy(os.Stdout, response.Body)
		}

	}
	/*
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			panic(err)
		}
		request, err := http.NewRequest("GET", "http://localhost:8080", nil)
		if err != nil {
			panic(err)
		}
		request.Write(conn)
		response, err := http.ReadResponse(bufio.NewReader(conn), request)
		if err != nil {
			panic(err)
		}
		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
	*/

}
