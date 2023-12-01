
FILE = 'input.txt'

sum = 0

for line in open(FILE):
    x = list(filter(lambda x: x.isnumeric(), line))
    sum += int(x[0]+x[-1])

print(sum)
