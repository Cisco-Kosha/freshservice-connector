package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kosha/freshservice-connector/pkg/httpclient"
	"github.com/kosha/freshservice-connector/pkg/models"
)

// getAllTickets godoc
// @Summary Get all tickets
// @Description List all tickets logged in the system
// @Tags tickets
// @Accept  json
// @Produce  json
// @Param page query string false "Page number"
// @Param allPages query boolean false "Collates all pages"
// @Param pageStart query integer false "First page to collate"
// @Param pageEnd query integer false "Last page to collate"
// @Success 200 {object} []models.Ticket
// @Failure      400  {object} string "bad request"
// @Failure      403  {object}  string "permission denied"
// @Failure      404  {object}  string "not found"
// @Failure      500  {object}  string "internal server error"
// @Router /api/v2/tickets [get]
func (a *App) getAllTickets(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	pageNum := r.FormValue("page")
	if pageNum == "" {
		pageNum = "1"
	}

	var tickets []*models.Tickets

	numPages, _ := httpclient.GetAllTickets(a.Cfg.GetFreshServiceURL(), a.Cfg.GetApiKey(), "1", true)
	pageStart, pageEnd, err := getPageRange(r, numPages)

	if err != nil {
		a.Log.Errorf("Error getting page range", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	//get page data
	for i := pageStart; i <= pageEnd; i++ {
		_, t := httpclient.GetAllTickets(a.Cfg.GetFreshServiceURL(), a.Cfg.GetApiKey(), strconv.Itoa(i), false)

		for _, ticket := range t.Tickets {
			status := ticket.Status.(float64)
			ticket.Status = models.Status(status).String()

			priority := ticket.Priority.(float64)
			ticket.Priority = models.Priority(priority).String()

			source := ticket.Source.(float64)
			ticket.Source = models.Source(source).String()
		}

		tickets = append(tickets, t)
	}

	respondWithJSON(w, http.StatusOK, tickets)
}

// getAllTicketsMetadata godoc
// @Summary Get number of pages
// @Description Get page metadata for endpoint
// @Tags tickets
// @Accept  json
// @Produce  json
// @Success 200
// @Router /api/v2/tickets/metadata [get]
func (a *App) getAllTicketsMetadata(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	numPages, _ := httpclient.GetAllTickets(a.Cfg.GetFreshServiceURL(), a.Cfg.GetApiKey(), "1", true)
	respondWithJSON(w, http.StatusOK, numPages)
}

// getSingleTicket godoc
// @Summary Get single ticket
// @Description List single ticket based on the ticket id logged in the system
// @Tags tickets
// @Accept  json
// @Produce  json
// @Param id path string false "Enter ticket id"
// @Success 200 {object} models.Ticket
// @Failure      400  {object} string "bad request"
// @Failure      403  {object}  string "permission denied"
// @Failure      404  {object}  string "not found"
// @Failure      500  {object}  string "internal server error"
// @Router /api/v2/tickets/{id} [get]
func (a *App) getSingleTicket(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	t := httpclient.GetSingleTicket(a.Cfg.GetFreshServiceURL(), id, a.Cfg.GetApiKey())

	status := t.Ticket.Status.(float64)
	t.Ticket.Status = models.Status(status).String()

	priority := t.Ticket.Priority.(float64)
	t.Ticket.Priority = models.Priority(priority).String()

	source := t.Ticket.Source.(float64)
	t.Ticket.Source = models.Source(source).String()

	respondWithJSON(w, http.StatusOK, t)
}

// createTicket godoc
// @Summary Create new ticket
// @Description Create single ticket in the system
// @Tags tickets
// @Accept  json
// @Produce  json
// @Param text body models.Ticket false "Enter ticket properties"
// @Success 200 {object} string
// @Failure      400  {object} string "bad request"
// @Failure      403  {object}  string "permission denied"
// @Failure      404  {object}  string "not found"
// @Failure      500  {object}  string "internal server error"
// @Router /api/v2/tickets [post]
func (a *App) createTicket(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	var t models.Ticket
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		a.Log.Errorf("Error parsing json payload", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	_, err := httpclient.CreateTicket(a.Cfg.GetFreshServiceURL(), a.Cfg.GetApiKey(), &t)
	if err != nil {
		a.Log.Errorf("Error creating a ticket", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, "OK")
}

// deleteTicket godoc
// @Summary Delete single ticket
// @Description Delete single ticket based on the ticket id logged in the system
// @Tags tickets
// @Accept  json
// @Produce  json
// @Param id path string true "Enter ticket id"
// @Success 200 {object} object
// @Failure      400  {object} string "bad request"
// @Failure      403  {object}  string "permission denied"
// @Failure      404  {object}  string "not found"
// @Failure      500  {object}  string "internal server error"
// @Router /api/v2/tickets/{id} [delete]
func (a *App) deleteTicket(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	t, err := httpclient.RemoveTicket(a.Cfg.GetFreshServiceURL(), id, a.Cfg.GetApiKey())
	if err != nil {
		a.Log.Errorf("Error deleting a ticket", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success", "resp": t})
}
