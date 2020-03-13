package domain

//-----------------------------------------------------------------------------
// User
//-----------------------------------------------------------------------------

// User represents a user or organization
type User struct {
	ID              int64  `json:"-" gorm:"primary_key;autoincrement"`
	Username        string `json:"username" gorm:"unique;not null"`
	AvatarURL       string `json:"avatar_url,omitempty"`
	ProfileHTML     string `json:"profile_html,omitempty"`
	ProfileMarkdown string `json:"profile_markdown,omitempty"`
	Boxes           []*Box `json:"boxes,omitempty"`
}
