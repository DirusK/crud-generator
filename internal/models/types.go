package models

// A Type represents a field type.
type Type string

// List of possible field types.
const (
	TypeBool    Type = "bool"
	TypeByte    Type = "byte"
	TypeEnum    Type = "enum"
	TypeString  Type = "string"
	TypeInt8    Type = "int8"
	TypeInt16   Type = "int16"
	TypeInt32   Type = "int32"
	TypeInt     Type = "int"
	TypeInt64   Type = "int64"
	TypeUint8   Type = "uint8"
	TypeUint16  Type = "uint16"
	TypeUint32  Type = "uint32"
	TypeUint    Type = "uint"
	TypeUint64  Type = "uint64"
	TypeFloat32 Type = "float32"
	TypeFloat64 Type = "float64"
	TypeUUID    Type = "uuid.UUID"
	TypeDecimal Type = "decimal.Decimal"
	TypeTime    Type = "time.Time"
	TypeCoins   Type = "coins.Coins"
)

// types store all types for validation.
var types = map[Type]struct{}{
	TypeBool:    {},
	TypeByte:    {},
	TypeEnum:    {},
	TypeString:  {},
	TypeInt8:    {},
	TypeInt16:   {},
	TypeInt32:   {},
	TypeInt:     {},
	TypeInt64:   {},
	TypeUint8:   {},
	TypeUint16:  {},
	TypeUint32:  {},
	TypeUint:    {},
	TypeUint64:  {},
	TypeFloat32: {},
	TypeFloat64: {},
	TypeUUID:    {},
	TypeDecimal: {},
	TypeTime:    {},
	TypeCoins:   {},
}

// TypesString store all types in string representations.
var TypesString = []string{
	TypeBool.String(),
	TypeByte.String(),
	TypeEnum.String(),
	TypeString.String(),
	TypeInt8.String(),
	TypeInt16.String(),
	TypeInt32.String(),
	TypeInt.String(),
	TypeInt64.String(),
	TypeUint8.String(),
	TypeUint16.String(),
	TypeUint32.String(),
	TypeUint.String(),
	TypeUint64.String(),
	TypeFloat32.String(),
	TypeFloat64.String(),
	TypeUUID.String(),
	TypeDecimal.String(),
	TypeTime.String(),
	TypeCoins.String(),
}

// ToType converts string to type and validates it.
func ToType(t string) (Type, bool) {
	result := Type(t)
	return result, result.Validate()
}

// Validate validates current type.
func (t Type) Validate() bool {
	_, ok := types[t]
	return ok
}

// String returns string representation of type.
func (t Type) String() string {
	return string(t)
}
