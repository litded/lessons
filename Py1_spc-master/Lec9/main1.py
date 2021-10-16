"""
D5: H solution
"""
valid = set(["да", "Да", "ДА", "дА"])

message = input() # Входная строка
if ((message[0] + message[-1])  in valid) or ((message[-1] + message[0])  in valid):
    print("СОГЛАСЕН")
else:
    print("НЕ СОГЛАСЕН")