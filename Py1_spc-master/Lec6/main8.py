"""
Пользователь вводит несколько чисел (натуральные) , до тех пор копа не будет введено
первое отрицательное целое число. Требуется найти минимальный и максимальный из введенных элементов
"""

current_max = None
current_min = None 

is_first = True # Отвечает на вопрос - это первое вводимое число, или нет

while True:
    num = int(input())
    if num < 0:
        break 

    if is_first:
        current_max = num 
        current_min = num 
        is_first = False

    if num > current_max:
        current_max = num 
    
    if num < current_min:
        current_min = num 

print("Current max:", current_max)
print("Current min:", current_min)