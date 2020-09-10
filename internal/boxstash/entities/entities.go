package entities

import "time"

type About struct {
	Build   string    `json:"build_number"`
	Commit  string    `json:"commit"`
	Date    time.Time `json:"date"`
	Version string    `json:"version"`
}

type Box struct {
	ID 					int64 	  `json:"-" db:"id"`
	Name                string    `json:"name" db:"name"`
	Username            string    `json:"username" db:"username"`
	Private             *bool      `json:"is_private" db:"is_private"`
	CreatedAt           time.Time `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time `json:"updated_at" db:"updated_at"`
	ShortDescription    *string   `json:"short_description,omitempty" db:"short_description"`
	Description    		*string   `json:"description,omitempty" db:"description"`
	DescriptionHTML     *string   `json:"description_html,omitempty" db:"description_html"`
	DescriptionMarkdown *string   `json:"description_markdown,omitempty" db:"description_markdown"`
	Tag                 *string   `json:"tag,omitempty" db:"tag"`
	Downloads           int64     `json:"downloads" db:"downloads"`
	CurrentVersion      *Version  `json:"current_version,omitempty" db:"current_version,omitempty"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Users struct {
	ID 				int64  `json:"-" db:"id"`
	Username        string `json:"username" db:"username"`
	AvatarURL       string `json:"avatar_url" db:"avatar_url"`
	ProfileHTML     string `json:"profile_html" db:"profile_html"`
	ProfileMarkdown string `json:"profile_markdown" db:"profile_markdown"`
	Boxes           []Box  `json:"boxes" db:"boxes"`
}

type Provider struct {
	ID 			 int64     `json:"-,omitempty" db:"id"`
	Name         string    `json:"name" db:"name"`
	Hosted       bool      `json:"hosted" db:"hosted"`
	HostedToken  string    `json:"hosted_token" db:"hosted_token"`
	OriginalURL  string    `json:"original_url" db:"original_url"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	DownloadURL  string    `json:"download_url" db:"download_url"`
	Checksum     string    `json:"checksum" db:"checksum"`
	ChecksumType string    `json:"checksum_type" db:"checksum_type"`
}

type Version struct {
	ID 				   int64 	  `json:"-,omitempty" db:"id"`
	Version            string     `json:"version" db:"version"`
	Status             string     `json:"status" db:"status"`
	DescriptionHTML    string     `json:"description_html" db:"description_html"`
	DesciptionMarkdown string     `json:"description_markdown" db:"description_markdown"`
	CreatedAt          time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at" db:"updated_at"`
	Number             string     `json:"number" db:"number"`
	ReleaseURL         string     `json:"release_url" db:"release_url"`
	RevokeURL          string     `json:"revoke_url" db:"revoke_url"`
	Providers          []Provider `json:"providers" db:"providers"`
}
