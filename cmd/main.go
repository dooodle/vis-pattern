package main

import (
	"bytes"
	"flag"
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

var serve = flag.Bool("serve", false, "serve")
var port = flag.String("http", ":8080", "port")

var currentFormat = Csv
var sparqlUrl = "http://178.62.59.88:31392/mondial"
var keyLimit = "300"

func main() {
	flag.Parse()
	if !*serve {
		writeBarchartMapping(os.Stdout)
		writeScatterMapping(os.Stdout)
		writeBubbleMapping(os.Stdout)
		writeWeakLineMapping(os.Stdout)
		writeO2mCircleMapping(os.Stdout)
		os.Exit(0)
	}
	// serve
	// could add in a caching layer here...
	http.HandleFunc("/mondial/basic/bar", func(w http.ResponseWriter, r *http.Request) {
		log.Println("serving mondial basic bar chart mappings")
		writeBarchartMapping(w)
	})

	http.HandleFunc("/mondial/basic/scatter", func(w http.ResponseWriter, r *http.Request) {
		log.Println("serving mondial basic scatter mappings")
		writeScatterMapping(w)
	})

	http.HandleFunc("/mondial/basic/bubble", func(w http.ResponseWriter, r *http.Request) {
		log.Println("serving mondial basic bubble mappings")
		writeBubbleMapping(w)
	})

	http.HandleFunc("/mondial/weak/line", func(w http.ResponseWriter, r *http.Request) {
		log.Println("serving mondial weak line mappings")
		writeWeakLineMapping(w)
	})

	http.HandleFunc("/mondial/o2m/circle", func(w http.ResponseWriter, r *http.Request) {
		log.Println("serving mondial o2m circle packing mappings")
		writeO2mCircleMapping(w)
	})

	http.HandleFunc("/mondial/m2m/chord", func(w http.ResponseWriter, r *http.Request) {
		log.Println("serving mondial m2m chord packing mappings")
		writeM2mChordMapping(w)
	})

	log.Println("starting server at", *port)
	if err := http.ListenAndServe(*port, nil); err != nil {
		log.Fatal(err)
	}

}


var ChordLimit = "150"
func writeM2mChordMapping(w io.Writer) {
	b := `SELECT ?entity ?m1 ?m2 ?measure 
    WHERE {
	?entity <http://dooodle/predicate/hasMany2ManyKey> ?key ;
		    <http://dooodle/predicate/hasColumn> ?measure .   
	?key <http://dooodle/predicate/hasManyKey> ?m1 .
	?key <http://dooodle/predicate/hasManyKey> ?m2 .
	?m1 <http://dooodle/predicate/numDistinct> ?n1 .
	?m2 <http://dooodle/predicate/numDistinct> ?n2 .
	?measure <http://dooodle/predicate/hasDimension> <http://dooodle/dimension/scalar> .
	FILTER (?measure != ?m1)
	FILTER (?measure != ?m2)
	FILTER (?m1 != ?m2)
	FILTER (?n1 < `+ ChordLimit + `)
	FILTER (?n2 < `+ ChordLimit + `)
	}
	LIMIT 1000`

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

func writeO2mCircleMapping(w io.Writer) {
	b := `SELECT ?entity ?one ?many WHERE {
	?entity  <http://dooodle/predicate/hasOne2ManyKey> ?key .
	?key  <http://dooodle/predicate/hasOneKey> ?one .
	?key <http://dooodle/predicate/hasManyKey> ?many .
	?one <http://dooodle/predicate/numDistinct> ?numOne .
	?many <http://dooodle/predicate/numDistinct> ?numMany .	
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

func writeBarchartMapping(w io.Writer) {
	b := `SELECT ?entity ?key ?scalar WHERE {
  ?entity <http://dooodle/predicate/hasColumn> ?scalar .
  ?entity <http://dooodle/predicate/hasSingleKey> ?key .
  ?key <http://dooodle/predicate/numDistinct> ?num .
  ?scalar <http://dooodle/predicate/hasDimension> <http://dooodle/dimension/scalar> .
  FILTER (?num > 0 && ?num <= ` + keyLimit + `)
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
	b := `SELECT ?entity ?key ?scalarA ?scalarB WHERE {
  ?entity <http://dooodle/predicate/hasKey> ?key .
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
	b := `SELECT ?entity ?key ?scalarA ?scalarB ?scalarC WHERE {
  ?entity <http://dooodle/predicate/hasKey> ?key .
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

func writeWeakLineMapping(w io.Writer) {
	b := `SELECT ?entity ?strong ?weak ?measure WHERE {
    ?entity <http://dooodle/predicate/hasCompoundKey> ?key .
    ?key <http://dooodle/predicate/hasStrongKey> ?strong .
    ?key <http://dooodle/predicate/hasWeakKey> ?weak .
    ?entity <http://dooodle/predicate/hasColumn> ?measure .
    ?weak <http://dooodle/predicate/hasDataType> ?dt .
    ?measure <http://dooodle/predicate/hasDimension> <http://dooodle/dimension/scalar> .
    FILTER (?dt = <http://dooodle/dataType/numeric> || ?dt = <http://dooodle/dataType/int4>)
   }`
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
