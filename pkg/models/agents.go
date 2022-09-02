package models

import "time"

type Agents struct {
	Agents []*Agent `json:"agents,omitempty"`
}

type SingleAgent struct {
	Agent *Agent `json:"agent,omitempty"`
}

type Agent struct {
	Active                                    bool        `json:"active,omitempty"`
	Address                                   interface{} `json:"address,omitempty"`
	AutoAssignStatusChangedAt                 interface{} `json:"auto_assign_status_changed_at,omitempty"`
	AutoAssignTickets                         bool        `json:"auto_assign_tickets,omitempty"`
	BackgroundInformation                     interface{} `json:"background_information,omitempty"`
	CanSeeAllTicketsFromAssociatedDepartments bool        `json:"can_see_all_tickets_from_associated_departments,omitempty"`
	CreatedAt                                 time.Time   `json:"created_at,omitempty"`
	CustomFields                              struct {
	} `json:"custom_fields,omitempty"`
	DepartmentIds      []interface{} `json:"department_ids,omitempty"`
	Email              string        `json:"email,omitempty"`
	ExternalID         interface{}   `json:"external_id,omitempty"`
	FirstName          string        `json:"first_name,omitempty"`
	HasLoggedIn        bool          `json:"has_logged_in,omitempty"`
	ID                 int64         `json:"id,omitempty"`
	JobTitle           interface{}   `json:"job_title,omitempty"`
	Language           string        `json:"language,omitempty"`
	LastActiveAt       *time.Time    `json:"last_active_at,omitempty"`
	LastLoginAt        *time.Time    `json:"last_login_at,omitempty"`
	LastName           string        `json:"last_name,omitempty"`
	LocationID         interface{}   `json:"location_id,omitempty"`
	MobilePhoneNumber  interface{}   `json:"mobile_phone_number,omitempty"`
	Occasional         bool          `json:"occasional,omitempty"`
	ReportingManagerID interface{}   `json:"reporting_manager_id,omitempty"`
	RoleIds            []int64       `json:"role_ids,omitempty"`
	Roles              []struct {
		RoleID          int64         `json:"role_id,omitempty"`
		AssignmentScope string        `json:"assignment_scope,omitempty"`
		Groups          []interface{} `json:"groups,omitempty"`
	} `json:"roles,omitempty"`
	Scopes struct {
		Ticket   interface{} `json:"ticket,omitempty"`
		Problem  interface{} `json:"problem,omitempty"`
		Change   interface{} `json:"change,omitempty"`
		Release  interface{} `json:"release,omitempty"`
		Asset    interface{} `json:"asset,omitempty"`
		Solution interface{} `json:"solution,omitempty"`
		Contract interface{} `json:"contract,omitempty"`
	} `json:"scopes,omitempty"`
	ScoreboardLevelID         interface{}   `json:"scoreboard_level_id,omitempty"`
	ScoreboardPoints          interface{}   `json:"scoreboard_points,omitempty"`
	Signature                 interface{}   `json:"signature,omitempty"`
	TimeFormat                string        `json:"time_format,omitempty"`
	TimeZone                  string        `json:"time_zone,omitempty"`
	UpdatedAt                 *time.Time    `json:"updated_at,omitempty"`
	VipUser                   bool          `json:"vip_user,omitempty"`
	WorkPhoneNumber           string        `json:"work_phone_number,omitempty"`
	GroupIds                  []interface{} `json:"group_ids,omitempty"`
	MemberOf                  []interface{} `json:"member_of,omitempty"`
	ObserverOf                []interface{} `json:"observer_of,omitempty"`
	MemberOfPendingApproval   []interface{} `json:"member_of_pending_approval,omitempty"`
	ObserverOfPendingApproval []interface{} `json:"observer_of_pending_approval,omitempty"`
}
