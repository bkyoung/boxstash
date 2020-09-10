package http

import (
	endpoint "boxstash/internal/endpoint"
	"context"
	"encoding/json"
	"errors"
	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	http1 "net/http"
)

// makeAboutHandler creates the handler logic
func makeAboutHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/about").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.AboutEndpoint, decodeAboutRequest, encodeAboutResponse, options...)))
}

// decodeAboutRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAboutRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.AboutRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeAboutResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAboutResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeNewAuthTokenHandler creates the handler logic
func makeNewAuthTokenHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/new-auth-token").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.NewAuthTokenEndpoint, decodeNewAuthTokenRequest, encodeNewAuthTokenResponse, options...)))
}

// decodeNewAuthTokenRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeNewAuthTokenRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.NewAuthTokenRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeNewAuthTokenResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeNewAuthTokenResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeValidateAuthTokenHandler creates the handler logic
func makeValidateAuthTokenHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/validate-auth-token").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.ValidateAuthTokenEndpoint, decodeValidateAuthTokenRequest, encodeValidateAuthTokenResponse, options...)))
}

// decodeValidateAuthTokenRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeValidateAuthTokenRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.ValidateAuthTokenRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeValidateAuthTokenResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeValidateAuthTokenResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteAuthTokenHandler creates the handler logic
func makeDeleteAuthTokenHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/delete-auth-token").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.DeleteAuthTokenEndpoint, decodeDeleteAuthTokenRequest, encodeDeleteAuthTokenResponse, options...)))
}

// decodeDeleteAuthTokenRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteAuthTokenRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.DeleteAuthTokenRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteAuthTokenResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteAuthTokenResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeReadOrganizationHandler creates the handler logic
func makeReadOrganizationHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/read-organization").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.ReadOrganizationEndpoint, decodeReadOrganizationRequest, encodeReadOrganizationResponse, options...)))
}

// decodeReadOrganizationRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeReadOrganizationRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.ReadOrganizationRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeReadOrganizationResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeReadOrganizationResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeSearchHandler creates the handler logic
func makeSearchHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/search").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.SearchEndpoint, decodeSearchRequest, encodeSearchResponse, options...)))
}

// decodeSearchRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSearchRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.SearchRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeSearchResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSearchResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeReadBoxHandler creates the handler logic
func makeReadBoxHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/api/v1/box/{username}/{name}").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}),
		handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.ReadBoxEndpoint, decodeReadBoxRequest, encodeReadBoxResponse, options...)))
}

// decodeReadBoxRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeReadBoxRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	req := endpoint.ReadBoxRequest{
		Name: vars["name"],
		Username: vars["username"],
	}
	//err := json.NewDecoder(r.Body).Decode(&req)
	return req, nil
}

// encodeReadBoxResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeReadBoxResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCreateBoxHandler creates the handler logic
func makeCreateBoxHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/api/v1/boxes").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}),
		handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.CreateBoxEndpoint, decodeCreateBoxRequest, encodeCreateBoxResponse, options...)))
}

// decodeCreateBoxRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateBoxRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.CreateBoxRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateBoxResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateBoxResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateBoxHandler creates the handler logic
func makeUpdateBoxHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("PUT").Path("/api/v1/box/{username}/{name}").Handler(handlers.CORS(handlers.AllowedMethods([]string{"PUT"}),
		handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.UpdateBoxEndpoint, decodeUpdateBoxRequest, encodeUpdateBoxResponse, options...)))
}

// decodeUpdateBoxRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateBoxRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	req := endpoint.UpdateBoxRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	req.Box["name"] = vars["name"]
	req.Box["username"] = vars["username"]
	return req, err
}

// encodeUpdateBoxResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateBoxResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteBoxHandler creates the handler logic
func makeDeleteBoxHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("DELETE").Path("/api/v1/box/{username}/{name}").Handler(handlers.CORS(handlers.AllowedMethods([]string{"DELETE"}),
		handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.DeleteBoxEndpoint, decodeDeleteBoxRequest, encodeDeleteBoxResponse, options...)))
}

// decodeDeleteBoxRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteBoxRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	req := endpoint.DeleteBoxRequest{}
	//err := json.NewDecoder(r.Body).Decode(&req)
	req.Name = vars["name"]
	req.Username = vars["username"]
	return req, nil
}

// encodeDeleteBoxResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteBoxResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeReadVersionHandler creates the handler logic
func makeReadVersionHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/read-version").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.ReadVersionEndpoint, decodeReadVersionRequest, encodeReadVersionResponse, options...)))
}

// decodeReadVersionRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeReadVersionRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.ReadVersionRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeReadVersionResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeReadVersionResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCreateVersionHandler creates the handler logic
func makeCreateVersionHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/create-version").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.CreateVersionEndpoint, decodeCreateVersionRequest, encodeCreateVersionResponse, options...)))
}

// decodeCreateVersionRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateVersionRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.CreateVersionRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateVersionResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateVersionResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateVersionHandler creates the handler logic
func makeUpdateVersionHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/update-version").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.UpdateVersionEndpoint, decodeUpdateVersionRequest, encodeUpdateVersionResponse, options...)))
}

// decodeUpdateVersionRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateVersionRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.UpdateVersionRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateVersionResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateVersionResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteVersionHandler creates the handler logic
func makeDeleteVersionHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/delete-version").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.DeleteVersionEndpoint, decodeDeleteVersionRequest, encodeDeleteVersionResponse, options...)))
}

// decodeDeleteVersionRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteVersionRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.DeleteVersionRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteVersionResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteVersionResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeReleaseVersionHandler creates the handler logic
func makeReleaseVersionHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/release-version").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.ReleaseVersionEndpoint, decodeReleaseVersionRequest, encodeReleaseVersionResponse, options...)))
}

// decodeReleaseVersionRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeReleaseVersionRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.ReleaseVersionRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeReleaseVersionResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeReleaseVersionResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeRevokeVersionHandler creates the handler logic
func makeRevokeVersionHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/revoke-version").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.RevokeVersionEndpoint, decodeRevokeVersionRequest, encodeRevokeVersionResponse, options...)))
}

// decodeRevokeVersionRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeRevokeVersionRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.RevokeVersionRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeRevokeVersionResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeRevokeVersionResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeReadProviderHandler creates the handler logic
func makeReadProviderHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/read-provider").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.ReadProviderEndpoint, decodeReadProviderRequest, encodeReadProviderResponse, options...)))
}

// decodeReadProviderRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeReadProviderRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.ReadProviderRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeReadProviderResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeReadProviderResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCreateProviderHandler creates the handler logic
func makeCreateProviderHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/create-provider").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.CreateProviderEndpoint, decodeCreateProviderRequest, encodeCreateProviderResponse, options...)))
}

// decodeCreateProviderRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateProviderRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.CreateProviderRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateProviderResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateProviderResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateProviderHandler creates the handler logic
func makeUpdateProviderHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/update-provider").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.UpdateProviderEndpoint, decodeUpdateProviderRequest, encodeUpdateProviderResponse, options...)))
}

// decodeUpdateProviderRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateProviderRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.UpdateProviderRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateProviderResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateProviderResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteProviderHandler creates the handler logic
func makeDeleteProviderHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/delete-provider").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.DeleteProviderEndpoint, decodeDeleteProviderRequest, encodeDeleteProviderResponse, options...)))
}

// decodeDeleteProviderRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteProviderRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.DeleteProviderRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteProviderResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteProviderResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http1.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http1.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http1.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
