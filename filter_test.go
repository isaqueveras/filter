package filter_test

import (
	"testing"

	"github.com/isaqueveras/filter"
)

func TestFilters(t *testing.T) {
	query := "SELECT * FROM events"

	t.Run("TestFlagIn", func(t *testing.T) {
		filters := map[string][]string{
			"id": {"1", "2", "3"},
		}

		query := filter.Build(filters, query, map[string]filter.Filter{
			"id": filter.Cond("id", filter.FlagIn),
		})

		if query != "SELECT * FROM events WHERE id IN ('1','2','3')" {
			t.Fail()
		}
	})

	t.Run("TestFlagEq", func(t *testing.T) {
		filters := map[string][]string{
			"name": {"Steve"},
		}

		query := filter.Build(filters, query, map[string]filter.Filter{
			"name": filter.Cond("name = '?'::VARCHAR", filter.FlagEq),
		})

		if query != "SELECT * FROM events WHERE name = 'Steve'::VARCHAR" {
			t.Fail()
		}
	})

	t.Run("TestFlagNotEq", func(t *testing.T) {
		filters := map[string][]string{
			"not_id": {"1"},
		}

		query := filter.Build(filters, query, map[string]filter.Filter{
			"not_id": filter.Cond("id = ?", filter.FlagNotEq),
		})

		if query != "SELECT * FROM events WHERE NOT id = 1" {
			t.Fail()
		}
	})

	t.Run("TestFlagEq_Filter_Date", func(t *testing.T) {
		filters := map[string][]string{
			"date_gte": {"2023-03-01T03:00:00.000Z"},
			"date_lte": {"2023-05-01T03:00:00.000Z"},
		}

		query := filter.Build(filters, query, map[string]filter.Filter{
			"date_gte": filter.Cond("created_at::TIMESTAMPTZ >= '?'::TIMESTAMPTZ", filter.FlagEq),
			"date_lte": filter.Cond("created_at::TIMESTAMPTZ <= '?'::TIMESTAMPTZ", filter.FlagEq),
		})

		if query != "SELECT * FROM events WHERE created_at::TIMESTAMPTZ >= '2023-03-01T03:00:00.000Z'::TIMESTAMPTZ AND created_at::TIMESTAMPTZ <= '2023-05-01T03:00:00.000Z'::TIMESTAMPTZ" {
			t.Fail()
		}
	})
}
