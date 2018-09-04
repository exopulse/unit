package unit

import "testing"

func TestSize_Humanize(t *testing.T) {
	tests := []struct {
		name string
		s    Size
		want string
	}{
		{"zero", Size(0), "0"},
		{"bytes", Size(50 * B), "50 B"},
		{"kilobytes", Size(1 * KB), "1 KB"},
		{"megabytes", Size(2 * MB), "2 MB"},
		{"gigabytes", Size(3 * GB), "3 GB"},
		{"terabytes", Size(4 * TB), "4 TB"},
		{"petabytes", Size(5 * PB), "5 PB"},
		{"exabytes", Size(6 * EB), "6 EB"},
		{"fractional", Size(1600 * B), "1.6 KB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Humanize(); got != tt.want {
				t.Errorf("Size.Humanize() = %v, want %v", got, tt.want)
			}
		})
	}
}
