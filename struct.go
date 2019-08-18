package jhstructmapper

import (
	"fmt"
	"reflect"
)

func ParseWithMapTag(source interface{}, target interface{}) error {
	return Parse(source, target, false)
}

func ParseSameFieldName(source interface{}, target interface{}) error {
	return Parse(source, target, true)
}

func Parse (source interface{}, target interface{}, isMapSameFieldName bool) error {

	srv := reflect.Indirect(reflect.ValueOf(source))
	trv := reflect.Indirect(reflect.ValueOf(target))

	for i := 0; i < trv.NumField(); i++ {

		tValue := trv.Field(i)
		tFieldName := trv.Type().Field(i).Name
		mapField := trv.Type().Field(i).Tag.Get("map")
		if !isMapSameFieldName && mapField == "" {
			continue
		}

		isMapped := false

		for j := 0; j < srv.NumField(); j++ {

			sFieldName := srv.Type().Field(j).Name
			sValue := srv.Field(j)
			if mapField == sFieldName || (isMapSameFieldName && tFieldName == sFieldName) {
				copyValue(sValue, tValue)
				isMapped = true
			}
		}

		if !isMapped {
			err := fmt.Errorf("jhstructmapper : no map for field %s (Field \"%s\" not found)", tFieldName, mapField)
			return err
		}
	}
	return nil
}

func copyValue(sValue reflect.Value, tValue reflect.Value) {
	switch sValue.Kind() {
	case reflect.Ptr:
		v := reflect.New(sValue.Elem().Type()) // allocate a new pointer of the same type as the sValue's pointer field is pointing to, in this case 'Meta'
		v.Elem().Set(sValue.Elem())

		tValue.Set(v)
	default:
		tValue.Set(sValue)
	}
}