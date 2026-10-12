package go_resume_questionnaire

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// TestExamplesUnmarshal checks that the published examples decode into the generated types.
func TestExamplesUnmarshal(t *testing.T) {
	// Each path is an example document and the generated value that should accept it.
	cases := []struct {
		path string
		dest any
	}{
		{path: "examples/export.minimal.json", dest: &ResumeQuestionnaireExportV1{}},
		{path: "examples/export.full.json", dest: &ResumeQuestionnaireExportV1{}},
		{path: "examples/fillable-15.json", dest: &ResumeQuestionnaireExportV1{}},
		{path: "examples/linked-projects/riverbank-senior-engineer-2021-present.json", dest: &ResumeQuestionnaireExportV1{}},
		{path: "examples/linked-projects/projects/ci-migration.json", dest: &ResumeQuestionnaireProjectFragmentV1{}},
		{path: "examples/linked-projects/projects/ledger-read-api.json", dest: &ResumeQuestionnaireProjectFragmentV1{}},
		{path: "examples/linked-projects/projects/refund-hardening.json", dest: &ResumeQuestionnaireProjectFragmentV1{}},
	}

	// Decode every example through the generated unmarshaler.
	for _, tc := range cases {
		// Read the example from the schema submodule.
		data, err := os.ReadFile(filepath.Join("resume-questionnaire", tc.path))
		if err != nil {
			t.Fatalf("read %s: %v", tc.path, err)
		}

		// Required fields and schemaVersion consts are enforced by the generated code.
		if err := json.Unmarshal(data, tc.dest); err != nil {
			t.Fatalf("unmarshal %s: %v", tc.path, err)
		}
	}
}
