# go-resume-questionnaire

This repository generates Go types from the resume questionnaire JSON Schemas using `go-jsonschema`.

The questionnaire package is a submodule at `resume-questionnaire`. The generated files are committed to this repo:

- `export.go` (generated from `resume-questionnaire/schemas/resume-questionnaire-export.schema.json`)
- `project.go` (generated from `resume-questionnaire/schemas/resume-project-fragment.schema.json`)
- `questionnaire.go` (generated from `resume-questionnaire/schemas/resume-questionnaire-v2.schema.json`)

## Prerequisites

- Go (see `go.mod`)
- `make`

## Regenerating schema code

```bash
make schema_questionnaire
```

## Pinning the `go-jsonschema` generator version

The generator is a Go tool in `go.mod`. To bump it:

```bash
go get -tool github.com/atombender/go-jsonschema@vX.Y.Z
make schema_questionnaire
```
