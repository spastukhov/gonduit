package requests

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/uber/gonduit/entities"
)

// ManiphestSearchRequest represents a request to maniphest.search API method.
type ManiphestSearchRequest struct {
	// QueryKey is builtin or saved query to use. It is optional and sets initial constraints.
	QueryKey    string             `json:"queryKey,omitempty"`
	// Constraints contains additional filters for results. Applied on top of query if provided.
	Constraints *SearchConstraints `json:"constraints,omitempty"`
	// Attachments specified what additional data should be returned with each result.
	Attachments *SearchAttachments `json:"attachments,omitempty"`

	*entities.Cursor
	Request
}

// SearchAttachments contains fields that specify what additional data should be returned with sesarch results.
type SearchAttachments struct {
	// Subscribers if true instructs server to return subscribers list for each task.
	Subscribers bool `json:"subscribers,omitempty"`
	// Columns requests to get the workboard columns where an object appears.
	Columns     bool `json:"columns,omitempty"`
	// Projects requests to get information about projects.
	Projects    bool `json:"projects,omitempty"`
}

// ManiphestRequestSearchOrder describers how results should be ordered.
type ManiphestRequestSearchOrder struct {
	// Builtin is the name of predefined order to use.
	Builtin string
	// Order is list of columns to use for sorting, e.g. ["color", "-name", "id"],
	Order   []string
}

// UnmarshalJSON parses JSON  into an instand of ManiphestRequestSearchOrder type.
func (o *ManiphestRequestSearchOrder) UnmarshalJSON(data []byte) error {
	if o == nil {
		return errors.New("maniphest search order is nil")
	}
	if jerr := json.Unmarshal(data, &o.Builtin); jerr == nil {
		return nil
	}

	return json.Unmarshal(data, &o.Order)
}

// MarshalJSON creates JSON our of ManiphestRequestSearchOrder instance.
func (o *ManiphestRequestSearchOrder) MarshalJSON() ([]byte, error) {
	if o == nil {
		return nil, errors.New("maniphest search order is nil")
	}
	if o.Builtin != "" {
		return json.Marshal(o.Builtin)
	}
	if len(o.Order) > 0 {
		return json.Marshal(o.Order)
	}

	return nil, nil
}

// SearchConstraints describes search criteria for request.
type SearchConstraints struct {
	// IDs - search for objects with specific IDs.
	IDs            []string        `json:"ids,omitempty"`
	// PHIDs - search for objects with specific PHIDs.
	PHIDs          []string        `json:"phids,omitempty"`
	// AssignedTo - search for tasks owned by a user from a list.
	AssignedTo     []string        `json:"assigned,omitempty"`
	// Authors - search for tasks with given authors.
	Authors        []string        `json:"authorPHIDs,omitempty"`
	// Statuses - search for tasks with given statuses.
	Statuses       []string        `json:"statuses,omitempty"`
	// Priorities - search for tasks with given priorities.
	Priorities     []string        `json:"priorities,omitempty"`
	// Subtypes - search for tasks with given subtypes.
	Subtypes       []string        `json:"subtypes,omitempty"`
	// Column PHIDs ??? - no doc on phab site.
	ColumnPHIDs    []string        `json:"columnPHIDs,omitempty"`
	// OpenParents - search for tasks that have parents in open state.
	OpenParents    *bool           `json:"hasParents,omitempty"`
	// OpenSubtasks - search for tasks that have child tasks in open state.
	OpenSubtasks   *bool           `json:"hasSubtasks,omitempty"`
	// ParentIDs - search for children of these parents.
	ParentIDs      []int           `json:"parentIDs,omitempty"`
	// SubtaskIDs - Search for tasks that have these children.
	SubtaskIDs     []int           `json:"subtaskIDs,omitempty"`
	// CreatedAfter - search for tasks created after given date.
	CreatedAfter   *entities.Epoch `json:"createdStart,omitempty"`
	// CreatedBefore - search for tasks created before given date.
	CreatedBefore  *entities.Epoch `json:"createdEnd,omitempty"`
	// ModifiedAfter - search for tasks modified after given date.
	ModifiedAfter  *entities.Epoch `json:"modifiedStart,omitempty"`
	// ModifiedBefore - search for tasks modified before given date.
	ModifiedBefore *entities.Epoch `json:"modifiedEnd,omitempty"`
	// ClosedAfter - search for tasks closed after given date.
	ClosedAfter    *entities.Epoch `json:"closedStart,omitempty"`
	// ClosedBefore - search for tasks closed before given date.
	ClosedBefore   *entities.Epoch `json:"closedEnd,omitempty"`
	// ClosedBy - search for tasks closed by people with given PHIDs.
	ClosedBy       []string        `json:"closerPHIDs,omitempty"`
	// Query - find objects matching a fulltext search query. See "Search User Guide" in the documentation for details.
	Query          string          `json:"query,omitempty"`
	// Subscribers - search for objects with certain subscribers.
	Subscribers    []string        `json:"subscribers,omitempty"`
	// Projects - search for objects tagged with given projects.
	Projects       []string        `json:"projects,omitempty"`
	// Spaces - search for objects in certain spaces.
	Spaces         []string        `json:"spaces,omitempty"`
}

