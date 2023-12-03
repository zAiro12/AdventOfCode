from curses import keyname
from curses.ascii import isdigit


def cercanum(i, j):
    tempi = int(i)
    tempj = int(j)
    tot = []
    tmp = ""

    
    tempi -= 1
    num = pos[str(tempi)+ ":" + str(tempj)]
    if num.isdigit():
        #caso 2
        tmp = num
        while True:
            tempj-=1
            try:
                num = pos[str(tempi) + ":" + str(tempj)]
            except KeyError:
                tempj = j
                break
            if num.isdigit():
                tmp = num + tmp
            else:
                tempj = j
                break
        
        
        while True:
            tempj+=1
            try:
                num = pos[str(tempi) + ":" + str(tempj)]
            except KeyError:
                tempj = j
                break
            if num.isdigit():
                tmp = tmp + num
            else:
                tempj = j
                break
        if tmp.isdigit():
            tot.append(int(tmp))
        tmp = ""
    else:
        #caso 1
        while True:
            tempj-=1
            try:
                num = pos[str(tempi) + ":" + str(tempj)]
            except KeyError:
                if tmp.isdigit():
                    tot.append(int(tmp))
                tempj = j
                tmp = ""
                break
            if num.isdigit():
                tmp = num + tmp
            else:
                if tmp.isdigit():
                    tot.append(int(tmp))
                tempj = j
                tmp = ""
                break
        
        #caso 3
        while True:
            tempj+=1
            try:
                num = pos[str(tempi) + ":" + str(tempj)]
            except KeyError:
                if tmp.isdigit():
                    tot.append(int(tmp))
                tempj = j
                tmp = ""
                break
            if num.isdigit():
                tmp = tmp + num
            else:
                if tmp.isdigit():
                    tot.append(int(tmp))
                tempj = j
                tmp = ""
                break


    tempi +=2
    num = pos[str(tempi)+ ":" + str(tempj)]
    if num.isdigit():
        #caso 8
        tmp = num
        while True:
            tempj-=1
            try:
                num = pos[str(tempi) + ":" + str(tempj)]
            except KeyError:
                tempj = j
                break
            if num.isdigit():
                tmp = num + tmp
            else:
                tempj = j
                break
        
        
        while True:
            tempj+=1
            try:
                num = pos[str(tempi) + ":" + str(tempj)]
            except KeyError:
                tempj = j
                break
            if num.isdigit():
                tmp = tmp + num
            else:
                tempj = j
                break

        if tmp.isdigit():
            tot.append(int(tmp))
        tmp = ""
    else:
        #caso 7
        while True:
            tempj-=1
            try:
                num = pos[str(tempi) + ":" + str(tempj)]
            except KeyError:
                if tmp.isdigit():
                    tot.append(int(tmp))
                tempj = j
                tmp = ""
                break
            if num.isdigit():
                tmp = num + tmp
            else:
                if tmp.isdigit():
                    tot.append(int(tmp))
                tempj = j
                tmp = ""
                break
        
        #caso 9
        while True:
            tempj+=1
            try:
                num = pos[str(tempi) + ":" + str(tempj)]
            except KeyError:
                if tmp.isdigit():
                    tot.append(int(tmp))
                tempj = j
                tmp = ""
                break
            if num.isdigit():
                tmp = tmp + num
            else:
                if tmp.isdigit():
                    tot.append(int(tmp))
                tempj = j
                tmp = ""
                break

    tempi -= 1

    #caso 4
    while True:
            tempj-=1
            try:
                num = pos[str(tempi) + ":" + str(tempj)]
            except KeyError:
                if tmp.isdigit():
                    tot.append(int(tmp))
                tempj = j
                tmp = ""
                break
            if num.isdigit():
                tmp = num + tmp
            else:
                if tmp.isdigit():
                    tot.append(int(tmp))
                tempj = j
                tmp = ""
                break
    #caso 6
    while True:
            tempj+=1
            try:
                num = pos[str(tempi) + ":" + str(tempj)]
            except KeyError:
                if tmp.isdigit():
                    tot.append(int(tmp))
                tempj = j
                tmp = ""
                break
            if num.isdigit():
                tmp += num
            else:
                if tmp.isdigit():
                    tot.append(int(tmp))
                tempj = j
                tmp = ""
                break


    return tot

file = open("input.txt")

totale = 0

i = 0
j = 0
pos = {}
syb = {}

# lettura input e salvataggio
for line in file:
    for char in line:
        if char != "\n":
            pos[str(i)+ ":" + str(j)] = char
            if char == "*":
                syb[str(i)+ ":" + str(j)] = True
        j+=1
    i+=1
    j=0

#partendo da ogni sibolo guardo tutti i numeri adiacenti
for k in syb:
    i = k.split(":")[0]
    j = k.split(":")[1]
    ris = cercanum(int(i), int(j))
    if len(ris) == 2:
        totale += ris[0] * ris[1]

print(totale)
    