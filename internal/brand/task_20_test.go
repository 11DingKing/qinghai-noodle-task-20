package brand

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQinghaiBrandTask20(t *testing.T) {
	now := time.Now()
	s := NewService(NewRegistry(), func() time.Time { return now })
	r := Recall{AffectedStores: []string{"s1", "s2"}, Acknowledged: []string{"s1", "s2"}, CreatedAt: now.Add(-time.Hour)}
	_, err := s.CloseRecall(context.Background(), r)
	require.NoError(t, err)
}
