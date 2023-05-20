package project

import "go-to-cloud/internal/models/scm"

type SourceCodeBranch struct {
	Branches []scm.Branch `json:"branches"`
}
