"""
Задача - десериализовать json файл.
"""
import json 

from_json = None 
with open("data.json", "r") as fhand:
    from_json = json.load(fhand)

print(from_json)
print("Type:", type(from_json))
if type(from_json) == dict:
    print("Name of person:", from_json["name"])
    for key, val in from_json.items():
        print("Key:", key, "Val:", val)
