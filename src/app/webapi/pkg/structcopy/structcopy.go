package structcopy

import (
	"fmt"
	"reflect"
)

// ByTag will copy values from one struct to another struct based on tags.
func ByTag(src interface{}, srcTag string, dst interface{}, dstTag string) (err error) {
	// Ensure a pointer is passed in.
	vsrc := reflect.ValueOf(src)
	if vsrc.Kind() != reflect.Ptr {
		return fmt.Errorf("src type not pointer - expected 'struct pointer' but got '%v'", vsrc.Kind())
	}

	// Ensure a struct is passed in.
	vs := reflect.Indirect(reflect.ValueOf(src))
	if vs.Kind() != reflect.Struct {
		return fmt.Errorf("src type not struct - expected 'struct pointer' but got '%v pointer'", vs.Kind())
	}

	// Ensure a pointer is passed in.
	vdst := reflect.ValueOf(dst)
	if vdst.Kind() != reflect.Ptr {
		return fmt.Errorf("dst type not pointer - expected 'struct pointer' but got '%v'", vdst.Kind())
	}

	// Ensure a struct is passed in.
	vd := reflect.Indirect(reflect.ValueOf(dst))
	if vd.Kind() != reflect.Struct {
		return fmt.Errorf("dst type not struct - expected 'struct pointer' but got '%v pointer'", vd.Kind())
	}

	// Loop through each field.
	keysSrc := vs.Type()
	keysDst := vd.Type()
	for jD := 0; jD < vd.NumField(); jD++ {
		fieldD := vd.Field(jD)
		tagD := keysDst.Field(jD).Tag
		for jS := 0; jS < vs.NumField(); jS++ {
			fieldS := vs.Field(jS)
			tagS := keysSrc.Field(jS).Tag

			// If the tags match, copy the value from src to dst field.
			if tagS.Get(srcTag) == tagD.Get(dstTag) {
				if fieldS.Type() != fieldD.Type() {
					return fmt.Errorf("field types do not match - src type '%v' for tag '%v' do not match dst type '%v' for tag '%v'",
						fieldS.Type(), tagS.Get(srcTag), fieldD.Type(), tagD.Get(dstTag))
				}
				vd.Field(jD).Set(fieldS)
			}
		}
	}
	return
}
