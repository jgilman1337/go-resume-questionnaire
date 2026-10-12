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
schema_questionnaire: SCHEMA_PKG := resume-questionnaire
schema_questionnaire:
	$(GO_JSONSCHEMA) -e \
		--schema-package=$(REPO_ROOT)/questionnaire.schema.json=$(SCHEMA_GO_PKG) \
		--schema-output=$(REPO_ROOT)/questionnaire.schema.json=$(SCHEMA_AUTOGEN_DIR)/questionnaire.go \
		--schema-root-type=$(REPO_ROOT)/questionnaire.schema.json=Questionnaire \
		$(SCHEMA_DIR)/$(SCHEMA_PKG)/questionnaire.schema.json
	# go-jsonschema hardcodes gopkg.in/yaml.v3. Rewrite it to the YAML org fork.
	sed -i 's|gopkg.in/yaml.v3|go.yaml.in/yaml/v4|g' \
		$(SCHEMA_AUTOGEN_DIR)/questionnaire.go
