using System;
using System.Linq;

class question5
{
    public static string[] make_dogList()
    {
        string[] dogs = new[] { "Chihuahua", "Dachshund", "Golden Retriever", "Pomeranian", "Toy Poodle", "Shiba Inu" };
        return dogs;
    }
    
    public static void Main()
    {
        string[] dog = make_dogList();
        Random random = new Random();
        string choiced_dog = dog[random.Next(dog.Count())];
        Console.WriteLine(choiced_dog);
    }
}

// mcs question5.cs & mono question5.exe