"""
Строковый тип str()
Строка - последовательность символов , заключенная в кавычки.
"""
message = 'Hello world!'
word = "this"
letter = 'a'
empty = ""

# Конкатенация
word1 = "Hello"
word2 = 'word'
total = word1 + " " +  word2 + "!"
print(total)
print("Мое имя " + "Боб!")

word3 = "Na"
multy_concat = word3 * 10  # word3 + word3 + word3 + ...
print(multy_concat + " Batman")

# Подсчет длины строки
print("Len:", len(message))

