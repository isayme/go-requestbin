package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	gdig "github.com/isayme/go-gdig"
	logger "github.com/isayme/go-logger"
	"github.com/isayme/go-requestbin/app/conf"
	"github.com/isayme/go-requestbin/app/router"
	"github.com/r3labs/sse/v2"
)

func main() {
	config := conf.Get()

	if config.Logger.Level != "" {
		logger.SetLevel(config.Logger.Level)
	}

	r := mux.NewRouter()

	err := gdig.Invoke(func(request *router.Request, sseServer *sse.Server) error {
		r.HandleFunc("/api/sse", func(w http.ResponseWriter, r *http.Request) {
			sseServer.ServeHTTP(w, r)
		})

		r.HandleFunc("/api/{slug:[0-9a-zA-Z]+}/inspect", request.ListRequests)

		r.HandleFunc("/{slug:[0-9a-zA-Z]+}", request.RecordRequest)

		return nil
	})
	if err != nil {
		panic(err)
	}

	spa := spaHandler{staticPath: "public", indexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)

	// {
	// 	webProxy, err := newWebProxy("http://127.0.0.1:5173")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	r.PathPrefix("/").Handler(webProxy)
	// }

	addr := fmt.Sprintf(":%d", config.HTTP.Port)
	logger.Debugf("listen %s ...", addr)
	http.ListenAndServe(addr, r)
}

func newWebProxy(targetHost string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}

	return httputil.NewSingleHostReverseProxy(url), nil
}

// spaHandler implements the http.Handler interface, so we can use it
// to respond to HTTP requests. The path to the static directory and
// path to the index file within that static directory are used to
// serve the SPA in the given static directory.
type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Join internally call path.Clean to prevent directory traversal
	path := filepath.Join(h.staticPath, r.URL.Path)

	// check whether a file exists or is a directory at the given path
	fi, err := os.Stat(path)
	if os.IsNotExist(err) || fi.IsDir() {
		// file does not exist or path is a directory, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	}

	if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static file
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}
