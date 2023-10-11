package main

import (
	"fmt"
	"sort"
)

// Поиск индекса элемента: Напишите функцию, которая находит индекс первого вхождения заданного значения в массиве.
// Если значение не найдено, верните -1.

// Удаление элемента из массива: Напишите функцию, которая удаляет заданный элемент из массива и возвращает новый массив

// Сумма элементов массива: Напишите функцию, которая принимает на вход массив целых чисел и возвращает их сумму.
// Поиск максимального элемента: Напишите функцию, которая находит максимальный элемент в массиве и возвращает его значение.
// Разворот массива: Напишите функцию, которая разворачивает массив на месте (без использования дополнительного массива).

//func findDuplicates(arr []int) []int {
//	duplicates := []int{}
//	seen := make(map[int]bool)
//
//	for _, num := range arr {
//		if seen[num] {
//			duplicates = append(duplicates, num)
//		} else {
//			seen[num] = true
//		}
//	}
//	return duplicates
//}
//func main() {
//	var size int
//	fmt.Print("Введите размер массива: ")
//	fmt.Scan(&size)
//
//	array := make([]int, size)
//	for i := 0; i < size; i++ {
//		fmt.Printf("Введите элементы массива %d: ", i+1)
//		fmt.Scan(&array[i])
//	}
//
//	newArray := findDuplicates(array)
//	fmt.Printf("Mассив дубликатов: %d ", newArray)
//
//}

//func CountFalseValueInMap(mp map[int]bool) int {
//	count := 0
//
//	for _, num := range mp {
//		if num == false {
//			count++
//		}
//	}
//	return count
//}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
func main() {

	var s = []int{12, 23, 34}
	fmt.Println("s - ", s)
	var sn = make([]int, len(s))
	fmt.Println("sn - ", sn)

	//copy(sn, s)
	//fmt.Println("Скопировали значения s в sn - ", sn)
	//copy(s, sn)
	//fmt.Println("Скопировали значения sn в s - ", sn)

	var n = copy(sn, s) // var n is the number of copied elements
	//fmt.Println("n - ", n)
	//fmt.Println("sn - ", sn)

	sn[0] = 0  // Here, we assign a new value to the slice elements in order
	sn[1] = 11 // to see whether it is a proper copy or the same slice

	fmt.Println(n)  // 3
	fmt.Println(s)  // [12 23 34] - the initial slice with no changes
	fmt.Println(sn) // [ 0 11 34] - the copied slice with modified elements

	ss := []int{8, 8, 17, 4}
	sort.Ints(ss)
	fmt.Println("сортировка: ", ss)

	//array := []int{45, 5, 78, 34, 47, 2, -5, 34}
	//fmt.Println("неотсортированный массив: ", array)
	//bubbleSort(array)
	//fmt.Println("Отсортированый массив: ", array)

	//s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//s2 := s[3:5]
	//s21 := s[3:]
	//s22 := s[:5]
	//s23 := s[:]
	//fmt.Println(s2)
	//fmt.Println(s21)
	//fmt.Println(s22)
	//fmt.Println(s23)
	//fmt.Println()
	//
	//arr := []int{7, 2}
	//arr2 := arr
	//fmt.Println("arr: ", arr)
	//fmt.Println("arr2: ", arr2)
	//fmt.Println()
	//
	//arr2[0] = 42
	//fmt.Println("Меняем у arr2 значение элемента под индексом [0] на 42")
	//fmt.Println("arr2: ", arr2)
	//fmt.Println("У среза arr значение элемента под индексом [0] тоже поменяется на 42")
	//fmt.Println("arr: ", arr)
	//fmt.Println()
	//
	//fmt.Println("Добавляем с помощью ключевого слова append переменной arr2 элементы 3, 4, 5, 6, 7, 8")
	//arr2 = append(arr2, 3, 4, 5, 6, 7, 8)
	//fmt.Println("arr2: ", arr2)
	//fmt.Println("После добавления в срез arr2 нескольких элементов произошла перелокация, создался новый срез\nЗначения среза arr остались прежнеми")
	//fmt.Println("arr: ", arr)
	//
	//fmt.Println()
	//arr2[0] = 1
	//fmt.Println("В переменной arr2 меняем значение элемента под индексом [0] на 1")
	//
	//fmt.Println("arr2: ", arr2)
	//arr[0] = 5
	//fmt.Println("arr: ", arr)

	//a := make([]byte, 1, 2)
	//fmt.Println("Срез а: ", a)
	//fmt.Printf("Длина среза: %d, Емкость среза: %d", len(a), cap(a))
	//fmt.Println()
	//a = append(a, 1)
	//fmt.Println("Срез а: ", a)
	//fmt.Printf("Длина среза: %d, Емкость среза: %d", len(a), cap(a))
	//fmt.Println()
	//a = append(a, 2)
	//fmt.Println("Срез а: ", a)
	//fmt.Printf("Длина среза: %d, Емкость среза: %d", len(a), cap(a))

}
