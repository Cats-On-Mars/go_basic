package main

import (
	"fmt"
	"reflect"
)

type Human struct {
	Name string "姓名"
	Age  int    "年龄"
}

func (h *Human) GetName() string {
	fmt.Println("GetName called",h.Name)
	return h.Name
}

func (h *Human) SetName(name string) {
	fmt.Println("SetName called:", name)
	h.Name = name
}

func (h *Human) Eat(food string, grams int) (power int) {
	fmt.Println("Eat called:", food, grams)
	return 5 * grams
}
func main() {
	h := Human{"bill", 60}
	//getObjTypeValue(h)
	//getObjFields(h)
	//getObjMethods(&h)


	modifyFieldValue(&h)
	callMethods(&h)
	fmt.Println("after:h=", h)
}


//获取任意对象的类型和值
func getObjTypeValue(obj interface{}) {
	oType := reflect.TypeOf(obj)
	oKind := oType.Kind()
	fmt.Println(oType, oKind)

	oValue := reflect.ValueOf(obj)
	fmt.Println(oValue)
}


//获得所有属性
func getObjFields(obj interface{}) {
	oType := reflect.TypeOf(obj)
	oValue := reflect.ValueOf(obj)

	fieldsCount := oType.NumField()
	for i := 0; i < fieldsCount; i++ {

		//从对象值中获取第i个属性的值，进而值的“正射”形式
		fValue := oValue.Field(i).Interface()
		structField := oType.Field(i)
		fmt.Println(structField.Name, structField.Type, structField.Tag, fValue)
	}
}

//获取所有方法
func getObjMethods(obj interface{}) {
	oType := reflect.TypeOf(obj)
	fmt.Println(oType.NumMethod())
	for i := 0; i < oType.NumMethod(); i++ {
		method := oType.Method(i)
		fmt.Println(method.Name, method.Type)
	}
}


//修改属性的值
func modifyFieldValue(objPtr interface{}) {
	//得到【指针的Value】
	oPtrValue := reflect.ValueOf(objPtr)
	//得到【指针所指向的值（结构体）的Value】
	oValue := oPtrValue.Elem()
	fmt.Println(oValue)

	//遍历所有属性的值
	for i := 0; i < oValue.NumField(); i++ {

		//根据序号拿到【属性的Value】
		fValue := oValue.Field(i)
		//拿到属性值的原生类型
		fKind := fValue.Kind()
		//fmt.Println(fKind)

		//根据不同的原生类型设置为不同的值
		switch fKind {
		case reflect.String:
			fValue.SetString("张全蛋")
		case reflect.Int:
			fValue.SetInt(99)
		case reflect.Bool:
			fValue.SetBool(false)
		default:
			fmt.Println("设置为其他值...")
		}
	}

}


//动态调用结构体的方法
func callMethods(objPtr interface{}) {

	//要通过对象的oType拿取方法的参数列表(oType.In(i))
	oType := reflect.TypeOf(objPtr)
	//要通过对象的oValue拿取方法的具体实现(methodValue)
	oValue := reflect.ValueOf(objPtr)

	//根据方法的数量进行遍历
	for i := 0; i < oType.NumMethod(); i++ {

		//预定义要传值的参数切片[]Value
		args := make([]reflect.Value, 0) //因为call函数要求参数类型是[]value ， 所以创建一个value类型的切片 下面的append也是因为这样

		//从对象的oType拿到当前的方法的methodType
		methodType := oType.Method(i).Type

		//根据参数个数进行遍历
		//为每个参数乱怼一个同种类型反射值，丢入参数列表
		//内层循环走完时，当前方法的参数列表就形成了
		for j:=0;j<methodType.NumIn();j++{

			//从方法的methodType获取当前参数argType
			argType := methodType.In(j)
			//再获取参数类型artType的原生类型
			argKind := argType.Kind()

			//根据不同的参数原生类型乱怼相同类型的值
			switch argKind {

			case reflect.String:
				//获取字符串"一坨翔"的反射值Value，丢入参数列表
				args = append(args,reflect.ValueOf("一坨翔"))     //取的是参数的值而不是类型

			case reflect.Int:
				args = append(args,reflect.ValueOf(100))
			case reflect.Bool:
				args = append(args,reflect.ValueOf(false))
			}
		}

		//从对象值oValue中获取当前方法值methodValue
		methodValue := oValue.Method(i)

		//使用参数列表+方法值，反射调用当前方法
		methodValue.Call(args)
	}
}