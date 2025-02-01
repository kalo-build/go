package godef

import "strings"

// GoTypeDerived represents a defined type that is derived from another type (ie. `type Nationality string`).
type GoTypeDerived struct {
	// PackagePath is the fully qualified import path
	PackagePath string

	// Name is the name of the type (e.g., "Nationality")
	Name string

	// BaseType is the underlying type (e.g., "string", "int", ...)
	BaseType GoType
}

func (t GoTypeDerived) IsPrimitive() bool {
	return t.BaseType.IsPrimitive()
}

func (t GoTypeDerived) IsMap() bool {
	return t.BaseType.IsMap()
}

func (t GoTypeDerived) IsArray() bool {
	return t.BaseType.IsArray()
}

func (t GoTypeDerived) IsStruct() bool {
	return t.BaseType.IsStruct()
}

func (t GoTypeDerived) IsInterface() bool {
	return t.BaseType.IsInterface()
}

func (t GoTypeDerived) IsPointer() bool {
	return t.BaseType.IsPointer()
}

func (t GoTypeDerived) GetImports() []string {
	if t.PackagePath == "" {
		return nil
	}
	return []string{t.PackagePath}
}

func (t GoTypeDerived) GetSyntax() string {
	packageName := t.getPackageName()
	if packageName == "" {
		return t.GetSyntaxLocal()
	}
	return packageName + "." + t.Name
}

func (t GoTypeDerived) GetSyntaxLocal() string {
	return t.Name
}

func (t GoTypeDerived) DeepClone() GoTypeDerived {
	return GoTypeDerived{
		PackagePath: t.PackagePath,
		Name:        t.Name,
		BaseType:    DeepCloneGoType(t.BaseType),
	}
}

func (t GoTypeDerived) getPackageName() string {
	if t.PackagePath == "" {
		return ""
	}
	packagePieces := strings.Split(t.PackagePath, "/")
	return packagePieces[len(packagePieces)-1]
}
