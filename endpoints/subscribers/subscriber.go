package subscriber

import(
//  "github.com/madetomeasure/api/lib"
  "net/http"
   _ "github.com/jinzhu/gorm"
    _ "github.com/lib/pq"
  "time"
)

type Subscriber struct {
  Id int64
  CreatedAt time.Time
  UpdatedAt time.Time
  Name      string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
//  subscriber = Subscriber{}
//  db.DB().First(&subscriber)
//  render.Perform(&subscriber)
  w.WriteHeader(http.StatusNotImplemented)
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

