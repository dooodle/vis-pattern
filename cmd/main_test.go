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

//TestWeakLine tests that the SPARQL query encapsulated by the method to get all the matches
//for the Weak Line Visualisation Pattern returns an entity and two columns in that order
func TestWeakLine(t *testing.T) {
	buf := bytes.Buffer{}
	writeWeakLineMapping(&buf)
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

//TestBubble tests that the SPARQL query encapsulated by the method to get all the matches
//for the Bubble  Visualisation Pattern returns an entity and four columns in that order
func TestBubble(t *testing.T) {
	buf := bytes.Buffer{}
	writeBubbleMapping(&buf)
	rows := bytes.Split(buf.Bytes(), []byte("\n"))
	tuple := bytes.Split(rows[1],[]byte(","))
	sub, col1, col2, col3, col4 := tuple[0],tuple[1],tuple[2],tuple[4],tuple[4]

	if !entity.Match(sub) {
		t.Errorf("expected IRI %s for entity but got %s",entity.String() ,sub)
	}
	if !column.Match(col1) {
		t.Errorf("expected IRI %s for column but got %s",column.String() ,col1)
	}
	if !column.Match(col2) {
		t.Errorf("expected IRI %s for column but got %s",column.String() ,col2)
	}
	if !column.Match(col3) {
		t.Errorf("expected IRI %s for column but got %s",column.String() ,col3)
	}
	if !column.Match(col4) {
		t.Errorf("expected IRI %s for column but got %s",column.String() ,col4)
	}
}

//TestChord tests that the SPARQL query encapsulated by the method to get all the matches
//for the Chord Visualisation Pattern returns an entity and three columns in that order
func TestChord(t *testing.T) {
	buf := bytes.Buffer{}
	writeM2mChordMapping(&buf)
	rows := bytes.Split(buf.Bytes(), []byte("\n"))
	tuple := bytes.Split(rows[1],[]byte(","))
	sub, col1, col2, col3 := tuple[0],tuple[1],tuple[2],tuple[3]

	if !entity.Match(sub) {
		t.Errorf("expected IRI %s for entity but got %s",entity.String() ,sub)
	}
	if !column.Match(col1) {
		t.Errorf("expected IRI %s for column but got %s",column.String() ,col1)
	}
	if !column.Match(col2) {
		t.Errorf("expected IRI %s for column but got %s",column.String() ,col2)
	}
	if !column.Match(col3) {
		t.Errorf("expected IRI %s for column but got %s",column.String() ,col3)
	}
}

//TestCircle tests that the SPARQL query encapsulated by the method to get all the matches
//for the Circle Packing  Visualisation Pattern returns an entity and two columns in that order
func TestCircle(t *testing.T) {
	buf := bytes.Buffer{}
	writeO2mCircleMapping(&buf)
	rows := bytes.Split(buf.Bytes(), []byte("\n"))
	tuple := bytes.Split(rows[1],[]byte(","))
	sub, col1, col2 := tuple[0],tuple[1],tuple[2]

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

//TestScatter tests that the SPARQL query encapsulated by the method to get all the matches
//for the Scatter  Visualisation Pattern returns an entity and three columns in that order
func TestScatter(t *testing.T) {
	buf := bytes.Buffer{}
	writeScatterMapping(&buf)
	rows := bytes.Split(buf.Bytes(), []byte("\n"))
	tuple := bytes.Split(rows[1],[]byte(","))
	sub, col1, col2, col3 := tuple[0],tuple[1],tuple[2],tuple[3]

	if !entity.Match(sub) {
		t.Errorf("expected IRI %s for entity but got %s",entity.String() ,sub)
	}
	if !column.Match(col1) {
		t.Errorf("expected IRI %s for column but got %s",column.String() ,col1)
	}
	if !column.Match(col2) {
		t.Errorf("expected IRI %s for column but got %s",column.String() ,col2)
	}
	if !column.Match(col3) {
		t.Errorf("expected IRI %s for column but got %s",column.String() ,col3)
	}
}



