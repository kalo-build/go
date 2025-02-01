package godef

import "github.com/kaloseia/clone"

type Enum struct {
	Type    GoTypeDerived
	Entries []EnumEntry
}

func (s Enum) DeepClone() Enum {
	return Enum{
		Type:    s.Type.DeepClone(),
		Entries: clone.DeepCloneSlice(s.Entries),
	}
}
