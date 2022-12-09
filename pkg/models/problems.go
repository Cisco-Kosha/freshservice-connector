package models

import "time"

type Problems struct {
	Problems []*Problem `json:"problems,omitempty"`
}

type SingleProblem struct {
	Problem *Problem `json:"problem,omitempty"`
}

type Problem struct {
	ID              int64         `json:"id,omitempty" format:"int64"`
	AgentID         interface{}   `json:"agent_id,omitempty"`
	RequesterID     int64         `json:"requester_id,omitempty" format:"int64"`
	Description     string        `json:"description,omitempty"`
	DescriptionText string        `json:"description_text,omitempty"`
	DueBy           *time.Time    `json:"due_by,omitempty"`
	Subject         string        `json:"subject,omitempty"`
	GroupID         interface{}   `json:"group_id,omitempty"`
	Priority        interface{}   `json:"priority,omitempty"`
	Impact          int           `json:"impact,omitempty"`
	Status          interface{}   `json:"status,omitempty"`
	KnownError      bool          `json:"known_error,omitempty"`
	DepartmentID    interface{}   `json:"department_id,omitempty"`
	Category        string        `json:"category,omitempty"`
	SubCategory     string        `json:"sub_category,omitempty"`
	ItemCategory    string        `json:"item_category,omitempty"`
	Assets          []interface{} `json:"assets,omitempty"`
	CreatedAt       *time.Time    `json:"created_at,omitempty"`
	UpdatedAt       time.Time     `json:"updated_at,omitempty"`
}
