## func

```
func newCard() string {
	return "Five of Diamonds"
}
```

- newCard -> Define a function called "newCard"
- string -> When executed, this function will return a value of type 'string'

## Array and Slice in Go

- Array:
  - Fixed length list of things
- Slice:
  - An array that can growth or shrink
  - Every element in a slice must be of same type

## Slice:

```
cards := []string{"Ace of Diamonds",newCard()}
cards = append(cards, "Six of spades")
```

## for loop:

```
for index, card := range cards{
    fmt.Println(card)
}
```

- index : index of this element in the array
- card : Current card we're iterating over
- range cards : Take the slice of 'cards' and loop over it
- fmt.Println(card) : Run this one time for each card in the slice

## receiver:

```
func (d deck) print() {
  for i, card := range d {
    fmt.Println(i, card)
  }
}
```

Any variable of type "deck" now gets access to the "print" method

- d : The actual copy of the deck we're working with is available in the funciton as a variable called 'd'

- deck : Every variable of type 'deck' can call this function on itself.

## Write and Read File

- func WriteFile(filename string, data []byte, perm os.FileMode) error
  - data []byte -> 必須將 data 轉為 byte 檔
  - perm os.FileMode -> 權限 0666 表示所有人都可以讀寫

* func ReadFile(filename string) ([]byte, error)
  - `byteSlice err := ioutil.ReadFile(filename)`
  * err -> Value of type 'error'. If nothing went wrong, it will have a value of 'nil'
