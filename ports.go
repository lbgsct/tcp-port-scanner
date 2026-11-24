package portscan

import "sort"

type PortSet struct{ entries []portEntry }
type portEntry struct{ first, last int }

func Ports(ports ...int) PortSet {
	entries := make([]portEntry, len(ports))
	for i, port := range ports {
		entries[i] = portEntry{port, port}
	}
	return PortSet{entries}
}

func Range(first, last int) PortSet { return PortSet{[]portEntry{{first, last}}} }

func Combine(sets ...PortSet) PortSet {
	var entries []portEntry
	for _, set := range sets {
		entries = append(entries, set.entries...)
	}
	return PortSet{entries}
}

func (p PortSet) expand() ([]uint16, error) {
	if len(p.entries) == 0 {
		return nil, ErrNoPorts
	}
	seen := make(map[int]struct{})
	for _, entry := range p.entries {
		if entry.first > entry.last {
			return nil, ErrInvalidRange
		}
		if entry.first < 1 || entry.last > 65535 {
			return nil, ErrInvalidPort
		}
		for port := entry.first; port <= entry.last; port++ {
			seen[port] = struct{}{}
		}
	}
	ports := make([]uint16, 0, len(seen))
	for port := range seen {
		ports = append(ports, uint16(port))
	}
	sort.Slice(ports, func(i, j int) bool { return ports[i] < ports[j] })
	return ports, nil
}
