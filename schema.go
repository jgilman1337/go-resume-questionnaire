package go_resume_questionnaire

import (
	"embed"
	"encoding/json"
	"fmt"
	"strings"
)

// SchemaFS holds the JSON Schema documents this package is generated from.
//
//go:embed resume-questionnaire/questionnaire.schema.json
var SchemaFS embed.FS

const (
	// QuestionnaireSchemaPath is questionnaire.schema.json inside SchemaFS.
	QuestionnaireSchemaPath = "resume-questionnaire/questionnaire.schema.json"

	localRefPrefix = "#/definitions/"
)

// QuestionnaireJSONSchema returns questionnaire.schema.json with local $ref values inlined.
func QuestionnaireJSONSchema() (map[string]any, error) {
	doc, err := readJSONObject(QuestionnaireSchemaPath)
	if err != nil {
		return nil, err
	}

	defs, _ := doc["definitions"].(map[string]any)
	if defs == nil {
		defs = map[string]any{}
	}

	resolved, err := resolveRefs(doc, defs)
	if err != nil {
		return nil, err
	}

	schema, ok := resolved.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("%s is not an object", QuestionnaireSchemaPath)
	}

	return schema, nil
}

// readJSONObject loads an embedded JSON object.
func readJSONObject(name string) (map[string]any, error) {
	raw, err := SchemaFS.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", name, err)
	}

	var doc map[string]any
	if err := json.Unmarshal(raw, &doc); err != nil {
		return nil, fmt.Errorf("decode %s: %w", name, err)
	}
	if doc == nil {
		return nil, fmt.Errorf("%s is empty", name)
	}

	return doc, nil
}

// resolveRefs replaces JSON Schema $ref nodes with the referenced definition.
func resolveRefs(node any, defs map[string]any) (any, error) {
	switch typed := node.(type) {
	case map[string]any:
		// A $ref node is replaced by the definition it names.
		ref, ok := typed["$ref"].(string)
		if ok {
			return resolveRefNode(typed, ref, defs)
		}

		out := make(map[string]any, len(typed))
		for key, child := range typed {
			resolved, err := resolveRefs(child, defs)
			if err != nil {
				return nil, err
			}
			out[key] = resolved
		}
		return out, nil

	case []any:
		out := make([]any, len(typed))
		for i, child := range typed {
			resolved, err := resolveRefs(child, defs)
			if err != nil {
				return nil, err
			}
			out[i] = resolved
		}
		return out, nil

	default:
		return node, nil
	}
}

// resolveRefNode inlines a $ref and overlays sibling keywords such as description.
func resolveRefNode(node map[string]any, ref string, defs map[string]any) (any, error) {
	name, err := definitionName(ref)
	if err != nil {
		return nil, err
	}

	def, ok := defs[name]
	if !ok {
		return nil, fmt.Errorf("unknown schema ref %q", ref)
	}

	resolved, err := resolveRefs(def, defs)
	if err != nil {
		return nil, err
	}

	base, ok := resolved.(map[string]any)
	// A non-object definition has no sibling keywords to merge.
	if !ok {
		return resolved, nil
	}

	out := make(map[string]any, len(base)+len(node)-1)
	for key, value := range base {
		out[key] = value
	}
	for key, value := range node {
		// $ref is the pointer that was just inlined.
		if key == "$ref" {
			continue
		}
		out[key] = value
	}

	return out, nil
}

// definitionName extracts the definition key from a local $ref.
func definitionName(ref string) (string, error) {
	if strings.HasPrefix(ref, localRefPrefix) {
		return strings.TrimPrefix(ref, localRefPrefix), nil
	}

	return "", fmt.Errorf("unsupported schema ref %q", ref)
}
