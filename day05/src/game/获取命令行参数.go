package game

import (
	"flag"
	"os"
)

func GetArgs() (name string,age int64,money float64,readBook bool){

	flag.StringVar(&name,"name","","姓名")
	flag.Int64Var(&age,"age",0,"年龄")
	flag.Float64Var(&money,"money",0,"财富")
	flag.BoolVar(&readBook,"readBook",false,"爱好阅读吗")

	flag.Parse()

	return
}

func GetArgs2() (name string,age int64,money float64,readBook bool){
	namePtr := flag.String("name","","姓名")
	agePtr := flag.Int64("age",0,"年龄")
	moneyPtr := flag.Float64("money",0,"财富")
	readBookPtr := flag.Bool("readBook",false,"爱好阅读吗")

	flag.Parse()

	return *namePtr,*agePtr,*moneyPtr,*readBookPtr
}

func GetArgs3() []string{
	return os.Args
}