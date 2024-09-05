package main

import (
	"fmt"
	variables "golesson/01_variables"
)

// errors "golesson/14_error_handling"
// strings "golesson/15_string_functions"

//"golesson/conditions"
//"golesson/loops"

func main() {
	variables.Variable()
	//fmt.Println(variables.Hello("Chris"))
	//variables.TypeCasting()
	//consoleread.ConsoleRead()
	//others.Times()
	//consoleread.ReadFromConsole()
	//conditions.WhileInGo()
	//fmt.Print()
	//conditions.Condition()
	//conditions.ConditionTwo()
	//conditions.Workshop1()
	//loops.Loops()
	//loops.Workshop2()
	//loops.Workshop3()
	//loops.Workshop4()
	//arrays.MultiArrays()
	//slices.Slice()
	//functions.SayHello("hanoi")
	//fmt.Println(functions.Sum(20, 30))
	// fmt.Println(functions.DortIslem(10, 6))
	// result1, result2, result3, _ := functions.DortIslem(10, 6) // istemediğin yere "_" koyman yeterli, koymamak olmaz.
	// fmt.Println("sum:", result1)
	// fmt.Println("abstraction:", result2)
	// fmt.Println("multiply:", result3)
	//fmt.Println("division:", result4)
	//fmt.Println(functions.SumVariadic(1, 2, 30))
	//numbers := []int{1, 2, 3, 4, 5}
	//fmt.Println(functions.SumVariadic(numbers...)) // variadic olanları 3 nokta ile gönderiyoruz.
	//maps.Map()
	//maps.Range()
	// message := "Hi"
	// pointers.SayHello(&message)
	// println(message)
	//fmt.Println(functions.NamedReturn(2, 3))
	// number := 20
	// pointers.Pointer(&number) // yollarken and operator
	// fmt.Println("number in the main", number)
	// numbers := []int{1, 2, 3} // arrayler bellekteki adres ile gönderilir, int ten farklıdır.
	// pointers.Pointer2(numbers)
	// fmt.Println("numbers[0] in the main", numbers[0]) // burada da 100 döner
	//structs.Struct()
	//structs.Struct2()
	// go goroutines.GoRoutine() // asenkron çalışmayı sağlar. "go"
	// go goroutines.GoRoutine2()
	// time.Sleep(3 * time.Second)
	// fmt.Println("main bitti")
	// ciftSayiCn := make(chan int) // sayı değil, içinde sayı taşıyan channel...
	// tekSayiCn := make(chan int)
	// go channels.GoRoutine(ciftSayiCn)
	// go channels.GoRoutine2(tekSayiCn)

	// ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn
	// multiply := ciftSayiToplam * tekSayiToplam

	// //multiply := ciftSayiCn * tekSayiCn // bu hata verir, channelları direkt int diyemezsin çünkü sayı değil, içinde sayı taşıyan channel...
	// fmt.Println("Multiply: ", multiply)
	//interfaces.Demo1()
	//interfaces.CreditCalculatorTest()
	//defers.RunDefer()
	//errors.ErrorHandling()
	//errors.TestDogrula()
	//errors.TahminTest()
	//fmt.Println(errors.TahminEt2(99))
	//strings.StringTest()
	//restful.GetRestful()
	//restful.PostRestful()
	//restful.GetAllProducts()
	// restful.AddProduct()
	// products, _ := restful.GetAllProducts()
	// fmt.Println(products)
	// for i := 0; i < len(products); i++ {
	// 	fmt.Println(products[i])
	// }

	//structs.
	//anonim fonksiyonlar
	addFunc := func(numbers ...int) (numNumbers, sum int) {
		for _, number := range numbers {
			sum += number
		}
		numNumbers = len(numbers)
		return
	}
	numNumbers, sum := addFunc(2, 3)
	fmt.Printf("numNumbers is %v and sum is %v\n", numNumbers, sum)

}
