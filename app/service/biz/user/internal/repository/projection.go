package repository

type UserAggregateCacheDTO struct {
	SchemaVersion int     `json:"schema_version"`
	UserID        int64   `json:"user_id"`
	AccessHash    int64   `json:"access_hash"`
	Phone         string  `json:"phone,omitempty"`
	FirstName     string  `json:"first_name,omitempty"`
	LastName      string  `json:"last_name,omitempty"`
	Username      string  `json:"username,omitempty"`
	PhotoID       int64   `json:"photo_id,omitempty"`
	ContactIDs    []int64 `json:"contact_ids,omitempty"`
}
