package domain

import (
    "encoding/json"
    "time"
)

//-----------------------------------------------------------------------------
// Box
//-----------------------------------------------------------------------------

// Box represents a vagrant box
type Box struct {
    ID                  int64      `json:"-" db:"id"`
    Name                string     `json:"name" db:"name"`
    UserID              int64      `json:"-" db:"user_id"`
    Username            string     `json:"username" db:"username"`
    Private             bool       `json:"is_private" db:"is_private"`
    CreatedAt           int64      `json:"created_at,omitempty" db:"created_at"`
    UpdatedAt           int64      `json:"updated_at,omitempty" db:"updated_at"`
    ShortDescription    string     `json:"short_description,omitempty" db:"short_description"`
    Description         string     `json:"description,omitempty" db:"description"`
    DescriptionHTML     string     `json:"description_html,omitempty" db:"description_html"`
    DescriptionMarkdown string     `json:"description_markdown,omitempty" db:"description_markdown"`
    Tag                 string     `json:"tag,omitempty" db:"tag"`
    Downloads           int64      `json:"downloads" db:"downloads"`
    CurrentVersion      *Version   `json:"current_version,omitempty" db:"current_version"`
    Versions            []*Version `json:"versions,omitempty" db:"-"`
}

// ToParams converts the Box struct to a set of named query parameters
func (b *Box) ToParams() map[string]interface{} {
    return map[string]interface{}{
        "id":                   b.ID,
        "name":                 b.Name,
        "user_id":              b.UserID,
        "username":             b.Username,
        "is_private":           b.Private,
        "created_at":           b.CreatedAt,
        "updated_at":           b.UpdatedAt,
        "short_description":    b.ShortDescription,
        "description":          b.Description,
        "description_html":     b.DescriptionHTML,
        "description_markdown": b.DescriptionMarkdown,
        "tag":                  b.Tag,
        "downloads":            b.Downloads,
    }
}

// CreatedTimestamps sets timestamps correctly on creation
func (b *Box) CreatedTimestamps() {
    now := time.Now().Unix()
    b.CreatedAt = now
    b.UpdatedAt = now
}

// UpdatedTimestamps sets timestamps correctly on update
func (b *Box) UpdatedTimestamps() {
    b.UpdatedAt = time.Now().Unix()
}

// MarshalJSON converts outgoing timestamps to RFC-3339 format
func (b *Box) MarshalJSON() ([]byte, error) {
    type Alias Box
    return json.Marshal(struct {
        CreatedAt string `json:"created_at,omitempty"`
        UpdatedAt string `json:"updated_at,omitempty"`
        *Alias
    }{
        CreatedAt: ToInternetTime(b.CreatedAt),
        UpdatedAt: ToInternetTime(b.UpdatedAt),
        Alias:     (*Alias)(b),
    })
}

// UnmarshalJSON converts incoming timestamps to UNIX format
func (b *Box) UnmarshalJSON(data []byte) error {
    type Alias Box
    tmp := &struct {
        CreatedAt string `json:"created_at,omitempty"`
        UpdatedAt string `json:"updated_at,omitempty"`
        *Alias
    }{
        Alias: (*Alias)(b),
    }

    if err := json.Unmarshal(data, &tmp); err != nil {
        return err
    }
    b.CreatedAt = ToUnixTime(tmp.CreatedAt)
    b.UpdatedAt = ToUnixTime(tmp.UpdatedAt)
    return nil
}
