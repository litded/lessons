"""
Подсчет суммы вводимых значений
"""
summ = 0 # Переменная-счетчик - накапливает сумму вводимых значений

while True:
    num = int(input())
    if num < 0:
        break 
    else:
        summ += num 

print("Sum:", summ)
