package models

import "time"

type Tickets struct {
	Tickets []*Ticket `json:"tickets,omitempty"`
}

type SingleTicket struct {
	Ticket *Ticket `json:"ticket,omitempty"`
}

type Ticket struct {
	Subject                string        `json:"subject,omitempty"`
	GroupID                interface{}   `json:"group_id,omitempty"`
	DepartmentID           int64         `json:"department_id,omitempty" format:"int64"`
	DepartmentName         string        `json:"department_name,omitempty"`
	Category               interface{}   `json:"category,omitempty"`
	SubCategory            interface{}   `json:"sub_category,omitempty"`
	ItemCategory           interface{}   `json:"item_category,omitempty"`
	RequesterID            int64         `json:"requester_id,omitempty" format:"int64"`
	ResponderID            interface{}   `json:"responder_id,omitempty"`
	DueBy                  *time.Time    `json:"due_by,omitempty"`
	FrEscalated            bool          `json:"fr_escalated,omitempty"`
	Deleted                bool          `json:"deleted,omitempty"`
	Spam                   bool          `json:"spam,omitempty"`
	EmailConfigID          interface{}   `json:"email_config_id,omitempty"`
	FwdEmails              []interface{} `json:"fwd_emails,omitempty"`
	ReplyCcEmails          []interface{} `json:"reply_cc_emails,omitempty"`
	CcEmails               []interface{} `json:"cc_emails,omitempty"`
	IsEscalated            bool          `json:"is_escalated,omitempty"`
	FrDueBy                *time.Time    `json:"fr_due_by,omitempty"`
	ID                     int64         `json:"id,omitempty" format:"int64"`
	Priority               interface{}   `json:"priority,omitempty"`
	Status                 interface{}   `json:"status,omitempty"`
	Source                 interface{}   `json:"source,omitempty"`
	CreatedAt              *time.Time    `json:"created_at,omitempty"`
	UpdatedAt              *time.Time    `json:"updated_at,omitempty"`
	RequestedForID         int64         `json:"requested_for_id,omitempty" format:"int64"`
	ToEmails               interface{}   `json:"to_emails,omitempty"`
	Type                   string        `json:"type,omitempty"`
	Description            string        `json:"description,omitempty"`
	DescriptionText        string        `json:"description_text,omitempty"`
	CustomFields           *struct{}     `json:"custom_fields,omitempty"`
	AssociatedTicketsCount interface{}   `json:"associated_tickets_count,omitempty"`
	Attachments            []interface{} `json:"attachments,omitempty"`
	SourceAdditionalInfo   interface{}   `json:"source_additional_info,omitempty"`
	TwitterId              string        `json:"twitter_id,omitempty"`
	Tags                   []string      `json:"tags,omitempty"`
	Requester              *Requester    `json:"requester,omitempty"`
	Stats                  *Stats        `json:"stats,omitempty"`
	NrDueBy                interface{}   `json:"nr_due_by,omitempty"`
	NrEscalated            bool          `json:"nr_escalated,omitempty"`
}

type Requester struct {
	ID     int64  `json:"id,omitempty" format:"int64"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
	Mobile int64  `json:"mobile,omitempty"`
	Phone  string `json:"phone,omitempty"`
}

type Stats struct {
	AgentRespondedAt     interface{} `json:"agent_responded_at,omitempty"`
	RequesterRespondedAt interface{} `json:"requester_responded_at,omitempty"`
	FirstRespondedAt     interface{} `json:"first_responded_at,omitempty"`
	StatusUpdatedAt      *time.Time  `json:"status_updated_at,omitempty"`
	ReopenedAt           interface{} `json:"reopened_at,omitempty"`
	ResolvedAt           interface{} `json:"resolved_at,omitempty"`
	ClosedAt             interface{} `json:"closed_at,omitempty"`
	PendingSince         interface{} `json:"pending_since,omitempty"`
}

type SearchResults struct {
	Tickets []*Ticket `json:"tickets,omitempty"`
	Total   int       `json:"total,omitempty"`
}

type OpenTickets struct {
	OpenTickets []*Ticket `json:"tickets,omitempty"`
	Total       int       `json:"total,omitempty"`
}

type PendingTickets struct {
	PendingTickets []*Ticket `json:"tickets,omitempty"`
	Total          int       `json:"total,omitempty"`
}

type ResolvedTickets struct {
	ResolvedTickets []*Ticket `json:"tickets,omitempty"`
	Total           int       `json:"total,omitempty"`
}

type ClosedTickets struct {
	ClosedTickets []*Ticket `json:"tickets,omitempty"`
	Total         int       `json:"total,omitempty"`
}

type AllTickets struct {
	OpenTickets     OpenTickets     `json:"open_tickets,omitempty"`
	PendingTickets  PendingTickets  `json:"pending_tickets,omitempty"`
	ResolvedTickets ResolvedTickets `json:"resolved_tickets,omitempty"`
	ClosedTickets   ClosedTickets   `json:"closed_tickets,omitempty"`
}

type Conversations struct {
	Conversations []*Conversation `json:"conversations,omitempty"`
}

type Conversation struct {
	ID           int           `json:"id,omitempty"`
	Body         string        `json:"body,omitempty"`
	BodyText     string        `json:"body_text,omitempty"`
	Incoming     bool          `json:"incoming,omitempty"`
	Private      bool          `json:"private,omitempty"`
	CreatedAt    time.Time     `json:"created_at,omitempty"`
	UpdatedAt    time.Time     `json:"updated_at,omitempty"`
	UserID       int64         `json:"user_id,omitempty" format:"int64"`
	SupportEmail interface{}   `json:"support_email,omitempty"`
	Source       int64         `json:"source,omitempty" format:"int64"`
	TicketID     int64         `json:"ticket_id,omitempty" format:"int64"`
	ToEmails     []interface{} `json:"to_emails,omitempty"`
	FromEmail    interface{}   `json:"from_email,omitempty"`
	CcEmails     []interface{} `json:"cc_emails,omitempty"`
	BccEmails    interface{}   `json:"bcc_emails,omitempty"`
	Attachments  []interface{} `json:"attachments,omitempty"`
}

type Status int

const (
	Open Status = iota + 2
	Pending
	Resolved
	Closed
)

func (s Status) String() string {
	switch s {
	case Open:
		return "open"
	case Pending:
		return "pending"
	case Resolved:
		return "resolved"
	case Closed:
		return "closed"
	}
	return "unknown"
}

type Priority int

const (
	Low Priority = iota + 1
	Medium
	High
	Urgent
)

func (p Priority) String() string {
	switch p {
	case Low:
		return "low"
	case Medium:
		return "medium"
	case High:
		return "high"
	case Urgent:
		return "urgent"
	}
	return "unknown"
}

type Source int

const (
	Email Source = iota + 1
	Portal
	Phone
	Chat
	FeedBackWidget
	OutboundEmail
)

func (s Source) String() string {
	switch s {
	case Email:
		return "email"
	case Portal:
		return "portal"
	case Phone:
		return "phone"
	case Chat:
		return "chat"
	case FeedBackWidget:
		return "feedback widget"
	case OutboundEmail:
		return "outbound email"
	}
	return "unknown"
}
