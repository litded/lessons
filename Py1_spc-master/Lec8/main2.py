"""
D3 : R solution
"""
n = int(input()) # Сколько слов уже было названо
words = set() # Множество названных слов

for _ in range(n):
    word = input()
    words.add(word)

# Контрольное слово
new_word = input()
if new_word in words:
    print("НЕ ЗАСЧИТАНО")
else:
    print("ЗАСЧИТАНО")