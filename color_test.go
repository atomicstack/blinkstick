package blinkstick

import (
	"image/color"
	"testing"
)

func TestApplyBrightness(t *testing.T) {
	cases := []struct {
		name       string
		in         color.Color
		brightness int
		want       color.RGBA
	}{
		{"zero brightness is black", color.RGBA{R: 255, G: 255, B: 255}, 0, color.RGBA{R: 0, G: 0, B: 0}},
		{"full brightness is unchanged", color.RGBA{R: 255, G: 128, B: 0}, 100, color.RGBA{R: 255, G: 128, B: 0}},
		{"half brightness halves channels", color.RGBA{R: 200, G: 100, B: 50}, 50, color.RGBA{R: 100, G: 50, B: 25}},
		{"over-100 is clamped to full", color.RGBA{R: 255, G: 255, B: 255}, 250, color.RGBA{R: 255, G: 255, B: 255}},
		{"negative is clamped to zero", color.RGBA{R: 255, G: 255, B: 255}, -50, color.RGBA{R: 0, G: 0, B: 0}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := applyBrightness(tc.in, tc.brightness).(color.RGBA)
			if !ok {
				t.Fatalf("expected RGBA, got %T", got)
			}
			if got.R != tc.want.R || got.G != tc.want.G || got.B != tc.want.B {
				t.Errorf("got %+v, want %+v", got, tc.want)
			}
		})
	}
}

func TestColorList_IsNonEmpty(t *testing.T) {
	if len(ColorList()) == 0 {
		t.Fatal("expected ColorList to be non-empty")
	}
}
