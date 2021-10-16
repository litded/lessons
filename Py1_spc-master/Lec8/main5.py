"""
D5 : A solution
"""

to_park = set() # из дома в парк
to_home = set() # из парка в дом

# Наполняем множество по пути от дома до парка
while True:
    current_bus = input()
    if current_bus == "":
        break 
    else:
        to_park.add(int(current_bus))

# Наполняем множество по пути из парка в дома
while True:
    current_bus = input()
    if current_bus == "":
        break
    else:
        to_home.add(int(current_bus))

total = to_park.intersection(to_home)
print(len(total))