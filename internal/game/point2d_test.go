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
		{Point2D{X: 1.0, Y: 0}},
		{Point2D{X: 0, Y: 1}},
		{Point2D{X: .1, Y: .2}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%f,%f", tt.pt.X, tt.pt.Y), func(t *testing.T) {
			if tt.pt.UnitVec().Len() != 1.0 {
				t.Errorf("want 1.0, got %v", tt.pt.UnitVec().Len())
			}
		})
	}
}

func TestPoint2D_Multiply_ScalesLength(t *testing.T) {
	tests := []struct {
		pt      Point2D
		factor  float64
		wantLen float64
	}{
		{Point2D{X: 3, Y: 4}, 2, 10},    // length 5 * 2 = 10
		{Point2D{X: 1, Y: 0}, 3, 3},     // length 1 * 3 = 3
		{Point2D{X: 0, Y: 1}, 0.5, 0.5}, // length 1 * 0.5 = 0.5
		{Point2D{X: 0, Y: 0}, 5, 0},     // length 0 * 5 = 0
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v*%v", tt.pt, tt.factor), func(t *testing.T) {
			newpt := tt.pt.Copy()
			newpt.Multiply(tt.factor)
			gotLen := newpt.Len()
			if (gotLen-tt.wantLen) > 1e-9 || (tt.wantLen-gotLen) > 1e-9 {
				t.Errorf("Multiply(%v, %v) length = %v, want %v", tt.pt, tt.factor, gotLen, tt.wantLen)
			}
		})
	}
}
