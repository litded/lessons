"""
Динамическое наполнение <= Список - изменяемый объект
"""

nums = []

# Конкатенация с другим списком
old_nums = [10, 20, 30]
nums = nums + old_nums

print("Concat:", nums)

# Добавление в конец списка (.append())
val = 20000
val2 = 30000
nums.append(val)
nums.append(val2)
print("appended:", nums)

# Добавление в произвольное место (.insert())
val3 = -40000000
nums.insert(2, val3)
print("inserted:", nums)
