package conduit

import "context"

// A Handler represents the endpoint of a pipeline's execution.
type Handler[Req, Res any] interface {
    // Handle handles the given request and returns the appropriate response.
    Handle(context.Context, Req) (Res, error)
}

// HandlerFunc provides a function alternative to the Handler interface.
type HandlerFunc[Req, Res any] func(context.Context, Req) (Res, error)

// Handle implements Handler.
func (fn HandlerFunc[Req, Res]) Handle(
    ctx context.Context, req Req,
) (Res, error) {
    return fn(ctx, req)
}
