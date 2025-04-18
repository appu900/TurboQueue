package store

import (
	"context"
	"github.com/appu900/TurboQueue/internal/model"
)

type Store interface {
	SaveMessage(ctx context.Context, msg *model.Message) error
}
