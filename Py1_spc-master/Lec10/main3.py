"""
D5 : N solution
"""

n = int(input())
nums = []
for _ in range(n):
    current_num = int(input())
    nums.append(current_num)

a = int(input())
b = int(input())

summ = 0
for i in range(a - 1, b):
    summ += nums[i]

print(summ)