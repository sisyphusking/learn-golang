package main

import (
	"errors"
	"fmt"
	"reflect"
)

var kIterateTimeErr error = errors.New("call time too small")
var kFuncReturnSizeZero error = errors.New("zero return value")
var kFirstValueNotErrorErr error = errors.New(
	"first return value shuld be error")

// function: 这个函数,必须第一个返回值是error,如果error是nil表示执行成功了
//           parameters,参数列表,可以是空,nil; returnValues,返回值列表,也可以是空,nil
func Retry(function interface{}, parameters []interface{}, times int) (
	err error, returnValues []interface{}) {
	if times <= 0 {
		return kIterateTimeErr, returnValues
	}
	var inputValue []reflect.Value
	var funcReValue []reflect.Value
	var errInterface interface{}
	var tmpErr error
	var ok bool

	if len(parameters) > 0 {
		inputValue = make([]reflect.Value, len(parameters))
		for i := range parameters {
			inputValue[i] = reflect.ValueOf(parameters[i])
		}
	}

	for times > 0 {
		funcReValue = reflect.ValueOf(function).Call(inputValue)
		if len(funcReValue) == 0 {
			return kFuncReturnSizeZero, returnValues
		}
		if !funcReValue[0].IsNil() {
			errInterface = funcReValue[0].Interface()
			tmpErr, ok = errInterface.(error)
			if !ok {
				return kFirstValueNotErrorErr, returnValues
			}
		} else {
			tmpErr = nil
		}
		if tmpErr == nil {
			if len(funcReValue) > 1 {
				returnValues = make([]interface{}, len(funcReValue)-1)
				for i := range funcReValue {
					if i == 0 {
						continue
					}
					returnValues[i-1] = funcReValue[i].Interface()
				}
			}
			return nil, returnValues
		}

		times--
	}

	return tmpErr, returnValues
}

// ---- test code ----
var G int

type reV1 struct {
	ReS1 string
	ReI1 int
}

type InV1 struct {
	InS1 string
	InI1 int
}

func TestFunc(in InV1) (err error, re reV1) {
	G++
	if G < 5 {
		return fmt.Errorf("error ocus"), re
	}

	re.ReI1 = 10 + in.InI1
	re.ReS1 = "test" + in.InS1
	return nil, re
}

func TestFunc2() (err error, re reV1) {
	G++
	if G < 5 {
		return fmt.Errorf("error ocus"), re
	}

	re.ReI1 = 10
	re.ReS1 = "test2"
	return nil, re
}

func main() {
	var in InV1
	in.InI1 = 5
	in.InS1 = " data"
	var re reV1

	err, valueVec := Retry(TestFunc, []interface{}{in}, 4)
	if err == nil {
		re = valueVec[0].(reV1)
		fmt.Println(re)
	} else {
		fmt.Printf("err:%s\n", err.Error())
	}

	err, valueVec = Retry(TestFunc, []interface{}{in}, 5)
	if err == nil {
		re = valueVec[0].(reV1)
		fmt.Println(re)
	} else {
		fmt.Printf("err:%s\n", err.Error())
	}

	err, valueVec = Retry(TestFunc2, nil, 5)
	if err == nil {
		re = valueVec[0].(reV1)
		fmt.Println(re)
	} else {
		fmt.Printf("err:%s\n", err.Error())
	}
}

