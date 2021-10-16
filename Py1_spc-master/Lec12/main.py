words = ["Alice", "Bob", "banana"]
# Создадим пустую строку
ans = ""
for i in range(len(words)):
    if i == (len(words) - 1):
        ans += words[i]
    else:
        ans += words[i] + "# #"

print(ans)
# Использование стандартного метода .join()
print("# #".join(words))