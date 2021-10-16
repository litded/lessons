"""
Top 1 Vlasov's interview question
"""

a_list = [] # set(), dict()


def changer():
    a_list.append(200)

print("Origin:", a_list)
changer()
print("After changer call:", a_list)