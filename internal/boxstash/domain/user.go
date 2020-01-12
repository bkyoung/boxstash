package domain

//-----------------------------------------------------------------------------
// User
//-----------------------------------------------------------------------------

// User represents a user or organization
type User struct {
	ID              int64  `json:"-" db:"id"`
	Username        string `json:"username" db:"username"`
	AvatarURL       string `json:"avatar_url,omitempty" db:"avatar_url"`
	ProfileHTML     string `json:"profile_html,omitempty" db:"profile_html"`
	ProfileMarkdown string `json:"profile_markdown,omitempty" db:"profile_markdown"`
	Boxes           []*Box `json:"boxes,omitempty" db:"-"`
}

// ToParams converts the Box struct to a set of named query parameters
func (u *User) ToParams() map[string]interface{} {
    return map[string]interface{}{
        "id":                   u.ID,
        "username":             u.Username,
        "avatar_url":           u.AvatarURL,
        "profile_html":         u.ProfileHTML,
        "profile_markdown":     u.ProfileMarkdown,
    }
}

