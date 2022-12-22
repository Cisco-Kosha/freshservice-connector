package models

import "time"

type Departments struct {
	Departments []*Department `json:"departments,omitempty"`
}

type Department struct {
	Description  string `json:"description,omitempty"`
	CustomFields struct {
	} `json:"custom_fields,omitempty"`
	ID          int64         `json:"id,omitempty"`
	Name        string        `json:"name,omitempty"`
	CreatedAt   time.Time     `json:"created_at,omitempty"`
	UpdatedAt   time.Time     `json:"updated_at,omitempty"`
	PrimeUserID interface{}   `json:"prime_user_id,omitempty"`
	HeadUserID  interface{}   `json:"head_user_id,omitempty"`
	Domains     []interface{} `json:"domains,omitempty"`
}

type SingleDepartment struct {
	Department Department `json:"department,omitempty"`
}
