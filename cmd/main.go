package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	//GET http://host:port/dataset?query=..
	//POST http://178.62.59.88:31392/mondial/query

	//POST http://host:port/dataset
	//Content-type: application/sparql-query
//	b := `SELECT ?subject ?predicate ?object
//WHERE {
//  ?subject ?predicate ?object
//}
//LIMIT 25`

	b := `SELECT * WHERE {
  ?sub <http://dooodle/predicate/hasColumn> ?obj .
  ?sub <http://dooodle/predicate/hasColumn> ?key .
  ?key <http://dooodle/predicate/numDistinct> ?num .
  ?obj <http://dooodle/predicate/hasDimension> <http://dooodle/dimension/scalar> .
  FILTER (?num > 0 && ?num <= 100)
} 
LIMIT 200`

	body := bytes.NewReader([]byte(b))

	//req, err := http.NewRequest("POST", "http://178.62.59.88:31392/mondial/query", body)
	req, err := http.NewRequest("POST", "http://178.62.59.88:31392/mondial", body)
	req.Header.Set("Content-type","application/sparql-query")
	//req.Header.Set("Accept", "text/csv")
	//req.Header.Set("Accept","application/sparql-results+json")
	req.Header.Set("Accept","text/tab-separated-values")
	//req.Header.Set("Accept", "application/sparql-results+xml")
	if err != nil {
		log.Fatal(err)
	}

	//req.SetBasicAuth("","")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	//data := bytes.Buffer{}
	//io.Copy(&data,resp.Body)
	io.Copy(os.Stdout,resp.Body)

}
