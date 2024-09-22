package utils

import (
	"errors"
	"reflect"

	"log"
)

// GenericMapper maps the values from the DTO to the corresponding entity
func GenericMapper(input interface{}, output interface{}) error {
	inputVal := reflect.ValueOf(input)
	outputVal := reflect.ValueOf(output)

	if inputVal.Kind() != reflect.Ptr || outputVal.Kind() != reflect.Ptr {
		return errors.New("input and output must be pointers")
	}

	inputElem := inputVal.Elem()
	outputElem := outputVal.Elem()

	if inputElem.Kind() != reflect.Struct || outputElem.Kind() != reflect.Struct {
		return errors.New("input and output must be structs")
	}

	for i := 0; i < inputElem.NumField(); i++ {
		inputField := inputElem.Type().Field(i)
		inputFieldValue := inputElem.Field(i)

		outputField := outputElem.FieldByName(inputField.Name)
		if outputField.IsValid() && outputField.CanSet() && outputField.Type() == inputFieldValue.Type() {
			outputField.Set(inputFieldValue)
		}
	}

	return nil
}

func ConvertDTOtoEntity(input interface{}, output interface{}) error {

	// Usar la función genérica para mapear el DTO a la entidad
	err := GenericMapper(input, output)
	if err != nil {
		log.Println("Error al intentar mapear el DTO a Entity {}", err)
		log.Println("DTO: ", input)
		return err
	}
	return nil
}
