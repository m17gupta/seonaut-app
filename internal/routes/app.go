package routes

import (
	"fmt"
	"log"
	"net/http"
	"github.com/stjudewashere/seonaut/internal/models"
	"github.com/stjudewashere/seonaut/internal/services"
	"mime"
	
)

func init() {
    // Ensure .css gets the right Content-Type everywhere
    _ = mime.AddExtensionType(".css", "text/css; charset=utf-8")
}

// PageView is the data structure used to render the html templates.
type PageView struct {
	PageTitle string
	User      models.User
	Data      interface{}
	Refresh   bool
}

// NewServer sets up the HTTP server routes and starts the HTTP server.
func NewServer(container *services.Container) {
	// Create a new ServeMux for better control over routing
	mux := http.NewServeMux()

	// Handle static files
	fileServer := http.FileServer(http.Dir("./web/static"))
	mux.Handle("GET /css/",   http.StripPrefix("/css/",   http.FileServer(http.Dir("./web/css"))))
	mux.Handle("GET /resources/", http.StripPrefix("/resources", fileServer))
	mux.Handle("GET /robots.txt", fileServer)
	mux.Handle("GET /favicon.ico", fileServer)

	// Health check endpoint for Render (with CORS)
	mux.HandleFunc("GET /health", CORSHandler(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "healthy", "service": "seonaut"}`))
	}))

	// API endpoints with CORS support
	mux.HandleFunc("OPTIONS /api/", CORSHandler(func(w http.ResponseWriter, r *http.Request) {
		// Handle preflight requests for all API endpoints
		w.WriteHeader(http.StatusOK)
	}))

	// Crawler routes
	crawlHandler := crawlHandler{container}
	mux.HandleFunc("GET /crawl/start", CORSHandler(container.CookieSession.Auth(crawlHandler.startHandler)))
	mux.HandleFunc("GET /crawl/stop", CORSHandler(container.CookieSession.Auth(crawlHandler.stopHandler)))
	mux.HandleFunc("GET /crawl/live", CORSHandler(container.CookieSession.Auth(crawlHandler.liveCrawlHandler)))
	mux.HandleFunc("GET /crawl/auth", CORSHandler(container.CookieSession.Auth(crawlHandler.authGetHandler)))
	mux.HandleFunc("POST /crawl/auth", CORSHandler(container.CookieSession.Auth(crawlHandler.authPostHandler)))
	mux.HandleFunc("GET /crawl/ws", CORSHandler(container.CookieSession.Auth(crawlHandler.wsHandler)))

	// Dashboard route
	dashboardHandler := dashboardHandler{container}
	mux.HandleFunc("GET /dashboard", CORSHandler(container.CookieSession.Auth(dashboardHandler.indexHandler)))

	// URL explorer route
	explorerHandler := explorerHandler{container}
	mux.HandleFunc("GET /explorer", CORSHandler(container.CookieSession.Auth(explorerHandler.indexHandler)))

	// Data export routes
	exportHandler := exportHandler{container}
	mux.HandleFunc("GET /export", CORSHandler(container.CookieSession.Auth(exportHandler.indexHandler)))
	mux.HandleFunc("GET /export/csv", CORSHandler(container.CookieSession.Auth(exportHandler.csvHandler)))
	mux.HandleFunc("GET /export/sitemap", CORSHandler(container.CookieSession.Auth(exportHandler.sitemapHandler)))
	mux.HandleFunc("GET /export/resources", CORSHandler(container.CookieSession.Auth(exportHandler.resourcesHandler)))
	mux.HandleFunc("GET /export/wazc", CORSHandler(container.CookieSession.Auth(exportHandler.waczHandler)))

	// Issues routes
	issueHandler := issueHandler{container}
	mux.HandleFunc("GET /issues", CORSHandler(container.CookieSession.Auth(issueHandler.indexHandler)))
	mux.HandleFunc("GET /issues/view", CORSHandler(container.CookieSession.Auth(issueHandler.viewHandler)))

	// Project routes
	projectHandler := projectHandler{container}
	mux.HandleFunc("GET /", CORSHandler(container.CookieSession.Auth(projectHandler.indexHandler)))
	mux.HandleFunc("GET /project/add", CORSHandler(container.CookieSession.Auth(projectHandler.addGetHandler)))
	mux.HandleFunc("POST /project/add", CORSHandler(container.CookieSession.Auth(projectHandler.addPostHandler)))
	mux.HandleFunc("GET /project/edit", CORSHandler(container.CookieSession.Auth(projectHandler.editGetHandler)))
	mux.HandleFunc("POST /project/edit", CORSHandler(container.CookieSession.Auth(projectHandler.editPostHandler)))
	mux.HandleFunc("GET /project/delete", CORSHandler(container.CookieSession.Auth(projectHandler.deleteHandler)))

	// Resource route
	resourceHandler := resourceHandler{container}
	mux.HandleFunc("GET /resources", CORSHandler(container.CookieSession.Auth(resourceHandler.indexHandler)))

	// Archive Handler
	archiveHandler := archiveHandler{container}
	mux.HandleFunc("GET /archive", CORSHandler(container.CookieSession.Auth(archiveHandler.archiveHandler)))
	mux.HandleFunc("GET /archive/download", CORSHandler(container.CookieSession.Auth(archiveHandler.downloadHandler)))

	// User routes
	userHandler := userHandler{container}
	mux.HandleFunc("GET /signup", CORSHandler(userHandler.signupGetHandler))
	mux.HandleFunc("POST /signup", CORSHandler(userHandler.signupPostHandler))
	mux.HandleFunc("GET /signin", CORSHandler(userHandler.signinGetHandler))
	mux.HandleFunc("POST /signin", CORSHandler(userHandler.signinPostHandler))
	mux.HandleFunc("GET /account", CORSHandler(container.CookieSession.Auth(userHandler.editGetHandler)))
	mux.HandleFunc("POST /account", CORSHandler(container.CookieSession.Auth(userHandler.editPostHandler)))
	mux.HandleFunc("GET /account/delete", CORSHandler(container.CookieSession.Auth((userHandler.deleteGetHandler))))
	mux.HandleFunc("POST /account/delete", CORSHandler(container.CookieSession.Auth((userHandler.deletePostHandler))))
	mux.HandleFunc("GET /signout", CORSHandler(container.CookieSession.Auth(userHandler.signoutHandler)))

	// Support SEOnaut
	supportHandler := supportHandler{container}
	mux.HandleFunc("GET /support-seonaut", CORSHandler(container.CookieSession.Auth(supportHandler.handleSupportSEOnaut)))

	// Replay routes
	replayHandler := replayHandler{container}
	mux.HandleFunc("GET /replay", CORSHandler(container.CookieSession.Auth(replayHandler.proxyHandler)))

	// Register API routes for frontend integration
	RegisterAPIRoutes(mux, container)

	fmt.Printf("Starting server at %s on port %d...\n", container.Config.HTTPServer.Server, container.Config.HTTPServer.Port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", container.Config.HTTPServer.Server, container.Config.HTTPServer.Port), mux)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
