"""
Функции в Python - это объекты 1-го рода.
Это значит, что :
* функции можно присваивать переменным, +
* функции можно помещать в коллекции, +
* функции можно возвращать в качестве возвр.значения, +
* функции можно принимать в качестве параметра, +
* функции можно определять где угодно +
"""

def add(a:int, b:int):
    """
    a + b 
    """
    return a + b 

def sub(a:int, b:int):
    """
    a - b 
    """
    return a - b 

# 1) Присваивание переменным
varaible_add = add 
varaible_sub = sub 

print("Variable_add(10, 20):", varaible_add(10, 20))
print("Variable_sub(20, 10):", varaible_sub(20, 10))

# 1.1) Тип переменных с функциями
print("Type:", type(varaible_add))

# 1.2) Значение переменных с функцией
print(varaible_add)

# 1.3) Размер в памяти
print("Varaible_add size:", varaible_add.__sizeof__(), "Bytes")

# 1.4) Тоже самое справедливо для lambda
variable_lambda = lambda x, y: x * y 

# 2) Помещение функций в коллекции
funcs_list = [add, sub, variable_lambda]
for func in funcs_list:
    print(func(20, 30))


# 3) Возврат функции в качестве значения другой функции
def born_func(sign="+"):
    if sign == "+":
        return lambda x, y: x + y 
    elif sign == "-":
        return lambda x, y: x - y 

child_plus = born_func(sign="+") 
print("Child+ make:", child_plus(10, 20))
child_minus = born_func(sign="-")
print("Child- make:", child_minus(10 ,20))


# 4) Передача функции в качестве параметра в другую функцию
def outer_func(inner_func):
    print("Before call innerfunc")
    inner_func()
    print("After call innerfunc")

def hello():
    print("Hello!!!!!!")

outer_func(hello)


# 5) Можно ли определить одну функцию в теле другой?
print("=========================================================")
def outer():
    def inner():
        print("INNER FUNCT!")
    inner()

outer()
