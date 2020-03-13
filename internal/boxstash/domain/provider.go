package domain

import "time"

//-----------------------------------------------------------------------------
// Provider
//-----------------------------------------------------------------------------

// Provider represents a single provider for a version (e.g. virtualbox)
type Provider struct {
    ID          int64      `json:"-" gorm:"primary_key"`
    Name        string     `json:"name,omitempty" gorm:"not null,unique_index:idx_provider"`
    Hosted      bool       `json:"hosted,omitempty"`
    HostedToken string     `json:"hosted_token,omitempty"`
    OriginalURL string     `json:"original_url,omitempty"`
    CreatedAt   time.Time  `json:"created_at,omitempty"`
    UpdatedAt   time.Time  `json:"updated_at,omitempty"`
    DownloadURL string     `json:"download_url,omitempty"`
    VersionID   int64      `json:"-" gorm:"not null,unique_index:idx_provider"`
}
