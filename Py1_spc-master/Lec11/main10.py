"""
D7 : D solution
"""

n = int(input())
birth_book = {} # {"month1" : ["name1", "name2", ....]}

for _ in range(n):
    message = input() # "Витя 12 дек"
    infos = message.split() # ["Витя", "12", "дек"]
    name = infos[0]
    month = infos[2]

    if month in birth_book.keys():
        birth_book[month].append(name)
    else:
        birth_book[month] = [name]

print(birth_book)