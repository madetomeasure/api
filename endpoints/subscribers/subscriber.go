package subscriber
import("net/http")

func RequestHandler(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusNotImplemented)
}
