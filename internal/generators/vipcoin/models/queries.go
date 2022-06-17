package models

import (
	"strings"

	"github.com/iancoleman/strcase"
)

func (e Entity) TableName() string {
	return e.Table
}

func (e Entity) InsertFields() string {
	f := make([]string, 0, len(e.Fields))
	for _, field := range e.Fields[1:] { // first field is id
		f = append(f, strcase.ToSnake(field.Name))
	}

	return strings.Join(f, ", ")
}

func (e Entity) SelectFields() string {
	f := make([]string, 0, len(e.Fields))
	for _, field := range e.Fields {
		f = append(f, strcase.ToSnake(field.Name))
	}

	return strings.Join(f, ", \n\t\t\t")
}

func (e Entity) InsertValues() string {
	v := make([]string, 0, len(e.Fields))
	for _, field := range e.Fields[1:] { // first field is id
		v = append(v, ":"+strcase.ToSnake(field.Name))
	}

	return strings.Join(v, ", ")
}

func (e Entity) UpdateQuery() string {
	v := make([]string, 0, len(e.Fields))
	for _, field := range e.Fields[1:] { // first field is id
		v = append(v, strcase.ToSnake(field.Name)+" = :"+strcase.ToSnake(field.Name))
	}

	return strings.Join(v, ",\n\t\t\t")
}
