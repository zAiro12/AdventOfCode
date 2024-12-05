import re

file = open('input.txt')
# file = open('2024/day03/prova.txt')

totale = 0
pattern = r'mul\(\d+,\d+\)'

for line in file:
    line = line.strip()
    matches = re.findall(pattern, line)
    for match in matches:
        num1, num2 = match[4:-1].split(',')
        totale += int(num1) * int(num2)
print(totale)