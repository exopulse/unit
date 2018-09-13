package unit

import (
	"testing"
)

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

func TestSize_String(t *testing.T) {
	tests := []struct {
		name string
		s    Size
		want string
	}{
		{"zero", Size(0), "0"},
		{"bytes", Size(50 * B), "50B"},
		{"kilobytes", Size(1 * KB), "1KB"},
		{"megabytes", Size(2 * MB), "2MB"},
		{"gigabytes", Size(3 * GB), "3GB"},
		{"terabytes", Size(4 * TB), "4TB"},
		{"petabytes", Size(5 * PB), "5PB"},
		{"fractional", Size(1600 * B), "1.6KB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("Size.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSize_ParseSize(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    Size
		wantErr bool
	}{
		{"default", "0", Size(0), false},
		{"bytes-default", "10", Size(10), false},
		{"bytes", "10B", Size(10), false},
		{"kilobytes", "10K", Size(10 * KB), false},
		{"megabytes", "20MB", Size(20 * KB), false},
		{"gigabytes", "30 G", Size(30 * GB), false},
		{"terabytes", "40 TB", Size(40 * TB), false},
		{"petabytes", "50  PB", Size(50 * PB), false},
		{"negative", "-20M", Size(0), true},
		{"invalid", "inv M", Size(0), true},
		{"fractional", "1.5 M", Size(1536), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseSize(tt.str)

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("ParseSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSize_ConversionInt(t *testing.T) {
	s := Size(1024)

	if int(s) != 1024 {
		t.Error("conversion failed")
	}
}
