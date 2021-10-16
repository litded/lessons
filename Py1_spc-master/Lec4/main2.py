password = input()

if (len(password) < 10) and ("qwe" in password):
    print("Weak password")
    print("Suggest: change pass")

print("Done!")