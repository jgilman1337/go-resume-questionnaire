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

// TestSampleQuestionnaires checks that the job and project blanks decode.
func TestSampleQuestionnaires(t *testing.T) {
	// Each sample is a fill-in questionnaire for one resume entry kind.
	paths := []string{
		"examples/job.sample.json",
		"examples/project.sample.json",
	}

	// Decode every sample through the generated unmarshaler.
	for _, path := range paths {
		// Read the sample from the schema submodule.
		data, err := os.ReadFile(filepath.Join("resume-questionnaire", path))
		if err != nil {
			t.Fatalf("read %s: %v", path, err)
		}

		var questionnaire Questionnaire
		if err := json.Unmarshal(data, &questionnaire); err != nil {
			t.Fatalf("unmarshal %s: %v", path, err)
		}

		// A blank form still has an id and its question list.
		if questionnaire.Id == "" || len(questionnaire.Questions) != 12 {
			t.Fatalf("%s: id %q, %d questions", path, questionnaire.Id, len(questionnaire.Questions))
		}
	}
}
