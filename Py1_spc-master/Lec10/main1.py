"""
Список (list) - упорядоченная (индексируемая) коллекция, способная хранить элементы ЛЮБОГО типа.
"""

# Инициализация пустого списка
a_lst = []
b_lst = list() 

# Инициализация с элементами
nums = [10, 20, 30, 40, 2, 4, 8, 7, "Hello", False]
print("Type:", type(nums), "Len:", len(nums), nums)

# Проверка вхождения элемента
if 10 in nums:
    print("10 in ", nums)

# Перебор элементов циклом for через индекс
for i in range(len(nums)):
    print("Id:", i, "nums[Id]:", nums[i])

# Перебор элементов напрямую
for elem in nums[1:5]:
    print(elem)