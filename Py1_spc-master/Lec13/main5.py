"""
Создадим функцию для подсчета факториала
D7 : I solution
"""
# n! = n * (n - 1) * (n - 2) * .... * 1
# 1! = 1
# 0! = 1

def factorial(n:int):
    """
    Функция для расчета n! по следующим правилам:
    * n - целое неотрицательное число
    * n! = n * (n-1) * (n-2) * ... 1
    * 4! = 4 * 3 * 2 * 1
    * 1! = 1
    * 0! =1
    """
    if n <= 1:
        return 1 
    else:
        res = 1
        for i in range(2, n + 1):
            res *= i 
        return res 

def combination(n:int, m:int):
    """
    """
    res = factorial(n) / (factorial(m) * (factorial(n - m)))
    return int(res)


N = int(input())
M = int(input())
print(combination(N,M))

# def print_factorial(from_:int, to:int):
#     """
#     Функция распечатывает факториалы чисел от from
#     до to включительно с шагом в 1.
#     """
#     for i in range(from_, to +1):
#         print("n:", i, "n!:", factorial(i))

# print_factorial(1, 15)