from word2number import w2n


file = open("input.txt")

totale = 0

for line in file:
    primo = 0
    ultimo = 0
    flag = True

    for i in line:

        if i.isdigit():
            if flag:
                primo = int(i)
                flag = False
            ultimo = int(i)
        else:
            s = ""
            for j in line:
                if not j.isdigit():
                    s += j
                    try:
                        if flag:
                            primo = w2n.word_to_num(s)
                            flag = False
                        ultimo = w2n.word_to_num(s)
                        s = ""
                    except:
                        continue
                else:
                    try:
                        ultimo = w2n.word_to_num(s)
                    except:
                        break

            
        line = line[1:]
    totale += int(str(primo) + str(ultimo))                
print(totale)