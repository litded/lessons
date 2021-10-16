"""
Возвращаемые значения или оператор return.
"""
def get_15():
    result = 10
    print("Inside get_15()")
    return result + 5 # В этой точке программы происходит вычисление переменной result
                  # а затем происходит ОСТАНОВКА ВЫПОЛНЕНИЯ ФУНКЦИИ и передача управления
                  # ВЫЗВАВШЕЙ ЭТУ ФУНКЦИЮ части кода, с передачей вычисленного значения.
    # print("After return")

print("Before get_15() call")
ans = get_15()
print("After get_15() call")
print("Result:", ans)
