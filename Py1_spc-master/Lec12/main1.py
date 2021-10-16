"""
Списочные выражения
"""
nums = [ 1, 2, 3, 4, 5, 6, 7, 8]

# for i in range(len(nums)):
#     nums[i] = str(nums[i])

print(", ".join([str(x) for x in nums]))


### Простейший синтаксис
# Задача - создадим список натуральных чисел от 1 до 100 включительно
nums = []
for i in range(1, 101):
    nums.append(i)

# Для i из дипазона от 1 до 101 добавь i
new_nums = [i for i in range(1, 101)]

# Задача - создадим список со следующим правилом: четные числа возводим в квадрат, нечетные - в куб
nums = []
for i in range(1 ,101):
    if i % 2 ==0 :
        nums.append(i ** 2)
    else:
        nums.append(i ** 3)
# Для i из диапазона от 1 до 101 добавить i**2 если i - четное, добавить i**3  в противном случае
new_nums = [i**2 if i%2 ==0 else i**3 for i in range(1, 101)]
print(nums == new_nums)


# Следствием наличия условного оператора в списочном выражении (list comprehension)
one = 1
ans = (100 if one == 1 else 200)
print(ans)