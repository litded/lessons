"""
Работа с модулем стандартной библиотеки
Разделенная область имен
"""
import math  as m 


print("Pi:", m.pi)
print("E:", m.e)
for i in range(1, 10):
    print("n=", i, "n!=", m.factorial(i))

"""
Импортирование с совмещением имен
"""
from math import pi, e, factorial
print("Pi:", pi)
print("E:", e)
for i in range(1, 10):
    print("n=", i, "n!=", factorial(i))