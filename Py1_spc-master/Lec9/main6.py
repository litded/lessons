"""
D5 : L solution
"""

message = input()
initial_length = len(message)
for i in range(initial_length):
    if i == 0:
        print(message)
    
    if i % 2 == 0:
        message = message[2:]
        print(message)
    elif i % 2 != 0:
        message = message[:-2]
        print(message)

    if (len(message) <= 2):
        break 