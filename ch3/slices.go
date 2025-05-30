package main

import (
	"fmt"
	"slices"
)

func IReallyNeedToFindAWayToNameAnotherMain() {
	fmt.Println("Slices")

	// Срезы в Go - это динамические массивы
	// Они могут изменять свой размер во время выполнения программы.
	// Их длина не является составной частью типа, как в случае с массивами.

	// Инициализация среза | Важно подметить []int - срез, [...]int - массив
	s1 := []int{10, 20, 30} // Слайс(срез) с 3 элементами

	// анлгично массиву можно создать срез нулевыми индексами или указать какие элементы не нулевые
	s2 := []int{0: 10, 1: 20, 2: 30}
	s3 := []int{1, 5: 4, 6, 10: 100, 15} // [1, 0, 0, 0, 0, 4, 6, 0, 0, 0, 100, 15]

	// Имитация многомерного среза
	var s4 [][]int
	// Для чтения и записи элементов используются [], прямо как и в массивах
	// Очевидно, надеюсь, что и выходить за пределы среза не получится и будет ошибка
	// либо компиляции, либо рантайма

	// Инициализация среза без использования литерала
	var s5 []int
	// нулевым значением среза является nil, а не пустой массив
	// nil - это специальное значение, которое указывает на отсутствие значения.
	// как не типизированные числовые константы, значение nil не имеет типа,
	// что позволяет присваивать или сравнивать его с любым типом данных.
	// Срез равный nil не содержит элементов.
	// Срез является не сравнительным типом. Попытки узнать равны ли два среза приведут к ошибке
	// компиляции. Их можно сравнивать только с nil.

	fmt.Println(s1, s2, s3, s4, s5)

	// С версии 1.21 пакет slices стандартной библиотеки включает в себя две функции
	// для сравнения срезов: Equal и EqualFunc.
	// 1-я принимает два среза и возвращает true, если у них одинаковая длина и элементы равны.
	// Для этого требуется, чтобы элементы срезов были сравнимыми.
	// 2-я позволяет передать функцию для определения равенства, и не требует,
	// чтобы элементы были сравнимыми. Она возвращает true, если все элементы равны.

	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}
	s := []string{"a", "b", "c"}

	fmt.Println("Equal:", slices.Equal(x, y)) // true
	fmt.Println("Equal:", slices.Equal(x, z)) // false
	// fmt.Println("Equal:", slices.Equal(x, s)) // не компилируется, т.к. разные типы

	// До этой версии Go для сравнения срезов использовался пакет reflect, функция DeepEqual.
	// Это предназначается для тестирования, а не для сравнения срезов в реальных приложениях.
	// НЕ НУЖНО ЕЁ ИСПОЛЬЗОВАТЬ, она медленнее и менее безопасна
	fmt.Println(s)

	// Встроенная функция len возвращает длину среза. Если срез равен nil, то len(s) == 0.
	// len существует, так как выполняемые действия невозможно осуществить с помощью функций,
	// написанных обычным образом.

	// Встроенная функция append используется для добавления элементов в срез.
	// Она принимает срез и один или несколько элементов, которые нужно добавить.
	// var x []int
	// x = append(x, 1) // добавляет 1 в конец среза x
	// Принимает минимум два параметра: срез и элемент, который нужно добавить.
	// Возвращает новый срез, тип которого равен типу первого параметра.
	// За один раз можно добавить несколько элементов, передав их через запятую.

	// Один срез добавляется в другой срез с помощью ... (троеточие).
	// y := []int{1, 2, 3}
	// x = append(x, y...) // добавляет элементы y в конец среза x

	// Написать append без переменной, которая будет хранить результат - нельзя, ошибка компиляции.
	// В Go используется передача по значению, а не по ссылке.
	// Когда мы передаём параметр в функцию, Go создаёт копию передаваемого значения.
	// Поэтому, если мы просто добавим элемент в срез, то он не изменится, т.к.
	// append получает копию среза, а затем создаёт новый срез и возвращает его,
	// а не изменяет существующий.

	// Для создания пустого среза
	// (который не равен nil, и все элементы проинициализированы стандартными значениями)
	// используется встроенная функция make.
	// Она принимает три параметра: тип, длину и ёмкость.

	b := make([]int, 5) // создаёт срез длиной 5 и ёмкостью 5 (т.к. не указана явно ёмкость)
	fmt.Println(b)      // [0 0 0 0 0]
	// Частая ошибка после создания среза с помощью make - попытка добавить элементы в него,
	// через append.

	// b = append(b, 1, 2, 3)
	// в этом случае у нас итак есть 5 элементов, и мы добавляем ещё 3 элемента В КОНЕЦ
	// поэтому после make достаточно обращаться к элементам через [индекс]
	// append только если необходимо увеличить и добавить в конец элемент

	// Можно например создать срез у которого длина 0 а вместимость 10.
	// Тогда у нас будет срез не равный nil, и при этом не имеющий даже нулевых значений
	// Таким образом в него можно будет добавлять элементы через append
	// (Я хз что быстрее, создать срез с нулевыми значениями и изменять индексы, или
	// создать срез определённой ёмкости и добавлять элементы через append)

	// В версии Go 1.21 добавлена встроенная функция clear, которая принимает срез,
	// устанавливает значения всех элементов в дефолтное/нулевое значение.
	// Длина среза остаётся той же

	// Срез среза
	// Обозначается как
	zz := z[1:3] // ":" символ среза среза (не опечатка), первое число индекс
	// ВКЛЮЧАЕМОГО элемента с которого начинается срез
	// второе число, индекс НЕ ВКЛЮЧАЕМОГО элемента до которого идёт срез
	// то есть здесь у нас новый срез с первого элемента (который включён)
	// до третьего элемента (который не включён, включён только предыдущий)

	// Также можно писать [:3] - если не указать первое число, то оно по стандарту 0
	// Если не указать второе число [1:], то оно будет до длины

	// ВАЖНО: при создании среза от среза, НЕ СОЗДАЁТСЯ КОПИЯ
	// срезы на основе среза лишь ссылаются на основной срез (bro wtf)
	// поэтому использование функции append может привести к неожиданному поведению
	// для избежания таких ситуаций, нужно использовать полное выражение среза
	// [startIndex:EndingIndex:Capacity] - [0:5:10] - 0 стартовый индекс, 5 - индекс конца
	// 10 - Вместимость
	// таким образом использование функции append приведёт к созданию нового независимого среза

	// Функция Copy
	// Если нужно создать срез чтоб он не зависел от исходного среза, то можно использовать
	// Copy
	// Принимает 2 параметра: 1)куда копируем срез 2)исходный срез
	// Копируется из исходного в целевой максимально возможное кол-во элементов,
	// которое определяется размером меньшего среза, и возвращает кол-во скопированных эл-ов.
	// Ёмкость срезов не имеет значения, важна лишь их длина.
	// Можно скопировать подмножество среза (Которое через ":")
	// Можно скопировать так:
	xx := []int{1, 2, 3, 4}
	yy := make([]int, 2)
	copy(y, x[2:])
	// Результат функции copy никуда не присваивается
	// Если не требуется кол-во скопированных переменных, то его можно не присваивать переменной

	// x := []int{1,2,3,4}
	// num := copy(x[:3],x[1:])
	// В данном случае последние три элемента среза X копируются в первые три элемента
	// этого среза
	// В итоге там будет [2 3 4 4] а в num будет 3

	// Функцию copy можно использовать и с массивами, путём взятия среза в массиве
	// Массив может выступать и в качестве источника и в качестве цели копирования

	// МАССИВЫ МОЖНО ТОЧНО ТАКЖЕ ПРЕОБРАЗОВЫВАТЬ В СРЕЗЫ
	// ТОЧНО ТАКЖЕ, СРЕЗ МАССИВА БУДЕТ ССЫЛАТЬСЯ НА МАССИВ, Т.Е. ПАМЯТЬ ОБЩАЯ
	fmt.Println(xx, yy, zz)

	// СРЕЗЫ МОЖНО ПРЕОБРАЗОВАТЬ ОБРАТНО В МАССИВЫ (хспде это такая огромная тема, как тут всё запомнить то)

	xSlice := []int{1, 2, 3, 4}
	xArray := [4]int(xSlice) // Преобразование типов, в данном случае слайс в массив
	smallArray := [2]int(xSlice)
	xSlice[0] = 10
	fmt.Println(xArray, xSlice, smallArray)
	// xSlice [10 2 3 4]
	// xArray [1 2 3 4]
	// smallArray [1 2]

	// Размер массива должен быть известен во время компиляции
	// Использование [...] в преобразовании типа среза в тип массив приводит к ошибке во время компиляции
	// Размер массива может быть меньше среза, но не может быть больше.
	// Если указать размер массива превышающий ДЛИНУ среза, то будет ошибка во время компиляции

	// panicArray := [5]int(xSlice)
	// fmt.Println(panicArray) // Ошибка во время рантайма

}
