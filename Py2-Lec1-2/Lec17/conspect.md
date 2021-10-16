## Лекция 17. Всего по-немногу про хранение объектов

Темы на сегодня:
* объектно-реляционные принципы хранения
* убер-коллекции (модуль collections)
* Структуры данных
* Алгоритмы. Введение

***Задача:*** хотим узнать, каким образом можно сохранить информацию про объекты за пределами работающего кода?
### Шаг 1. Объектно-реляционный принцип хранения данных.
***ОРпх*** - принцип, согласно которому, объекты сохраняются совместимо со своей внутренней структурой.

Рассмотрим 2 варианта данного принцпа:
* Использование сериализации
* Использование табличных представлений


### Шаг 2. Использование сериализации для ОРпх
***Сериализация*** - это способ представления объекта в виде байтовой последовательности.
***Десериализация*** - это способ сборки объекта из байтовой последовательности.

***Существует*** 2 достаточно популярных механизмов сериализации :
* низкоуровненвая сериализация ```pickle```
* высокоуровненвая сериализация ```json```

### Шаг 3. Высокоуровневая сериализация ```json```
```
data = [
    {"id":1, "name":"bob"},
    {"id":501, "name" :"alice"},
    {"id":None, "name" : "petya"},
]

# Хотелка - хотим сохранить этот объект data с его структурой
import json 
"""
JavaScript Object Notation
"""

# 1. Нативная сераилизация

with open("data.json", "w") as fhand:
    json.dump(data, fhand, indent=4)

```

На этапе ```.dump``` происходят следующие действия:
* Представление объекта ```data``` в виде байтовой последовательности
* Затем байтовая последовательность транслируется в ```.js``` код
* После чего данный код записывается в ```.json``` файл


Теперь в обратном порядке - ***десериализация***:
```
import json
# 2. Нативная десериализация
new_data = None
with open("data.json", "r") as fhand:
    new_data = json.load(fhand)

print("Read from json:", new_data)
print("Type:", type(new_data))
```
На этапе ```.load``` происходят следующие действия:
* Все, что написано в файле ```data.json``` представляем в виде байтовой последовательности
* Затем байтовая последовательность транслируется в ```.py``` код
* После чего объект, получившийся на прошлом этапе, помещается в какую-либо переменную


### Шаг 4. Низкоуровневая сериализация ```pickle```
Теперь главная задача - это посмотреть на саму байтовую последовательность, и может быть мы даже захотим
эту последовательность сохранять в исходном виде.

```
import pickle 
"""
Модуль низкоуровневой сериализации в Python
"""

data = [
    {"id":1, "name":True},
    {"id":501, "name" :"alice"},
    {"id":None, "name" : "petya"},
]

# 3. Сериализация в байтовый файл
with open("data.pickle", "wb") as fhand:
    pickle.dump(data, fhand)


# 4. Десериализация из байтового файла
new_data = None 
with open("data.pickle", "rb") as fhand:
    new_data = pickle.load(fhand)

print(new_data)
```
Сериализация и десериализация при помощи модуля ```pickle``` позволяет конвертировать (и реконвертировать) объект с сохраняем его вложенной структуры по отношению к байтовым последовательностям на прямую (не создается дополнительного файла, содержащего код другого ЯП).

***Какие преимущества в работе с сериализованными объектами***?:
* Гораздо ускоренное время обращения как к байтовому файлу, так и к его содержимому
* Сохранение структуры объекта 


### Шаг 5. А что можно сериализовать?
* примитивы (базовые типы, базоовые коллекции)
* пользовательские обхекты
* функции
* классы


### Шаг 6.  Использование табличных представлений
***Предоложение***: давайте использовать для хранения объектов реляционные базы данных. Это очень удобно, модностильнопопулярно (хотя > 60 лет этой технологии). Добавочное классное свойство - способность взаимодействовать с объектами - через ```SQL```.

***База данных*** - это набор структурированных (информационных) представлений.

***С базами данных неразрывно связана некая система менеджмента данных*** - это ***СУБД*** (система управления базами данных).

***СУБД*** бывают:
* реляционные
* столбчатые
* колончатые
* нереляционные
* и прочие.

Будем рассматривать реляционные СУБД. ***РСУБД*** - это сисетма управления базами данных основанная на принципе отношений между таблицами.

Данные в реляционных СУБД располагаются в виде наборов таблиц , между которым (чаще всего) сущесвуют связи (реляции, отношений).

### Шаг 7. Хотелка
* Посмотреть на основные ```SQL``` команды (запросы)
* Попробовать выполнить их с использованием простейшей РСУБД
* Научить наши объекты (в контексте Python) взаимодействовать с такими базами данных

***Замечание***:
* Learning SQL . Книга, Алан Болье
* Введение в системы баз данных Книга, Кристофер Дейт

### Шаг 7.1 Основные запросы на языке SQL
***Запрос на яызке SQL*** - это команда кмпилятору (СУБД) что-нибудь сделать с данными или таблицами.
***Запросов бывает 2 семейства:***
* запросы на изменение/создание/сброс структуры хранения данных 
* запросы на изменение/создание/сброс данных в таблицах

```
/* This is comment on SQL */

/* Хотим хранить книги */
/* Книга представляет из себя следующий набор */

/* Название строка(текст) */
/* Автор строка(текст)*/
/* Количество страниц (целое число)*/
/* Цена (вещественное число)*/

/* Первый запрос на создание таблицы для хранения информации про книги */
CREATE TABLE IF NOT EXISTS books (title TEXT, author TEXT, pages INT, price REAL);

/* Несколько запросов по работе с данными в таблице books */
/* Допустим у меня есть книга ("LOTR:1", "J.J.Tolkin", 750, 14.99) */
INSERT INTO books (title, author, pages, price) VALUES("LOTR:1", "J.J.Tolkin", 750, 14.99);

/* Теперь хочу прочитать из таблицы */
SELECT *  FROM books;

/* Теперь я понял, что ошибься в цене товара, хочу поменять */
UPDATE books SET price = 149.99 WHERE title = "LORT:1";

/* Теперь удалим какую-нибудь книжку */
DELETE FROM books WHERE title = "LOTR:1";


```

### Шаг 7.2 Существующие РСУБД
* ```PostgreSQL``` - круто мощно молодежно (бесплатно)
* ```MySQL``` - круто мощно молодежно (бесплатно)
* ```SQLite``` - относительно круто, не мощно, но молодежно (СРАЗУ ВШИТО В ИНТЕРПРЕТАТОР PYTHON).

Попробуем теперь сосредоточить свое внимание на ```sqlite```.

### Шаг 7.3 Подключение sqlite и создание базы данных