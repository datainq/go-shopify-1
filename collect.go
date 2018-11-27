package goshopify

import (
	"fmt"
	"time"
)

const collectsBasePath = "admin/collects"

// CollectService is an interface for interfacing with the collect endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/products/collect
type CollectService interface {
	Create(Collect) (*Collect, error)
	List(interface{}) ([]Collect, error)
	Count(interface{}) (int, error)
}

// CollectServiceOp handles communication with the collect related methods of
// the Shopify API.
type CollectServiceOp struct {
	client *Client
}

// Collect represents a Shopify collect
type Collect struct {
	ID           int        `json:"id,omitempty"`
	CollectionID int        `json:"collection_id,omitempty"`
	ProductID    int        `json:"product_id,omitempty"`
	Featured     bool       `json:"featured,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	Position     int        `json:"position,omitempty"`
	SortValue    string     `json:"sort_value,omitempty"`
}

// Represents the result from the collects/X.json endpoint
type CollectResource struct {
	Collect *Collect `json:"collect"`
}

// Represents the result from the collects.json endpoint
type CollectsResource struct {
	Collects []Collect `json:"collects"`
}

func (s *CollectServiceOp) Create(collect Collect) (*Collect, error) {
	path := fmt.Sprintf("%s.json", collectsBasePath)
	wrappedData := CollectResource{Collect: &collect}
	resource := new(CollectResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Collect, err
}

// List collects
func (s *CollectServiceOp) List(options interface{}) ([]Collect, error) {
	path := fmt.Sprintf("%s.json", collectsBasePath)
	resource := new(CollectsResource)
	err := s.client.Get(path, resource, options)
	return resource.Collects, err
}

// Count collects
func (s *CollectServiceOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", collectsBasePath)
	return s.client.Count(path, options)
}
