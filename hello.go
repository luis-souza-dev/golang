package main

import (
	"encoding/json"
	"fmt"
	"os"

	"rsc.io/quote"
)

func main() {
    fmt.Println(quote.Go())
    a := "";
    var b string;

    a = "'1'";
    b = "'2'";
    fmt.Println(a, "b", b);
    var letters []string;

    for i := 0; i <= 17; i++ {

        if i % 2 == 0 {
            letters = append(letters, fmt.Sprint(i));
        }
    }
    fmt.Println(letters);
    fmt.Println(len(letters));
    fmt.Println(cap(letters));

    letters = append(letters, "18");
    fmt.Println(len(letters));
    fmt.Println(cap(letters));
    printMap();

}

func printMap() {
    userMap := make(map[string]any);
    userMap["userId"] = 1;
    userMap["userName"] = "Luis Teste";
    userMap["userEmail"] = "luisotavioldesouza1@gmail.com";
    userMap["userRole"] = "Admin";
    userMap["userPicture"] = "abc.jpg";
    encoded,b := json.Marshal(userMap);
    fmt.Println(encoded, "-", b);
    os.Stdout.Write(encoded);
}