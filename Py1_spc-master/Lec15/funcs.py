"""
Модуль содержащий набор убер-важных функций
"""
def add(a:int, b:int):
    """
    a + b -> c:int 
    """
    return a + b 

def sub(a:int, b:int):
    """
    a - b -> c:int
    """
    return a - b 

def mult(a:int, b:int):
    """
    a * b -> c:int
    """
    return a * b 

DB_URI = "//mysql//config//127.2.2.33"


if __name__ == "__main__":
    print("============FUNCS WORKING================")
    result = add(10, 20) * sub(30, 10) / mult(20, 30)
    print("Result:", result)
    print("URI:", DB_URI)

