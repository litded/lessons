"""
Кортежи используют для передачи.
Списки используют для обработки.
"""
# Блок обработки
nums = [10, 20, 30, 40]
nums.append(-10000)
nums.sort()
# Блок передачи
nums_tpl = tuple(nums) # Явное преобразование коллекций
print("Type:", type(nums_tpl))
# ..........
# Блок обработки
nums_lst = list(nums_tpl)
nums.append(10000000)
print("Type:", type(nums_lst))