package tele

// MiddlewareFunc represents a middleware processing function,
// which get called before the endpoint group or specific handler.
type MiddlewareFunc func(HandlerFunc) HandlerFunc

// Group is a separated group of handlers, united by the general middleware.
type Group struct {
	b          *Handle
	middleware []MiddlewareFunc
}

// Use adds middleware to the chain.
func (g *Group) Use(middleware ...MiddlewareFunc) {
	g.middleware = append(g.middleware, middleware...)
}

// Handle adds endpoint handler to the bot, combining group's middleware
// with the optional given middleware.
func (g *Group) Handle(endpoint interface{}, h HandlerFunc, m ...MiddlewareFunc) {
	g.b.Handle(endpoint, h, append(g.middleware, m...)...)
}
