package main

import "fmt"

func main() {
	//	fmt.Println("Объявление массива")
	//	var arrya [4]int
	//
	//	fmt.Println(arrya)
	//	fmt.Println()
	//

	var arrya11 [7]int
	fmt.Println(arrya11)
	arrya11[0] = 55
	arrya11[6] = 555
	fmt.Println(arrya11)

	const arr = 5 //Обратите внимание, что вы можете использовать константы для указания размера массива,
	// но вы не можете использовать переменные:
	//var arr1 = 5

	var a [arr]int
	//var b [arr1]int //Will throw a compilation error (Выдаст ошибку компиляции)

	fmt.Println(a)
	//	fmt.Println()
	//
	//	fmt.Println("Декларация и назначение элементов")
	//	//Вы можете сразу объявить массив и присвоить значения его элементам:
	//
	//	var arr2 = [4]int{4, 5, 7, 8}
	//	fmt.Println(arr2)
	//
	//	var arr3 = [4]int{4, 3}
	//	fmt.Println(arr3)
	//	fmt.Println()
	//

	var arr11 = [3]int{1, 2, 3}
	fmt.Println(arr11)
	arr11[0] = 22
	arr11[1] = 33
	arr11[2] = 44
	fmt.Println(arr11)

	var arr12 = [3]int{0: 3, 2: 1}
	fmt.Println(arr12)
	arr12[0] = 55
	arr12[1] = 777
	arr12[2] = 888
	fmt.Println(arr12)
	var arr13 = [4]int{
		0: 4,
		3: 1,
	}
	fmt.Println(arr13)
	arr13[0] = 44
	arr13[1] = 1010
	arr13[2] = 1111
	arr13[3] = 2222
	fmt.Println(arr13)

	var arr1313 = [4]int{arr13[0], arr13[1], arr13[2], arr13[3]}
	fmt.Println(arr1313[0])
	fmt.Println(arr1313[1])
	fmt.Println(arr1313[2])
	fmt.Println(arr1313[3])

	fmt.Println(arr13[0] + arr1313[0])

	//	//Кроме того, вы можете объявить массив и присвоить значения определенным элементам,
	//	//используя пару index: value:
	//	fmt.Println("Объявим массив и присвоим значения определенным элементам")
	//	fmt.Println("используя пару index: value:")
	//	var arr4 = [5]int{1: 3, 4: 45}
	//	fmt.Println(arr4)
	//
	//	//Имейте в виду, что если у вас есть многострочное задание,
	//	//вам нужно поставить запятую в конце каждой строки:
	//
	//	var arr44 = [5]int{
	//		0: 4,
	//		3: 55,
	//	}
	//	fmt.Println(arr44)
	//	fmt.Println()
	//	fmt.Println("Индексация")
	//
	//	//Индексация
	//	//Теперь давайте посмотрим, как мы можем использовать массивы.
	//	//Прежде всего, нам нужно будет получить доступ к элементу в массиве.
	//	//Для этого мы должны указать имя массива, а затем индекс требуемого элемента в пределах []
	//
	//	var arr5 = [4]int{2, 4, 8, 10}
	//	fmt.Println("Элемент массива arr5 под индексом 3: arr5[3] = ", arr5[3])
	//	fmt.Println()
	//
	//	//arr5[3]эквивалентно обычной переменной. Следовательно, мы можем присвоить ему новое значение:
	//	fmt.Println("arr5[3]эквивалентно обычной переменной. Следовательно, мы можем присвоить ему новое значение:")
	//	arr5[3] = 45
	//	fmt.Printf("arr5[3] = 45\nНовое значение равно %d\n", arr5[3])
	//	fmt.Println()
	//
	//	//В отличие от объявления, индексация может использовать переменные.
	//	//Имейте в виду, что программа Go будет паниковать, если индекс выйдет за границу массива во время выполнения:
	//	fmt.Println("В отличие от объявления, индексация может использовать переменные.")
	//
	//	var i = 3
	//	fmt.Println("var i = 3\nПод индексом [3] в нашем массиве находиться значение = ", arr5[i])
	//	fmt.Println()
	//
	//	fmt.Println("Назначение массива")
	//	//Назначение массива
	//	//Массив сам по себе является типом с двумя свойствами: размером и типом элементов.
	//	//Чтобы мы могли назначать arr6, arr7 они должны иметь одинаковый размер и один и тот же тип элементов,
	//	//как здесь:
	//	var arr6 = [4]int{3, 5, 7, 8}
	//	var arr7 [4]int
	//
	//	arr7 = arr6
	//	fmt.Println("Массиву arr7 назначены значения массива arr6:", arr7)
	//
	//	var r = [5]int{4, 4: 55}
	//	var t [5]int
	//
	//	t = r
	//	fmt.Println("Массиву t назначены значения массива r:", t)
	//	fmt.Println()
	//
	//	//Многомерные массивы
	//	//Как мы упоминали ранее, массив сам по себе является типом.
	//	//Поэтому мы можем снова проделать тот же трюк и создать массив массивов:
	//
	//	fmt.Println("Многомерные массивы")
	//	var multiArrya [3][4]int
	//	fmt.Println(multiArrya)
	//
	//	var multiArrya1 = [3][4]int{
	//		{4, 5, 6, 7},
	//		{1, 2, 4, 55},
	//		{57, 4, 55, 67},
	//	}
	//	fmt.Println(multiArrya1)
	//	fmt.Println()
	//
	//	var multiArrya2 = [3][4]int{
	//		1: {
	//			2: 5,
	//			3: 77,
	//		},
	//	}
	//	fmt.Println(multiArrya2)
	//	fmt.Println()
	//
	//	var h = multiArrya1[1] // Will assign [1 2 4 55] to the v variable
	//	fmt.Println(h)
	//	fmt.Println()
	//
	//	var h1 = multiArrya1[1][3] //Will assign 55 to the v variable
	//	fmt.Println(h1)
	//	fmt.Println()
	//
	//	var ee = [2][3]string{
	//		{"Super", "Ilyas", "Svetlana"},
	//		{"Valihan", "Danya", "Davlat"},
	//	}
	//	var ff [2][3]string
	//
	//	ff = ee // OK
	//	fmt.Println(ff)
	//
}

//func main() {
//	// DO NOT delete this line! It takes the input numbers and generates the 3x3 matrix `A`.
//	var A = [3][3]int{
//		{1, 2, 3},
//		{4, 5, 6},
//		{7, 8, 9},
//	}
//
//	// Declare and initialize the transposed version of the 3x3 matrix `A` below:
//	var transposedA = [3][3]int{
//		{A[0][0], A[1][0], A[2][0]},
//		{A[0][1], A[1][1], A[2][1]},
//		{A[0][2], A[1][2], A[2][2]},
//	}
//
//	// Print each row of the 3x3 `transposedA` matrix below:
//	fmt.Println(transposedA[0])
//	fmt.Println(transposedA[1])
//	fmt.Println(transposedA[2])
//}

//func main() {
//	var a = [3][3]int{
//		{1, 2, 3},
//		{4, 5, 6},
//		{7, 8, 9},
//	}
//
//	var aa = [3][3]int{
//		{a[2][2], a[1][2], a[0][2]},
//		{a[2][1], a[1][1], a[0][1]},
//		{a[2][0], a[1][0], a[0][0]},
//	}
//	fmt.Println(aa[0])
//	fmt.Println(aa[1])
//	fmt.Println(aa[2])
//}
