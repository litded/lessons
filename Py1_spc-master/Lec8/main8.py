"""
D5 : E solution
"""

total = set() # Множество студентов, которые были во все дни семестра на данный момент
n_days = int(input()) # Количество дней в семестре

for i in range(n_days):
    # Каждая итерация - это отдельный день в семестре

    # Считываем кол-во студентов в i-ый день
    stud_amount = int(input())
    stds = set() # Множество студентов , которые пришли в i-ый день
    for _ in range(stud_amount):
        student = input()
        stds.add(student)
    
    if i == 0:
        # Это значит - это первый в семестре
        # Значит все, кто пришел в первый день - были на данный момент,
        # на всех парах в этом семестре
        total = total.union(stds)
    else:
        # ЭТО уже не первый день, значит начинаем пересекать
        total = total.intersection(stds)

print(len(total))