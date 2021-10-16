"""
Работа с User IO/(Input/Output) (с консольным вводом-выводом).
Input (Stream) - входной поток - поток, откуда программа ПОЛУЧАЕТ данные
Output (Stream) - выходной поток - поток, КУДА ВЫВОДЯТ данные
"""

"""
Рассмотрим базовые принципы работы с выходным потоком на примере работы команда print()
"""
print("Hello world!")
"""
1-ая особенность - по завершению работы команда print() выполняет перенос на новую строчку автоматически!
"""
print("First line") # "First line" + '\n'
print("Second line")

"""
2-ая - особенность - данные могут быть сохранены в переменных
"""
message = "My message"
print(message)

"""
3-ая особенность - при перечислении данные выводятся разделенные пробельным символом
"""
name = "John"
lastname = "Brown"

# Использую перечисление
print(name, lastname) # print("John", "Brown")
# Использую перечисление с добавлением литерала
print("Имя:", name, "и", "Фамилия:", lastname) # print("John", "и", "Brown")


"""
4-ая особенность (не совсем про print()) - двойные кавычки равносильны одинарным
"""

new_name = "Lev"
new_lastname = 'Tolstoy'
book_title = "Title : 'War and Peace' at .... year"

print(new_name, "и", new_lastname, "Wrote:", book_title)


