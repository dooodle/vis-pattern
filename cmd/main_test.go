package main

import (
	"bytes"
	"regexp"
	"testing"
)

//Regex for the various IRI's defined for conceptual model
var entity = regexp.MustCompile("http://dooodle/entity/[[:word:]]+")
var column = regexp.MustCompile("http://dooodle/entity/[[:word:]]+/column/[[:word:]]+")

//TestBar tests that the SPARQL query encapsulated by the method to get all the matches
//for the Barchart Visualisation Pattern returns an entity and two columns in that order
func TestBar(t *testing.T) {
	buf := bytes.Buffer{}
	writeBarchartMapping(&buf)
	rows := bytes.Split(buf.Bytes(), []byte("\n"))
	triple := bytes.Split(rows[1],[]byte(","))
	sub, col1, col2 := triple[0],triple[1],triple[2]

	if !entity.Match(sub) {
		t.Errorf("expected IRI %s for entity but got %s",entity.String() ,sub)
	}
	if !column.Match(col1) {
		t.Errorf("expected IRI %s for column but got %s",column.String() ,col1)
	}
	if !column.Match(col2) {
		t.Errorf("expected IRI %s for column but got %s",column.String() ,col2)
	}
}
