password = input()
password2 = input()

if len(password) < 8 :
    print("Пароль слишком короткий")
elif ("qwe" in password) or ("123" in password):
    print("Пароль слишком слабый")
elif password != password2:
    print("Пароли не совпадают")
else:
    print("Все ок!")

"""
age = 10
if age >= 18:
    print("Welcome")
else:
    print("Not Welcome")
"""

age = 15
weight = 20

if age > 10:  # age > 10 and weight > 15
    if weight > 15:
        print("Cool")
    else:
        print("Not cool")
else:
    print("Not not cool")