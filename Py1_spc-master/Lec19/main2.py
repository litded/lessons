def my_generator(n:int):
    """
    Создает диапазон от n до 0 включительно
    """
    while n >= 0:
        yield n 
        n -= 1

gen10 = my_generator(10)
gen200 = my_generator(200)
gen2000000000000000000000000 = my_generator(2000000000000000000000000)

print("gen 10 size:", gen10.__sizeof__(), "Bytes")
print("gen 200 size:", gen200.__sizeof__(), "Bytes")
print("gen 2000000000000000000000000 size:", gen2000000000000000000000000.__sizeof__(), "Bytes")

for i in gen10:
    print("gen10:", i)

for i in gen200:
    print("gen200:", i)

for i in gen2000000000000000000000000:
    print("gen20000000000000:", i)