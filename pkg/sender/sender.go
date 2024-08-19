package sender

import (
	"context"
	"errors"
	"fmt"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
	"math/rand"
	"strconv"
	"time"
)

type Sender struct{}

func New() *Sender {
	return &Sender{}
}

func (s *Sender) SendEmail(ctx context.Context, recipient string, message string) error {
	// Имитация отправки сообщения
	duration := time.Duration(rand.Int63n(3000)) * time.Millisecond
	time.Sleep(duration)

	// Имитация неуспешной отправки сообщения
	errorProbability := 0.1
	if rand.Float64() < errorProbability {
		return errors.New("internal error")
	}

	fmt.Printf("send message '%s' to '%s'\n", message, recipient)

	return nil
}

func (s *Sender) AsyncSendEmails(ctx context.Context, subscriptions []domain.Subscription) {
	go func() {
		for _, subscription := range subscriptions {
			_ = s.SendEmail(ctx, subscription.UserEmail, strconv.Itoa(int(subscription.HouseID)))
		}
	}()
}
