nums_set = set([10, 20, 30, 40, 50])
lst = [10, 20, 30, 40, 50]
tpl = (10, 20, 30, 40, 50)

print("Set:", nums_set, "Size:", nums_set.__sizeof__(), "Bytes")
print("List:", lst, "Size:", lst.__sizeof__(), "Bytes")
print("Tuple:", tpl, "Size:", tpl.__sizeof__(), "Bytes")