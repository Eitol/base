package main

import (
	"fmt"
	"github.com/cabify/go-couchdb"
	"net/http"
	"net/url"
)

type roundTripperFunc func(*http.Request) (*http.Response, error)

func main() {
	httpClient := &http.Client{}
	u, _ := url.Parse("http://127.0.0.1:5984")
	c := couchdb.NewClient(u, httpClient, couchdb.BasicAuth("admin", "123456"))
	err := c.Ping() // trigger round trip
	if err != nil {
		fmt.Printf("%v", err)
	}
	names, err := c.AllDBs()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Print(names)
}
