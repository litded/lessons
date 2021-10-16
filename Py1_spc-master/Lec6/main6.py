"""
Подсчет произведения значений
"""

mult_result = 1 # Накапливает произведение вводимых значений

while True:
    num = int(input())
    if num < 0:
        break
    else:
        mult_result *= num 

print("Total multiply:", mult_result)