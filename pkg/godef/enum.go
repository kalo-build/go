package godef

import "github.com/kaloseia/clone"

type Enum struct {
	Name    string
	Type    GoTypeDerived
	Entries []EnumEntry
}

func (s Enum) DeepClone() Enum {
	return Enum{
		Name:    s.Name,
		Type:    s.Type.DeepClone(),
		Entries: clone.DeepCloneSlice(s.Entries),
	}
}
