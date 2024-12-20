package rating

import (
	"context"
	"errors"

	"movieexample.com/rating/internal/repository"
	"movieexample.com/rating/pkg/model"
)

// ErrNotFound is returned when no rating are found for a record
var ErrNotFound = errors.New("ratings not found for a record")

type ratingRepository interface {
	Get(ctx context.Context, recordId model.RecordId, recordType model.RecordType) ([]model.Rating, error)
	Put(ctx context.Context, recordID model.RecordId, recordType model.RecordType, rating *model.Rating) error
}

// Controller defines a rating service controller
type Controller struct {
	repo ratingRepository
}

// New creates a rating service controller
func New(repo ratingRepository) *Controller {
	return &Controller{repo: repo}
}

// GetAggregatedRating returns the aggregated rating for a record or ErrNotFound if there are ratings for it
func (c *Controller) GetAggregatedRating(ctx context.Context, recordId model.RecordId, recordType model.RecordType) (float64, error) {
	ratings, err := c.repo.Get(ctx, recordId, recordType)
	if err != nil && err == repository.ErrNotFound {
		return 0, ErrNotFound
	} else if err != nil {
		return 0, err
	}
	sum := float64(0)
	for _, r := range ratings {
		sum += float64(r.Value)
	}
	return sum / float64(len(ratings)), nil
}

// PutRating writes a rating for a given record.
func (c *Controller) PutRating(ctx context.Context, recordId model.RecordId, recordType model.RecordType, rating *model.Rating) error {

	return c.repo.Put(ctx, recordId, recordType, rating)
}
