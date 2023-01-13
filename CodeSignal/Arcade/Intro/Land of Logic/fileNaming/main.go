package main

import "fmt"
import "regexp"

func fileNaming(names []string) []string {
    namesMap := make(map[string]bool)

    newNames := []string{}
    for _, name := range(names) {
        cand := name
        _, exists := namesMap[cand]

        // while the name already exists, keep looking
        i := 0
        for exists {
            i++
            cand = name + fmt.Sprintf("(%d)", i)
            _, exists = namesMap[cand]
        }

        namesMap[cand] = true
        newNames = append(newNames, cand)
    }

    return newNames
}

func main() {
	fmt.Printf("%v", fileNaming(new string {"doc",
                                          "doc",
                                          "image",
                                          "doc(1)",
                                          "doc"}))
}
