package go_resume_questionnaire

import "testing"

// TestQuestionnaireJSONSchemaLoads checks the embedded questionnaire schema inlines local refs.
func TestQuestionnaireJSONSchemaLoads(t *testing.T) {
	schema, err := QuestionnaireJSONSchema()
	if err != nil {
		t.Fatal(err)
	}

	if schema["type"] != "object" {
		t.Fatalf("type: %#v", schema["type"])
	}

	props, ok := schema["properties"].(map[string]any)
	if !ok {
		t.Fatal("properties is missing")
	}

	role, ok := props["role"].(map[string]any)
	if !ok {
		t.Fatal("role is missing")
	}
	enum, ok := role["enum"].([]any)
	if !ok {
		t.Fatalf("role.enum: %#v", role["enum"])
	}

	found := false
	for _, value := range enum {
		if value == "work" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("role.enum: %#v", enum)
	}

	meta, ok := props["meta"].(map[string]any)
	if !ok {
		t.Fatal("meta is missing")
	}
	if _, ok := meta["$ref"]; ok {
		t.Fatal("meta still has a $ref")
	}
}
