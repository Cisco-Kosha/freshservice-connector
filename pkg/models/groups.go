package models

import "time"

type Groups struct {
	Groups []*Group `json:"groups,omitempty"`
}

type SingleGroup struct {
	Group *Group `json:"group,omitempty"`
}

type Group struct {
	ID                       int           `json:"id,omitempty"`
	Name                     string        `json:"name,omitempty"`
	Description              string        `json:"description,omitempty"`
	BusinessHoursID          interface{}   `json:"business_hours_id,omitempty"`
	EscalateTo               int           `json:"escalate_to,omitempty"`
	UnassignedFor            string        `json:"unassigned_for,omitempty"`
	AgentIds                 []int         `json:"agent_ids,omitempty"`
	Members                  []int         `json:"members,omitempty"`
	Observers                []int         `json:"observers,omitempty"`
	Restricted               bool          `json:"restricted,omitempty"`
	ApprovalRequired         bool          `json:"approval_required,omitempty"`
	Leaders                  []interface{} `json:"leaders,omitempty"`
	MembersPendingApproval   []interface{} `json:"members_pending_approval,omitempty"`
	LeadersPendingApproval   []interface{} `json:"leaders_pending_approval,omitempty"`
	ObserversPendingApproval []interface{} `json:"observers_pending_approval,omitempty"`
	AutoTicketAssign         bool          `json:"auto_ticket_assign,omitempty"`
	CreatedAt                *time.Time    `json:"created_at,omitempty"`
	UpdatedAt                *time.Time    `json:"updated_at,omitempty"`
}

type Specification struct {
	ApiKey     string `json:"api_key,omitempty"`
	DomainName string `json:"domain_name,omitempty"`
}
