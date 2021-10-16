"""
1) int() float() -> числовые типы, неявно приводит к float() целые числа
2) str() -> строковый тип, в языке нет типа char()
3) bool() -> логический тип, неявно приводится к int() если есть арифметика
4) NoneType() -> это тип "без значения". Нужен для того, чтобы отслеживать
возврат из функций.


Общее свойство всех базовых типов - они неизменяемы.
"""

a_int = 10
print("Old value:", a_int)
a_int = a_int + 20
print("New value:", a_int)


"""
Динамическая типизация, это когда вот так
"""
a_str = "Message"
a_str = 10

# Строгая типиация (С++)
# string name = "Message";
# name = 20;