package main

import (
	"bufio"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type etcdresponse struct {
	Node etcdNode
}

type etcdNode struct {
	Nodes []node
}

type node struct {
	Value string
	Key   string
}

type server struct {
	Host    string
	Service string
}

func main() {
	var err error
	r, err := http.Get("http://172.17.42.1:4001/v2/keys/endpoints")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var n etcdresponse
	json.Unmarshal(body, &n)
	log.Println(n)
	locations := make([]server, 0)
	for _, endpoint := range n.Node.Nodes {
		keyparts := strings.Split(endpoint.Key, "/")
		log.Println(keyparts)
		s := server{
			Host:    endpoint.Value,
			Service: keyparts[len(keyparts)-1],
		}
		locations = append(locations, s)
	}
	tmpl, err := template.ParseFiles("nginx.template")
	if err != nil {
		log.Fatal(err)
	}
	// open output file
	fo, err := os.Create("/usr/share/nginx/autoproxy.conf")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	// make a write buffer
	w := bufio.NewWriter(fo)
	err = tmpl.Execute(w, locations)
	if err != nil {
		log.Fatal(err)
	}
	if err = w.Flush(); err != nil {
		panic(err)
	}
}
