"""
Удаление элементов из списка.
"""

nums = [10, 20, 30, 40, 50, -1, 1000]
print("Origin:", nums)
# Удаление по индексу
"""
Для удаления по индексу можно воспользоваться двумя способами:
1) Если понадобится удаленное значение получить и продолжить с ним работу,
то используй .pop(id)
2) Если удаленное значение не нужно получать, можно восопльзоваться del
"""
deleted_num = nums.pop(0)
print("After .pop(0):", nums)
print("Deleted elem:", deleted_num)

last_elem = nums.pop() # По умолчанию удаляет последний элемент
print("After .pop():", nums)
print("Last elem:", last_elem)

"""
Мораль. Перед удалением по индексу убедитесь что этот индекс есть в списке
"""
deleted_id = 3
if len(nums) > deleted_id:
    print("Can delete")
    nums.pop(deleted_id)
    print("After deleting:", nums)

print("Before del:", nums)
del nums[2] # Оператор передачи объекта сборщику мусора ("деструктор")
print("After del:", nums)

print("==========================================================")
# Удаление по значению
"""
Для, удаления элемента по значению нужно пользоваться методом remove()
1) Удаляется первое вхождение элемента в список
2) Перед удалением нужно убедиться, что существует хотя бы один такой
элемент в списке
"""

names = ["Alice", "Bob", "Jack", "George", "Bob"]
name_to_delete = "Bob"
print("Before deleting:", names)
names.remove(name_to_delete)
print("After deleting:", names)

# Проверка перед удалением
if "Jack" in names:
    names.remove("Jack")
    print("After safe delete:", names)