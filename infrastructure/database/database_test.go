package database_test

import (
	"bytes"
	"flag"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"go-bootcamp/infrastructure/database"
)

var update = flag.Bool("update", false, "update golden files")

func TestRead(t *testing.T) {
	var cases = map[string]struct{ Name string }{
		"read file": {"db-read"},
	}

	for k, c := range cases {
		content := getByteContent(c.Name, k, t)
		expected := getGoldenFile(c.Name, k, content, t)
		if !bytes.Equal(content, expected) {
			t.Errorf(
				"[case '%s']: db.Read('%s.csv')\nexpected\n-----\n%s\n-----\nbut got\n-----\n%s\n-----\n",
				k, c.Name, expected, content)
		}
	}
}

func TestWrite(t *testing.T) {
	var cases = map[string]struct {
		Name  string
		Value []string
	}{
		"write file": {"db-write", []string{"1", "bulbasaur"}},
	}

	for k, c := range cases {
		content := writeContent(c.Name, k, c.Value, t)
		expected := getGoldenFile(c.Name, k, content, t)
		if !bytes.Equal(content, expected) {
			t.Errorf(
				"[case '%s']: db.Write('%s.csv', '%s')\nexpected\n-----\n%s\n-----\nbut got\n-----\n%s\n-----\n",
				k, c.Name, c.Value, expected, content)
		}
	}
}

func TestWriteAll(t *testing.T) {
	var cases = map[string]struct {
		Name  string
		Value [][]string
	}{
		"write file": {"db-write-all", [][]string{
			{"1", "bulbasaur"},
			{"2", "ivysaur"},
			{"3", "venosaur"},
		}},
	}

	for k, c := range cases {
		content := writeAllContent(c.Name, k, c.Value, t)
		expected := getGoldenFile(c.Name, k, content, t)
		if !bytes.Equal(content, expected) {
			t.Errorf(
				"[case '%s']: db.WriteAll('%s.csv', '%s')\nexpected\n-----\n%s\n-----\nbut got\n-----\n%s\n-----\n",
				k, c.Name, c.Value, expected, content)
		}
	}
}

func getByteContent(name, k string, t *testing.T) []byte {
	actual := filepath.Join("testdata", name+".csv")
	result, err := database.New(actual).Read()
	if err != nil {
		t.Fatalf(
			"[case '%s']: db.Read(%s), unexpected error reading file: '%s'",
			k, actual, err)
	}
	rows := []string{}
	for _, row := range result {
		rows = append(rows, strings.Join(row, ","))
	}

	return []byte(strings.Join(rows, "\n"))
}

func readFile(name, k string, t *testing.T) []byte {
	bytes, err := ioutil.ReadFile(name)
	if err != nil {
		t.Fatalf(
			"[case '%s']: unexpected error reading file '%s': '%s'",
			k, name, err)
	}
	return bytes
}

func writeFile(name, k string, content []byte, t *testing.T) {
	err := ioutil.WriteFile(name, content, 0644)
	if err != nil {
		t.Fatalf(
			"[case '%s']: unexpected error writing file (%s): '%s'",
			k, name, err)
	}
}

func getGoldenFile(name, k string, content []byte, t *testing.T) []byte {
	golden := filepath.Join("testdata", name+".golden")
	if *update {
		writeFile(golden, k, content, t)
	}
	expected := readFile(golden, k, t)
	return expected
}

func contentThenReset(name, k string, t *testing.T) []byte {
	content := readFile(name, k, t)
	writeFile(name, k, []byte{}, t)
	return content
}

func writeContent(name, k string, value []string, t *testing.T) []byte {
	actual := filepath.Join("testdata", name+".csv")
	err := database.New(actual).Write(value)
	if err != nil {
		t.Fatalf(
			"[case '%s']: db.Write(%s, %s), unexpected error writting file: '%s'",
			k, actual, value, err)
	}
	content := contentThenReset(actual, k, t)
	return content
}

func writeAllContent(name, k string, value [][]string, t *testing.T) []byte {
	actual := filepath.Join("testdata", name+".csv")
	err := database.New(actual).WriteAll(value)
	if err != nil {
		t.Fatalf(
			"[case '%s']: db.Write(%s, %s), unexpected error writting file: '%s'",
			k, actual, value, err)
	}
	content := contentThenReset(actual, k, t)
	return content
}
