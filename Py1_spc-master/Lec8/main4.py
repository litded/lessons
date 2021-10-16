"""
Операции над двумя множествами
"""

a_set = set([1,2,3,4])
b_set = set([3,4,5,6])

# Пересечение
intersect = a_set.intersection(b_set) 
print("Intersect:", intersect)

# Объединение
uni = a_set.union(b_set)
print("Union:", uni)

# Разность
ab = a_set.difference(b_set)
print("A - B:", ab)

ba = b_set.difference(a_set)
print("B - A:", ba)

# Симметрическая разность
sdiff = a_set.symmetric_difference(b_set)
print("A^B:", sdiff)