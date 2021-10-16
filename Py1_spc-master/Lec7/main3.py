"""
D3 : N solution
"""

n = int(input())
factorial = 1 
if n <= 1:
    # 0! = 1
    # 1! = 1
    print(factorial)
else:
    # 4! = 1 * 2 * 3 * 4
    for i in range(2, n + 1):
        factorial *= i 

    print(factorial)
