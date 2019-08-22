package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
)

type Format int

const (
	Csv Format = iota
	Tsv
	Json
	Xml
)

var currentFormat = Csv
var sparqlUrl = "http://178.62.59.88:31392/mondial"
var keyLimit = "100"

func main() {
	writeBarchartMapping(os.Stdout)
	writeScatterMapping(os.Stdout)
	writeBubbleMapping(os.Stdout)
}

func writeBarchartMapping(w io.Writer) {
	b := `SELECT ?entity ?key ?scalar WHERE {
  ?entity <http://dooodle/predicate/hasColumn> ?scalar .
  ?entity <http://dooodle/predicate/hasColumn> ?key .
  ?key <http://dooodle/predicate/numDistinct> ?num .
  ?scalar <http://dooodle/predicate/hasDimension> <http://dooodle/dimension/scalar> .
  FILTER (?num > 0 && ?num <= `+ keyLimit +`)
  FILTER (?key != ?scalar)
} 
LIMIT 200`
	body := bytes.NewReader([]byte(b))
	req, err := http.NewRequest("POST", sparqlUrl, body)
	req.Header.Set("Content-type", "application/sparql-query")
	req.Header.Set("Accept", mimeFormat(currentFormat))
	if err != nil {
		log.Fatal(err)
	}
	//req.SetBasicAuth("","")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	io.Copy(w, resp.Body)
}

func writeScatterMapping(w io.Writer) {
	b := `SELECT ?entity ?scalarA ?scalarB WHERE {
  ?entity <http://dooodle/predicate/hasColumn> ?scalarA .
  ?entity <http://dooodle/predicate/hasColumn> ?scalarB .
  ?scalarA <http://dooodle/predicate/hasDimension> <http://dooodle/dimension/scalar> .
  ?scalarB <http://dooodle/predicate/hasDimension> <http://dooodle/dimension/scalar> .
  FILTER (?scalarA != ?scalarB)
} 
LIMIT 200`
	body := bytes.NewReader([]byte(b))
	req, err := http.NewRequest("POST", sparqlUrl, body)
	req.Header.Set("Content-type", "application/sparql-query")
	req.Header.Set("Accept", mimeFormat(currentFormat))
	if err != nil {
		log.Fatal(err)
	}
	//req.SetBasicAuth("","")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	io.Copy(w, resp.Body)
}

func writeBubbleMapping(w io.Writer) {
	b := `SELECT ?entity ?scalarA ?scalarB ?scalarC WHERE {
  ?entity <http://dooodle/predicate/hasColumn> ?scalarA .
  ?entity <http://dooodle/predicate/hasColumn> ?scalarB .
  ?entity <http://dooodle/predicate/hasColumn> ?scalarC .
  ?scalarA <http://dooodle/predicate/hasDimension> <http://dooodle/dimension/scalar> .
  ?scalarB <http://dooodle/predicate/hasDimension> <http://dooodle/dimension/scalar> .
  ?scalarC <http://dooodle/predicate/hasDimension> <http://dooodle/dimension/scalar> .
  FILTER (?scalarA != ?scalarB)
  FILTER (?scalarA != ?scalarC)
  FILTER (?scalarB != ?scalarC)
} 
LIMIT 200`
	body := bytes.NewReader([]byte(b))
	req, err := http.NewRequest("POST", sparqlUrl, body)
	req.Header.Set("Content-type", "application/sparql-query")
	req.Header.Set("Accept", mimeFormat(currentFormat))
	if err != nil {
		log.Fatal(err)
	}
	//req.SetBasicAuth("","")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	io.Copy(w, resp.Body)
}

func mimeFormat(cf Format) string {
	switch cf {
	case Csv:
		return "text/csv"
	case Json:
		return "application/sparql-results+json"
	case Tsv:
		return "text/tab-separated-values"
	case Xml:
		return "application/sparql-results+xml"
	default:
		return "application/sparql-results+xml"
	}
}
