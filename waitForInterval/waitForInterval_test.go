package waitForInterval

import (
	"testing"
	"time"
)

func Test_waitForInterval(t *testing.T) {
	type args struct {
		max      time.Duration
		interval time.Duration
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{max: 5 * time.Second, interval: 1 * time.Second}, "waited"},
		{"test2", args{max: 4 * time.Second, interval: 1 * time.Second}, "waited"},
		{"test3", args{max: 3 * time.Second, interval: 1 * time.Second}, "waited"},
		{"test4", args{max: 2 * time.Second, interval: 1 * time.Second}, "waited"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waitForInterval(tt.args.max, tt.args.interval); got != tt.want {
				t.Errorf("waitForInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test1"},
		{"test2"},
		{"test3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
