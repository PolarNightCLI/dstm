package dstini

import (
	"reflect"
	"strings"

	"gopkg.in/ini.v1"
)

type DSTConfig interface {
	ClusterConfig | ShardConfig
}

func LoadINI[T DSTConfig](path string, obj *T) {
	confFile, err := ini.Load(path)
	if err != nil {
		panic(err)
	}

	rootType := reflect.TypeOf(*obj)
	rootVal := reflect.ValueOf(obj).Elem()
	for i := 0; i < rootType.NumField(); i++ {
		secType := rootType.Field(i)
		secVal := rootVal.Field(i)
		for j := 0; j < secType.Type.NumField(); j++ {
			t := secType.Type.Field(j)
			v := secVal.Field(j)

			secName := strings.ToUpper(secType.Name)
			iniKey := t.Tag.Get("ini")
			newVal := confFile.Section(secName).Key(iniKey)

			vType := v.Type().String()
			switch vType {
			case "string":
				v.SetString(newVal.MustString(""))
			case "uint":
				if t.Name == "Id" {
					v.SetUint(newVal.MustUint64(1))
				} else {
					v.SetUint(newVal.MustUint64(0))
				}
			case "bool":
				v.SetBool(newVal.MustBool())
			}
		}
	}
}
