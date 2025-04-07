package main

import (
	"strings"
	"testing"
)

func TestParseMusmarkFile(t *testing.T) {
	content := `
template
===
@@@mustache
Hello, {{name}}!
@@@

data
===
@@@yaml
name: World
@@@
`
	content = strings.ReplaceAll(content, "@@@", "```")
	tmpl, data, err := parseMusmarkFile(content)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedTmpl := "Hello, {{name}}!"
	expectedData := "name: World"

	if tmpl != expectedTmpl {
		t.Errorf("Expected template %q, got %q", expectedTmpl, tmpl)
	}

	if data != expectedData {
		t.Errorf("Expected data %q, got %q", expectedData, data)
	}
}

func TestParseMusmarkFileInvalidFormat(t *testing.T) {
	content := `
template
===
@@@mustache
Hello, {{name}}!
@@@

data
===
INVALID YAML
`
	content = strings.ReplaceAll(content, "@@@", "```")
	_, _, err := parseMusmarkFile(content)
	if err == nil {
		t.Fatal("Expected an error for invalid format, got none")
	}
}

func TestRenderMusmarkTemplate(t *testing.T) {
	tmpl := "Hello, {{name}}!"
	data := "name: World"

	result, err := renderMusmarkTemplate(tmpl, data)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := "Hello, World!"
	if result != expectedResult {
		t.Errorf("Expected rendered result %q, got %q", expectedResult, result)
	}
}
