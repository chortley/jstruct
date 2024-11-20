package generator

import (
	"encoding/json"
	"fmt"
	"strings"
)

func GenerateStruct(jsonData, structName string) (string, error) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		return "", fmt.Errorf("invalid JSON: %w", err)
	}

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("type %s struct {\n", structName))

	for key, value := range data {
		fieldName := toCamelCase(key)
		fieldType, nestedStruct := inferGoType(value, key)

		if nestedStruct != "" {
			nestedStructDefinition, err := GenerateStruct(toJSON(value), nestedStruct)
			if err != nil {
				return "", err
			}
			builder.WriteString(fmt.Sprintf("    %s %s `json:\"%s\"`\n", fieldName, nestedStruct, key))
			builder.WriteString(nestedStructDefinition) // Include the nested struct definition here
		} else {
			builder.WriteString(fmt.Sprintf("    %s %s `json:\"%s\"`\n", fieldName, fieldType, key))
		}
	}
	builder.WriteString("}")
	return builder.String(), nil
}

func toCamelCase(input string) string {
	parts := strings.Split(input, "_")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}

func inferGoType(value interface{}, key string) (string, string) {
	switch value.(type) {
	case string:
		return "string", ""
	case float64:
		return "float64", ""
	case bool:
		return "bool", ""
	case nil:
		return "interface{}", ""
	case map[string]interface{}:
		structName := fmt.Sprintf("%sStruct", toCamelCase(key)) 
		return "interface{}", structName 
	default:
		return "interface{}", ""
	}
}

func toJSON(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(data)
}

