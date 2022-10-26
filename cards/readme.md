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
