"""
Работа с файлами.
"""
# fhand = open("input.txt", "r")
# print("Descriptor:", fhand)
# content = fhand.read()
# print(content)
# fhand.close()

with open("input.txt", "r") as fhand:
    content = fhand.read() # Считывание файла целиком в виде одной строки
    content2 = fhand.read() # ""
    print(content)
    print(content2)


with open("input.txt", "r") as fhand:
    lines = [line.strip() for line in fhand.readlines()]
    print(lines)

with open("input.txt", "r") as fhand:
    line = fhand.readline() # "Hello world\n"
    print(line)
    line2 = fhand.readline()
    print(line2)