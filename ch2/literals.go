package main

import "fmt"

func literals() {
	fmt.Println("Literatures")

	// Знаковые целые числа
	var signedInt8 int8 = 127
	var signedInt16 int16 = 32_767
	var signedInt32 int32 = 2_147_483_647
	var signedInt64 int64 = 9_223_372_036_854_775_807
	// На копмьютерах с 64 битной архитектурой int это 64 бита,
	// на 32 битной архитектуре int это 32 бита

	// Беззнаковые целые числа
	var unsignedInt8 uint8 = 255 //byte is an alias for uint8
	// Байт это псевдоним для uint8, поэтому лучше использовать byte вместо uint8
	// Байт и uint8 могут использоваться друг с другом в логических операциях

	var unsignedInt16 uint16 = 65_535
	var unsignedInt32 uint32 = 4_294_967_295
	var unsignedInt64 uint64 = 18_446_744_073_709_551_615
	//Значение по умолчанию всех целых чисел = 0

	//				======= Рекомендации по применению ======
	// Если работаем с файлами двоичного формата или сетевым протоколом которые
	// используют целые числа определённой размерности или знака, выбираем соответствующий int тип.

	// Если пишем библиотечную функцию, которая должна работать с любым целочисленным типом,
	// пишем две функции, одна из которых будет использовать для параметров тип int64,
	// а другая uint64.

	// Стандартные +,-,*,/,% работают с любыми  целыми числами
	// Побитовые &,^,|,<<,>> операции работают с любыми целыми числами.
	// Стандартные операции сравнения <,>,<=,>=,==,!= работают с любыми целыми числами.

	//				====== Числа с плавающей точкой ======
	// float32 - 7 знаков после запятой, float64 - 15 знаков после запятой
	// float32 - 4 байта, float64 - 8 байт
	// Все стандартые операции работают с любыми числами с плавающей точкой кроме % и <<,>>
	// Деление на 0 возвращает +Inf, -Inf, деление 0 на 0 = NaN (Not a Number)

	// Лучше не сравнивать числа с плавающей точкой на равенство с помощью операторов == и !=
	// Лучше определить максимально допустимое отклонение(эпсилон) и посмотреть не превышает ли его
	// разница между числами

	// 				====== Руны ======
	var myRune rune = 'A' // rune - это псевдоним для int32
	// rune - это 4 байта, 32 бита, 1 символ в UTF-8 кодировке

	fmt.Println("Инты: ", signedInt8, signedInt16, signedInt32, signedInt64, unsignedInt8, unsignedInt16, unsignedInt32, unsignedInt64)
	fmt.Println("Руны: ", myRune)

	// 				====== Строки ======
	// Строка - это последовательность байтов (рун), которая заканчивается нулевым байтом
	// Строки в Go неизменяемые, то есть нельзя изменить строку после её создания
	var myString string = "Hello, World!"
	fmt.Println("Строка: ", myString)

	//				====== Преобразование типов ======
	// Преобразование типов в Go неявное
	// То есть нельзя преобразовать int в float без явного указания
	// Всегда нужно явно указывать преобразование типов

	var myInt32 int32 = 10
	var myInt64 int64 = 20

	var sum int64 = myInt64 + int64(myInt32) // явное преобразование типов
	fmt.Println("Сумма: ", sum)
}

// В Go отсутсвует перегрузка функций
func MyFunctionInt(a int64) {
	fmt.Println("I AM DOING SOMETHING ", a)
}

// Ошибка статического анализатора кода и компилятора

// func MyFunctionUint(a uint64) {
// 	fmt.Println("I AM DOING SOMETHING ", a, b)
// }
