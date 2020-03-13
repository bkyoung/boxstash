package domain

import (
    "time"
)

//-----------------------------------------------------------------------------
// Box
//-----------------------------------------------------------------------------

// Box represents a vagrant box
type Box struct {
    ID                  int64      `json:"-" gorm:"primary_key"`
    Name                string     `json:"name" gorm:"not null,unique_index:idx_box_name"`
    UserID              int64      `json:"-" gorm:"not null"`
    Username            string     `json:"username" gorm:"not null,unique_index:idx_box_name"`
    Private             bool       `json:"is_private"`
    CreatedAt           time.Time  `json:"created_at,omitempty"`
    UpdatedAt           time.Time  `json:"updated_at,omitempty"`
    ShortDescription    string     `json:"short_description,omitempty"`
    Description         string     `json:"description,omitempty"`
    DescriptionHTML     string     `json:"description_html,omitempty"`
    DescriptionMarkdown string     `json:"description_markdown,omitempty"`
    Tag                 string     `json:"tag,omitempty"`
    Downloads           int64      `json:"downloads"`
    CurrentVersion      *Version   `json:"current_version,omitempty"`
    Versions            []*Version `json:"versions,omitempty"`
}
