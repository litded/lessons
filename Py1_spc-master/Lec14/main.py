"""
Про анонимные функции
"""
def add(a:int, b:int):
    """
    This func for ......
    """
    return a + b 


# Синтаксис анонимных функций (лямбда-выражения)
myfunc = lambda a, b : a + b 

print(add(10, 20))
print(myfunc(10, 20))
