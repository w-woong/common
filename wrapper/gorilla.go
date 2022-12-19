package wrapper

import (
	"github.com/gorilla/mux"
)

// CollectRoutes collect api endpoints' path and its methods
func CollectRoutes(r *mux.Router) (map[string][]string, error) {
	m := make(map[string][]string)
	var resErr error
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err1 := route.GetPathTemplate()
		if err1 != nil {
			resErr = err1
		}
		methods, err2 := route.GetMethods()
		if err2 != nil {
			resErr = err2
		}

		if err1 == nil && err2 == nil {
			m[path] = methods
		}

		return nil
	})

	return m, resErr
}
