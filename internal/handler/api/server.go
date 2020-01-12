package api

import (
	"context"
	"crypto/tls"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
)

type loggerKey struct{}

// ListenAndServe starts an HTTP server listening on the network
func (s Server) ListenAndServe(ctx context.Context) error {
	if s.LetsEncrypt {
		return s.listenAndServeLetsEncrypt(ctx)
	} else if s.Key != "" {
		return s.listenAndServeTLS(ctx)
	}
	return s.listenAndServe(ctx)
}

func (s Server) listenAndServe(ctx context.Context) error {
	var g errgroup.Group
	s1 := &http.Server{
		Addr:    s.Addr,
		Handler: s.Handler,
	}
	s.logger.WithFields(logrus.Fields{
		"func": "listenAndServe",
		"addr": s.Addr,
		"proto": "http",
		"host": s.Host}).Info("starting HTTP server")
	g.Go(func() error {
		select {
		case <-ctx.Done():
			return s1.Shutdown(ctx)
		}
	})
	g.Go(func() error {
		return s1.ListenAndServe()
	})
	return g.Wait()
}

func (s Server) listenAndServeTLS(ctx context.Context) error {
	s.logger.WithFields(logrus.Fields{
		"func": "listenAndServe",
		"addr": s.Addr,
		"proto": "https",
		"host": s.Host}).Info("starting HTTPS server")
	var g errgroup.Group
	s1 := &http.Server{
		Addr:    ":http",
		Handler: http.HandlerFunc(redirect),
	}
	s2 := &http.Server{
		Addr:    ":https",
		Handler: s.Handler,
	}
	g.Go(func() error {
		return s1.ListenAndServe()
	})
	g.Go(func() error {
		return s2.ListenAndServeTLS(
			s.Cert,
			s.Key,
		)
	})
	g.Go(func() error {
		select {
		case <-ctx.Done():
			s1.Shutdown(ctx)
			s2.Shutdown(ctx)
			return nil
		}
	})
	return g.Wait()
}

func (s Server) listenAndServeLetsEncrypt(ctx context.Context) error {
	s.logger.WithFields(logrus.Fields{
		"func": "listenAndServe",
		"addr": s.Addr,
		"proto": "https",
		"host": s.Host}).Info("starting HTTPS server using LetsEncrypt")
	var g errgroup.Group

	c := cacheDir()
	m := &autocert.Manager{
		Email:      s.Email,
		Cache:      autocert.DirCache(c),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(s.Host),
	}
	s1 := &http.Server{
		Addr:    ":http",
		Handler: m.HTTPHandler(s.Handler),
	}
	s2 := &http.Server{
		Addr:    ":https",
		Handler: s.Handler,
		TLSConfig: &tls.Config{
			GetCertificate: m.GetCertificate,
			NextProtos:     []string{"h2", "http/1.1"},
			MinVersion:     tls.VersionTLS12,
		},
	}
	g.Go(func() error {
		return s1.ListenAndServe()
	})
	g.Go(func() error {
		return s2.ListenAndServeTLS("", "")
	})
	g.Go(func() error {
		select {
		case <-ctx.Done():
			s1.Shutdown(ctx)
			s2.Shutdown(ctx)
			return nil
		}
	})
	return g.Wait()
}

func redirect(w http.ResponseWriter, req *http.Request) {
	target := strings.Join([]string{"https://", req.Host, req.URL.Path}, "")
	http.Redirect(w, req, target, http.StatusTemporaryRedirect)
}

func cacheDir() string {
	const base = "letsencrypt"
	if xdg := os.Getenv("XDG_CACHE_HOME"); xdg != "" {
		return filepath.Join(xdg, base)
	}
	return filepath.Join(os.Getenv("HOME"), ".cache", base)
}
