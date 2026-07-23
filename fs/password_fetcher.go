package fs

import (
	"fmt"
	"strings"
)

// PassFetcherMap is a map of matchers to password fetcher commands.
type PassFetcherMap struct {
	Star   SpaceSepList
	Param  map[string]SpaceSepList
	Remote map[string]map[string]SpaceSepList
}

func (m *PassFetcherMap) Set(s string) (err error) {
	list := strings.Split(s, ",")
	if len(list)%2 != 0 {
		return fmt.Errorf("password fetcher should consist of comma-separated pairs")
	}
	for len(list) != 0 {
		matcher := list[0]
		fetcher := SpaceSepList{}
		fetcher.Set(list[1])
		list = list[2:]
		matcherParts := strings.Split(matcher, ".")
		if len(matcherParts) > 2 {
			return fmt.Errorf("matcher should contain at most one dot")
		} else if len(matcherParts) == 0 {
			return fmt.Errorf("matcher should not be empty")
		} else if len(matcherParts) == 2 {
			remote := matcherParts[0]
			param := matcherParts[1]
			if m.Remote == nil {
				m.Remote = map[string]map[string]SpaceSepList{}
			}
			if m.Remote[remote] == nil {
				m.Remote[remote] = map[string]SpaceSepList{}
			}
			m.Remote[remote][param] = fetcher
		} else if matcher == "*" {
			m.Star = fetcher
		} else {
			if m.Param == nil {
				m.Param = map[string]SpaceSepList{}
			}
			m.Param[matcher] = fetcher
		}
	}
	return nil
}
