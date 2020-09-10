package service

type NewBoxRequest struct {
	Username         string `json:"username"`
	Name             string `json:"name"`
	ShortDescription *string `json:"short_description"`
	Description      *string `json:"description"`
	Private          *bool   `json:"is_private"`
}
