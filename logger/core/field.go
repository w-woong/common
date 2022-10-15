package core

type FieldType uint8

const (
	StringType FieldType = iota
	BytesType
	Int32Type
	Int64Type
	UnknownType
)

type Field struct {
	Type        FieldType
	Name        string
	ValueString string
	ValueBytes  []byte
	ValueInt32  int32
	ValueInt64  int64
	ValueAny    interface{}
}

// type Fields []Field

func WithStringField(fieldName string, fieldValue string) Field {
	return Field{
		Type:        StringType,
		Name:        fieldName,
		ValueString: fieldValue,
	}
}

func WithBytesField(fieldName string, fieldValue []byte) Field {
	return Field{
		Type:       BytesType,
		Name:       fieldName,
		ValueBytes: fieldValue,
	}
}

func WithInt32Field(fieldName string, fieldValue int32) Field {
	return Field{
		Type:       Int32Type,
		Name:       fieldName,
		ValueInt32: fieldValue,
	}
}

func WithInt64Field(fieldName string, fieldValue int64) Field {
	return Field{
		Type:       Int64Type,
		Name:       fieldName,
		ValueInt64: fieldValue,
	}
}
