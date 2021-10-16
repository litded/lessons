"""
Словарь - это множество пар КЛЮЧ:ЗНАЧЕНИЕ
"""
# Инициализация пустого словаря
famous = dict()
famous = {}

# Инициализация словаря с элементами
famous = {1952 : "Fortran", 2015 : "Kotlin release"}
print("Famous dictionary:", famous)

# Обращение к элементам - так же как и у индексируемых коллекций. Только вместо индекса
# ключ
print("Key:", 1952, "Value:", famous[1952])
print("Key:", 2015, "Value:", famous[2015])

# Добавление новых ПАР в словарь famous
famous[1997] = "Python release 1.0"
famous[1998] = "Java release"
famous[1965] = "Objective C"
famous[2010] = "GoLang release"
famous[2014] = "Rust release"

print("Total dict:", famous)



