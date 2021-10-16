"""
Сортировка списка.
"""

nums = [20, 10, -1, 2, 3, -100, 4, 500, 6]

# Сортировка копии
new_nums = sorted(nums, reverse=True)
print("Origin:", nums)
print("Sorted copy:", new_nums)

# Сортировка с изменением первоначальной структуры (in-place )
nums.sort(reverse=True)
print("after .sort():", nums)

