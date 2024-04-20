package conduit

// A Layer represents a middleware in the pipeline chain.
type Layer[Req, Res any] interface {
    // Layer wraps the next handler.
    Process(next HandlerFunc[Req, Res]) HandlerFunc[Req, Res]
}

// LayerFunc provides a function alternative to the Layer interface.
type LayerFunc[Req, Res any] func(next HandlerFunc[Req, Res]) HandlerFunc[Req, Res]

// Process implements Layer.
func (fn LayerFunc[Req, Res]) Process(next HandlerFunc[Req, Res]) HandlerFunc[Req, Res] {
    return fn(next)
}
