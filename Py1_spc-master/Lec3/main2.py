"""
Логический тип bool()
"""
a_bool = True
b_bool = False

# Логические операции
print("and:", a_bool and b_bool)
print("or:", a_bool or b_bool)
print("not:", not a_bool)
print((a_bool and (not b_bool)) or (True and True))

# Арифметические операции провоцируют логический тип НЕЯВНО приводится к целым числам
# True -> 1
# False -> 0
print(a_bool + a_bool ** b_bool)