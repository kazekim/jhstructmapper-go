package jhstructmapper

import (
	"fmt"
	"reflect"
)

func ParseStruct (source interface{}, target interface{}) error {

	srv := reflect.Indirect(reflect.ValueOf(source))
	trv := reflect.Indirect(reflect.ValueOf(target))

	for i := 0; i < trv.NumField(); i++ {

		tValue := trv.Field(i)
		mapField := trv.Type().Field(i).Tag.Get("map")
		if mapField == "" {
			continue
		}

		isMapped := false

		for j := 0; j < srv.NumField(); j++ {

			fieldName := srv.Type().Field(j).Name
			sValue := srv.Field(j)
			if mapField == fieldName {

				switch sValue.Kind() {
				case reflect.Ptr:
					v := reflect.New(sValue.Elem().Type()) // allocate a new pointer of the same type as the sValue's pointer field is pointing to, in this case 'Meta'
					v.Elem().Set(sValue.Elem())

					tValue.Set(v)
				default:
					tValue.Set(sValue)
				}
				isMapped = true
			}
		}

		if !isMapped {
			err := fmt.Errorf("jhstructmapper : no map for field %s (Field \"%s\" not found)", trv.Type().Field(i).Name, mapField)
			return err
		}
	}
	return nil
}

