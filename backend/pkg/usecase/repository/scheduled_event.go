package repository

import (
	"context"

	"github.com/ayumi-sk2720/animal_healthcare_app/backend/pkg/domain/model"
)

type ScheduledEvent interface {
	Store(ctx context.Context, event *model.ScheduledEvent) error
}
