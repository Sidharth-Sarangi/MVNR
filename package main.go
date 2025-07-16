package main
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func titleCase(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func main() {
	file, err := os.Open("TS33128Payloads.asn")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	reStructStart := regexp.MustCompile(`^(\w+)\s+::=\s+SEQUENCE\s*{`)
	reField := regexp.MustCompile(`^\s*([\w-]+)\s+\[(\d+)\]\s+([\w-]+)(\s+OPTIONAL)?\s*,?$`)
	inStruct := false
	structName := ""
	var fields []string

	for scanner.Scan() {
		line := scanner.Text()
		if !inStruct {
			if m := reStructStart.FindStringSubmatch(line); m != nil {
				inStruct = true
				structName = m[1]
				fields = []string{}
			}
		} else {
			if strings.TrimSpace(line) == "}" {
				// Emit the struct
				fmt.Printf("type %s struct {\n", structName)
				for _, f := range fields {
					fmt.Println(f)
				}
				fmt.Println("}\n")
				inStruct = false
				structName = ""
				continue
			}
			if m := reField.FindStringSubmatch(line); m != nil {
				fieldName := titleCase(m[1])
				fieldType := m[3]
				isOptional := strings.Contains(m[4], "OPTIONAL")
				if isOptional {
					fields = append(fields, fmt.Sprintf("    %s *%s", fieldName, fieldType))
				} else {
					fields = append(fields, fmt.Sprintf("    %s %s", fieldName, fieldType))
				}
			}
		}
	}
}
