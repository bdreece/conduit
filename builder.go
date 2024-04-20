package conduit

type Builder[Req, Res any] struct {
    layers []Layer[Req, Res]
}

func (b *Builder[Req, Res]) Use(layer Layer[Req, Res]) *Builder[Req, Res] {
    b.layers = append(b.layers, layer)
    return b
}

func (b *Builder[Req, Res]) Handle(handler Handler[Req, Res]) Pipe[Req, Res] {
    return &pipe[Req, Res]{b.layers, handler}
}

func New[Req, Res any]() *Builder[Req, Res] {
    return &Builder[Req, Res]{
        layers: make([]Layer[Req, Res], 0),
    }
}
