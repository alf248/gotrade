package forms

import (
	"errors"
	"strings"
)

type VectorSearch struct {
	Search string `json:"search,omitempty"` // a search query
	Max    int    `json:"max,omitempty"`    // the max amount of results
	Page   int    `json:"page,omitempty"`   // search pagination
	SortBy string `json:"sortBy,omitempty"`
	SortUp bool   `json:"sortUp,omitempty"`
	Active bool   `json:"active,omitempty"`
}

func (o *VectorSearch) Curate() error {
	return nil
}

type VectorAction struct {
	Action string `json:"action"`
}

func (o *VectorAction) Curate() error {

	actions := []string{"delete", "pause"}

	for _, action := range actions {
		if o.Action == action {
			return nil
		}
	}

	return errors.New("action must be one of " + strings.Join(actions, " "))
}
