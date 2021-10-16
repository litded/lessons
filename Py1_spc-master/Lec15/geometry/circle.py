"""
Модуль для расчета периметра и площади окружности
"""
import math as m 

def perimeter(r:float):
    """
    2pir
    """
    return 2 * m.pi * r 


def area(r:float):
    """
    pir^2
    """
    return m.pi * r ** 2