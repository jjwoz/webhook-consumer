package consumer

import "context"

type repository interface {
	ConsumeMessage(ctx context.Context) ()
}
