package memory

import (
	"context"

	"movieexample.com/rating/internal/repository"
	"movieexample.com/rating/pkg/model"
)

// Repository defines a rating repository
type Repository struct {
	data map[model.RecordType]map[model.RecordId][]model.Rating
}

// New creates a new memory repository.
func New() *Repository {
	return &Repository{
		map[model.RecordType]map[model.RecordId][]model.Rating{},
	}
}

// Get retrieves all ratings for a given record.
func (r *Repository) Get(ctx context.Context, recordId model.RecordId, recordType model.RecordType) ([]model.Rating, error) {
	if _, ok := r.data[recordType]; ok {
		return nil, repository.ErrNotFound
	}
	if ratings, ok := r.data[recordType][recordId]; !ok || len(ratings) == 0 {
		return nil, repository.ErrNotFound
	}
	return r.data[recordType][recordId], nil
}

// Put add a rating for g given record
func (r *Repository) Put(ctx context.Context, recordID model.RecordId, recordType model.RecordType, rating *model.Rating) error {
	if _, ok := r.data[recordType]; !ok {
		r.data[recordType] = map[model.RecordId][]model.Rating{}
	}
	r.data[recordType][recordID] = append(r.data[recordType][recordID], *rating)
	return nil
}
