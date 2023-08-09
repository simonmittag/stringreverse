package main

import "testing"

func TestStringReverse(t *testing.T) {
	var tests = []struct {
		Sample   string
		Reversed string
	}{
		{"anna", "anna"},
		{"desserts", "stressed"},
		{"mary had a little lamb", "bmal elttil a dah yram"},
	}

	for _, tt := range tests {
		t.Run(tt.Sample, func(t *testing.T) {
			want := tt.Reversed
			got := FirstReverse(tt.Sample)
			if want != got {
				t.Errorf("want %v got %v", want, got)
			}
		})
	}
}
