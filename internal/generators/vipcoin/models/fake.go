package models

import (
	"fmt"
	"strings"

	"crud-generator-gui/internal/generators/vipcoin/helpers"
	"crud-generator-gui/internal/models"
)

var (
	fakeEnum = `%s(gofakeit.RandomString(
			[]string{
				%s
			},
		))`

	fakeOption = `
// FakeWith{{.FieldNameCamel}} setups custom {{.FieldNameCamel}} to fake {{.NameCamel}}.
func FakeWith{{.FieldNameCamel}}({{.FieldNameLowerCamel}} {{.FieldType}}) FakeOption {
	return func({{.NameLowerCamel}} *{{.NameCamel}}) {
		{{.NameLowerCamel}}.{{.FieldNameCamel}} = {{.FieldNameLowerCamel}}
	}
}
`
)

func fakeValue(field Field) string {
	switch field.Type {
	case models.TypeEnum:
		return fmt.Sprintf(fakeEnum, field.NameCamel(true), field.EnumStringArray())
	default:
		return defaultFakeValue[field.Type]
	}
}

var defaultFakeValue = map[models.Type]string{
	models.TypeString:  "gofakeit.SentenceSimple()",
	models.TypeBool:    "gofakeit.Bool()",
	models.TypeByte:    "gofakeit.Uint8()",
	models.TypeEnum:    "",
	models.TypeInt8:    "gofakeit.Int8()",
	models.TypeInt16:   "gofakeit.Int16()",
	models.TypeInt32:   "gofakeit.Int32()",
	models.TypeInt:     "int(gofakeit.Int32())",
	models.TypeInt64:   "gofakeit.Int64()",
	models.TypeUint8:   "gofakeit.Uint8()",
	models.TypeUint16:  "gofakeit.Uint16()",
	models.TypeUint32:  "gofakeit.Uint32()",
	models.TypeUint:    "uint(gofakeit.Uint32())",
	models.TypeUint64:  "gofakeit.Uint64()",
	models.TypeFloat32: "gofakeit.Float32()",
	models.TypeFloat64: "gofakeit.Float64()",
	models.TypeUUID:    "uuid.New()",
	models.TypeDecimal: "decimal.NewFromFloat(gofakeit.Float64())",
	models.TypeTime:    "gofakeit.Date()",
	models.TypeCoins:   "coins.FromFloat(gofakeit.Float64())",
}

func (e Entity) FakeModel() string {
	var result []string

	for _, field := range e.Fields {
		result = append(result, fmt.Sprintf("%s: %s,", field.NameCamel(true), fakeValue(field)))
	}

	return strings.Join(result, "\n")
}

func (e Entity) FakeOptions() string {
	var result []string

	for _, field := range e.Fields {
		var fieldType string

		switch field.Type {
		case models.TypeEnum:
			fieldType = field.NameCamel(true)
		default:
			fieldType = field.Type.String()
		}

		result = append(result, helpers.ExecuteTemplateFromString(fakeOption, struct {
			FieldNameCamel      string
			FieldNameLowerCamel string
			FieldType           string
			NameLowerCamel      string
			NameCamel           string
		}{
			FieldNameCamel:      field.NameCamel(true),
			FieldNameLowerCamel: field.NameLowerCamel(true),
			FieldType:           fieldType,
			NameLowerCamel:      e.NameLowerCamel(),
			NameCamel:           e.NameCamel(),
		}))
	}

	return strings.Join(result, "\n\n")
}
