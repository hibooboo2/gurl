package main

import "time"

// Channel ..
type Channel struct {
	ID          int           `json:"id,string" db:"id"`
	AccountID   int           `json:"-" db:"account_id"`
	FounderID   int           `json:"founder_id,string" db:"founder_id"`
	Name        string        `json:"name" db:"name"`
	Slug        string        `json:"slug"`
	Description string        `json:"description" db:"description"`
	Created     time.Time     `json:"created"`
	Updated     *time.Time    `json:"updated"`
	Links       *ChannelLinks `json:"links"`
	Deleted     *time.Time    `json:"-" db:"deleted"`
}

// ChannelLinks ..
type ChannelLinks struct {
	SelfURL     string `json:"self_url"`
	ActivityURL string `json:"activity_url"`
	DocsURL     string `json:"docs_url"`
	EventsURL   string `json:"events_url"`
	FilesURL    string `json:"files_url"`
	InvitesURL  string `json:"invites_url"`
	SettingsURL string `json:"settings_url"`
	StatusURL   string `json:"status_url"`
	TasksURL    string `json:"tasks_url"`
	UsersURL    string `json:"users_url"`
}
