"""
Срезы строк.
"""
message = "Hello#world!"
#word = message[0] + message[1] + message[2] + message[3] + message[4]
word = message[0:5:1] # <- срез. [start : stop[ : step]
print("Total:", word)

# У срезов шаг по умолчанию также равен 1
new_word = message[0:5]
print("New word with step = 1:", new_word)

# Срез - это объект, того же типа, что и исходник
print("Type:", type(new_word))

# По умолчанию - start = 0
start_word = message[:5]
print("Start = 0:", start_word)

# В срезах можно опускать границу stop (от 6  до конца с шагом в 1)
stop_word = message[6::1]
print("Stop word=", stop_word)

# Сделать срез:копию
copy = message[:]

# Можно срезать по отрицательным индексам
word_2 = message[:-2]
print(word_2)

# Вернемся к исходным строкам
for letter in message:
    print("Message elem:", letter)

# По срезу можно итерироваться
for letter in message[0:8:2]:
    print("Elem:", letter)
