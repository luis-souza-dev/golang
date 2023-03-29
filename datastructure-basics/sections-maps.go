package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
    var letters []string;
    var lettersNumbers []int;

    for i := 0; i <= 17; i++ {
        if i % 2 == 0 {
            letters = append(letters, "Even");
        } else {
            letters = append(letters, "Odd");
        }
    }

    
    fmt.Println(letters);
    fmt.Println(len(letters));
    fmt.Println(cap(letters));
    userJson, err := printMap();

    for a, i := range letters {
        fmt.Println();
        fmt.Printf("variable here: %#v and the other one is here: %#v", a, i);
        lettersNumbers = append(lettersNumbers, a);
    }

    sum(lettersNumbers...);
    os.Stdout.Write(userJson);
    fmt.Println(err);

    fmt.Println();
    test := intSeq();
    fmt.Println(test());
    fmt.Println(test());
    fmt.Println(test());
    fmt.Println(test());

    newTest := intSeq();
    fmt.Println(newTest());

    refMap := map[string]int{"idade": 1, "nome": 2};

    refMapPtr := &refMap;

    refMapNoPtr := refMap;

    // using pointers with int and string(?) can make it easier and faster to do some calcs idk
    i := 10;
    n := &i;
    p := i;
    fmt.Println(i, "i");
    *n = 90;
    fmt.Println(i, "i2");
    fmt.Println(p, "p");
    p = p + 2;
    fmt.Println(i, "i3");
    fmt.Println(p, "p2");


    // maps dont require pointers to update the original reference, much like javascript.
    fmt.Println(refMap);
    refMapNoPtr["idade"] = 2;
    fmt.Println(refMap);
    (*refMapPtr)["idade"] = 10;
    fmt.Println(refMap);

    
}

func printMap()([]byte, error){
    userMap := make(map[string]any);
    userMap["userId"] = 1;
    userMap["userName"] = "Luis Teste";
    userMap["userEmail"] = "luisotavioldesouza1@gmail.com";
    userMap["userRole"] = "Admin";
    userMap["userPicture"] = "abc.jpg";
    return json.Marshal(userMap);
}

func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func intSeq() func() int {
    i := 10
    return func() int {
        i++
        return i
    }
}
