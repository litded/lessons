evens = 0 # Накапливает количество четных чисел
odds = 0 # Накапливает количество нечетных чисел

while True:
    num = int(input())
    if num < 0:
        break
    else:
        if num%2 ==0:
            evens += 1
        else:
            odds += 1

print("Total evens:", evens)
print("Total odds:", odds)