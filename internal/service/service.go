package service

import (
	"boxstash/internal/boxstash/entities"
	"boxstash/internal/repository"
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/opentracing/opentracing-go"
	"strings"
)

var (
	build,
	commit,
	date,
	version string
)

// BoxstashService describes the service's interface, from which endpoints are derived
type BoxstashService interface {
	About(ctx context.Context) (entities.About, error)
	NewAuthToken(ctx context.Context, creds entities.Credentials) (token string, success bool, errors []error)
	ValidateAuthToken(ctx context.Context, token string) (success bool, errors []error)
	DeleteAuthToken(ctx context.Context, token string) (success bool, errors []error)
	ReadOrganization(ctx context.Context, org string) (organization entities.Users, success bool, errors []error)
	Search(ctx context.Context) (boxes []entities.Box, success bool, errors []error)
	ReadBox(ctx context.Context, username, name string) (box entities.Box, success bool, errors []error)
	CreateBox(ctx context.Context, new NewBoxRequest) (entities.Box, bool, []error)
	UpdateBox(ctx context.Context, u map[string]interface{}) (entities.Box, bool, []error)
	DeleteBox(ctx context.Context, username, name string) (entities.Box, bool, []error)
	ReadVersion(ctx context.Context, username, name, version string) (entities.Version, bool, []error)
	CreateVersion(ctx context.Context, username, name string, version entities.Version) (entities.Version, bool, []error)
	UpdateVersion(ctx context.Context, version entities.Version) (entities.Version, bool, []error)
	DeleteVersion(ctx context.Context, username, name, version string) (entities.Version, bool, []error)
	ReleaseVersion(ctx context.Context, username, name, version string) (entities.Version, bool, []error)
	RevokeVersion(ctx context.Context, username, name, version string) (entities.Version, bool, []error)
	ReadProvider(ctx context.Context, username, name, version, provider string) (entities.Provider, bool, []error)
	CreateProvider(ctx context.Context, username, name, version string, provider entities.Provider) (entities.Provider, bool, []error)
	UpdateProvider(ctx context.Context, username, name, version, provider string) (entities.Provider, bool, []error)
	DeleteProvider(ctx context.Context, username, name, version, provider string) (entities.Provider, bool, []error)
}

type boxstashService struct {
	store  boxstashServiceInteractor
	logger log.Logger
	tracer opentracing.Tracer
}

type boxstashServiceInteractor struct {
	box repository.BoxRepository
}

// NewBoxstashServiceInteractor returns a service interactor which contains interactors for each repository type
func NewBoxstashServiceInteractor(box repository.BoxRepository) boxstashServiceInteractor {
	return boxstashServiceInteractor{box: box}
}

func (b *boxstashService) About(ctx context.Context) (e0 entities.About, e1 error) {
	// TODO implement the business logic of About
	return e0, e1
}

func (b *boxstashService) NewAuthToken(ctx context.Context, creds entities.Credentials) (token string, success bool, errors []error) {
	// TODO implement the business logic of NewAuthToken
	return token, success, errors
}

func (b *boxstashService) ValidateAuthToken(ctx context.Context, token string) (success bool, errors []error) {
	// TODO implement the business logic of ValidateAuthToken
	return success, errors
}

func (b *boxstashService) DeleteAuthToken(ctx context.Context, token string) (success bool, errors []error) {
	// TODO implement the business logic of DeleteAuthToken
	return success, errors
}

func (b *boxstashService) ReadOrganization(ctx context.Context, org string) (organization entities.Users, success bool, errors []error) {
	// TODO implement the business logic of ReadOrganization
	return organization, success, errors
}

func (b *boxstashService) Search(ctx context.Context) (boxes []entities.Box, success bool, errors []error) {
	// TODO implement the business logic of Search
	return boxes, success, errors
}

// ReadBox coordinates retrieval of the requested box record from the database
func (b *boxstashService) ReadBox(ctx context.Context, username string, name string) (box entities.Box, success bool, errors []error) {
	box, err := b.store.box.ReadBoxByName(ctx, username, name)
	if err != nil {
		errors = append(errors, err)
		return box, false, errors
	}

	return box, true, []error{}
}

// CreateBox coordinates creation of the requested box in the database, as a new record
func (b *boxstashService) CreateBox(ctx context.Context, new NewBoxRequest) (box entities.Box, success bool, errors []error) {
	success = true
	box.Username = new.Username
	box.Name = new.Name
	box.Private = new.Private
	box.ShortDescription = new.ShortDescription
	box.Description = new.Description

	box_id, err := b.store.box.CreateBox(ctx, box)
	if err != nil {
		b.logger.Log("error", "error creating new box", "component", "boxstash.internal.service.CreateBox", "err", "database insert unsuccessful")
		errors = append(errors, err)
		success = false
	}
	bx, err := b.store.box.ReadBoxByID(ctx, box_id)
	if err != nil {
		b.logger.Log("error", "error retrieving new box record", "component", "boxstash.internal.service.CreateBox", "err",
			"database query unsuccessful")
		errors = append(errors, err)
		success = false
	}
	return bx, success, errors
}

// UpdateBox coordinates updating allowed attributes of the specified box in the database
func (b *boxstashService) UpdateBox(ctx context.Context, u map[string]interface{}) (box entities.Box, success bool, errors []error) {
	if len(u) == 0 {
		b.logger.Log("error", "request contained no attributes to update", "component", "boxstash.internal.service.UpdateBox")
		errors = append(errors, fmt.Errorf("request contained no attributes to update"))
	}
	username, ok := u["username"]
	if !ok {
		b.logger.Log("error", "request missing required attribute 'username'", "component", "boxstash.internal.service.UpdateBox")
		errors = append(errors, fmt.Errorf("request missing required attribute 'username'"))
	} else if username == "" {
		b.logger.Log("error", "request contained empty required attribute 'username'", "component", "boxstash.internal.service.UpdateBox")
		errors = append(errors, fmt.Errorf("request contained empty required attribute 'username'"))
	}
	name, ok := u["name"]
	if !ok {
		b.logger.Log("error", "request missing required attribute 'name'", "component", "boxstash.internal.service.UpdateBox")
		errors = append(errors, fmt.Errorf("request missing required attribute 'name'"))
	} else if name == "" {
		b.logger.Log("error", "request contained empty required attribute 'name'", "component", "boxstash.internal.service.UpdateBox")
		errors = append(errors, fmt.Errorf("request contained empty required attribute 'name'"))
	}

	if len(errors) > 0 {
		b.logger.Log("error", "too many errors encountered to continue", "component", "boxstash.internal.service.UpdateBox")
		return entities.Box{}, false, errors
	}

	allowed_updates := []string{
		"name",
		"username",
		"short_description",
		"description",
		"is_private",
	}

	updates := make(map[string]interface{})
	for _, k := range allowed_updates {
		v, ok := u[k]
		if ok {
			updates[k] = v
			delete(u, k)
		}
	}

	if len(u) > 0 {
		b.logger.Log("warn", "request contained more attributes than allowed, ignoring extras", "component", "boxstash.internal.service.UpdateBox")
		attrs := []string{}
		for k := range u {
			attrs = append(attrs, k)
		}
		list := strings.Join(attrs, ", ")
		err := fmt.Errorf("the following attributes cannot be updated through the API and were ignored: %s", list)
		errors = append(errors, err)
	}

	if len(updates) == 0 {
		b.logger.Log("error", "request contained no attributes to update", "updates", fmt.Sprintf("%+v", updates), "component",
			"boxstash.internal.service.UpdateBox")
		err := fmt.Errorf("request contained no attributes to update")
		errors = append(errors, err)
		return box, false, errors
	}

	box_id, err := b.store.box.UpdateBox(ctx, updates)
	if err != nil {
		errors = append(errors, err)
		return box, false, errors
	}
	box, err = b.store.box.ReadBoxByID(ctx, box_id)
	if err != nil {
		errors = append(errors, err)
		return box, false, errors
	}
	return box, true, errors
}

// DeleteBox coordinates removal of a vagrant box definition from the database
func (b *boxstashService) DeleteBox(ctx context.Context, username string, name string) (box entities.Box, success bool, errors []error) {
	tag := fmt.Sprintf("%s/%s", username, name)
	b.logger.Log("trace", "request received to delete box", "box", tag, "component",
		"boxstash.internal.service.DeleteBox")
	box, err := b.store.box.DeleteBox(ctx, username, name)
	if err != nil {
		b.logger.Log("error", "failed to delete requested box record", "box", tag, "err", err, "component",
			"boxstash.internal.service.DeleteBox")
		errors = append(errors, err)
		success = false
	} else {
		success = true
	}
	return box, success, errors
}

func (b *boxstashService) ReadVersion(ctx context.Context, username string, name string, version string) (e0 entities.Version, b1 bool, e2 []error) {
	// TODO implement the business logic of ReadVersion
	return e0, b1, e2
}

func (b *boxstashService) CreateVersion(ctx context.Context, username string, name string, version entities.Version) (e0 entities.Version, b1 bool, e2 []error) {
	// TODO implement the business logic of CreateVersion
	return e0, b1, e2
}

func (b *boxstashService) UpdateVersion(ctx context.Context, version entities.Version) (e0 entities.Version, b1 bool, e2 []error) {
	// TODO implement the business logic of UpdateVersion
	return e0, b1, e2
}

func (b *boxstashService) DeleteVersion(ctx context.Context, username string, name string, version string) (e0 entities.Version, b1 bool, e2 []error) {
	// TODO implement the business logic of DeleteVersion
	return e0, b1, e2
}

func (b *boxstashService) ReleaseVersion(ctx context.Context, username string, name string, version string) (e0 entities.Version, b1 bool, e2 []error) {
	// TODO implement the business logic of ReleaseVersion
	return e0, b1, e2
}

func (b *boxstashService) RevokeVersion(ctx context.Context, username string, name string, version string) (e0 entities.Version, b1 bool, e2 []error) {
	// TODO implement the business logic of RevokeVersion
	return e0, b1, e2
}

func (b *boxstashService) ReadProvider(ctx context.Context, username string, name string, version string, provider string) (e0 entities.Provider, b1 bool, e2 []error) {
	// TODO implement the business logic of ReadProvider
	return e0, b1, e2
}

func (b *boxstashService) CreateProvider(ctx context.Context, username string, name string, version string, provider entities.Provider) (e0 entities.Provider, b1 bool, e2 []error) {
	// TODO implement the business logic of CreateProvider
	return e0, b1, e2
}

func (b *boxstashService) UpdateProvider(ctx context.Context, username string, name string, version string, provider string) (e0 entities.Provider, b1 bool, e2 []error) {
	// TODO implement the business logic of UpdateProvider
	return e0, b1, e2
}

func (b *boxstashService) DeleteProvider(ctx context.Context, username string, name string, version string, provider string) (e0 entities.Provider, b1 bool, e2 []error) {
	// TODO implement the business logic of DeleteProvider
	return e0, b1, e2
}

// NewBoxstashService returns an instance of the BoxstashService injected with everything the service needs to perform its duties (db, logger,
// tracer, etc)
func NewBoxstashService(s boxstashServiceInteractor, l log.Logger, t opentracing.Tracer) BoxstashService {
	return &boxstashService{
		store:  s,
		logger: l,
		tracer: t,
	}
}

// New returns a BoxstashService with all of the expected middleware wired in.
func New(middleware []Middleware, i boxstashServiceInteractor, l log.Logger, t opentracing.Tracer) BoxstashService {
	var svc BoxstashService = NewBoxstashService(i, l, t)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
