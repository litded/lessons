"""
Задача частотного анализа.
"""
message = "Hello Bob Bob Hello bob hello bob bob bob Hello bob"

# 1. Необходимо построить список слов
words = message.split() # ["Hello", "Bob", "Bob", "Hello", ...]

# 2. Создать словарь, который будет хранить в себе информацию в виде
# count = {"word" : count}
counter = {}

for word in words:
    if word in counter.keys():
        counter[word] += 1
    else:
        counter[word] = 1

print(counter)