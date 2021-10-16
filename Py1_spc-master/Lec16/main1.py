"""
Запись в файл
"""
message = "This is my message!"

# "w" - перезапись файла
# "a" - добавление в конец файла
with open("output.txt", "a") as fhand:
    fhand.write("\n" + message)