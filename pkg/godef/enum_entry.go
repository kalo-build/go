package godef

type EnumEntry struct {
	Name  string
	Value any
}

func (e EnumEntry) DeepClone() EnumEntry {
	return EnumEntry{
		Name:  e.Name,
		Value: e.Value,
	}
}
