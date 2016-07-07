package parcom

import (
	"errors"
	"fmt"
	"reflect"
)

// Map is a parser that maps the result of the previous parser with a function.
func Map(parser Parser, fns ...interface{}) Parser {
	fn := func(i interface{}) interface{} {
		for _, f := range fns {
			i = mapFunc(f, i)
		}
		return i
	}
	return func(in string) (string, interface{}, bool) {
		remaining, value, ok := parser(in)
		if ok {
			value = fn(value)
		}
		return remaining, value, ok
	}
}

func convert(inV reflect.Value, outT reflect.Type) reflect.Value {
	if !inV.IsValid() {
		return reflect.Zero(outT)
	}

	inT := inV.Type()
	if inT.ConvertibleTo(outT) {
		return inV.Convert(outT)
	} else if inT.Kind() == reflect.Interface {
		return convert(inV.Elem(), outT)
	} else if inT.Kind() == reflect.Slice && outT.Kind() == reflect.Slice {
		outET, outLen := outT.Elem(), inV.Len()

		sliceT := reflect.SliceOf(outET)
		sliceV := reflect.MakeSlice(sliceT, outLen, outLen)

		for i := 0; i < outLen; i++ {
			iV := convert(inV.Index(i), outET)
			sliceV.Index(i).Set(iV)
		}
		return sliceV
	}
	panic(fmt.Errorf("parcom: cannot convert %s to %s", inT, outT))
}

func mapFunc(f, i interface{}) interface{} {
	// Get value and type of the mapping function.
	fV := reflect.ValueOf(f)
	fT := fV.Type()

	// Check the mapping function's type.
	if fT.Kind() != reflect.Func || fT.NumIn() == 0 || fT.NumOut() != 1 {
		panic(errors.New("parcom: invalid map function"))
	}

	// Convert i into a valid parameter for the mapping function and call the
	// mapping function.
	var retVs []reflect.Value
	iV := reflect.ValueOf(i)
	if nIn := fT.NumIn(); nIn == 1 {
		pT := fT.In(0)
		pV := convert(iV, pT)
		retVs = fV.Call([]reflect.Value{pV})
	} else if iV.Kind() == reflect.Slice || iV.Kind() == reflect.Array {
		pVs := make([]reflect.Value, nIn)
		for i := 0; i < (nIn - 1); i++ {
			pVs[i] = convert(iV.Index(i), fT.In(i))
		}
		i := nIn - 1
		if nIn == iV.Len() {
			pVs[i] = convert(iV.Index(i), fT.In(i))
			retVs = fV.Call(pVs)
		} else if nIn < iV.Len() && fT.IsVariadic() {
			if fT.In(i) == reflect.TypeOf([]interface{}{}) {
				pVs[i] = iV.Slice(i, iV.Len())
			} else {
				pVs[i] = convert(iV.Slice(i, iV.Len()), fT.In(i))
			}
			retVs = fV.CallSlice(pVs)
		} else {
			panic(fmt.Errorf("parcom: too many arguments"))
		}
	} else {
		panic(fmt.Errorf("parcom: cannot use %s as multiple arguments", iV.Type()))
	}

	// Return.
	return retVs[0].Interface()
}
