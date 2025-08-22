package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/stjudewashere/seonaut/internal/services"
)

type apiHandler struct {
	container *services.Container
}

// APIResponse represents a standard API response
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

// RegisterAPIRoutes registers all API routes for the frontend
func RegisterAPIRoutes(mux *http.ServeMux, container *services.Container) {
	apiHandler := apiHandler{container}
	
	// Authentication API routes
	mux.HandleFunc("POST /api/auth/signup", CORSHandler(apiHandler.signupAPIHandler))
	mux.HandleFunc("POST /api/auth/signin", CORSHandler(apiHandler.signinAPIHandler))
	mux.HandleFunc("POST /api/auth/signout", CORSHandler(container.CookieSession.Auth(apiHandler.signoutAPIHandler)))
	mux.HandleFunc("GET /api/auth/user", CORSHandler(container.CookieSession.Auth(apiHandler.getCurrentUserAPIHandler)))
	
	// Project API routes
	mux.HandleFunc("GET /api/projects", CORSHandler(container.CookieSession.Auth(apiHandler.getProjectsAPIHandler)))
	mux.HandleFunc("POST /api/projects", CORSHandler(container.CookieSession.Auth(apiHandler.createProjectAPIHandler)))
	mux.HandleFunc("GET /api/projects/{id}", CORSHandler(container.CookieSession.Auth(apiHandler.getProjectAPIHandler)))
	mux.HandleFunc("PUT /api/projects/{id}", CORSHandler(container.CookieSession.Auth(apiHandler.updateProjectAPIHandler)))
	mux.HandleFunc("DELETE /api/projects/{id}", CORSHandler(container.CookieSession.Auth(apiHandler.deleteProjectAPIHandler)))
	
	// Crawl API routes
	mux.HandleFunc("POST /api/projects/{id}/crawl/start", CORSHandler(container.CookieSession.Auth(apiHandler.startCrawlAPIHandler)))
	mux.HandleFunc("POST /api/projects/{id}/crawl/stop", CORSHandler(container.CookieSession.Auth(apiHandler.stopCrawlAPIHandler)))
	mux.HandleFunc("GET /api/projects/{id}/crawl/status", CORSHandler(container.CookieSession.Auth(apiHandler.getCrawlStatusAPIHandler)))
	
	// Issues API routes
	mux.HandleFunc("GET /api/projects/{id}/issues", CORSHandler(container.CookieSession.Auth(apiHandler.getIssuesAPIHandler)))
	mux.HandleFunc("GET /api/projects/{id}/issues/{issue_id}", CORSHandler(container.CookieSession.Auth(apiHandler.getIssueAPIHandler)))
	
	// Dashboard/Analytics API routes
	mux.HandleFunc("GET /api/projects/{id}/dashboard", CORSHandler(container.CookieSession.Auth(apiHandler.getDashboardDataAPIHandler)))
	mux.HandleFunc("GET /api/projects/{id}/pages", CORSHandler(container.CookieSession.Auth(apiHandler.getPagesAPIHandler)))
	
	// Export API routes
	mux.HandleFunc("GET /api/projects/{id}/export/csv", CORSHandler(container.CookieSession.Auth(apiHandler.exportCSVAPIHandler)))
	mux.HandleFunc("GET /api/projects/{id}/export/sitemap", CORSHandler(container.CookieSession.Auth(apiHandler.exportSitemapAPIHandler)))
}

// Helper function to send JSON response
func (h *apiHandler) sendJSONResponse(w http.ResponseWriter, statusCode int, response APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// Helper function to get project ID from URL
func (h *apiHandler) getProjectIDFromURL(r *http.Request) (int, error) {
	idStr := r.PathValue("id")
	return strconv.Atoi(idStr)
}

// Authentication API handlers
func (h *apiHandler) signupAPIHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.sendJSONResponse(w, http.StatusMethodNotAllowed, APIResponse{
			Success: false,
			Error:   "Method not allowed",
		})
		return
	}
	
	// TODO: Implement signup logic
	h.sendJSONResponse(w, http.StatusNotImplemented, APIResponse{
		Success: false,
		Error:   "Not implemented yet",
	})
}

func (h *apiHandler) signinAPIHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.sendJSONResponse(w, http.StatusMethodNotAllowed, APIResponse{
			Success: false,
			Error:   "Method not allowed",
		})
		return
	}
	
	// TODO: Implement signin logic
	h.sendJSONResponse(w, http.StatusNotImplemented, APIResponse{
		Success: false,
		Error:   "Not implemented yet",
	})
}

func (h *apiHandler) signoutAPIHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement signout logic
	h.sendJSONResponse(w, http.StatusNotImplemented, APIResponse{
		Success: false,
		Error:   "Not implemented yet",
	})
}

func (h *apiHandler) getCurrentUserAPIHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement get current user logic
	h.sendJSONResponse(w, http.StatusNotImplemented, APIResponse{
		Success: false,
		Error:   "Not implemented yet",
	})
}

// Project API handlers
func (h *apiHandler) getProjectsAPIHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement get projects logic
	h.sendJSONResponse(w, http.StatusNotImplemented, APIResponse{
		Success: false,
		Error:   "Not implemented yet",
	})
}

func (h *apiHandler) createProjectAPIHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement create project logic
	h.sendJSONResponse(w, http.StatusNotImplemented, APIResponse{
		Success: false,
		Error:   "Not implemented yet",
	})
}

func (h *apiHandler) getProjectAPIHandler(w http.ResponseWriter, r *http.Request) {
	projectID, err := h.getProjectIDFr omURL(r)
	if err != nil {
		h.sendJSONResponse(w, http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "Invalid project ID",
		})
		return
	}
	
	// TODO: Implement get project logic using projectID
	h.sendJSONResponse(w, http.StatusNotImplemented, APIResponse{
		Success: false,
		Error:   "Not implemented yet",
	})
}

func (h *apiHandler) updateProjectAPIHandler(w http.ResponseWriter, r *http.Request) {
       _, err := h.getProjectIDFromURL(r)
       if err != nil {
	       h.sendJSONResponse(w, http.StatusBadRequest, APIResponse{
		       Success: false,
		       Error:   "Invalid project ID",
	       })
	       return
       }
	
	// TODO: Implement update project logic using projectID
	h.sendJSONResponse(w, http.StatusNotImplemented, APIResponse{
		Success: false,
		Error:   "Not implemented yet",
	})
}

func (h *apiHandler) deleteProjectAPIHandler(w http.ResponseWriter, r *http.Request) {
       _, err := h.getProjectIDFromURL(r)
       if err != nil {
	       h.sendJSONResponse(w, http.StatusBadRequest, APIResponse{
		       Success: false,
		       Error:   "Invalid project ID",
	       })
	       return
       }
	
	// TODO: Implement delete project logic using projectID
	h.sendJSONResponse(w, http.StatusNotImplemented, APIResponse{
		Success: false,
		Error:   "Not implemented yet",
	})
}

// Crawl API handlers
func (h *apiHandler) startCrawlAPIHandler(w http.ResponseWriter, r *http.Request) {
       _, err := h.getProjectIDFromURL(r)
       if err != nil {
	       h.sendJSONResponse(w, http.StatusBadRequest, APIResponse{
		       Success: false,
		       Error:   "Invalid project ID",
	       })
	       return
       }
	
	// TODO: Implement start crawl logic using projectID
	h.sendJSONResponse(w, http.StatusNotImplemented, APIResponse{
		Success: false,
		Error:   "Not implemented yet",
	})
}

func (h *apiHandler) stopCrawlAPIHandler(w http.ResponseWriter, r *http.Request) {
       _, err := h.getProjectIDFromURL(r)
       if err != nil {
	       h.sendJSONResponse(w, http.StatusBadRequest, APIResponse{
		       Success: false,
		       Error:   "Invalid project ID",
	       })
	       return
       }
	
	// TODO: Implement stop crawl logic using projectID
	h.sendJSONResponse(w, http.StatusNotImplemented, APIResponse{
		Success: false,
		Error:   "Not implemented yet",
	})
}

func (h *apiHandler) getCrawlStatusAPIHandler(w http.ResponseWriter, r *http.Request) {
       _, err := h.getProjectIDFromURL(r)
       if err != nil {
	       h.sendJSONResponse(w, http.StatusBadRequest, APIResponse{
		       Success: false,
		       Error:   "Invalid project ID",
	       })
	       return
       }
	
	// TODO: Implement get crawl status logic using projectID
	h.sendJSONResponse(w, http.StatusNotImplemented, APIResponse{
		Success: false,
		Error:   "Not implemented yet",
	})
}

// Issues API handlers
func (h *apiHandler) getIssuesAPIHandler(w http.ResponseWriter, r *http.Request) {
       _, err := h.getProjectIDFromURL(r)
       if err != nil {
	       h.sendJSONResponse(w, http.StatusBadRequest, APIResponse{
		       Success: false,
		       Error:   "Invalid project ID",
	       })
	       return
       }
	
	// TODO: Implement get issues logic using projectID
	h.sendJSONResponse(w, http.StatusNotImplemented, APIResponse{
		Success: false,
		Error:   "Not implemented yet",
	})
}

func (h *apiHandler) getIssueAPIHandler(w http.ResponseWriter, r *http.Request) {
       _, err := h.getProjectIDFromURL(r)
       if err != nil {
	       h.sendJSONResponse(w, http.StatusBadRequest, APIResponse{
		       Success: false,
		       Error:   "Invalid project ID",
	       })
	       return
       }
	
	// TODO: Implement get specific issue logic using projectID and issue_id
	h.sendJSONResponse(w, http.StatusNotImplemented, APIResponse{
		Success: false,
		Error:   "Not implemented yet",
	})
}

// Dashboard API handlers
func (h *apiHandler) getDashboardDataAPIHandler(w http.ResponseWriter, r *http.Request) {
       _, err := h.getProjectIDFromURL(r)
       if err != nil {
	       h.sendJSONResponse(w, http.StatusBadRequest, APIResponse{
		       Success: false,
		       Error:   "Invalid project ID",
	       })
	       return
       }
	
	// TODO: Implement get dashboard data logic using projectID
	h.sendJSONResponse(w, http.StatusNotImplemented, APIResponse{
		Success: false,
		Error:   "Not implemented yet",
	})
}

func (h *apiHandler) getPagesAPIHandler(w http.ResponseWriter, r *http.Request) {
       _, err := h.getProjectIDFromURL(r)
       if err != nil {
	       h.sendJSONResponse(w, http.StatusBadRequest, APIResponse{
		       Success: false,
		       Error:   "Invalid project ID",
	       })
	       return
       }
	
	// TODO: Implement get pages logic using projectID
	h.sendJSONResponse(w, http.StatusNotImplemented, APIResponse{
		Success: false,
		Error:   "Not implemented yet",
	})
}

// Export API handlers
func (h *apiHandler) exportCSVAPIHandler(w http.ResponseWriter, r *http.Request) {
       _, err := h.getProjectIDFromURL(r)
       if err != nil {
	       h.sendJSONResponse(w, http.StatusBadRequest, APIResponse{
		       Success: false,
		       Error:   "Invalid project ID",
	       })
	       return
       }
	
	// TODO: Implement CSV export logic using projectID
	h.sendJSONResponse(w, http.StatusNotImplemented, APIResponse{
		Success: false,
		Error:   "Not implemented yet",
	})
}

func (h *apiHandler) exportSitemapAPIHandler(w http.ResponseWriter, r *http.Request) {
	projectID, err := h.getProjectIDFromURL(r)
	if err != nil {
		h.sendJSONResponse(w, http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "Invalid project ID",
		})
		return
	}
	
	// TODO: Implement sitemap export logic using projectID
	h.sendJSONResponse(w, http.StatusNotImplemented, APIResponse{
		Success: false,
		Error:   "Not implemented yet",
	})
}
