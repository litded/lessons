"""
Замыкание + стероиды = декораторы

Декоратор (фактически) - замыкание, построенное на основе параметрической передачи функции.
Декоратор (физически) - способ добавления новых свойств (нового функционала) , существующим функциям, не меняя
их реализации (функциональным наследованием).

Отдельный паттерн проектирования - "декоратор" (refactoring.guru.ru)
Декоратор часто применяется в другом паттерне - "адаптер".
"""
def my_decorator(func):
    def wrapper():
        print("New functionality before func call")
        func()
        print("New functionality after func call")
    return wrapper 

@my_decorator
def hello():
    print("Hello world!")

# hello = my_decorator(hello)
hello()