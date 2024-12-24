package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"movieexample.com/rating/pkg/model"
)

// Gateway defines an HTTP gateway for a rating  service
type Gateway struct {
	addr string
}

// New create a new HTTP gateway for a rating service
func New(addr string) *Gateway {
	return &Gateway{addr: addr}
}

/*
	GetAggregatedRating return the aggreated rating for a record or ErrNotFound

if there are no rating for it
*/
func (g *Gateway) GetAggregatedRating(ctx context.Context, recordID model.RecordId, recordType model.RecordType) (float64, error) {
	req, err := http.NewRequest(http.MethodGet, g.addr+"/rating", nil)
	if err != nil {
		return 0, err
	}
	req = req.WithContext(ctx)
	values := req.URL.Query()
	values.Add("id", string(recordID))
	values.Add("type", fmt.Sprintf("%v", recordType))
	req.URL.RawQuery = values.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, nil
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return 0, fmt.Errorf("non-2xx response : %v", resp)
	}
	var v float64
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return 0, err
	}
	return v, nil
}

// PutRating write a rating
func (g *Gateway) PutRating(ctx context.Context, recordId model.RecordId, recordType model.RecordType, rating *model.
	Rating) error {
	req, err := http.NewRequest(http.MethodPost, g.addr+"/reting", nil)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)
	values := req.URL.Query()
	values.Add("id", string(recordType))
	values.Add("type", fmt.Sprintf("%v", recordType))
	values.Add("userId", string(rating.UserID))
	values.Add("value ", fmt.Sprintf("%v", rating.Value))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("non-2xx response : %v", resp)
	}
	return nil
}
