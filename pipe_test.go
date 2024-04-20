package conduit

import (
	"context"
	"testing"
	"time"
)

var sleepHandler HandlerFunc[int, int] = func(ctx context.Context, req int) (int, error) {
	time.Sleep(time.Duration(req) * time.Millisecond)
	return req, nil
}

var doubleLayer LayerFunc[int, int] = func(
    next HandlerFunc[int, int],
) HandlerFunc[int, int] {
    return func(ctx context.Context, req int) (int, error) {
        return next(ctx, req * 2)
    }
}

func TestBasicExecute(t *testing.T) {
    req := 5
	res, err := New[int, int]().
		Handle(sleepHandler).
		Execute(context.Background(), req)

	if err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

    if res != req {
        t.Errorf("expected %d, got %d\n", req, res)
    }
}

func TestLayeredExecute(t *testing.T) {
    req := 5
	res, err := New[int, int]().
        Use(doubleLayer).
		Handle(sleepHandler).
		Execute(context.Background(), req)

	if err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

    if res != req * 2 {
        t.Errorf("expected %d, got %d\n", req, res)
    }
}
