package valueobject_test

import (
	valueobject "domain-driven-design/part01introduction_to_domain_driven_design/03entities_value_objects_and_aggregates/02value_object"
	"testing"
)

func Test_Point(t *testing.T) {
	a := valueobject.NewPoints(1, 1)
	b := valueobject.NewPoints(1, 1)

	if a != b {
		t.Fatal("a and b were equal")
	}
}
