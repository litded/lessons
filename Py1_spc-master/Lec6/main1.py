"""
Условное бесконечный цикл - заранее неизвестно, сколько итераций потребуется, но обязательно есть
условие остановки выполнения итераций.
"""

while True:
    password = input()
    if len(password) < 8:
        print("Try again")
        continue # continue - оператор, который останавливает выполнение текущей итерации 
              # и передает управление СЛЕДУЮЩЕЙ ИТЕРАЦИИ
    else:
        print("Ok!")
        break # break - оператор, который останавливает выполнение текущей итерации 
              # и передает управление ИНСТРУКЦИИ, СТОЯЩЕЙ ЗА ТЕЛОМ ЦИКЛА

print("Outside while loop!")