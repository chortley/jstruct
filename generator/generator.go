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
		fieldType := inferGoType(value)
		builder.WriteString(fmt.Sprintf("    %s %s `json:\"%s\"`\n", fieldName, fieldType, key))
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

func inferGoType(value interface{}) string {
	switch v := value.(type) {
	case string:
		return "string"
	case float64:
		if v == float64(int(v)) {
			return "int"
		}
		return "float64"
	case bool:
		return "bool"
	case []interface{}:
		return "[]interface{}"
	case map[string]interface{}:
		return "map[string]interface{}"
	default:
		return "interface{}"
	}
}


