file = open("input.txt")

totale = 0

for line in file:
    primo = 0
    ultimo = 0
    flag = True
   
    for digit in line:
        if digit.isdigit():
            if flag:
                primo = int(digit)
                flag = False
            ultimo = int(digit)

    totale += int(str(primo) + str(ultimo))
    
print(totale)