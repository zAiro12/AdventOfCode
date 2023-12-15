using System.Text.RegularExpressions;

internal class Program
{
    public record Node{
        public string? L;
        public string? R;
    }
    private static void Main(string[] args){
        Console.WriteLine($"parte a: {A(args[0])}");
        Console.WriteLine($"parte b: {B(args[0])}");
    }

    public static int A(string input){
        StreamReader sr = new(input);

        string? path = sr.ReadLine();

        sr.ReadLine();

        string? rawLine = sr.ReadLine();

        Regex regex = new("(...) = \\((...), (...)\\)", RegexOptions.Compiled | RegexOptions.IgnoreCase);
        Dictionary<string, Node> albero = new();
        
        while (rawLine != null){
            Node tmpNode = new();
            Match match = regex.Match(rawLine);
            tmpNode.L = Convert.ToString(match.Groups[2]);
            tmpNode.R = Convert.ToString(match.Groups[3]);
            albero[Convert.ToString(match.Groups[1])] = tmpNode;
            rawLine = sr.ReadLine();
        }

        string cursore = "AAA";
        int totale = 1;
        for (int i = 0; i <= path.Length; i++){
            if (i == path.Length){
                i = 0;
            }
            if (path[i] == 'R'){
                if (albero[cursore].R == "ZZZ"){
                    break;
                }else{
                    totale++;
                    cursore = albero[cursore].R;
                }
            }else{
                if (albero[cursore].L == "ZZZ"){
                    break;
                }else{
                    totale++;
                    cursore = albero[cursore].L;
                }
            }
        }

        return totale;
    }
    public static float B(string input){
        StreamReader sr = new(input);

        string? path = sr.ReadLine();

        sr.ReadLine();

        string? rawLine = sr.ReadLine();

        Regex regex = new("(...) = \\((...), (...)\\)", RegexOptions.Compiled | RegexOptions.IgnoreCase);
        Dictionary<string, Node> albero = new();
        
        List<string> cursori = new();
        while (rawLine != null){
            Node tmpNode = new();
            Match match = regex.Match(rawLine);
            tmpNode.L = Convert.ToString(match.Groups[2]);
            tmpNode.R = Convert.ToString(match.Groups[3]);
            albero[Convert.ToString(match.Groups[1])] = tmpNode;
            if (Convert.ToString(match.Groups[1])[2] == 'A'){
                cursori.Add(Convert.ToString(match.Groups[1]));
            }
            rawLine = sr.ReadLine();
        }
        
        int totale = 0;

        for (int i = 0; i <= path.Length; i++){
            int tmptot = 0;
            totale++;
            for (int el = 0; el<cursori.Count; el++){
                if (i == path.Length){
                    i = 0;
                }

                if (path[i] == 'R'){
                    if (albero[cursori[el]].R[2] == 'Z'){
                        tmptot++;
                    }
                    cursori[el] = albero[cursori[el]].R;

                }else{
                    if (albero[cursori[el]].L[2] == 'Z'){
                        tmptot++;
                    }
                    cursori[el] = albero[cursori[el]].L;
                }
                // Console.WriteLine($"{path[i]} | {string.Join(" ", cursori)} | ");
            }
            if (tmptot==cursori.Count){
                break;
            }

        }
        // Console.WriteLine(string.Join(" ", totali));

        return totale;
    }
}