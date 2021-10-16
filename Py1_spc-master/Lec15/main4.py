"""
Как пользоваться уже готовыми библиотеками/фреймворками/модулями
которых нет в составе стандартной библиотеки?

Утилита pip - (Python Install Packages manager) - менеджер пакетов в питоне.

1) Версия pip : pip --version (pip3 для тех кто Ubuntu/MacOs)
2) Посмотреть список всех уже установленных СТОРОННИХ пакетов/модулей/фреймворков
pip freeze
3) Для того, чтобы установить какой-нибудь пакет :
pip install <name1> <name2> <name3> .... <nameN>
Пример: pip install  Flask (ПОСМОТРИТЕ В WIKI КАК ОБХОДИТЬ ЗАЩИТУ СЕТИ)
"""
from flask import Flask 

"""
How to create setup.py module
"""
