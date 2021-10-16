"""
D3: O solution
"""
n = int(input())
summ = 0 

for i in range(n):
    # Запустить цикл n раз
    # 0, 1, 2, 3, 4
    current_val = int(input())
    summ += current_val * (-1) ** i
    # if i % 2 ==0:
    #     summ += current_val
    # else:
    #     summ -= current_val 

print(summ)

#     1 - 2 + 3 - 4 + 5
#     1 * (-1) **0   + 2 *(-1) ** (1) + 3* (-1) ** 2