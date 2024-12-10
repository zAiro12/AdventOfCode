def calcola_checksum(disk_map):
    string = ""
    file_id = 0
    for x, el in enumerate(disk_map):
        blocchi = int(el)
        if x % 2 == 0:
            string += str(file_id) * blocchi
            file_id = (file_id + 1) % 10
        else: 
            string += '.' * blocchi

    string_list = list(string)
    sx, dx = 0, len(string_list) - 1

    while sx < dx:
        while sx < len(string_list) and string_list[sx] != ".":
            sx += 1
        while dx >= 0 and string_list[dx] == ".":
            dx -= 1
        if sx < dx:
            string_list[sx], string_list[dx] = string_list[dx], string_list[sx]

    string = ''.join(string_list).rstrip('.')
    checksum = sum(int(el) * i for i, el in enumerate(string) if el != ".")
    
    return checksum

# Lettura del file di input
with open('input.txt') as file:
    disk_map = file.read().strip()

# Calcolo del checksum finale
checksum = calcola_checksum(disk_map)
print(f"Checksum: {checksum}")