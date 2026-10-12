### SIMPLE MAKEFILE TEMPLATE FOR GO PROJECTS ###

NAME := go-resume-questionnaire
SCHEMA_DIR := ./
SCHEMA_AUTOGEN_DIR := ./

# ----
# DO NOT EDIT BELOW THIS LINE
# ----
GO := go

# go-jsonschema is pinned as a tool in go.mod.
# See: https://github.com/atombender/go-jsonschema
GO_JSONSCHEMA := $(GO) tool go-jsonschema

.PHONY: schema_questionnaire
schema_questionnaire: REPO_ROOT := https://raw.githubusercontent.com/jgilman1337/resume-questionnaire/refs/heads/master
schema_questionnaire: SCHEMA_GO_PKG := go_resume_questionnaire
schema_questionnaire: SCHEMA_PKG := resume-questionnaire/schemas
schema_questionnaire:
	$(GO_JSONSCHEMA) -e \
		--schema-package=$(REPO_ROOT)/schemas/resume-questionnaire-export.schema.json=$(SCHEMA_GO_PKG) \
		--schema-output=$(REPO_ROOT)/schemas/resume-questionnaire-export.schema.json=$(SCHEMA_AUTOGEN_DIR)/export.go \
		--schema-root-type=$(REPO_ROOT)/schemas/resume-questionnaire-export.schema.json=ResumeQuestionnaireExportV1 \
		--schema-package=$(REPO_ROOT)/schemas/resume-project-fragment.schema.json=$(SCHEMA_GO_PKG) \
		--schema-output=$(REPO_ROOT)/schemas/resume-project-fragment.schema.json=$(SCHEMA_AUTOGEN_DIR)/project.go \
		--schema-root-type=$(REPO_ROOT)/schemas/resume-project-fragment.schema.json=ResumeQuestionnaireProjectFragmentV1 \
		--schema-package=$(REPO_ROOT)/schemas/resume-questionnaire-v2.schema.json=$(SCHEMA_GO_PKG) \
		--schema-output=$(REPO_ROOT)/schemas/resume-questionnaire-v2.schema.json=$(SCHEMA_AUTOGEN_DIR)/questionnaire.go \
		--schema-root-type=$(REPO_ROOT)/schemas/resume-questionnaire-v2.schema.json=ResumeQuestionnaireV2 \
		$(SCHEMA_DIR)/$(SCHEMA_PKG)/resume-questionnaire-export.schema.json \
		$(SCHEMA_DIR)/$(SCHEMA_PKG)/resume-project-fragment.schema.json \
		$(SCHEMA_DIR)/$(SCHEMA_PKG)/resume-questionnaire-v2.schema.json
