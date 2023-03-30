package pointers

import "fmt"

func Pointers() {
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