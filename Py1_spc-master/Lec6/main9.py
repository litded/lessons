"""
D3:K Pre-solution
"""

current_max = None # Максимальный пульс ОТОБРАННОГО кандидата
current_min = None  # Минимальный  пульс ОТОБРАННОГО кандидата
count = 0 # Сколько людей вообще отобрано
is_first = True # Флаг для первого отобранного человека (позволит определить мин и макс)

left_side = 100 
right_side = 140

while True:
    pulse = float(input())
    if pulse < 0:
        break 

    if pulse >= left_side and pulse <= right_side:
        # Пульс данной персоны прошел отбор
        count += 1
        # Вычислить мин и макс для отобранных персон
        """
        Code here
        """

print(count)
print(current_min, current_max)

