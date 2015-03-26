package subscriber

import (
	//  "github.com/madetomeasure/api/lib"
	"net/http"
	"time"
)

// Subscriber is the basis for every other model. This effectively
// Is what emails are and holds onto created at and updated at timestamps
type Subscriber struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string
}

// IndexHandler handles the index case of subscribers returning an array of users given a specific query.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//  subscriber = Subscriber{}
	//  db.DB().First(&subscriber)
	//  render.Perform(&subscriber)
	w.WriteHeader(http.StatusNotImplemented)
}

// CreateHandler handles the creation context of a subscriber creating a new user in batch or single mode
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// UpdateHandler updates an existing subscriber
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// DestroyHandler will destroy an existing subscriber
func DestroyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
