using System.ComponentModel;
using System.Text.Json;

internal class Program{
    private static void Main(string[] args){
        string arg;
        try{
            arg = args[0]; 
        }
        catch{
            arg = "prova.txt";
        } 
        Console.WriteLine($"parte A: {A(arg)}");
        Console.WriteLine($"parte B: {B(arg)}");
    }

    public static int A(string nameFile){
        StreamReader sr = new(nameFile);
        
        string? rawline = sr.ReadLine();
        
        int totale = 0;

        while (rawline != null){
            string[] split = rawline.Split(" ");
            
            List<int> numeriIniziali = [];
            foreach (string el in split){
                numeriIniziali.Add(Convert.ToInt32(el));
            }
            List<List<int>> mani = [];
            mani.Add(numeriIniziali);
            bool flag = true;
            do{
                List<int> tmp = [];
                // Console.WriteLine(mani[mani.Count-1].Count);
                List<int> ultimaMano = mani[mani.Count-1];
                for (int i=0; i < ultimaMano.Count-1; i++){
                    tmp.Add(ultimaMano[i+1]-ultimaMano[i]);
                }

                Dictionary<int, int> conta = [];
                foreach (int el in tmp){
                    try{
                        conta[el] ++;
                    }
                    catch{
                        conta[el] = 1;
                    }
                }

                mani.Add(tmp);
                if (conta.Count == 1){
                    flag = false;
                    List<int> zeri = [];
                    for (int i=0; i < mani[mani.Count-1].Count-1; i++){
                        zeri.Add(0);
                    }
                    mani.Add(zeri);
                }
            } while (flag);

            // Console.WriteLine(JsonSerializer.Serialize(mani));
            
            for (int i = mani.Count-1; i >= 0; i--){
                try{
                    //Console.WriteLine($"{mani[i][mani[i].Count-1]} + {mani[i+1][mani[i+1].Count-1]} = {mani[i][mani[i].Count-1] + mani[i+1][mani[i+1].Count-1]}");
                    mani[i].Add(mani[i][mani[i].Count-1] + mani[i+1][mani[i+1].Count-1]);
                }
                catch{
                    mani[i].Add(0);
                }
            }

            // Console.WriteLine(JsonSerializer.Serialize(mani));
            // Console.WriteLine("\n---------\n");
            totale+= mani[0][mani[0].Count-1];
            
            rawline = sr.ReadLine();
        }

        return totale;
    }
    public static int B(string nameFile){
        StreamReader sr = new(nameFile);
        
        string? rawline = sr.ReadLine();
        
        int totale = 0;

        while (rawline != null){
            string[] split = rawline.Split(" ");
            
            List<int> numeriIniziali = [];
            numeriIniziali.Add(0);
            foreach (string el in split){
                numeriIniziali.Add(Convert.ToInt32(el));
            }
            List<List<int>> mani = [];
            mani.Add(numeriIniziali);

            bool flag = true;
            do{
                List<int> tmp = [];
                List<int> ultimaMano = mani[mani.Count-1];
                for (int i=0; i < ultimaMano.Count-1; i++){
                    if (i==0){
                        tmp.Add(0);
                        continue;
                    }
                    tmp.Add(ultimaMano[i+1]-ultimaMano[i]);
                }

                Dictionary<int, int> conta = [];
                foreach (int el in tmp){
                    try{
                        conta[el] ++;
                    }
                    catch{
                        conta[el] = 1;
                    }
                }

                mani.Add(tmp);
                if (conta.Count == 1){
                    flag = false;
                }
            } while (flag);
            
            for (int i = mani.Count-2; i >= 0; i--){
                mani[i][0] = mani[i][1] - mani[i+1][0];
            }

            totale+= mani[0][0];
            
            rawline = sr.ReadLine();
        }

        return totale;
    }
}

internal class Directory<T1, T2>
{
}