"""
Строка - упорядоченная (индексируемая) коллекция, способная хранить элементы ОДНОГО типа.

Конкатенация осуществима только тогда, когда у коллекции есть голова и хвост (первый и последний элемент).
"""

message = "Hello world!"
print(message)
# Индексация
print("First letter:", message[0]) # Индекс самого первого элемента строки - 0
print("Second letter:", message[1]) # Индекс элемента, стоящего после `H`
print("Last letter:", message[len(message) - 1]) # last_id = len(str) - 1
# Существует упрощение
print("Last letter:", message[-1])
print("Prelast message:", message[-2])

# Проверка вхождения
if "ello" in message:
    print("YES")
else:
    print("NO")

# Перебор элементов строк
for i in range(0, len(message), 1):
    print("Id:", i, "Message[Id]:", message[i])
