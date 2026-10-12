# go-resume-questionnaire

This repository generates Go types from the resume questionnaire JSON Schema using `go-jsonschema`.

The questionnaire package is a submodule at `resume-questionnaire`. The generated file is committed to this repo:

- `questionnaire.go` (generated from `resume-questionnaire/questionnaire.schema.json`)

A questionnaire is resume `meta` (copied from the resume schema), the `id` of one work, project, or volunteer entry (or a stable id for the cover letter), a `role` of `project`, `work`, `volunteer`, or `coverletter`, and a `questions` array of `{ "question", "answer" }`.

YAML unmarshalling uses `go.yaml.in/yaml/v4`.

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
