"""
D5 : I solution
"""

prev_word = input()
while True:
    current_word = input()
    if prev_word[-1] == current_word[0]:
        # Все ок. Переопределим prev_word
        prev_word = current_word
    else:
        print(current_word)
        break 