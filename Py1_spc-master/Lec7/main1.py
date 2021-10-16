"""
Слова break/continue работают точно так же , как и у while
"""

for i in range(0, 100000000000000000000):
    if i > 20:
        break 
    else:
        print("Elem:", i)

"""
Вложенные циклы - достаточно эффективны в задачах полного перебора
"""
size =  10
for i in range(1, size + 1):
    for j in range(1, size + 1):
        print(i, "x", j ,"=", i * j)