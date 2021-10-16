"""
Строковые методы. Базовые.
"""

message = "hElLo WoRlD"

# Перевод строки в верхний регистр
up = message.upper() 
print(".upper():", up)

# Перевод строки в нижний регистр
low = message.lower()
print(".lower():", low)

# Выделение заглавной буквы
cap = message.capitalize()
print(".capitalize():", cap)

"""
В результате применения строковых методов - исходная строка НЕИЗМЕННА
"""
print("Origin:", message)
#message[3] = "Q"