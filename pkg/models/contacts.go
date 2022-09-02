package models

import "time"

type Contact struct {
	Active       bool   `json:"active,omitempty"`
	Address      string `json:"address,omitempty"`
	Description  string `json:"description,omitempty"`
	Email        string `json:"email,omitempty"`
	ID           int64  `json:"id,omitempty"`
	JobTitle     string `json:"job_title,omitempty"`
	Language     string `json:"language,omitempty"`
	Mobile       int64  `json:"mobile,omitempty"`
	Name         string `json:"name,omitempty"`
	Phone        string `json:"phone,omitempty"`
	TimeZone     string `json:"time_zone,omitempty"`
	TwitterID    string `json:"twitter_id,omitempty"`
	CustomFields struct {
	} `json:"custom_fields,omitempty"`
	FacebookID       string      `json:"facebook_id,omitempty"`
	CreatedAt        *time.Time  `json:"created_at,omitempty"`
	UpdatedAt        *time.Time  `json:"updated_at,omitempty"`
	CsatRating       interface{} `json:"csat_rating,omitempty"`
	PreferredSource  interface{} `json:"preferred_source,omitempty"`
	CompanyID        interface{} `json:"company_id,omitempty"`
	UniqueExternalID interface{} `json:"unique_external_id,omitempty"`
	FirstName        string      `json:"first_name,omitempty"`
	LastName         string      `json:"last_name,omitempty"`
	VisitorID        string      `json:"visitor_id,omitempty"`
	OrgContactID     int64       `json:"org_contact_id,omitempty"`
	OtherEmails      []string    `json:"other_emails,omitempty"`
}
