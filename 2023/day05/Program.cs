using System.Text.Json;

public class Program{

    public struct Conversione{
        public long inizio;
        public long fine;
        public long variazione;
    }

    public static void Main(){
        Console.WriteLine($"Parte A: {A()}");
        Console.WriteLine($"Parte B: {B()}");
    }

    public static long A(){
        StreamReader sr = new StreamReader("input.txt");

        string? line = sr.ReadLine();

        string[] strSeed = line.Split(": ")[1].Split(" ");
        List<long> seeds = new List<long>();

        foreach(string n in strSeed){
            seeds.Add(Convert.ToInt64(n));
        }

        List<List<Conversione>> insieme = new List<List<Conversione>>();


        line = sr.ReadLine();
        line = sr.ReadLine();
        insieme.Add(new List<Conversione>());
        int index = 0;
        line = sr.ReadLine();
        while (line != null){
            
            if (line == ""){
                index++;
                if (index == 7){
                    break;
                }
                insieme.Add(new List<Conversione>());
                line = sr.ReadLine();
                line = sr.ReadLine();
                continue;
            }

            string[] strNums = {};
            
            strNums = line.Split(" ");
            long[] nums = new long[3];
            
            for(int i=0; i<3; i++){
                nums[i] = Convert.ToInt64(strNums[i]);
            }
            insieme[index].Add(new Conversione{inizio = nums[1], fine = nums[1]+nums[2]-1, variazione = nums[0]-nums[1]});

            line = sr.ReadLine();
        }
        
        for (int i = 0; i < seeds.Count; i++){
            long seed = seeds[i];

            foreach (var map in insieme){
                foreach (var item in map){
                    
                    if (seed >= item.inizio && seed <= item.fine){
                        seed += item.variazione; 
                        break;
                    }
                    
                }
            }
            seeds[i] = seed;
        }
        return seeds.Min();
    }

    public static long B(){
        StreamReader sr = new StreamReader("input.txt");

        string? line = sr.ReadLine();

        string[] strSeed = line.Split(": ")[1].Split(" ");
        List<long> seeds = new List<long>();

        foreach(string n in strSeed){
            seeds.Add(Convert.ToInt64(n));
        }

        List<List<Conversione>> insieme = new List<List<Conversione>>();


        line = sr.ReadLine();
        line = sr.ReadLine();
        insieme.Add(new List<Conversione>());
        int index = 0;
        line = sr.ReadLine();
        while (line != null){
            
            if (line == ""){
                index++;
                if (index == 7){
                    break;
                }
                insieme.Add(new List<Conversione>());
                line = sr.ReadLine();
                line = sr.ReadLine();
                continue;
            }
            
            string[] strNums = line.Split(" ");
            long[] nums = new long[3];
            
            for(int i=0; i<3; i++){
                nums[i] = Convert.ToInt64(strNums[i]);
            }
            insieme[index].Add(new Conversione{inizio = nums[1], fine = nums[1]+nums[2]-1, variazione = nums[0]-nums[1]});

            line = sr.ReadLine();
        }

        for (int i = 0; i < seeds.Count; i+=2){
            long seedMin = seeds[i];
            long seedMax = seeds[i] + seeds[i+1];
            long min = long.MaxValue;

            for (long seed = seedMin; seed < seedMax; seed++){
                long analisi = seed;
                
                foreach (var map in insieme){
                    foreach (var item in map){
                        
                        if (analisi >= item.inizio && analisi <= item.fine){
                            analisi += item.variazione; 
                            break;
                        }
                        
                    }
                }
                if (analisi < min){
                    min = analisi;
                }
            }
            seeds[i]=min;
            seeds[i+1]=min;

        }
        return seeds.Min();
    }
}