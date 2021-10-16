"""
D5:B soltion
"""

m_cpp = int(input())
n_rust = int(input())

cpp = set()
rust = set()

for _ in range(m_cpp):
    lastname = input()
    cpp.add(lastname)

for _ in range(n_rust):
    lastname = input()
    rust.add(lastname)


total = cpp.symmetric_difference(rust)
print(len(total))