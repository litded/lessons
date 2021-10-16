"""
Типизированный except - это блок который реагирует только на исключения какого-то определенного типа
Блоков except может быть сколько угодно
Блоки except - блоки "обработчики".
Блоки обработчики выбираются по-очереди сверху-вниз.
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
except:
    # Голый обработчик отлавливает абсолютно все исключение
    print("ЛЮБОЕ ИСКЛЮЧЕНИЕ!")

print("Graceful shutdown!")
