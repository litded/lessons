"""
Общая конкатенация.
"""

message = "Hello"
total_str = "" # Аналогично summ = 0 при подсчете суммы
for i in range(0, len(message), 1):
    total_str += message[i] + "#"
print(total_str)