package filter

import "strings"

type flag int

const (
	FlagEq flag = 1 << iota
	FlagNotEq
	FlagIn
)

type Filter struct {
	value string
	flag  flag
}

// Build ...
func Build(params map[string][]string, query string, filters map[string]Filter) string {
	for filter := range filters {
		for param, values := range params {
			if filter == param {
				if ok := strings.Contains(strings.ToLower(query), "where"); ok {
					query += " AND "
				} else {
					query += " WHERE "
				}

				switch filters[filter].flag {
				case FlagEq:
					query += strings.Replace(filters[filter].value, "?", values[0], 1)

				case FlagNotEq:
					query += "NOT " + strings.Replace(filters[filter].value, "?", values[0], 1)

				case FlagIn:
					var v string
					for idx := range values {
						v += "'" + values[idx] + "',"
					}

					query += filters[filter].value + " IN (" + strings.TrimRight(v, ",") + ")"
				}
			}
		}
	}
	return query
}

// Cond ...
func Cond(value string, flag flag) Filter {
	return Filter{value: value, flag: flag}
}
