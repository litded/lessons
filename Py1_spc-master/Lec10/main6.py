"""
Ссылочность списков и их изменяемость
"""
nums = [10, 20, 30]
another_nums = nums[:] #.copy() # ref link
another_nums.append(100000)
print("nums:", nums)
print("another_nums:", another_nums)

"""
Списки изменяемы по индексам из-за механизма ссылочности
"""
nums[0] = -10
print("nums[0] = -10:", nums)