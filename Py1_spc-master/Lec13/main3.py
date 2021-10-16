"""
Возвращаемые значения. Часть 2.
"""
def get_name():
    name = "Alice"
    return name 

def get_void():
    print("This function without RETURN!")
    # Если функция не имеет при себе явно описанного блока return....
    # то функция по умолчанию возвращает None
    # Синтаксический сахар - можно написать голый return
    return 

if get_void() is None:
    # Проверка на  None
    print()

ans1 = get_name()
print("From get_name():", ans1)

ans2 = get_void()
print("From get_void():", ans2)

# """
# В Java
# """
# void MyFunc(){
# }