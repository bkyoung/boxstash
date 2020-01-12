package domain

import (
    "encoding/json"
    "time"
)

//-----------------------------------------------------------------------------
// Provider
//-----------------------------------------------------------------------------

// Provider represents a single provider for a version (e.g. virtualbox)
type Provider struct {
    ID          int64  `json:"-" db:"id"`
    Name        string `json:"name,omitempty" db:"name"`
    Hosted      bool   `json:"hosted,omitempty" db:"hosted"`
    HostedToken string `json:"hosted_token,omitempty" db:"hosted_token"`
    OriginalURL string `json:"original_url,omitempty" db:"original_url"`
    CreatedAt   int64  `json:"created_at,omitempty" db:"created_at"`
    UpdatedAt   int64  `json:"updated_at,omitempty" db:"updated_at"`
    DownloadURL string `json:"download_url,omitempty" db:"download_url"`
    VersionID   int64  `json:"-" db:"version_id"`
}

// ToParams converts the Provider struct to a set of named query parameters
func (p *Provider) ToParams() map[string]interface{} {
    return map[string]interface{}{
        "id":           p.ID,
        "name":         p.Name,
        "hosted":       p.Hosted,
        "hosted_token": p.HostedToken,
        "original_url": p.OriginalURL,
        "created_at":   p.CreatedAt,
        "updated_at":   p.UpdatedAt,
        "download_url": p.DownloadURL,
        "version_id":   p.VersionID,
    }
}

// CreatedTimestamps sets timestamps correctly on creation
func (p *Provider) CreatedTimestamps() {
    now := time.Now().Unix()
    p.CreatedAt = now
    p.UpdatedAt = now
}

// UpdatedTimestamps sets timestamps correctly on update
func (p *Provider) UpdatedTimestamps() {
    p.UpdatedAt = time.Now().Unix()
}

// MarshalJSON converts outgoing timestamps to RFC-3339 format
func (p *Provider) MarshalJSON() ([]byte, error) {
    type Alias Provider
    return json.Marshal(struct {
        CreatedAt string `json:"created_at,omitempty"`
        UpdatedAt string `json:"updated_at,omitempty"`
        *Alias
    }{
        CreatedAt: ToInternetTime(p.CreatedAt),
        UpdatedAt: ToInternetTime(p.UpdatedAt),
        Alias:     (*Alias)(p),
    })
}

// UnmarshalJSON converts incoming timestamps to UNIX format
func (p *Provider) UnmarshalJSON(data []byte) error {
    type Alias Provider
    tmp := &struct {
        CreatedAt string `json:"created_at,omitempty"`
        UpdatedAt string `json:"updated_at,omitempty"`
        *Alias
    }{
        Alias: (*Alias)(p),
    }

    if err := json.Unmarshal(data, &tmp); err != nil {
        return err
    }
    p.CreatedAt = ToUnixTime(tmp.CreatedAt)
    p.UpdatedAt = ToUnixTime(tmp.UpdatedAt)
    return nil
}
