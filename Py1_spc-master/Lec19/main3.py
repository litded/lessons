"""
Бесконечный цикл for
"""
def infinity(up=True):
    """
    Бесконечный генератор
    Если up = True : [0........+inf)
    Если up = False: [0, .... -inf)
    """
    if up:
        i = 0
        while True:
            yield i 
            i+= 1
    else:
        i = 0
        while True:
            yield i 
            i -= 1


inf_plus = infinity(False)
for i in inf_plus:
    print("Inf-:", i)

