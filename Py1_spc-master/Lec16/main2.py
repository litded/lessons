"""
Работа с json
"""
import json 
# Java Script Object Notation

# Сериализация в JSON - преобразование объекта в архитектуре языка в байтовую последовательность,
# с последующим превращеним в JS синтаксис.

# Десериализация из JSON - сериализация наоборот.

# Задача: сериализовать data
data = {
    "name" : "Bob",
    "lastname" : "London",
    "age" : 23,
    "balance" : 12400.5,
    "sex" : True ,
    "items" : ["item1", "item2", "item3"],
    "hasCar" : None ,
    "previous" : {
        "time" : 123123,
        "users" : ["user1", 'user2', 'user3'],
    }
}

with open("data.json", "w") as fhand:
    json.dump(data, fhand, indent=4)