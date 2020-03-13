package domain

import "time"

//-----------------------------------------------------------------------------
// Version
//-----------------------------------------------------------------------------

// Version represents a single version of a vagrant box
type Version struct {
    ID                  int64       `json:"-" gorm:"primary_key"`
    Version             string      `json:"version,omitempty" gorm:"not null,unique_index:idx_version"`
    Status              string      `json:"status,omitempty"`
    Description         string      `json:"description,omitempty"`
    DescriptionHTML     string      `json:"description_html,omitempty"`
    DescriptionMarkdown string      `json:"description_markdown,omitempty"`
    CreatedAt           time.Time   `json:"created_at,omitempty"`
    UpdatedAt           time.Time   `json:"updated_at,omitempty"`
    Number              string      `json:"number,omitempty"`
    ReleaseURL          string      `json:"release_url,omitempty"`
    RevokeURL           string      `json:"revoke_url,omitempty"`
    Providers           []*Provider `json:"providers,omitempty"`
    BoxID               int64       `json:"-" gorm:"not null,unique_index:idx_version"`
}
