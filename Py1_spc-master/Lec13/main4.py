"""
Параметры/аргументы явных функций.

Аннотации параметров функций - РЕКОМЕНДАЦИЯ по использованию явной функции. Аннотировать можно ЛЮБОЙ 
стандартный тип в языке Python.

Docstring - строка с документацией кода. Чаще всего, docstring сопровождает ВСЕ явные функции.
"""
# add: a - int, b - int
def add(a :int, b :int):
    """
    Это строка документации для функции add:
    a :int -  целочисленный параметр а
    b :int - целочисленный параметр b
    Возвращаемое значение: результат суммы а + b
    """
    result = a + b 
    return result 

# arg_a, arg_b = int(input()), int(input())
res1 = add(10, 10)
print("Res1:", res1)
print(add("Alice", "Bob"))
