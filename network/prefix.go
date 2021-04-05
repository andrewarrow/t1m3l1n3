package network

import "strings"

func UniverseSearchByPrefix(prefix string) *Universe {

	all := map[string]int{}
	for k, _ := range universes {
		tokens := strings.Split(k, "-")
		p := tokens[1]
		all[p]++
	}
	if all[prefix] > 1 {
		return nil
	}
	for k, v := range universes {
		tokens := strings.Split(k, "-")
		p := tokens[1]
		if p == prefix {
			return v
		}
	}
	return nil
}
