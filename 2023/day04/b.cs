using System.IO;
using System.Text.Json;

int totalScartchcard = 0;
Dictionary<string, int> istance = new Dictionary<string, int>();
StreamReader sr = new StreamReader("input.txt");
String? line = sr.ReadLine();

int cardId = 1;
while (line != null){
    try{
        istance.Add(Convert.ToString(cardId), 1);
    }
    catch{
        istance[Convert.ToString(cardId)]++;
    }
    string split = line.Split(":")[1];
    string[] nums = split.Split(" | ");

    string[] ownNums = nums[1].Split(" ");
    string[] winNums = nums[0].Split(" ");

    Dictionary<string, int> dic = new Dictionary<string, int>();

    foreach(string n in ownNums){
        try{
            if (n != ""){
                dic.Add(n, 0);
            }
        }
        catch {
            continue;
        }
    }

    foreach(string n in winNums){
        try{
            dic[n] += 1;
        }
        catch{
            continue;
        }
    }
    
    // Console.WriteLine(JsonSerializer.Serialize(dic));

    int tmptot = 0;
    foreach(var k in dic){
        if (k.Value > 0){
            tmptot++;
        }
    }

    for(int i = 1; i <= tmptot; i++){
        try{
            istance[Convert.ToString(cardId + i)] += istance[Convert.ToString(cardId)];
        }
        catch{
            istance.Add(Convert.ToString(cardId + i), istance[Convert.ToString(cardId)]);
        }
    }

    line = sr.ReadLine();
    cardId++;
}

foreach(var k in istance){
    totalScartchcard += k.Value;
}
Console.WriteLine(totalScartchcard);
sr.Close();