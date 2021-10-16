"""
Цикл for - range based for - цикл, основанный на диапазоне
Это значит, что для работы этого цикла - необходимо жестко задать границы итерирования.
"""

# 1) Переменная цикла определяется на этапе инициализации условного выражения внутри
# цикла
# 2) range() - "генератор диапазонов"
for i in range(0, 10, 1): # range(start, stop[, step)
    print("Elem:", i)


"""
Шаг по умолчанию
"""
print("step = 1")
for j in range(0, 10): # range(start, stop, step=1)
    print("Elem:", j)

"""
Начало по умолчанию
"""

print("start = 0")
for k in range(10): # range(start=0, stop, step=1)
    print("Elem:", k)

"""
start, stop,step - ЦЕЛОЧИСЛЕННЫЕ ЗНАЧЕНИЯ
"""






