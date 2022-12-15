## Iterating Over Maps

```
func printMap(c map[string]string){
    for color, hex := range c{

    }
}
```

- c : Argument name
- map[string]string : type off the map
- color : key for htis step through the loop
- hex : value for this step through the loop

## Difference Between Maps and Structs

General guidance : use a map whenever you are representing a collection of very colosely related properties

- Map
  - All keys must be the same type
  - Use to represent a collection of related properties
  - All values must be the same type
  - Don't need to know all the keys at compile time
  - Keys are indexed - we can iterate over them
  - Reference Type !

* Struct
  - Values can be of difference type
  - You need to know all the difference fields at compile time
  - Keys don't support indexing
  - Use to represent a "thing" with a lot of different properties
  - Value Type!
