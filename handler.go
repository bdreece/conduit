package conduit

import "context"

type Handler[Req, Res any] interface {
    Handle(context.Context, Req) (Res, error)
}

type HandlerFunc[Req, Res any] func(context.Context, Req) (Res, error)

func (fn HandlerFunc[Req, Res]) Handle(
    ctx context.Context, req Req,
) (Res, error) {
    return fn(ctx, req)
}

var _ Handler[any, any] = HandlerFunc[any, any](nil)
