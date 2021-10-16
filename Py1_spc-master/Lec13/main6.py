"""
Типы параметров в явных функциях Python
"""

def add(a:int, b:int, c:int):
    """
    Необходимо-позициональный параметр
    Функция с Required Positional Args
    """
    return a ** 2 + b ** 3 + c ** 4

print(add(2,3,4))
print(add(3,4,2))
print(add(4,3,2))
print(add(4, 3, 3))

def mult(a:int=10, b:int=20, c:int=30):
    """
    Функция с аргументами со значениям по умолчанию
    Default value args
    """
    return a + b + c

print(mult())
print(mult(1))
print(mult(1,2))
print(mult(1,2,3))



def sub(a:int, b:int, c:int=10):
    """
    Функция с комбинированными аргументами
    ОСОБЕННОСТЬ: default value args всегда идут ПОСЛЕ позициональых
    аргументов
    """
    return (a/b ) ** c 

print(sub(10, 20))
print(sub(10, 20, 3))
