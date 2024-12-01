file = open('input.txt')

totale = 0
dx = []
sx = []

for line in file:
    line = line.strip().split('   ')
    sx.append(line[0])
    dx.append(line[1])
dx.sort()
sx.sort()
for i in range(len(dx)):
    totale += abs(int(sx[i]) - int(dx[i]))
print(totale)