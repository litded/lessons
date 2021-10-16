"""
Общие операции над словарями.
Записать в словарь можно в любой момент (достаточно указать явно пару)
Считывать из словаря можно только тогда, когда ключ для считывания уже существует
"""
famous = {1952 : "Fortran", 2015 : "Kotlin release"}
famous[1997] = "Python release 1.0"
famous[1998] = "Java release"
famous[1965] = "Objective C"
famous[2010] = "GoLang release"
famous[2014] = "Rust release"

# Подсчет количества элементов коллекции
print("Len:", len(famous))

# Проверка вхождения (в ключи)
if 2010 in famous.keys():
    print("2010 in keys!")

# Проверка вхождения в значения
if "Objective C" in famous.values():
    print("Objective C in values!")

# Перебор элементов словаря (по ключам)
for key in famous.keys():
    print("Key:", key, "Value:", famous[key])

# Перебор элементов по значению
for val in famous.values():
    print("Value:", val)

# Перебор пар
for key, val in famous.items(): # [(key1, val1), (key2, val2), (key3, val3)]
    print("Key:", key, "Val:", val)
