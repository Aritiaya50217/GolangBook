package store

import (
	coffeeco "domain-driven-design/part02real_world_domain_driven_design/chapter05applying_domain_driven_design_to_a_monolithic_application/internal"

	"github.com/google/uuid"
)

type Store struct {
	ID              uuid.UUID
	Location        string
	ProductsForSale []coffeeco.Product
}
