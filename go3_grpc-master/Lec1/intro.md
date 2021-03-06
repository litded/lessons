# Тестовое задание
### На разработку API для решения квадратного уравнения

Ваше задание состоит в следующем. Вам необходимо разработать простое REST API, способное решать квадратное уравнение. Функционал выглядит следующим образом:
* ***POST /solve/a/b/c*** - запрос на добавление в БД  коэффициентов квадратного уравнения (вида ax^2 + bx + c = 0)
* ***GET /solution*** - возвращает JSON файл ПОСЛЕДНЕГО решенного уравнения. Ответ должен выглядеть следующим образом :
```
{
    "a" : value,
    "b" : value,
    "c" : value,
    "n_roots" : value
}
```
где ```"a", "b", "c"```- коэффициенты , получаемые запросом POST, а ```"n_roots"``` - количество корней уравнения с такими коэффициентами.

Ограничения:
a,b,c - целые числа
Гарантируется, что как минимум один из коэффициентов не равен нулю.

БД:
Локальная (слайс/мапа) или postgres на выбор.