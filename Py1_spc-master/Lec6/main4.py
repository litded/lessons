"""
Собственный валидатор паролей.
break/continue можно использовать ТОЛЬКО ВНУТРИ ТЕЛА ЦИКЛА
D3:I solution
"""

while True:
    password1 = input()
    password2 = input()

    if len(password1) < 8:
        print("Слишком короткий пароль")
    elif ("qwe" in password1) or ("123" in password1):
        print("Слишком простой пароль")
    elif password1 != password2:
        print("Пароли различаются")
    else:
        print("Пароль принят")
        break
