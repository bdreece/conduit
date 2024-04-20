package conduit

import "context"

// A Pipe represents a constructed command pipeline
type Pipe[Req, Res any] interface {
    // Execute invokes the command pipeline
    Execute(context.Context, Req) (Res, error)
}

type pipe[Req, Res any] struct {
    layers []Layer[Req, Res]
    handler Handler[Req, Res]
}

func (p *pipe[Req, Res]) Execute(ctx context.Context, req Req) (Res, error) {
    handler := p.handler
    for i := len(p.layers) - 1; i >= 0; i-- {
        handler = p.layers[i].Process(handler.Handle)
    }

    return handler.Handle(ctx, req)
}
