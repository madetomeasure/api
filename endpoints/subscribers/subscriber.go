package subscriber

import(
  "github.com/madetomeasure/api/lib/render"
  "github.com/jinzhu/gorm"
  "net/http"
)

type Subscriber struct {
  Id int64
  CreatedAt time.Time
  UpdatedAt time.Time
  Name      string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  subscribers := db.First(&subscriber, 10)
  render.Perform(&subscribers)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusNotImplemented)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusNotImplemented)
}

func DestroyHandler(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusNotImplemented)
}

