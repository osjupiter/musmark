package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/hoisie/mustache"
	"gopkg.in/yaml.v3"
)

func main() {
	// Change default behavior to updateMode and add resultMode flag
	resultMode := flag.Bool("result", false, "Output only the result")
	flag.BoolVar(resultMode, "r", false, "Output only the result")
	flag.Parse()

	args := flag.Args()

	var content []byte
	var err error

	switch {
	case len(args) >= 1:
		content, err = os.ReadFile(args[0])
	default:
		content, err = io.ReadAll(os.Stdin)
	}

	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	tmplStr, dataStr, err := parseMusmarkFile(string(content))
	if err != nil {
		log.Fatalf("Failed to parse file: %v", err)
	}

	output, err := renderMusmarkTemplate(tmplStr, dataStr)
	if err != nil {
		log.Fatalf("Failed to render template: %v", err)
	}

	// Check if input has existing result block
	resultRegex := regexp.MustCompile("(?s)result\\s*\\n===\\s*\\n```\\s*\\n.*?\\s*\\n```")
	hasResult := resultRegex.MatchString(string(content))

	var finalOutput string
	if *resultMode {
		// Output only the result
		finalOutput = output
	} else {
		if hasResult {
			// Replace existing result block
			finalOutput = string(content)
			finalOutput = resultRegex.ReplaceAllString(finalOutput, fmt.Sprintf("result\n===\n```\n%s\n```", output))
		} else {
			// Append new result block
			result := fmt.Sprintf("\n\nresult\n===\n```\n%s\n```", output)
			finalOutput = fmt.Sprintf("template\n===\n```mustache\n%s\n```\n\ndata\n===\n```yaml\n%s\n```%s", tmplStr, dataStr, result)
		}
	}

	// Write to file if specified, otherwise print to stdout
	if len(args) >= 1 {
		err = os.WriteFile(args[0], []byte(finalOutput), 0644)
		if err != nil {
			log.Fatalf("Failed to write output: %v", err)
		}
	} else {
		fmt.Println(finalOutput)
	}
}

func parseMusmarkFile(content string) (template string, yamlData string, err error) {
	tmplRegex := regexp.MustCompile("(?s)[tT]emplate\\s*\\n===\\s*\\n```mustache\\s*\\n(.*?)\\s*\\n```")
	dataRegex := regexp.MustCompile("(?s)[dD]ata\\s*\\n===\\s*\\n```yaml\\s*\\n(.*?)\\s*\\n```")

	tmplMatch := tmplRegex.FindStringSubmatch(content)
	dataMatch := dataRegex.FindStringSubmatch(content)

	if len(tmplMatch) < 2 || len(dataMatch) < 2 {
		return "", "", fmt.Errorf("Invalid musmark file format")
	}

	return strings.TrimSpace(tmplMatch[1]), strings.TrimSpace(dataMatch[1]), nil
}

func renderMusmarkTemplate(tmpl string, data string) (string, error) {
	var parsedData map[string]interface{}
	if err := yaml.Unmarshal([]byte(data), &parsedData); err != nil {
		return "", err
	}

	return mustache.Render(tmpl, parsedData), nil
}
