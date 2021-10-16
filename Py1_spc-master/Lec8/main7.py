"""
D5 : F solution
"""

n_serials_in_db = int(input())
m_requests  = int(input())
db = set()

for _ in range(n_serials_in_db):
    serial = input()
    db.add(serial)

for _ in range(m_requests):
    request = input()
    if request in db:
        print("ЕСТЬ")
    else:
        print("НЕТ В НАЛИЧИИ")