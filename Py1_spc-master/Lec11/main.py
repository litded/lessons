"""
Кортеж - упорядоченная (индексируемая) коллекция, споосбная хранить элементы ЛЮБОГО типа,
при этом НЕ ИЗМЕНЯЕМАЯ.
"""

# Инициализация пустого кортежа
a_tuple = () 
b_tuple = tuple()

# Инициализация кортежа с элементами
c_tuple = (10, 20, 30, 40, 50 , "Bob", "Vito")
print("Tuple:", c_tuple, "Type:", type(c_tuple))

# Инициализация кортежа с одним элементом
tpl = (1, ) # Постановка запятой показывает , что речь идет о перечислении
print("Tpl:", tpl, "Type:", type(tpl))

# Проверка вхождения
if 10 in c_tuple:
    print("10 in ", c_tuple)

# Итерирование через индексы
for i in range(len(c_tuple)):
    print("Id:", i, "Elem[id]:", c_tuple[i])

#  Итерирование по значениям
for elem in c_tuple:
    print("Elem:", elem)

# Поскольку кортеж - индексируемый - ему доступны операции срезов и конкатенации
slice_1 = c_tuple[2:5]
print("Slice:", slice_1, "Type:", type(slice_1), "Len:", len(slice_1))


concat = c_tuple + slice_1 * 10
print("Total concat:", concat)
