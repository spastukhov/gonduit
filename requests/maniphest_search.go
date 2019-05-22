package requests

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/uber/gonduit/entities"
)

// ManiphestSearchRequest represents a request to maniphest.search.
type ManiphestSearchRequest struct {
	QueryKey    string             `json:"queryKey,omitempty"`
	Constraints *SearchConstraints `json:"constraints,omitempty"`
	Attachments *SearchAttachments `json:"attachments,omitempty"`
	*entities.Cursor
	Request
}

type SearchAttachments struct {
	Subscribers bool `json:"subscribers,omitempty"`
	Columns     bool `json:"columns,omitempty"`
	Projects    bool `json:"projects,omitempty"`
}

type ManiphestRequestSearchOrder struct {
	Builtin string
	Order   []string
}

func (o *ManiphestRequestSearchOrder) UnmarshalJSON(data []byte) error {
	if o == nil {
		return errors.New("maniphest search order is nil")
	}
	if jerr := json.Unmarshal(data, &o.Builtin); jerr == nil {
		return nil
	}

	return json.Unmarshal(data, &o.Order)
}

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

type SearchConstraints struct {
	IDs            []string `json:"ids,omitempty"`
	PHIDs          []string `json:"phids,omitempty"`
	AssignedTo     []string `json:"assigned,omitempty"`
	Authors        []string `json:"authorPHIDs,omitempty"`
	Statuses       []string `json:"statuses,omitempty"`
	Priorities     []string `json:"priorities,omitempty"`
	Subtypes       []string `json:"subtypes,omitempty"`
	ColumnPHIDs    []string `json:"columnPHIDs,omitempty"`
	OpenParents    *bool     `json:"hasParents,omitempty"`
	OpenSubtasks   *bool     `json:"hasSubtasks,omitempty"`
	ParentIDs      []int `json:"parentIDs,omitempty"`
	SubtaskIDs     []int `json:"subtaskIDs,omitempty"`
	CreatedAfter   *entities.Epoch    `json:"createdStart,omitempty"`
	CreatedBefore  *entities.Epoch    `json:"createdEnd,omitempty"`
	ModifiedAfter  *entities.Epoch    `json:"modifiedStart,omitempty"`
	ModifiedBefore *entities.Epoch    `json:"modifiedEnd,omitempty"`
	ClosedAfter    *entities.Epoch    `json:"closedStart,omitempty"`
	ClosedBefore   *entities.Epoch    `json:"closedEnd,omitempty"`
	ClosedBy       []string `json:"closerPHIDs,omitempty"`
	Query          string   `json:"query,omitempty"`
	Subscribers    []string `json:"subscribers,omitempty"`
	Projects       []string `json:"projects,omitempty"`
	Spaces         []string `json:"spaces,omitempty"`
}

type ManiphestSearchResponse struct {
	Data   []*SearchResponseItem `json:"data"`
	Cursor entities.Cursor          `json:"cursor,omitempty"`
}

type ManiphestSearchAttachmentColumnBoardsColumn struct {
	ID   int    `json:"id"`
	Phid string `json:"phid"`
	Name string `json:"name"`
}

type ManiphestSearchAttachmentColumnBoardsColumns struct {
	Columns []*ManiphestSearchAttachmentColumnBoardsColumn `json:"columns,omitempty"`
}

type ManiphestSearchAttachmentColumnBoards struct {
	ColumnMap map[string]*ManiphestSearchAttachmentColumnBoardsColumns
	Columns []*ManiphestSearchAttachmentColumnBoardsColumn
}

func (b *ManiphestSearchAttachmentColumnBoards) UnmarshalJSON(data []byte) error {
	if b == nil {
		return errors.New("boards is nil")
	}
	//if b.ColumnMap == nil {
	//	b.ColumnMap = make(map[string]*ManiphestSearchAttachmentColumnBoardsColumns)
	//}
	jerr := json.Unmarshal(data, &b.ColumnMap)
	if jerr == nil {
		return nil
	}
	//fmt.Println(jerr.Error())
	return json.Unmarshal(data, &b.Columns)
}

type TaskDescription struct {
	Raw string `json:"raw"`
}

type SearchResponseItem struct {
	ID     int    `json:"id"`
	Type   string `json:"type"`
	Phid   string `json:"phid"`
	Fields struct {
		Name           string               `json:"name"`
		Description    *TaskDescription `json:"description"`
		AuthorPHID     string               `json:"authorPHID"`
		OwnerPHID      string               `json:"ownerPHID"`
		Status         SearchResultStatus   `json:"status"`
		Priority       SearchResultPriority `json:"priority"`
		Points         interface{}          `json:"points"`
		Subtype        string               `json:"subtype"`
		SpacePHID      string               `json:"spacePHID"`
		DateCreated    int                  `json:"dateCreated"`
		DateModified   int                  `json:"dateModified"`
		Policy         SearchResultPolicy   `json:"policy"`
		CustomTaskType string               `json:"custom.task_type"`
		CustomSeverity string               `json:"custom.severity"`
	} `json:"fields"`
	Attachments struct {
		Columns struct {
			Boards *ManiphestSearchAttachmentColumnBoards `json:"boards"`
		} `json:"columns"`
		Subscribers struct {
			SubscriberPHIDs    []string `json:"subscriberPHIDs"`
			SubscriberCount    int      `json:"subscriberCount"`
			ViewerIsSubscribed bool     `json:"viewerIsSubscribed"`
		} `json:"subscribers"`
		Projects struct {
			ProjectPHIDs []string `json:"projectPHIDs"`
		} `json:"projects"`
	} `json:"attachments"`
}

// SearchResultStatus represents a maniphest status as returned by maniphest.search
type SearchResultStatus struct {
	Value string `json:"value"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

// SearchResultPriority represents a priority for a maniphest item in a search result
type SearchResultPriority struct {
	Value       int     `json:"value"`
	Subpriority float64 `json:"subpriority"`
	Name        string  `json:"name"`
	Color       string  `json:"color"`
}

// SearchResultPolicy reflects the permission policy on a maniphest item in a search result
type SearchResultPolicy struct {
	View     string `json:"view"`
	Interact string `json:"interact"`
	Edit     string `json:"edit"`
}

// SearchResultColumn represents what workboard columns an item may be a member of
type SearchResultColumn struct {
	ID          int
	PHID        string
	Name        string
	ProjectPHID string
}
