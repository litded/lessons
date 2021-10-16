"""
Множество set() - неупорядоченная (неиндексируемая) коллекция,
способная хранить элементы ЛЮБОГО типа, при этом УНИКАЛИЗИРУЯ их.
"""

# Создадим пустое множество
empty = set()
print("empty:", empty)
# print("Len(empty):", len(empty))
# print("type(empty):", type(empty))

# Создадим множество с элементами
values = set([ 10,"Bob", "Bob", 20, 30, "Bob", "Alice", 40,10, 10, 10, "George"])
print("values:", values)
# print("len(values):", len(values))
# print("type(values):", type(values))

# Добавление элементов в о множество
new_empty = set()
word1 = "Alice"
word2 = "Bob"

new_empty.add(word1) # .add() - метод для добавления элемента в множество
new_empty.add(word2)
new_empty.add(word1)
new_empty.add(word2)

print("new_empty:", new_empty)

# Проверка, есть ли элемент во множестве?
characters = set(["Char1", "Char2", "Char3", "Char4"])
predict_char = "Char3"

# Супер классификатор для выявления победы. 1 - выиграл, 0 - проиграл
if predict_char in characters:
    print(1)
else:
    print(0)










