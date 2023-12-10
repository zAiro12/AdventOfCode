using System.ComponentModel.DataAnnotations;

namespace day06;

class Program{
    public struct Dato{
        public int time;
        public int distance;

    }

    public static void Main(){
        Console.WriteLine($"parte a: {A()}");
        Console.WriteLine($"parte b: {B()}");
    }
    public static int A(){
        StreamReader sr = new StreamReader("input.txt");
        
        string rawline = sr.ReadLine().Split(":")[1];
        List<int> times = new List<int>();
        foreach(var el in rawline.Split(" ")){
            if (el != ""){
                times.Add(Convert.ToInt16(el));
            }
        }
        
        rawline = sr.ReadLine().Split(":")[1];
        List<int> distance = new List<int>();
        foreach (var el in rawline.Split(" ")){
            if (el != ""){
                distance.Add(Convert.ToInt16(el));
            }
        }

        Dato[] dati = new Dato[times.Count];
        for (int i = 0; i <= times.Count-1; i++){
            dati[i].time = times[i];
            dati[i].distance = distance[i];
        }

        List<int> topTime = new List<int>();
        int index = 0;
        int maxTime = dati[dati.Length-1].time;
        topTime.Add(0);
        for (int i = 0; i<=maxTime; i++){
            // Console.Write($"i: {i} ");
            
            // controllo index
            if (i > dati[index].time && index < dati.Length){
                i = 0;
                topTime.Add(0);
                index++;
                // Console.Write($"nuovo index: {index} ");
            }

            // calcolo punteggio
            // Console.WriteLine($"point: {topTime[index]}");
            if ((i * (dati[index].time-i)) > dati[index].distance){
                // Console.Write($"| {i}x{dati[index].time-i}={i * (dati[index].time-i)} | ");
                topTime[index]++;
            }

        }

        int total = 1;
        foreach (var item in topTime){
            total *= item;
        }
        return total;
    }
    public static long B(){
        StreamReader sr = new StreamReader("input.txt");
        
        string rawline = sr.ReadLine().Split(":")[1];
        int time = Convert.ToInt32(string.Join("", rawline.Split(" ")));
        
        rawline = sr.ReadLine().Split(":")[1];
        long distance = Convert.ToInt64(string.Join("", rawline.Split(" ")));

        Console.WriteLine($"{time} {distance}");


        long total = 0;
        for (long i = 0; i<=time; i++){

            // calcolo punteggio
            if ((i * (time-i)) > distance){
                total++;
            }

        }

        return total;
    }

}
