package go_resume_questionnaire

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"go.yaml.in/yaml/v4"
)

// TestQuestionnaireUnmarshalJSON accepts meta, an entry id, and question-answer pairs.
func TestQuestionnaireUnmarshalJSON(t *testing.T) {
	// A minimal document matches the single questionnaire schema.
	input := []byte(`{
		"meta": {"version": "1.0.0"},
		"id": "refund-hardening",
		"questions": [
			{"question": "What did you ship?", "answer": "Cut refund latency."}
		]
	}`)

	var questionnaire Questionnaire
	if err := json.Unmarshal(input, &questionnaire); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	// The join key is the resume entry id.
	if questionnaire.Id != "refund-hardening" {
		t.Fatalf("id: got %q", questionnaire.Id)
	}

	// One answer pair is preserved.
	if len(questionnaire.Questions) != 1 || questionnaire.Questions[0].Question != "What did you ship?" {
		t.Fatalf("questions: got %#v", questionnaire.Questions)
	}
}

// TestQuestionnaireUnmarshalYAML decodes the same document with go.yaml.in/yaml/v4.
func TestQuestionnaireUnmarshalYAML(t *testing.T) {
	// YAML uses the same three fields as JSON.
	input := []byte(`
meta:
  version: "1.0.0"
id: riverbank
questions:
  - question: What was your role?
    answer: Senior engineer.
`)

	var questionnaire Questionnaire
	if err := yaml.Unmarshal(input, &questionnaire); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	// The join key is the resume entry id.
	if questionnaire.Id != "riverbank" {
		t.Fatalf("id: got %q", questionnaire.Id)
	}

	// One answer pair is preserved.
	if len(questionnaire.Questions) != 1 || questionnaire.Questions[0].Answer != "Senior engineer." {
		t.Fatalf("questions: got %#v", questionnaire.Questions)
	}
}

// TestExampleJSON decodes every example document against the generated schema.
func TestExampleJSON(t *testing.T) {
	root := filepath.Join("resume-questionnaire", "examples")

	// Walk the examples tree so new JSON files are checked without a hand-maintained list.
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, walkErr error) error {
		// A walk error stops the scan.
		if walkErr != nil {
			return walkErr
		}

		// Directories and non-JSON notes are not questionnaires.
		if entry.IsDir() || filepath.Ext(path) != ".json" {
			return nil
		}

		// Read the example from the schema submodule.
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read %s: %v", path, err)
		}

		var questionnaire Questionnaire
		if err := json.Unmarshal(data, &questionnaire); err != nil {
			t.Fatalf("unmarshal %s: %v", path, err)
		}

		// Every example needs a join key and at least one question.
		if questionnaire.Id == "" || len(questionnaire.Questions) == 0 {
			t.Fatalf("%s: id %q, %d questions", path, questionnaire.Id, len(questionnaire.Questions))
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
