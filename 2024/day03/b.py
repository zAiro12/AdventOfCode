import re

# file = open('input.txt')
file = open('2024/day03/input.txt')
# file = open('2024/day03/prova.txt')

totale = 0
string = ''
pattern = r'mul\(\d+,\d+\)'

for line in file:
    string += line.strip()

dont = string.split('don\'t()')
matches = re.findall(pattern, dont[0])
for match in matches:
    num1, num2 = match[4:-1].split(',')
    totale += int(num1) * int(num2)
for s in dont[1:]:
    do = s.split('do()')
    if len(do) == 1:
        continue
    for d in do[1:]:
        matches = re.findall(pattern, d)
        for match in matches:
            num1, num2 = match[4:-1].split(',')
            totale += int(num1) * int(num2)

print(totale)