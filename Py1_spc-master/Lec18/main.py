message = "Hello world"
"""
А насколько сильно производительность этих методов зависит от размера строки?
Производительность строковых методов в случае больших строк - скачкообразно меняется.

Предпочтительнее использовать модуль re
"""
import re 

# re - модуль для выполнения регулярных выражений (и так же простых строковых шаблонов)

"""
* re.match() - поиск в начале строки
* re.search() - поиск по всей строке
* re.findall() - поиск всех соответствий по строке
* re.split() - разделяет строку по паттерну
* re.sub() - находит подстроку и заменяет ее на другую
* re.compile() - захват паттерна
"""

# re.match(pattern, string) - функция ищет по шаблону pattern в самом начале строки string
sample = "Python is a programming language!"
result = re.match(r'Go', sample) # r'' - так обозначается строка в естественном представлении (с запретом экранирования символов)
print(result)


# re.search(pattern , string) - функция ищет по шаблону pattern во всей строке string 
sample = "Programming Python - awesome!"
result = re.search(r'Python', sample)
print(result)
print("Start:", result.start())
print("End:", result.end())



# re.findall(pattern, string) - функция ищет все соответствия по всей строке string  паттерну pattern
sample = "Code Code Code Notcode Code code Code"
result = re.findall(r'Code', sample)
print(result)


# re.split(pattern, string, maxsplit=0) - функция разбивает строку на список по pattern
sample = "Hello world"
result = re.split(r'o', sample, maxsplit=1)
print(result)


# re.sub(pattern, repl, string) - функция заменяет в строке string подстроку содержащую pattern на repl
sample = "GoLang is the best programming language!"
result = re.sub(r'GoLang', 'Python', sample)
print(result)

# re.compile(pattern) - создает объкт паттерна, и все остальные методы могут вызываться от него
pattern = re.compile('Python')
result = pattern.findall("Python is the best programming language!")
print(result)


"""
Вместо паттерна (в виде литералов) можно использовать регулярные выражения, состоящие из следующих символов:
The special characters are:
    "."      Matches any character except a newline.
    "^"      Matches the start of the string.
    "$"      Matches the end of the string or just before the newline at
             the end of the string.
    "*"      Matches 0 or more (greedy) repetitions of the preceding RE.
             Greedy means that it will match as many repetitions as possible.
    "+"      Matches 1 or more (greedy) repetitions of the preceding RE.
    "?"      Matches 0 or 1 (greedy) of the preceding RE.
    *?,+?,?? Non-greedy versions of the previous three special characters.
    {m,n}    Matches from m to n repetitions of the preceding RE.
    {m,n}?   Non-greedy version of the above.
    "\\"     Either escapes special characters or signals a special sequence.
    []       Indicates a set of characters.
             A "^" as the first character indicates a complementing set.
    "|"      A|B, creates an RE that will match either A or B.
    (...)    Matches the RE inside the parentheses.
             The contents can be retrieved or matched later in the string.
    (?aiLmsux) The letters set the corresponding flags defined below.
    (?:...)  Non-grouping version of regular parentheses.
    (?P<name>...) The substring matched by the group is accessible by name.
    (?P=name)     Matches the text matched earlier by the group named name.
    (?#...)  A comment; ignored.
    (?=...)  Matches if ... matches next, but doesn't consume the string.
    (?!...)  Matches if ... doesn't match next.
    (?<=...) Matches if preceded by ... (must be fixed length).
    (?<!...) Matches if not preceded by ... (must be fixed length).
    (?(id/name)yes|no) Matches yes pattern if the group with id/name matched,
                       the (optional) no pattern otherwise.

The special sequences consist of "\\" and a character from the list
below.  If the ordinary character is not on the list, then the
resulting RE will match the second character.
    \number  Matches the contents of the group of the same number.
    \A       Matches only at the start of the string.
    \Z       Matches only at the end of the string.
    \b       Matches the empty string, but only at the start or end of a word.
    \B       Matches the empty string, but not at the start or end of a word.
    \d       Matches any decimal digit; equivalent to the set [0-9] in
             bytes patterns or string patterns with the ASCII flag.
             In string patterns without the ASCII flag, it will match the whole
             range of Unicode digits.
    \D       Matches any non-digit character; equivalent to [^\d].
    \s       Matches any whitespace character; equivalent to [ \t\n\r\f\v] in
             bytes patterns or string patterns with the ASCII flag.
             In string patterns without the ASCII flag, it will match the whole
             range of Unicode whitespace characters.
    \S       Matches any non-whitespace character; equivalent to [^\s].
    \w       Matches any alphanumeric character; equivalent to [a-zA-Z0-9_]
             in bytes patterns or string patterns with the ASCII flag.
             In string patterns without the ASCII flag, it will match the
             range of Unicode alphanumeric characters (letters plus digits
             plus underscore).
             With LOCALE, it will match the set [0-9_] plus characters defined
             as letters for the current locale.
    \W       Matches the complement of \w.
    \\       Matches a literal backslash.

"""
