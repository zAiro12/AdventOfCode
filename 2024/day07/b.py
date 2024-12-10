file = open('input.txt')
# file = open('2024/day07/input.txt')
# file = open('2024/day07/prova.txt')

class Nodo:
    def __init__(self, valore):
        self.valore = valore
        self.sinistra = None
        self.destra = None
        self.concatena = None

class AlberoBinario:
    def __init__(self):
        self.radice = None

    def inserisci(self, valore, maxVal):
        if self.radice is None:
            self.radice = Nodo(valore)
        else:
            self._aggiungi_figli(self.radice, valore, maxVal)

    def _aggiungi_figli(self, nodo, numero, maxVal):
        nuovo_valore = nodo.valore + numero
        # if nuovo_valore <= maxVal:
        if nodo.sinistra is None:
            nodo.sinistra = Nodo(nuovo_valore)
        else:
            self._aggiungi_figli(nodo.sinistra, numero, maxVal)
        
        # Operazione di moltiplicazione
        nuovo_valore = nodo.valore * numero
        # if nuovo_valore <= maxVal:
        if nodo.destra is None:
            nodo.destra = Nodo(nuovo_valore)
        else:
            self._aggiungi_figli(nodo.destra, numero, maxVal)
        
        # Operazione di concatenazione
        concatenato = int(str(nodo.valore) + str(numero))
        # if concatenato <= maxVal:
        if nodo.concatena is None:
            nodo.concatena = Nodo(concatenato)
        else:
            self._aggiungi_figli(nodo.concatena, numero, maxVal)

    def controlla_figli(self, nodo, valore):
        if nodo is None:
            return False
        
        if nodo.sinistra is None and nodo.destra is None and nodo.concatena is None:  # È una foglia
            return nodo.valore == valore

        return (self.controlla_figli(nodo.sinistra, valore) or 
                self.controlla_figli(nodo.destra, valore) or
                self.controlla_figli(nodo.concatena, valore))

def stampa_albero(nodo, livello=0):
    if nodo is not None:
        stampa_albero(nodo.destra, livello + 1)
        print(' ' * 4 * livello + f"({livello})" +'->', nodo.valore)
        stampa_albero(nodo.sinistra, livello + 1)
        stampa_albero(nodo.concatena, livello + 1)

totale = 0
var = 0
for line in file:
    daRaggiungere, nums = line.strip().split(": ")
    daRaggiungere = int(daRaggiungere)
    nums = list(map(int, nums.split(" ")))
    
    albero = AlberoBinario()
    for n in nums:
        albero.inserisci(n, daRaggiungere)
    
    var += 1
    print(f"Albero {var}")
   
    if albero.controlla_figli(albero.radice, daRaggiungere):
        totale += daRaggiungere
print(totale)