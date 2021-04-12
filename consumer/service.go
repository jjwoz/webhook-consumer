package consumer

import "context"

type ConsumerService interface {
	ConsumeMessage(ctx context.Context, message *Message) error
}
