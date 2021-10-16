"""
Замыкание (closure) - способ конфигурации функций, при сокрытии реализаций для других пользователей
Замыкание негласно называют "функциональной инкапсуляцией".
"""

def add_outer(a:int):
    """
    Внешняя функция - параметр a конфигурирует внутренню функцию
    """
    def add_inner(b:int):
        """
        Внутренняя функция - знает про a  из вышестоящей области, принимает b
        """
        return a + b 

    return add_inner

add10 = add_outer(10)
#def add_inner(b:int):
#    return 10 + b
print("Add10(7):", add10(7))

add15 = add_outer(15)
#def add_inner(b:int):
#    return 15 + b 
print("Add15(7):", add15(7))
