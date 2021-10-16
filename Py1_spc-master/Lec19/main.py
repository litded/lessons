"""
Задача - 100000000000000 раз проитерироваться при помощи цикла for
"""
# 1-ый способ
print("range()")
print("range(1,1000000001) size:", range(1, 1000000001).__sizeof__(), "Bytes")
for i in range(1, 1000000001):
    print(i)
    break 

# 2-ой способ
print("list()")
nums = [0] * 1000000000
print("nums size:", nums.__sizeof__(), "Bytes")
for i in nums:
    break 

"""
Проблема №1 - список с большим количеством элементов впамять не вмещается, 
а range() с таким же количеством - умещается.
"""