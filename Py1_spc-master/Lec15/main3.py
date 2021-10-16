"""
Пакет - это набор модулей, написанных на языке Python.
Пакет может включать в себя только модули и другие пакеты.

Для того, чтобы подсказать интерпретатору, что перед ним пакет (а не случайная директория)
всегда помещайте в пакет модуль __init__.py . Данный модуль будет выполняться при первом обращении к пакету
. Чаще всего в него помещают логику по подготовке окружения к работе функционала пакет.

Воспользуемся пакетом geometry
"""
from geometry import rectangle , circle 
#from geometry import circle 

print("Perimeter(10,20):", rectangle.perimeter(10, 20))
print("Area(10, 20):", rectangle.area(10, 20))
print("Perimeter(10):", circle.perimeter(10))
print("Area(10):", circle.area(10))

"""
Для использования совмещенных имен
"""
from geometry.rectangle import perimeter
from geometry.circle import area 

print("Union namespaces")
print("Perimeter(10, 20):", perimeter(10, 20))
print("Area(10):", area(10))


"""
Дотянуться до foobuzz
"""
from geometry.foo.buzz import hello
hello()