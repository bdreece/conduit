package conduit

type Layer[Req, Res any] interface {
    Process(next HandlerFunc[Req, Res]) HandlerFunc[Req, Res]
}

type LayerFunc[Req, Res any] func(next HandlerFunc[Req, Res]) HandlerFunc[Req, Res]

func (fn LayerFunc[Req, Res]) Process(next HandlerFunc[Req, Res]) HandlerFunc[Req, Res] {
    return fn(next)
}
