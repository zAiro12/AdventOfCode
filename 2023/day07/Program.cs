using System.Text.Json;
using Microsoft.VisualBasic;

internal class Program{

    public record Mani{
        public string? Carte;
        public int[] CarteInt = new int[5];
        public int Val;
    }

    private static int Sort(Mani mano1, Mani mano2, int i)
    {
        if (i==5){
            return 0;
        }
        if (mano1.CarteInt[i] < mano2.CarteInt[i])
        {
            return -1;
        }
        else if (mano1.CarteInt[i] > mano2.CarteInt[i])
        {
            return 1;
        }
        else
        {
            return Sort(mano1, mano2, i+1);
        }
    }

    private static void Main(string[] args){
        Console.WriteLine($"parte a: {A()}");
        Console.WriteLine($"parte b: {B()}");
    }

    public static int A(){
        
        StreamReader sr = new("input.txt");
        string? rawLine = sr.ReadLine();

        List<Mani> cartaAlta = [];
        List<Mani> coppie = [];
        List<Mani> dcoppie = [];
        List<Mani> tris = [];
        List<Mani> full = [];
        List<Mani> poker4 = [];
        List<Mani> poker5 = [];
        
        while (rawLine != null){
            string[] split;
            Mani manotmp = new();
            
            split = rawLine.Split(" ");
            manotmp.Carte = split[0];
            manotmp.Val = Convert.ToInt32(split[1]);
    
            Dictionary<int, int> contaMano = [];
            
            int i = 0;
            foreach (var carta in manotmp.Carte.ToCharArray()){
                switch (carta){
                    case 'T':
                        try{
                            contaMano[10]++;
                        }
                        catch{
                            contaMano[10] = 1;
                        }
                        manotmp.CarteInt[i] = 10;
                    break;
                    case 'J':
                        try{
                            contaMano[11]++;
                        }
                        catch{
                            contaMano[11] = 1;
                        }
                        manotmp.CarteInt[i] = 11;
                    break;
                    case 'Q':
                        try{
                            contaMano[12]++;
                        }
                        catch{
                            contaMano[12] = 1;
                        }
                        manotmp.CarteInt[i] = 12;
                    break;
                    case 'K':
                        try{
                            contaMano[13]++;
                        }
                        catch{
                            contaMano[13] = 1;
                        }
                        manotmp.CarteInt[i] = 13;
                    break;
                    case 'A':
                        try{
                            contaMano[14]++;
                        }
                        catch{
                            contaMano[14] = 1;
                        }
                        manotmp.CarteInt[i] = 14;
                    break;

                    default:
                        try{
                            contaMano[Convert.ToInt16(Convert.ToString(carta))]++;
                        }
                        catch{
                            contaMano[Convert.ToInt16(Convert.ToString(carta))] = 1;
                        }
                        manotmp.CarteInt[i] = Convert.ToInt16(Convert.ToString(carta));
                    break;
                }
                i++;
            }
            
            switch (contaMano.Count){
                case 5:
                    cartaAlta.Add(manotmp);
                    break;
                
                case 4:
                    coppie.Add(manotmp);
                    break;
                
                case 3:
                    if (contaMano.ContainsValue(3)){
                        tris.Add(manotmp);
                    }else{
                        dcoppie.Add(manotmp);
                    }
                    break;
                
                case 2:
                    if (contaMano.ContainsValue(4)){
                        poker4.Add(manotmp);
                    }else{
                        full.Add(manotmp);
                    }
                    break;
                
                case 1:
                    poker5.Add(manotmp);
                    break;
            
            }

            // Console.WriteLine(string.Join(" ", cartaAlta));
            // Console.WriteLine(string.Join(" ", coppie));
            // Console.WriteLine(string.Join(" ", dcoppie));
            // Console.WriteLine(string.Join(" ", tris));
            // Console.WriteLine(string.Join(" ", full));
            // Console.WriteLine(string.Join(" ", poker4));
            // Console.WriteLine(string.Join(" ", poker5));
            
            rawLine = sr.ReadLine();
        }

        cartaAlta.Sort((mano1, mano2)=>{
            return Sort(mano1, mano2, 0);
        });
        coppie.Sort((mano1, mano2)=>{
            return Sort(mano1, mano2, 0);
        });
        dcoppie.Sort((mano1, mano2)=>{
            return Sort(mano1, mano2, 0);
        });
        tris.Sort((mano1, mano2)=>{
            return Sort(mano1, mano2, 0);
        });
        full.Sort((mano1, mano2)=>{
            return Sort(mano1, mano2, 0);
        });
        poker4.Sort((mano1, mano2)=>{
            return Sort(mano1, mano2, 0);
        });
        poker5.Sort((mano1, mano2)=>{
            return Sort(mano1, mano2, 0);
        });

            // Console.WriteLine(string.Join(" ", cartaAlta));
            // Console.WriteLine(string.Join(" ", coppie));
            // Console.WriteLine(string.Join(" ", dcoppie));
            // Console.WriteLine(string.Join(" ", tris));
            // Console.WriteLine(string.Join(" ", full));
            // Console.WriteLine(string.Join(" ", poker4));
            // Console.WriteLine(string.Join(" ", poker5));
        int totale = 0;
        int index = 1;

        foreach (var el in cartaAlta){
            totale += el.Val * index;
            index++;
        }
        foreach (var el in coppie){
            totale += el.Val * index;
            index++;
        }
        foreach (var el in dcoppie){
            totale += el.Val * index;
            index++;
        }
        foreach (var el in tris){
            totale += el.Val * index;
            index++;
        }
        foreach (var el in full){
            totale += el.Val * index;
            index++;
        }
        foreach (var el in poker4){
            totale += el.Val * index;
            index++;
        }
        foreach (var el in poker5){
            totale += el.Val * index;
            index++;
        }
        
        return totale;
    }

    public static int B(){
        StreamReader sr = new("input.txt");
        string? rawLine = sr.ReadLine();

        List<Mani> cartaAlta = [];
        List<Mani> coppie = [];
        List<Mani> dcoppie = [];
        List<Mani> tris = [];
        List<Mani> full = [];
        List<Mani> poker4 = [];
        List<Mani> poker5 = [];
        
        while (rawLine != null){
            string[] split;
            Mani manotmp = new();
            
            split = rawLine.Split(" ");
            manotmp.Carte = split[0];
            manotmp.Val = Convert.ToInt32(split[1]);
    
            Dictionary<int, int> contaMano = [];
            
            int i = 0;
            foreach (var carta in manotmp.Carte.ToCharArray()){
                switch (carta){
                    case 'T':
                        try{
                            contaMano[10]++;
                        }
                        catch{
                            contaMano[10] = 1;
                        }
                        manotmp.CarteInt[i] = 10;
                    break;
                    case 'J':
                        try{
                            contaMano[1]++;
                        }
                        catch{
                            contaMano[1] = 1;
                        }
                        manotmp.CarteInt[i] = 1;
                    break;
                    case 'Q':
                        try{
                            contaMano[12]++;
                        }
                        catch{
                            contaMano[12] = 1;
                        }
                        manotmp.CarteInt[i] = 12;
                    break;
                    case 'K':
                        try{
                            contaMano[13]++;
                        }
                        catch{
                            contaMano[13] = 1;
                        }
                        manotmp.CarteInt[i] = 13;
                    break;
                    case 'A':
                        try{
                            contaMano[14]++;
                        }
                        catch{
                            contaMano[14] = 1;
                        }
                        manotmp.CarteInt[i] = 14;
                    break;

                    default:
                        try{
                            contaMano[Convert.ToInt16(Convert.ToString(carta))]++;
                        }
                        catch{
                            contaMano[Convert.ToInt16(Convert.ToString(carta))] = 1;
                        }
                        manotmp.CarteInt[i] = Convert.ToInt16(Convert.ToString(carta));
                    break;
                }
                i++;
            }
            if (contaMano.ContainsKey(1) && contaMano[1] != 5){
                var maxKey = 0;
                foreach (var el in contaMano){
                    if (el.Key != 1 && (maxKey == 0 || el.Value > contaMano[maxKey])){
                        maxKey = el.Key;
                    }
                }

                contaMano[maxKey] += contaMano[1];
                contaMano.Remove(1);
            }
            
            switch (contaMano.Count){
                case 5:
                    cartaAlta.Add(manotmp);
                    break;
                
                case 4:
                    coppie.Add(manotmp);
                    break;
                
                case 3:
                    if (contaMano.ContainsValue(3)){
                        tris.Add(manotmp);
                    }else{
                        dcoppie.Add(manotmp);
                    }
                    break;
                
                case 2:
                    if (contaMano.ContainsValue(4)){
                        poker4.Add(manotmp);
                    }else{
                        full.Add(manotmp);
                    }
                    break;
                
                case 1:
                    poker5.Add(manotmp);
                    break;
            
            }

            
            rawLine = sr.ReadLine();
        }

        // Console.WriteLine($"carta alta: {string.Join(" ", cartaAlta)}");
        // Console.WriteLine($"coppie: {string.Join(" ", coppie)}");
        // Console.WriteLine($"dcoppie: {string.Join(" ", dcoppie)}");
        // Console.WriteLine($"tris: {string.Join(" ", tris)}");
        // Console.WriteLine($"full: {string.Join(" ", full)}");
        // Console.WriteLine($"poker4: {string.Join(" ", poker4)}");
        // Console.WriteLine($"poker5: {string.Join(" ", poker5)}");

        cartaAlta.Sort((mano1, mano2)=>{
            return Sort(mano1, mano2, 0);
        });
        coppie.Sort((mano1, mano2)=>{
            return Sort(mano1, mano2, 0);
        });
        dcoppie.Sort((mano1, mano2)=>{
            return Sort(mano1, mano2, 0);
        });
        tris.Sort((mano1, mano2)=>{
            return Sort(mano1, mano2, 0);
        });
        full.Sort((mano1, mano2)=>{
            return Sort(mano1, mano2, 0);
        });
        poker4.Sort((mano1, mano2)=>{
            return Sort(mano1, mano2, 0);
        });
        poker5.Sort((mano1, mano2)=>{
            return Sort(mano1, mano2, 0);
        });

            // Console.WriteLine(string.Join(" ", cartaAlta));
            // Console.WriteLine(string.Join(" ", coppie));
            // Console.WriteLine(string.Join(" ", dcoppie));
            // Console.WriteLine(string.Join(" ", tris));
            // Console.WriteLine(string.Join(" ", full));
            // Console.WriteLine(string.Join(" ", poker4));
            // Console.WriteLine(string.Join(" ", poker5));
        int totale = 0;
        int index = 1;

        foreach (var el in cartaAlta){
            totale += el.Val * index;
            index++;
        }
        foreach (var el in coppie){
            totale += el.Val * index;
            index++;
        }
        foreach (var el in dcoppie){
            totale += el.Val * index;
            index++;
        }
        foreach (var el in tris){
            totale += el.Val * index;
            index++;
        }
        foreach (var el in full){
            totale += el.Val * index;
            index++;
        }
        foreach (var el in poker4){
            totale += el.Val * index;
            index++;
        }
        foreach (var el in poker5){
            totale += el.Val * index;
            index++;
        }
        
        return totale;
    }
}