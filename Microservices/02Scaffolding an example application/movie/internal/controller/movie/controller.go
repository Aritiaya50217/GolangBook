package movie

import (
	"context"
	"errors"

	metadatamodel "movieexample.com/metadata/pkg/model"
	"movieexample.com/movie/internal/gateway"
	"movieexample.com/movie/pkg/model"
	ratingmodel "movieexample.com/rating/pkg/model"
)

// ErrNotFound is returned when the movie metadata is not found
var ErrNotFound = errors.New("movie metadata not found")

type ratingGeteway interface {
	GetAggregatedRating(ctx context.Context, recordId ratingmodel.RecordId, recordType ratingmodel.RecordType) (float64, error)
	PutRating(ctx context.Context, recordID ratingmodel.RecordId, recordType ratingmodel.RecordType, rating *ratingmodel.Rating) error
}

type metadataGateway interface {
	Get(ctx context.Context, id string) (*metadatamodel.Metadata, error)
}

// controller defines a movie service controller
type Controller struct {
	ratingGeteway   ratingGeteway
	metadataGateway metadataGateway
}

// New create a new movie service controller
func New(ratingGeteway ratingGeteway, metadataGateway metadataGateway) *Controller {
	return &Controller{
		ratingGeteway:   ratingGeteway,
		metadataGateway: metadataGateway,
	}
}

// Get returns the movie details including the  aggregated rating and movie metadata.
func (c *Controller) Get(ctx context.Context, id string) (*model.MovieDetails, error) {
	metadata, err := c.metadataGateway.Get(ctx, id)
	if err != nil && errors.Is(err, gateway.ErrNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}
	details := &model.MovieDetails{Metadata: *metadata}
	rating, err := c.ratingGeteway.GetAggregatedRating(ctx, ratingmodel.RecordId(id), ratingmodel.RecordTypeMovie)
	if err != nil && !errors.Is(err, gateway.ErrNotFound) {
		// just proceed in the case , it's ok not to have ratings yet

	} else if err != nil {
		return nil, err
	} else {
		details.Rating = &rating
	}

	return details, nil
}
