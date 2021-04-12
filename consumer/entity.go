package consumer

import "github.com/stripe/stripe-go"

type Message struct {
	stripe.Event
}
