package main

import (
	"awesomeProject/mathslice"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)
import "unicode/utf8"

// ООП и ГО
// Глобальные переменный вне функции, локальнеые внутри ЛООООЛ
//Как Go определяет, что сущность экспортируемая? Если имя переменной,
//константы или функции начинается с прописной буквы, то она экспортируемая.
//Если со строчной, то неэкспортируемая.

const (
	Black = iota // Глобальная экспортируемая
	Gray
	White
)

// счётчик обнуляется

const (
	_      = iota * 10
	yellow // Глобальная неэкспортируемая
	red
	green = iota * -1 // это присваивание не обнулит iota
	blue
)

type Person struct {
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	age         int       `json:"-"`
	lastVisited time.Time `json:"-"`
}

func main() { //Неэкспротная функция
	// В го есть три типа обьявления переменной
	// var name = expression
	// var name type = expression && var name type
	// name := expression
	fmt.Println("Hello Go")
	var russianStr string // Локальная, экспорт недоступен пишёв в пызду
	russianStr = "абц"

	println(utf8.RuneCountInString(russianStr))

	type Name string
	type Fruit string

	var fruit Fruit
	var name Name

	fruit = "Apple"
	name = Name(fruit) // type(variable) - привод типов данных

	println(name)
	println(name[0])         // выведет байт, тк в ГО стринг это массив байтов
	println(string(name[0])) // конвертация байта в букву

	// Вывод к константам и iota
	fmt.Println(Black, Gray, White)
	fmt.Println(yellow, red, green, blue)

	Export()
	fmt.Println("BuzzFizz")
	FizzBuzz()
	fmt.Println("End")
	fmt.Println("Composite")
	CompositeVariable()

	p := Person{
		Name:        "Alex",
		Email:       russianStr,
		age:         25,
		lastVisited: time.Time{},
	}

	UpdatePersonWithLastVisited(&p)

	fmt.Println(p.lastVisited)
	fmt.Println("end")
	fmt.Println("Arrays")
	ArrayVariable()
	fmt.Println("end")
	fmt.Println("Arrays")
	Maps()
	fmt.Println("end")

	fmt.Println("Exam")
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}
	fmt.Println(RemoveDuplicates(input))
	fmt.Println("end")

	fileJson, err := json.Marshal(p)
	if err != nil {
		log.Fatalln("unable marshal to json")
	}
	fmt.Println(string(fileJson))

	fmt.Println("Path")
	PrintAllFilesWithFilter(".", "main.go")
	fmt.Println("PrintAllFilesWithFuncFilter")
	PrintFilesWithFuncFilter(".", containsDot)
	fmt.Println("1")
	fmt.Println(Global)
	functionGlobal()
	fmt.Println("3")

	fmt.Println("Import")
	s := mathslice.Slice{1, 2, 3}
	fmt.Println(s)
	fmt.Println("Сумма слайса: ", mathslice.SumSlice(s))

	mathslice.MapSlice(s, func(i mathslice.Element) mathslice.Element {
		return i * 2
	})

	fmt.Println("Слайс, умноженный на два: ", s)

	fmt.Println("Сумма слайса: ", mathslice.SumSlice(s))

	fmt.Println("Свёртка слайса умножением ",
		mathslice.FoldSlice(s,
			func(x mathslice.Element, y mathslice.Element) mathslice.Element {
				return x * y
			},
			1))

	fmt.Println("Свёртка слайса сложением ",
		mathslice.FoldSlice(s,
			func(x mathslice.Element, y mathslice.Element) mathslice.Element {
				return x + y
			},
			0))

}

func UpdatePersonWithLastVisited(p *Person) {
	p.lastVisited = time.Now()
}

func Export() { //Экспортная функция
	answer := "Иди нахуй!"
	println(answer)
}

func FizzBuzz() {
	for i := 1; i < 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz!")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func CompositeVariable() {
	var localVar int = 5                            // переменная
	var localPointer *int = &localVar               // адрес переменной
	var localPointerToPointer **int = &localPointer // адресс адресса переменной
	fmt.Println(localVar, localPointer, localPointerToPointer, localPointerToPointer, *localPointerToPointer, **localPointerToPointer)

	type localStruct struct {
		IntField int
	}
	// Литерал localStruct{} создаёт в памяти переменную типа А. Затем от неё берётся указатель
	_ = &localStruct{
		IntField: 10,
	}
	var localPointerStruct *localStruct = new(localStruct) //  то же самое, что и &localStruct{}
	localPointerStruct.IntField = 123
	fmt.Println(localPointerStruct)
}

func ArrayVariable() {

	fmt.Println("Arrays")

	var lastWeekTemp [7]int // обьявление массива
	lastWeekTemp[2] = 13
	tempOnWednesday := lastWeekTemp[2] // обращение к массиву
	fmt.Println(tempOnWednesday)
	rgbColor := [...]uint8{255, 255, 128} // создание массива без обьявления длины
	fmt.Println(rgbColor)
	nextWeekTemp := [7]int{-3, 5, 7} // инициализация массива с значениями
	fmt.Println(nextWeekTemp)
	thisWeekTemp := [7]int{6: 11, 2: 3} // инициализация массива с значениями в разных местах массива
	fmt.Println(thisWeekTemp)

	var _ [4][7]int // обьявление многомерного массива
	var rgbImage [1080][1920][3]uint8

	var _ = rgbImage[2]       // 3-я строка в изображении
	var _ = rgbImage[2][3]    // 4-й пиксель в третьей строке изображения
	var _ = rgbImage[2][3][1] // значение синей компоненты (второй байт) 4-го пикселя в третьей строке изображения

	fmt.Println("Arrays Range")

	var weekTemp = [7]int{5, 4, 6, 8, 11, 9, 5}

	sumTemp := 0
	for i := 0; i < len(weekTemp); i++ {
		sumTemp += weekTemp[i]
	}
	average := sumTemp / len(weekTemp)

	fmt.Println(average)

	altSumTemp := 0
	for _, temp := range weekTemp {
		altSumTemp += temp
	}
	altAverage := altSumTemp / len(weekTemp)

	fmt.Println(altAverage)

	alterSumTemp := 0
	for i := range weekTemp {
		alterSumTemp += weekTemp[i]
	}
	alterAverage := alterSumTemp / len(weekTemp)

	fmt.Println(alterAverage)

	fmt.Println("Slice Create") // mySlice := make([]TypeOfElement, LenOfSlice, CapOfSlice)

	weekTempArr := [7]int{1, 2, 3, 4, 5, 6, 7} // создание слайсов на основе массива или слайса <- ТУТ
	workDaysSlice := weekTempArr[:5]
	weekendSlice := weekTempArr[5:]
	fromTuesdayToThursDaySlice := weekTempArr[1:4]
	weekTempSlice := weekTempArr[:]

	fmt.Println(workDaysSlice, len(workDaysSlice), cap(workDaysSlice))                                        // [1 2 3 4 5] 5 7
	fmt.Println(weekendSlice, len(weekendSlice), cap(weekendSlice))                                           // [6 7] 2 2
	fmt.Println(fromTuesdayToThursDaySlice, len(fromTuesdayToThursDaySlice), cap(fromTuesdayToThursDaySlice)) // [2 3 4] 3 6
	fmt.Println(weekTempSlice, len(weekTempSlice), cap(weekTempSlice))                                        // [1 2 3 4 5 6 7] 7 7

	workDaysSlice[4] = 28

	fmt.Println(workDaysSlice, len(workDaysSlice), cap(workDaysSlice))
	fmt.Println(weekendSlice, len(weekendSlice), cap(weekendSlice))
	fmt.Println(fromTuesdayToThursDaySlice, len(fromTuesdayToThursDaySlice), cap(fromTuesdayToThursDaySlice))
	fmt.Println(weekTempSlice, len(weekTempSlice), cap(weekTempSlice))

	fmt.Println("Slice Append(+)")

	s := make([]int, 4, 7) // [0 0 0 0], len = 4 cap = 7
	// 1. Создаём слайс s с базовым массивом на 7 элементов.
	// Четыре первых элемента будут доступны в слайсе.

	slice1 := append(s[:2], 2, 3, 4)
	fmt.Println(s, slice1) // [0 0 2 3] [0 0 2 3 4]
	// 2. Берём слайс из первых двух элементов s и добавляем к ним три элемента.
	// Так как суммарная длина полученного слайса (len == 5) меньше ёмкости s[:2] (cap == 7),
	// то базовый массив остаётся прежним.
	// Слайс s тоже изменился, но его длина осталась прежней

	slice2 := append(s[1:2], 7)
	fmt.Println(s, slice1, slice2) // [0 0 7 3] [0 0 7 3 4] [0 7]
	// 3. Здесь также базовый массив остаётся прежним, изменились все три слайса

	slice3 := append(s, slice1[1:]...)
	fmt.Println(len(slice3), cap(slice3)) // 8 14
	// 4. Длина s и slice1[1:] равна 4, длина нового слайса будет равна 8,
	// что больше ёмкости базового массива.
	// Будет создан новый базовый массив ёмкостью 14,
	// ёмкость нового базового массива подбирается автоматически
	// и зависит от текущего размера и количества добавленных элементов

	// 5. Легко проверить, что slice3 ссылается на новый базовый массив
	s[1] = 99
	fmt.Println(s, slice1, slice2, slice3)
	// [0 99 7 3] [0 99 7 3 4] [99 7] [0 0 7 3 0 7 3 4]

	fmt.Println("Slice Copy")

	var dest []int
	dest2, dest3 := make([]int, 3), make([]int, 5)
	src := []int{1, 2, 3, 4}
	copy(dest, src)
	copy(dest2, src)
	copy(dest3, src)
	fmt.Println(dest, dest2, dest3, src) // [] [1 2 3] [1 2 3 4 0] [1 2 3 4]

	fmt.Println("Slice Tips")

	tipOne := []int{1, 2, 3} // Удаление последнего элемента слайса
	if len(tipOne) != 0 {    // защищаемся от паники
		tipOne = tipOne[:len(tipOne)-1]
	}
	fmt.Println(tipOne) // [1 2]

	tipTwo := []int{1, 2, 3} // Удаление первого элемента слайса:
	if len(tipTwo) != 0 {    // защищаемся от паники
		tipTwo = tipTwo[1:]
	}
	fmt.Println(tipTwo) // [2 3]

	tipThree := []int{1, 2, 3, 4, 5} // Удаление элемента слайса с индексом i
	tipCounter := 2
	if len(tipThree) != 0 && tipCounter < len(tipThree)-1 { // защищаемся от паники
		tipThree = append(tipThree[:tipCounter], tipThree[tipCounter+1:]...)
	}
	fmt.Println(tipThree) // [1 2 4 5]

	s1 := []int{1, 2, 3} // Сравнение двух слайсов
	s2 := []int{1, 2, 4}
	s3 := []string{"1", "2", "3"}
	s4 := []int{1, 2, 3}
	fmt.Println(reflect.DeepEqual(s1, s2)) // false
	fmt.Println(reflect.DeepEqual(s1, s3)) // false
	fmt.Println(reflect.DeepEqual(s1, s4)) // true

	fmt.Println("mySlice test")

	mySlice := make([]int, 100)
	for i := range mySlice {
		mySlice[i] = i + 1
	}
	fmt.Println(mySlice)
	mySlice = append(mySlice[:10], mySlice[90:]...)
	fmt.Println(mySlice)
	for i, x := range mySlice[:] {
		mySlice[i], mySlice[len(mySlice)-1-i] = mySlice[len(mySlice)-1-i], x
	}
	fmt.Println(mySlice)

	fmt.Println("mySlice not i do")

	input := "The quick brown 狐 jumped over the lazy 犬"
	// Get Unicode code points.
	n := 0
	// создаём слайс рун
	runes := make([]rune, len(input))
	// добавляем руны в слайс
	for _, r := range input {
		runes[n] = r
		n++
	}
	// убираем лишние нулевые руны
	runes = runes[0:n]
	// разворачиваем
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	// преобразуем в строку UTF-8.
	output := string(runes)
	fmt.Println(output)
}

func Maps() {
	priceCourant := map[string]int{"хлеб": 50, "молоко": 100, "масло": 200, "колбаса": 500, "соль": 20, "огурцы": 200, "сыр": 600, "ветчина": 700, "буженина": 900, "помидоры": 250, "рыба": 300, "хамон": 1500}
	fmt.Println(priceCourant)
	for k := range priceCourant {
		if priceCourant[k] > 500 {
			fmt.Println(k)
		}
	}
	orderSlice := []string{"хлеб", "буженина", "сыр", "огурцы"}
	price := 0
	for _, v := range orderSlice {
		price += priceCourant[v]
	}
	fmt.Println(price)
}

func find(arr []int, k int) []int {
	// Создадим пустую map
	m := make(map[int]int)
	// будем складывать в неё индексы массива, а в качестве ключей использовать само значение
	for i, a := range arr {
		if j, ok := m[k-a]; ok { // если значение k-a уже есть в массиве, значит, arr[j] + arr[i] = k и мы нашли, то что нужно
			return []int{i, j}
		}
		// если искомого значения нет, то добавляем текущий индекс и значение в map
		m[a] = i
	}
	// не нашли пары подходящих чисел
	return nil
	// как можно заметить, алгоритм пройдётся по массиву всего один раз
	// если бы мы искали подходящее значение каждый раз через перебор массива, то пришлось бы сделать гораздо больше вычислений
}

func RemoveDuplicates(input []string) []string {
	m := make(map[string]int)
	for i, v := range input {
		m[v] = i
	}
	output := make([]string, len(m))
	counter := 0
	for v := range m {
		output[counter] = v
		counter++
	}
	return output
}

func PrintAllFilesWithFilter(path string, filter string) {
	// получаем список всех элементов в папке (и файлов, и директорий)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("unable to get list of files", err)
		return
	}
	//  проходим по списку
	for _, f := range files {
		// получаем имя элемента
		// filepath.Join — функция, которая собирает путь к элементу с разделителями
		var filename string
		filename = filepath.Join(path, f.Name())
		// печатаем имя элемента
		if strings.Contains(filename, filter) {
			fmt.Println(filename)
		}
		// если элемент — директория, то вызываем для него рекурсивно ту же функцию
		if f.IsDir() {
			PrintAllFilesWithFilter(filename, "")
		}
	}
}

func PrintAllFilesWithFilterClosure(path string, filter string) {
	// создаём переменную, содержащую функцию обхода
	// мы создаём её заранее, а не через оператор :=, чтобы замыкание могло сослаться на него
	var walk func(string)
	walk = func(path string) {
		// получаем список всех элементов в папке (и файлов, и директорий)
		files, err := ioutil.ReadDir(path)
		if err != nil {
			fmt.Println("unable to get list of files", err)
			return
		}
		//  проходим по списку
		for _, f := range files {
			// получаем имя элемента
			// filepath.Join — функция, которая собирает путь к элементу с разделителями
			filename := filepath.Join(path, f.Name())
			// печатаем имя элемента, если путь к нему содержит filter, который получим из внешнего контекста
			if strings.Contains(filename, filter) {
				fmt.Println(filename)
			}
			// если элемент — директория, то вызываем для него рекурсивно ту же функцию
			if f.IsDir() {
				walk(filename)
			}
		}
	}
	// теперь вызовем функцию walk
	walk(path)
}

func containsDot(s string) bool {
	return strings.Contains(s, ".")
}
func PrintFilesWithFuncFilter(path string, predicate func(string) bool) {
	var walk func(string)
	walk = func(path string) {
		// получаем список всех элементов в папке (и файлов, и директорий)
		files, err := ioutil.ReadDir(path)
		if err != nil {
			fmt.Println("unable to get list of files", err)
			return
		}
		//  проходим по списку
		for _, f := range files {
			// получаем имя элемента
			// filepath.Join — функция, которая собирает путь к элементу с разделителями
			filename := filepath.Join(path, f.Name())
			// печатаем имя элемента, если путь к нему содержит filter, который получим из внешнего контекста
			if predicate(filename) {
				fmt.Println(filename)
			}
			// если элемент — директория, то вызываем для него рекурсивно ту же функцию
			if f.IsDir() {
				walk(filename)
			}
		}
	}
	// теперь вызовем функцию walk
	walk(path)
}

type figures int

const (
	square   figures = iota // квадрат
	circle                  // круг
	triangle                // равносторонний треугольник
)

func area(f figures) (func(float64) float64, bool) {
	switch f {
	case square:
		return func(x float64) float64 { return x * x }, true
	case circle:
		return func(x float64) float64 { return 3.142 * x * x }, true
	case triangle:
		return func(x float64) float64 { return 0.433 * x * x }, true
	default:
		return nil, false
	}
}

var Global = 5

func functionGlobal() {
	defer func(g int) {
		Global = g
	}(Global)
	Global = 4
	fmt.Println("2")
	fmt.Println(Global)
}
