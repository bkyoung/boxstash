package domain

import (
    "encoding/json"
    "time"
)

//-----------------------------------------------------------------------------
// Version
//-----------------------------------------------------------------------------

// Version represents a single version of a vagrant box
type Version struct {
    ID                  int64       `json:"-" db:"id"`
    Version             string      `json:"version,omitempty" db:"version"`
    Status              string      `json:"status,omitempty" db:"status"`
    Description         string      `json:"description,omitempty" db:"description"`
    DescriptionHTML     string      `json:"description_html,omitempty" db:"description_html"`
    DescriptionMarkdown string      `json:"description_markdown,omitempty" db:"description_markdown"`
    CreatedAt           int64       `json:"created_at,omitempty" db:"created_at"`
    UpdatedAt           int64       `json:"updated_at,omitempty" db:"updated_at"`
    Number              string      `json:"number,omitempty" db:"number"`
    ReleaseURL          string      `json:"release_url,omitempty" db:"release_url"`
    RevokeURL           string      `json:"revoke_url,omitempty" db:"revoke_url"`
    Providers           []*Provider `json:"providers,omitempty" db:"-"`
    BoxID               int64       `json:"-" db:"box_id"`
}

// ToParams converts the Version struct to a set of named query parameters
func (v *Version) ToParams() map[string]interface{} {
    return map[string]interface{}{
        "id":                   v.ID,
        "version":              v.Version,
        "status":               v.Status,
        "created_at":           v.CreatedAt,
        "updated_at":           v.UpdatedAt,
        "description":          v.Description,
        "description_html":     v.DescriptionHTML,
        "description_markdown": v.DescriptionMarkdown,
        "number":               v.Number,
        "release_url":          v.ReleaseURL,
        "revoke_url":           v.RevokeURL,
        "box_id":               v.BoxID,
    }
}

// CreatedTimestamps sets timestamps correctly on creation
func (v *Version) CreatedTimestamps() {
    now := time.Now().Unix()
    v.CreatedAt = now
    v.UpdatedAt = now
}

// UpdatedTimestamps sets timestamps correctly on update
func (v *Version) UpdatedTimestamps() {
    v.UpdatedAt = time.Now().Unix()
}

// MarshalJSON converts outgoing timestamps to RFC-3339 format
func (v *Version) MarshalJSON() ([]byte, error) {
    type Alias Version
    return json.Marshal(struct {
        CreatedAt string `json:"created_at,omitempty"`
        UpdatedAt string `json:"updated_at,omitempty"`
        *Alias
    }{
        CreatedAt: ToInternetTime(v.CreatedAt),
        UpdatedAt: ToInternetTime(v.UpdatedAt),
        Alias:     (*Alias)(v),
    })
}

// UnmarshalJSON converts incoming timestamps to UNIX format
func (v *Version) UnmarshalJSON(data []byte) error {
    type Alias Version
    tmp := &struct {
        CreatedAt string `json:"created_at,omitempty"`
        UpdatedAt string `json:"updated_at,omitempty"`
        *Alias
    }{
        Alias: (*Alias)(v),
    }

    if err := json.Unmarshal(data, &tmp); err != nil {
        return err
    }
    v.CreatedAt = ToUnixTime(tmp.CreatedAt)
    v.UpdatedAt = ToUnixTime(tmp.UpdatedAt)
    return nil
}
