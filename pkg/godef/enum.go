package godef

import "github.com/kaloseia/clone"

type Enum struct {
	Package Package
	Name    string
	Type    GoTypeDerived
	Entries []EnumEntry
}

func (e Enum) DeepClone() Enum {
	return Enum{
		Package: e.Package,
		Name:    e.Name,
		Type:    e.Type.DeepClone(),
		Entries: clone.DeepCloneSlice(e.Entries),
	}
}
