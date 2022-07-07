package main

import (
    "fmt"
   "math/rand"
    "time"
)

func getRandomElement (slice[]string) string{
  randomNum := rand.Intn(len(slice))
  return slice[randomNum]
}

func main() {
  guestList := [5]string {"Shakira", "Ronaldo", "Britney", "Mbappe", "Jlo"}
  catStorage := [3]string{"vase", "cupboard", "closet"}

rand.Seed(time.Now().UnixNano())

// converstion to slices:
guestListSlice := guestList[:]
catStorageSlice := catStorage[:]
catnapper:= getRandomElement(guestListSlice)
catLocation := getRandomElement(catStorageSlice)
fmt.Println("The catnapper is:",catnapper, "and the cat is here: ",catLocation)

}