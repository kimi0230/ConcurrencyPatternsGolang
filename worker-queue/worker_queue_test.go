package workerqueue

import "testing"

func TestWorkQueue(t *testing.T) {
	tests := []struct {
		jobs  []string
		limit int
	}{
		{
			jobs:  []string{"a", "b", "c", "d", "e", "f"},
			limit: 5,
		},
	}

	for _, tt := range tests {
		t.Run("work queue", func(t *testing.T) {
			WorkQueue(tt.jobs, tt.limit)
		})
	}
}
