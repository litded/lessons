"""
Блоки else и finally.
Блок else - отрабатывает только тогда, когда внутри try не произошло исключительной ситуации.
Блок finally - отрабатывает в любом случае.
"""
def dangerous(n:int):
    """
    Это опасная фунция
    """
    return 1 / n 


try:
    n = int(input())
    dangerous(n)
except ZeroDivisionError as z_err:
    print("Ой! Деление на ноль:", z_err)
except ValueError as v_err:
    print("Такое нельзя привести к целому числу:", v_err)
except KeyboardInterrupt as k_err:
    print("Не используй Ctrl + C!")
except EOFError as e_err:
    print("Не искользуй Ctrl +Z!")
else:
    print("else block working!")
finally:
    print("finally block!")

print("Graceful shutdown!")
