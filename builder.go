package conduit

// A Builder provides pipeline orchestration methods used to construct a Pipe.
type Builder[Req, Res any] struct {
    layers []Layer[Req, Res]
}

// Use appends the given layer to the pipeline. Note that layers are invoked
// in the order they are declared.
func (b *Builder[Req, Res]) Use(layer Layer[Req, Res]) *Builder[Req, Res] {
    b.layers = append(b.layers, layer)
    return b
}

// Handle provides a given handler to the Builder, finalizing the construction
// of the Pipe. This handler will be invoked after all of the preceding layers.
func (b *Builder[Req, Res]) Handle(handler Handler[Req, Res]) Pipe[Req, Res] {
    return &pipe[Req, Res]{b.layers, handler}
}

// New creates a new Builder for the specified request and response types.
func New[Req, Res any]() *Builder[Req, Res] {
    return &Builder[Req, Res]{
        layers: make([]Layer[Req, Res], 0),
    }
}
