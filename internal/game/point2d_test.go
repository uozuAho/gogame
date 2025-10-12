package game

import (
	"fmt"
	"testing"
)

func TestUnitVec_len_is_1(t *testing.T) {
	tests := []struct {
		pt Point2D
	}{
		{Point2D{X: 3.3, Y: 499}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%f,%f", tt.pt.X, tt.pt.Y), func(t *testing.T) {
			if tt.pt.UnitVec().Len() != 1.0 {
				t.Errorf("want 1.0, got %v", tt.pt.UnitVec().Len())
			}
		})
	}
}
