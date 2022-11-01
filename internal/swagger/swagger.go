package swagger

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"rests.com/internal/taskstore"
)

var store = taskstore.New()

func TaskIdDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = store.DeleteTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}
