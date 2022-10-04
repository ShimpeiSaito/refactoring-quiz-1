using System;
using System.Linq;

class Question5
{
    public static string[] getDogList()
    {
        string[] dog_list = new[] { "Chihuahua", "Dachshund", "Golden Retriever", "Pomeranian", "Toy Poodle", "Shiba Inu" };
        return dog_list;
    }

    public static string choiceDog()
    {
        string[] dog_list = getDogList();
        
        Random random = new Random();
        return dog_list[random.Next(dog_list.Count())];
    }
    
    public static void Main()
    {
        string dog = choiceDog();
        System.Console.WriteLine(dog);
    }
}

// mcs question5_ans.cs & mono question5_ans.exe