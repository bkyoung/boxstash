package api

import (
	"boxstash/internal/boxstash/service"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Handler returns an http.Handler with all valid API routes/endpoints
func (s *serviceInteractor) Handler() http.Handler {
	r := chi.NewRouter()
	// TODO: Create a route for '/' that returns all categories/list of endpoints
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/boxes", s.CreateBox())
		r.Route("/box/{username}/{name}", func(r chi.Router) {
			r.Get("/", s.FindBox())
			r.Put("/", s.UpdateBox())
			r.Delete("/", s.DeleteBox())
			r.Post("/versions", s.CreateVersion())
			r.Get("/versions", s.ListVersions())
			r.Route("/version/{version}", func(r chi.Router) {
				r.Get("/", s.FindVersion())
				r.Put("/", s.UpdateVersion())
				r.Delete("/", s.DeleteVersion())
				r.Put("/release", s.ReleaseVersion())
				r.Put("/revoke", s.RevokeVersion())
				r.Post("/providers", s.CreateProvider())
				r.Get("/providers", s.ListProviders())
				r.Route("/provider/{provider}", func(r chi.Router){
					r.Get("/", s.FindProvider())
					r.Put("/", s.UpdateProvider())
					r.Delete("/", s.DeleteProvider())
				})
			})
		})
		r.Post("/users", s.CreateUser())
		r.Route("/user/{username}", func(r chi.Router) {
			r.Get("/", s.FindUser())
			r.Put("/", s.UpdateUser())
			r.Delete("/", s.DeleteUser())
		})
	})
	return r
}

// TODO: Add Acme and TLS support
// TODO: Add session support to requests (w/ session ids)
// Server is what we attach our http.Handler to for exposing boxstash functionality via HTTP
type Server struct {
	LetsEncrypt bool
	Email       string
	Cert        string
	Key         string
	Host        string
	Addr        string
	logger 		*logrus.Logger
	Handler     http.Handler
	interactor  *serviceInteractor
}

// New returns an initialized Server
func New(
	letsencrypt bool,
	email string,
	cert string,
	key string,
	host string,
	addr string,
	logger *logrus.Logger,
	handler http.Handler,
	interactor *serviceInteractor,
) *Server {
	return &Server{
		LetsEncrypt: letsencrypt,
		Email:       email,
		Cert:        cert,
		Key:         key,
		Host:        host,
		Addr:        addr,
		logger:		 logger,
		Handler:     handler,
		interactor:  interactor,
	}
}

// serviceInteractor intermediates between the api and the core application
type serviceInteractor struct {
	boxService service.BoxService
	logger *logrus.Logger
}

// ServiceInteractor provides simplified methods for the interactor to work with the application
type ServiceInteractor interface {
	CreateBox() http.HandlerFunc
	CreateVersion() http.HandlerFunc
	CreateProvider() http.HandlerFunc
	DeleteBox() http.HandlerFunc
	DeleteVersion() http.HandlerFunc
	DeleteProvider() http.HandlerFunc
	ListBoxes() http.HandlerFunc
	ListVersions() http.HandlerFunc
	ListProviders() http.HandlerFunc
	FindBox() http.HandlerFunc
	FindVersion() http.HandlerFunc
	FindProvider() http.HandlerFunc
	UpdateBox() http.HandlerFunc
	UpdateVersion() http.HandlerFunc
	UpdateProvider() http.HandlerFunc
	ReleaseVersion() http.HandlerFunc
	RevokeVersion() http.HandlerFunc
}

// NewInteractor returns a new serviceInteractor
func NewInteractor(s service.BoxService, logger *logrus.Logger) *serviceInteractor {
	return &serviceInteractor{
		boxService: s,
		logger: logger,
	}
}
