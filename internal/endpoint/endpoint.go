package endpoint

import (
	entities "boxstash/internal/boxstash/entities"
	service "boxstash/internal/service"
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
)

// AboutRequest collects the request parameters for the About method.
type AboutRequest struct{}

// AboutResponse collects the response parameters for the About method.
type AboutResponse struct {
	About entities.About `json:"about"`
	Error error          `json:"error"`
}

// MakeAboutEndpoint returns an endpoint that invokes About on the service.
func MakeAboutEndpoint(s service.BoxstashService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		e0, e1 := s.About(ctx)
		return AboutResponse{
			About: e0,
			Error: e1,
		}, nil
	}
}

// Failed implements Failer.
func (r AboutResponse) Failed() error {
	return r.Error
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// About implements Service. Primarily useful in a client.
func (e Endpoints) About(ctx context.Context) (e0 entities.About, e1 error) {
	request := AboutRequest{}
	response, err := e.AboutEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AboutResponse).About, response.(AboutResponse).Error
}

// NewAuthTokenRequest collects the request parameters for the NewAuthToken method.
type NewAuthTokenRequest struct {
	Creds entities.Credentials `json:"credentialss"`
}

// NewAuthTokenResponse collects the response parameters for the NewAuthToken method.
type NewAuthTokenResponse struct {
	Token   string  `json:"token"`
	Success bool    `json:"success"`
	Errors  []error `json:"errors"`
}

// MakeNewAuthTokenEndpoint returns an endpoint that invokes NewAuthToken on the service.
func MakeNewAuthTokenEndpoint(s service.BoxstashService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(NewAuthTokenRequest)
		token, success, errors := s.NewAuthToken(ctx, req.Creds)
		return NewAuthTokenResponse{
			Errors:  errors,
			Success: success,
			Token:   token,
		}, nil
	}
}

// ValidateAuthTokenRequest collects the request parameters for the ValidateAuthToken method.
type ValidateAuthTokenRequest struct {
	Token string `json:"token"`
}

// ValidateAuthTokenResponse collects the response parameters for the ValidateAuthToken method.
type ValidateAuthTokenResponse struct {
	Success bool    `json:"success"`
	Errors  []error `json:"errors"`
}

// MakeValidateAuthTokenEndpoint returns an endpoint that invokes ValidateAuthToken on the service.
func MakeValidateAuthTokenEndpoint(s service.BoxstashService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ValidateAuthTokenRequest)
		success, errors := s.ValidateAuthToken(ctx, req.Token)
		return ValidateAuthTokenResponse{
			Errors:  errors,
			Success: success,
		}, nil
	}
}

// DeleteAuthTokenRequest collects the request parameters for the DeleteAuthToken method.
type DeleteAuthTokenRequest struct {
	Token string `json:"token"`
}

// DeleteAuthTokenResponse collects the response parameters for the DeleteAuthToken method.
type DeleteAuthTokenResponse struct {
	Success bool    `json:"success"`
	Errors  []error `json:"errors"`
}

// MakeDeleteAuthTokenEndpoint returns an endpoint that invokes DeleteAuthToken on the service.
func MakeDeleteAuthTokenEndpoint(s service.BoxstashService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteAuthTokenRequest)
		success, errors := s.DeleteAuthToken(ctx, req.Token)
		return DeleteAuthTokenResponse{
			Errors:  errors,
			Success: success,
		}, nil
	}
}

// ReadOrganizationRequest collects the request parameters for the ReadOrganization method.
type ReadOrganizationRequest struct {
	Org string `json:"org"`
}

// ReadOrganizationResponse collects the response parameters for the ReadOrganization method.
type ReadOrganizationResponse struct {
	Organization entities.Users `json:"organization"`
	Success      bool           `json:"success"`
	Errors       []error        `json:"errors"`
}

// MakeReadOrganizationEndpoint returns an endpoint that invokes ReadOrganization on the service.
func MakeReadOrganizationEndpoint(s service.BoxstashService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReadOrganizationRequest)
		organization, success, errors := s.ReadOrganization(ctx, req.Org)
		return ReadOrganizationResponse{
			Errors:       errors,
			Organization: organization,
			Success:      success,
		}, nil
	}
}

// SearchRequest collects the request parameters for the Search method.
type SearchRequest struct{}

// SearchResponse collects the response parameters for the Search method.
type SearchResponse struct {
	Boxes   []entities.Box `json:"boxes"`
	Success bool           `json:"success"`
	Errors  []error        `json:"errors"`
}

// MakeSearchEndpoint returns an endpoint that invokes Search on the service.
func MakeSearchEndpoint(s service.BoxstashService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		boxes, success, errors := s.Search(ctx)
		return SearchResponse{
			Boxes:   boxes,
			Success: success,
			Errors:  errors,
		}, nil
	}
}

// ReadBoxRequest collects the request parameters for the ReadBox method.
type ReadBoxRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

// ReadBoxResponse collects the response parameters for the ReadBox method.
type ReadBoxResponse struct {
	Box     entities.Box `json:"box"`
	Success bool         `json:"success"`
	Errors  []error      `json:"errors"`
}

// MakeReadBoxEndpoint returns an endpoint that invokes ReadBox on the service.
func MakeReadBoxEndpoint(s service.BoxstashService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReadBoxRequest)
		box, success, errors := s.ReadBox(ctx, req.Username, req.Name)
		return ReadBoxResponse{
			Box:     box,
			Success: success,
			Errors:  errors,
		}, nil
	}
}

// CreateBoxRequest collects the request parameters for the CreateBox method.
type CreateBoxRequest struct {
	Box service.NewBoxRequest `json:"box"`
}

// CreateBoxResponse collects the response parameters for the CreateBox method.
type CreateBoxResponse struct {
	Box     entities.Box `json:"box"`
	Success bool         `json:"success"`
	Errors  []error      `json:"errors"`
}

// MakeCreateBoxEndpoint returns an endpoint that invokes CreateBox on the service.
func MakeCreateBoxEndpoint(s service.BoxstashService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateBoxRequest)
		box, success, errors := s.CreateBox(ctx, req.Box)
		return CreateBoxResponse{
			Box:     box,
			Success: success,
			Errors:  errors,
		}, nil
	}
}

// UpdateBoxRequest collects the request parameters for the UpdateBox method.
type UpdateBoxRequest struct {
	Box map[string]interface{} `json:"box"`
}

// UpdateBoxResponse collects the response parameters for the UpdateBox method.
type UpdateBoxResponse struct {
	Box     entities.Box `json:"box"`
	Success bool         `json:"success"`
	Errors  []error      `json:"errors"`
}

// MakeUpdateBoxEndpoint returns an endpoint that invokes UpdateBox on the service.
func MakeUpdateBoxEndpoint(s service.BoxstashService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateBoxRequest)
		box, success, errors := s.UpdateBox(ctx, req.Box)
		return UpdateBoxResponse{
			Box:     box,
			Success: success,
			Errors:  errors,
		}, nil
	}
}

// DeleteBoxRequest collects the request parameters for the DeleteBox method.
type DeleteBoxRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

// DeleteBoxResponse collects the response parameters for the DeleteBox method.
type DeleteBoxResponse struct {
	Box     entities.Box `json:"box"`
	Success bool         `json:"success"`
	Errors  []error      `json:"errors"`
}

// MakeDeleteBoxEndpoint returns an endpoint that invokes DeleteBox on the service.
func MakeDeleteBoxEndpoint(s service.BoxstashService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteBoxRequest)
		box, success, errors := s.DeleteBox(ctx, req.Username, req.Name)
		return DeleteBoxResponse{
			Box:     box,
			Success: success,
			Errors:  errors,
		}, nil
	}
}

// ReadVersionRequest collects the request parameters for the ReadVersion method.
type ReadVersionRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Version  string `json:"version"`
}

// ReadVersionResponse collects the response parameters for the ReadVersion method.
type ReadVersionResponse struct {
	Version entities.Version `json:"version"`
	Success bool             `json:"success"`
	Errors  []error          `json:"errors"`
}

// MakeReadVersionEndpoint returns an endpoint that invokes ReadVersion on the service.
func MakeReadVersionEndpoint(s service.BoxstashService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReadVersionRequest)
		e0, b1, e2 := s.ReadVersion(ctx, req.Username, req.Name, req.Version)
		return ReadVersionResponse{
			Success: b1,
			Version: e0,
			Errors:  e2,
		}, nil
	}
}

// CreateVersionRequest collects the request parameters for the CreateVersion method.
type CreateVersionRequest struct {
	Username string           `json:"username"`
	Name     string           `json:"name"`
	Version  entities.Version `json:"version"`
}

// CreateVersionResponse collects the response parameters for the CreateVersion method.
type CreateVersionResponse struct {
	Version entities.Version `json:"version"`
	Success bool             `json:"success"`
	Errors  []error          `json:"errors"`
}

// MakeCreateVersionEndpoint returns an endpoint that invokes CreateVersion on the service.
func MakeCreateVersionEndpoint(s service.BoxstashService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateVersionRequest)
		e0, b1, e2 := s.CreateVersion(ctx, req.Username, req.Name, req.Version)
		return CreateVersionResponse{
			Success: b1,
			Version: e0,
			Errors:  e2,
		}, nil
	}
}

// UpdateVersionRequest collects the request parameters for the UpdateVersion method.
type UpdateVersionRequest struct {
	Version entities.Version `json:"version"`
}

// UpdateVersionResponse collects the response parameters for the UpdateVersion method.
type UpdateVersionResponse struct {
	Version entities.Version `json:"version"`
	Success bool             `json:"success"`
	Errors  []error          `json:"errors"`
}

// MakeUpdateVersionEndpoint returns an endpoint that invokes UpdateVersion on the service.
func MakeUpdateVersionEndpoint(s service.BoxstashService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateVersionRequest)
		e0, b1, e2 := s.UpdateVersion(ctx, req.Version)
		return UpdateVersionResponse{
			Success: b1,
			Version: e0,
			Errors:  e2,
		}, nil
	}
}

// DeleteVersionRequest collects the request parameters for the DeleteVersion method.
type DeleteVersionRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Version  string `json:"version"`
}

// DeleteVersionResponse collects the response parameters for the DeleteVersion method.
type DeleteVersionResponse struct {
	Version entities.Version `json:"version"`
	Success bool             `json:"success"`
	Errors  []error          `json:"errors"`
}

// MakeDeleteVersionEndpoint returns an endpoint that invokes DeleteVersion on the service.
func MakeDeleteVersionEndpoint(s service.BoxstashService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteVersionRequest)
		e0, b1, e2 := s.DeleteVersion(ctx, req.Username, req.Name, req.Version)
		return DeleteVersionResponse{
			Success: b1,
			Version: e0,
			Errors:  e2,
		}, nil
	}
}

// ReleaseVersionRequest collects the request parameters for the ReleaseVersion method.
type ReleaseVersionRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Version  string `json:"version"`
}

// ReleaseVersionResponse collects the response parameters for the ReleaseVersion method.
type ReleaseVersionResponse struct {
	Version entities.Version `json:"version"`
	Success bool             `json:"success"`
	Errors  []error          `json:"errors"`
}

// MakeReleaseVersionEndpoint returns an endpoint that invokes ReleaseVersion on the service.
func MakeReleaseVersionEndpoint(s service.BoxstashService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReleaseVersionRequest)
		e0, b1, e2 := s.ReleaseVersion(ctx, req.Username, req.Name, req.Version)
		return ReleaseVersionResponse{
			Success: b1,
			Version: e0,
			Errors:  e2,
		}, nil
	}
}

// RevokeVersionRequest collects the request parameters for the RevokeVersion method.
type RevokeVersionRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Version  string `json:"version"`
}

// RevokeVersionResponse collects the response parameters for the RevokeVersion method.
type RevokeVersionResponse struct {
	Version entities.Version `json:"version"`
	Success bool             `json:"success"`
	Errors  []error          `json:"errors"`
}

// MakeRevokeVersionEndpoint returns an endpoint that invokes RevokeVersion on the service.
func MakeRevokeVersionEndpoint(s service.BoxstashService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RevokeVersionRequest)
		e0, b1, e2 := s.RevokeVersion(ctx, req.Username, req.Name, req.Version)
		return RevokeVersionResponse{
			Success: b1,
			Version: e0,
			Errors:  e2,
		}, nil
	}
}

// ReadProviderRequest collects the request parameters for the ReadProvider method.
type ReadProviderRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Version  string `json:"version"`
	Provider string `json:"provider"`
}

// ReadProviderResponse collects the response parameters for the ReadProvider method.
type ReadProviderResponse struct {
	Provider entities.Provider `json:"provider"`
	Success  bool              `json:"success"`
	Errors   []error           `json:"errors"`
}

// MakeReadProviderEndpoint returns an endpoint that invokes ReadProvider on the service.
func MakeReadProviderEndpoint(s service.BoxstashService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReadProviderRequest)
		e0, b1, e2 := s.ReadProvider(ctx, req.Username, req.Name, req.Version, req.Provider)
		return ReadProviderResponse{
			Success:  b1,
			Provider: e0,
			Errors:   e2,
		}, nil
	}
}

// CreateProviderRequest collects the request parameters for the CreateProvider method.
type CreateProviderRequest struct {
	Username string            `json:"username"`
	Name     string            `json:"name"`
	Version  string            `json:"version"`
	Provider entities.Provider `json:"provider"`
}

// CreateProviderResponse collects the response parameters for the CreateProvider method.
type CreateProviderResponse struct {
	Provider entities.Provider `json:"provider"`
	Success  bool              `json:"success"`
	Errors   []error           `json:"errors"`
}

// MakeCreateProviderEndpoint returns an endpoint that invokes CreateProvider on the service.
func MakeCreateProviderEndpoint(s service.BoxstashService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateProviderRequest)
		e0, b1, e2 := s.CreateProvider(ctx, req.Username, req.Name, req.Version, req.Provider)
		return CreateProviderResponse{
			Success:  b1,
			Provider: e0,
			Errors:   e2,
		}, nil
	}
}

// UpdateProviderRequest collects the request parameters for the UpdateProvider method.
type UpdateProviderRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Version  string `json:"version"`
	Provider string `json:"provider"`
}

// UpdateProviderResponse collects the response parameters for the UpdateProvider method.
type UpdateProviderResponse struct {
	Provider entities.Provider `json:"provider"`
	Success  bool              `json:"success"`
	Errors   []error           `json:"errors"`
}

// MakeUpdateProviderEndpoint returns an endpoint that invokes UpdateProvider on the service.
func MakeUpdateProviderEndpoint(s service.BoxstashService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateProviderRequest)
		e0, b1, e2 := s.UpdateProvider(ctx, req.Username, req.Name, req.Version, req.Provider)
		return UpdateProviderResponse{
			Success:  b1,
			Provider: e0,
			Errors:   e2,
		}, nil
	}
}

// DeleteProviderRequest collects the request parameters for the DeleteProvider method.
type DeleteProviderRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Version  string `json:"version"`
	Provider string `json:"provider"`
}

// DeleteProviderResponse collects the response parameters for the DeleteProvider method.
type DeleteProviderResponse struct {
	Provider entities.Provider `json:"provider"`
	Success  bool              `json:"success"`
	Errors   []error           `json:"errors"`
}

// MakeDeleteProviderEndpoint returns an endpoint that invokes DeleteProvider on the service.
func MakeDeleteProviderEndpoint(s service.BoxstashService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteProviderRequest)
		e0, b1, e2 := s.DeleteProvider(ctx, req.Username, req.Name, req.Version, req.Provider)
		return DeleteProviderResponse{
			Success:  b1,
			Provider: e0,
			Errors:   e2,
		}, nil
	}
}

// NewAuthToken implements Service. Primarily useful in a client.
func (e Endpoints) NewAuthToken(ctx context.Context, creds entities.Credentials) (token string, success bool, errors []error) {
	request := NewAuthTokenRequest{Creds: creds}
	response, err := e.NewAuthTokenEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(NewAuthTokenResponse).Token, response.(NewAuthTokenResponse).Success, response.(NewAuthTokenResponse).Errors
}

// ValidateAuthToken implements Service. Primarily useful in a client.
func (e Endpoints) ValidateAuthToken(ctx context.Context, token string) (success bool, errors []error) {
	request := ValidateAuthTokenRequest{Token: token}
	response, err := e.ValidateAuthTokenEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ValidateAuthTokenResponse).Success, response.(ValidateAuthTokenResponse).Errors
}

// DeleteAuthToken implements Service. Primarily useful in a client.
func (e Endpoints) DeleteAuthToken(ctx context.Context, token string) (success bool, errors []error) {
	request := DeleteAuthTokenRequest{Token: token}
	response, err := e.DeleteAuthTokenEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteAuthTokenResponse).Success, response.(DeleteAuthTokenResponse).Errors
}

// ReadOrganization implements Service. Primarily useful in a client.
func (e Endpoints) ReadOrganization(ctx context.Context, org string) (organization entities.Users, success bool, errors []error) {
	request := ReadOrganizationRequest{Org: org}
	response, err := e.ReadOrganizationEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ReadOrganizationResponse).Organization, response.(ReadOrganizationResponse).Success, response.(ReadOrganizationResponse).Errors
}

// Search implements Service. Primarily useful in a client.
func (e Endpoints) Search(ctx context.Context) (boxes []entities.Box, success bool, errors []error) {
	request := SearchRequest{}
	response, err := e.SearchEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SearchResponse).Boxes, response.(SearchResponse).Success, response.(SearchResponse).Errors
}

// ReadBox implements Service. Primarily useful in a client.
func (e Endpoints) ReadBox(ctx context.Context, username string, name string) (box entities.Box, success bool, errors []error) {
	request := ReadBoxRequest{
		Name:     name,
		Username: username,
	}
	response, err := e.ReadBoxEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ReadBoxResponse).Box, response.(ReadBoxResponse).Success, response.(ReadBoxResponse).Errors
}

// CreateBox implements Service. Primarily useful in a client.
func (e Endpoints) CreateBox(ctx context.Context, b service.NewBoxRequest) (box entities.Box, success bool, errors []error) {
	request := CreateBoxRequest{Box: b}
	response, err := e.CreateBoxEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateBoxResponse).Box, response.(CreateBoxResponse).Success, response.(CreateBoxResponse).Errors
}

// UpdateBox implements Service. Primarily useful in a client.
func (e Endpoints) UpdateBox(ctx context.Context, b map[string]interface{}) (box entities.Box, success bool, errors []error) {
	request := b
	response, err := e.UpdateBoxEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateBoxResponse).Box, response.(UpdateBoxResponse).Success, response.(UpdateBoxResponse).Errors
}

// DeleteBox implements Service. Primarily useful in a client.
func (e Endpoints) DeleteBox(ctx context.Context, username string, name string) (e0 entities.Box, b1 bool, e2 []error) {
	request := DeleteBoxRequest{
		Name:     name,
		Username: username,
	}
	response, err := e.DeleteBoxEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteBoxResponse).Box, response.(DeleteBoxResponse).Success, response.(DeleteBoxResponse).Errors
}

// ReadVersion implements Service. Primarily useful in a client.
func (e Endpoints) ReadVersion(ctx context.Context, username string, name string, version string) (e0 entities.Version, b1 bool, e2 []error) {
	request := ReadVersionRequest{
		Name:     name,
		Username: username,
		Version:  version,
	}
	response, err := e.ReadVersionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ReadVersionResponse).Version, response.(ReadVersionResponse).Success, response.(ReadVersionResponse).Errors
}

// CreateVersion implements Service. Primarily useful in a client.
func (e Endpoints) CreateVersion(ctx context.Context, username string, name string, version entities.Version) (e0 entities.Version, b1 bool, e2 []error) {
	request := CreateVersionRequest{
		Name:     name,
		Username: username,
		Version:  version,
	}
	response, err := e.CreateVersionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateVersionResponse).Version, response.(CreateVersionResponse).Success, response.(CreateVersionResponse).Errors
}

// UpdateVersion implements Service. Primarily useful in a client.
func (e Endpoints) UpdateVersion(ctx context.Context, version entities.Version) (e0 entities.Version, b1 bool, e2 []error) {
	request := UpdateVersionRequest{Version: version}
	response, err := e.UpdateVersionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateVersionResponse).Version, response.(UpdateVersionResponse).Success, response.(UpdateVersionResponse).Errors
}

// DeleteVersion implements Service. Primarily useful in a client.
func (e Endpoints) DeleteVersion(ctx context.Context, username string, name string, version string) (e0 entities.Version, b1 bool, e2 []error) {
	request := DeleteVersionRequest{
		Name:     name,
		Username: username,
		Version:  version,
	}
	response, err := e.DeleteVersionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteVersionResponse).Version, response.(DeleteVersionResponse).Success, response.(DeleteVersionResponse).Errors
}

// ReleaseVersion implements Service. Primarily useful in a client.
func (e Endpoints) ReleaseVersion(ctx context.Context, username string, name string, version string) (e0 entities.Version, b1 bool, e2 []error) {
	request := ReleaseVersionRequest{
		Name:     name,
		Username: username,
		Version:  version,
	}
	response, err := e.ReleaseVersionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ReleaseVersionResponse).Version, response.(ReleaseVersionResponse).Success, response.(ReleaseVersionResponse).Errors
}

// RevokeVersion implements Service. Primarily useful in a client.
func (e Endpoints) RevokeVersion(ctx context.Context, username string, name string, version string) (e0 entities.Version, b1 bool, e2 []error) {
	request := RevokeVersionRequest{
		Name:     name,
		Username: username,
		Version:  version,
	}
	response, err := e.RevokeVersionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(RevokeVersionResponse).Version, response.(RevokeVersionResponse).Success, response.(RevokeVersionResponse).Errors
}

// ReadProvider implements Service. Primarily useful in a client.
func (e Endpoints) ReadProvider(ctx context.Context, username string, name string, version string, provider string) (e0 entities.Provider, b1 bool, e2 []error) {
	request := ReadProviderRequest{
		Name:     name,
		Provider: provider,
		Username: username,
		Version:  version,
	}
	response, err := e.ReadProviderEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ReadProviderResponse).Provider, response.(ReadProviderResponse).Success, response.(ReadProviderResponse).Errors
}

// CreateProvider implements Service. Primarily useful in a client.
func (e Endpoints) CreateProvider(ctx context.Context, username string, name string, version string, provider entities.Provider) (e0 entities.Provider, b1 bool, e2 []error) {
	request := CreateProviderRequest{
		Name:     name,
		Provider: provider,
		Username: username,
		Version:  version,
	}
	response, err := e.CreateProviderEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateProviderResponse).Provider, response.(CreateProviderResponse).Success, response.(CreateProviderResponse).Errors
}

// UpdateProvider implements Service. Primarily useful in a client.
func (e Endpoints) UpdateProvider(ctx context.Context, username string, name string, version string, provider string) (e0 entities.Provider, b1 bool, e2 []error) {
	request := UpdateProviderRequest{
		Name:     name,
		Provider: provider,
		Username: username,
		Version:  version,
	}
	response, err := e.UpdateProviderEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateProviderResponse).Provider, response.(UpdateProviderResponse).Success, response.(UpdateProviderResponse).Errors
}

// DeleteProvider implements Service. Primarily useful in a client.
func (e Endpoints) DeleteProvider(ctx context.Context, username string, name string, version string, provider string) (e0 entities.Provider, b1 bool, e2 []error) {
	request := DeleteProviderRequest{
		Name:     name,
		Provider: provider,
		Username: username,
		Version:  version,
	}
	response, err := e.DeleteProviderEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteProviderResponse).Provider, response.(DeleteProviderResponse).Success, response.(DeleteProviderResponse).Errors
}
