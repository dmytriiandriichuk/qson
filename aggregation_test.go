package qson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAggregation_Match(t *testing.T) {
	tt := []struct {
		name     string
		query    Query
		expected M
	}{
		{
			name: "default",
			query: Queries(
				In("status", []string{"active", "pending"}),
				Same("user_id", "uuid_user"),
			),
			expected: M{
				"$match": M{
					"user_id": "uuid_user",
					"status":  M{"$in": []string{"active", "pending"}},
				},
			},
		},
	}

	for _, tc := range tt {
		var actual = make(M)
		AGGREGATION.Match(tc.query).Ensure(actual)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestAggregation_Aggregate(t *testing.T) {
	tt := []struct {
		name     string
		stages   []Stage
		expected MS
	}{
		{
			name: "match",
			stages: []Stage{
				AGGREGATION.Match(
					In("status", []string{"active", "pending"}),
					Same("user_id", "uuid_user"),
				),
			},
			expected: MS{
				{
					"$match": M{
						"user_id": "uuid_user",
						"status":  M{"$in": []string{"active", "pending"}},
					},
				},
			},
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.expected, AGGREGATION.Aggregate(tc.stages...))
	}
}
