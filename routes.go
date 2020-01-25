package routes

import (
	"regexp"
)

type HandlerFunc func(string)

type Router struct {
	routes []*entry
}

type entry struct {
	route   *Route
	handler HandlerFunc
}

func NewRouter() *Router {
	return &Router{}
}

type Route struct {
	*regexp.Regexp
}

func (router *Router) match(input string) HandlerFunc {
	for _, e := range router.routes {
		if match := e.route.Match(input); match != nil {
			return e.handler
		}
	}
	return nil
}

func (router *Router) AddRoute(pattern string, fn HandlerFunc) {
	router.routes = append(router.routes, &entry{route: NewRoute(pattern), handler: fn})
}

func NewRoute(pattern string) *Route {
	return &Route{regexp.MustCompile(pattern)}
}

type RouteMatch struct {
	Args   []string
	Kwargs map[string]string
}

func (r *Route) Match(target string) *RouteMatch {
	submatches := r.FindStringSubmatch(target)
	if submatches == nil {
		return nil
	}

	if len(submatches) == 1 {
		return new(RouteMatch)
	}

	m := new(RouteMatch)
	submatches = submatches[1:]
	for i, name := range r.SubexpNames()[1:] {
		if name == "" {
			m.Args = append(m.Args, submatches[i])
		} else {
			if m.Kwargs == nil {
				m.Kwargs = make(map[string]string)
			}
			m.Kwargs[name] = submatches[i]
		}
	}
	return m
}

func (router *Router) Exec(input string) HandlerFunc {
	m := router.match(input)
	return m
}
