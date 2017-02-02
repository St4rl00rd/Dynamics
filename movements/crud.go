package movements

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/st4rl00rd/dynamics/context"
)

// Index function handles the request to a Html View
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := context.New()
	IndexCtl(ctx, w)
}

// Show function handles the request to a Html View
func Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := context.New()
	paramID := ps.ByName("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		NotValidMovementID(w)
	} else {
		ShowCtl(ctx, w, id)
	}
}
