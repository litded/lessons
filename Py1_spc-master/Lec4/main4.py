"""
Оператор множественного выбора
"""
age = int(input())

if age < 14:
    print("Доступ закрыт.")
elif age >= 14 and age < 18:
    print("Неплатежеспособный")
    print("Демо-версия")
elif age >= 18 and age < 25:
    print("Стандарт")
else:
    print("Привелегии")

"""
Вложенный условный оператор
"""

age = 15
weight = 50

if age > 10 and age < 20:
    print("COOL")
    if weight < 60:
        print("Very cool!")
    print("COOL DONE!")

