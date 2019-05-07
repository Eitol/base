package config_extractor

// Hector Oliveros - 2019
// hector.oliveros.leon@gmail.com

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"os"
	"reflect"
	"strconv"
)

const EnvTagName string = "env"
const DefaultCase caseType = ScreamingSnake

type caseType string

const (
	ScreamingSnake = caseType("ANY_KIND_OF_STRING")
	Snake          = caseType("any_kind_of_string")
	Kebab          = caseType("any-kind-of-string")
	ScreamingKebab = caseType("ANY-KIND-OF-STRING")
	Camel          = caseType("AnyKindOfString")
	LowerCamel     = caseType("anyKindOfString")
)

type ExtractorOptions struct {
	TakeNameIfNoTag bool
	EnvNameCaseType caseType
}

func (e *ExtractorOptions) mergeWithDefault() {
	if e.EnvNameCaseType == caseType("") {
		e.EnvNameCaseType = DefaultCase
	}
}

type ExtractorArgs struct {
	Options ExtractorOptions
	Configs []interface{}
}

func Extract(args ExtractorArgs) error {
	args.Options.mergeWithDefault()
	for i := 0; i < len(args.Configs); i += 2 {
		err := extractByIdx(i, args)
		if err != nil {
			return err
		}
	}
	return nil
}

func extractByIdx(idx int, args ExtractorArgs) error {
	//defConfig := args.DefaultValues[idx]
	v := reflect.ValueOf(args.Configs[idx]).Elem()
	vptr := v
	// if its a pointer, resolve its value
	if v.Kind() == reflect.Ptr {
		vptr = reflect.Indirect(v)
	}

	for i := 0; i < vptr.NumField(); i++ {
		f := v.Field(i)
		// make sure that this field is defined, and can be changed.
		if !f.IsValid() || !f.CanSet() {
			continue
		}
		envName, haveTagEnv := v.Type().Field(i).Tag.Lookup(EnvTagName)
		if !haveTagEnv && args.Options.TakeNameIfNoTag {
			prefix := ""
			if args.Configs[idx+1] != "" {
				var ok bool
				prefix, ok = args.Configs[idx+1].(string)
				if !ok {
					panic("ERROR: Invalid configurarion. Expected string")
				}
			}
			envName = changeCase(args.Options.EnvNameCaseType, prefix+v.Type().Field(i).Name)
		}
		envVal, exist := os.LookupEnv(envName)
		if !exist || envVal == "" {
			continue
		}
		err := setValue(f, envVal)
		if err != nil {
			return err
		}
	}
	return nil
}

func setValue(field reflect.Value, value interface{}) error {
	switch field.Kind() {
	case reflect.Int, reflect.Int64:
		valueInt, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			return fmt.Errorf("invalid integer: %v of type %v", value, reflect.TypeOf(value))
		}
		field.SetInt(int64(valueInt))
	case reflect.String:
		field.SetString(value.(string))
	}
	return nil
}

func changeCase(caseType_ caseType, envName string) string {
	m := map[caseType]func(string) string{
		ScreamingSnake: strcase.ToScreamingSnake,
		Snake:          strcase.ToSnake,
		Kebab:          strcase.ToKebab,
		ScreamingKebab: strcase.ToScreamingKebab,
		Camel:          strcase.ToCamel,
		LowerCamel:     strcase.ToLowerCamel,
	}
	f, ok := m[caseType_]
	if !ok {
		return envName
	}
	return f(envName)
}
