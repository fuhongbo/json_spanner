package wildcard

import (
	"strings"
)

type MatchResult struct {
	Result []string
}

func Match(topic, wildcard string) *MatchResult {
	if topic == wildcard {
		return &MatchResult{Result: []string{}}
	}
	if wildcard == "#" {
		return &MatchResult{Result: []string{topic}}
	}
	t := strings.Split(topic, "/")
	w := strings.Split(wildcard, "/")
	i := 0
	results := make([]string, 0)
	for lt := len(t); i < lt; i++ {
		if len(w) <= i {
			return nil
		} else if w[i] == "+" {
			results = append(results, t[i])
		} else if w[i] == "#" {
			results = append(results, strings.Join(t[i:], "/"))
			return &MatchResult{Result: results}
		} else if w[i] != t[i] {
			return nil
		}
	}
	if len(w) > i && w[i] == "#" {
		i++
	}
	if i == len(w) {
		return &MatchResult{Result: results}
	}
	return nil
}
