## Struct

Data structure. Collection of properties that are related together.

## &variable

Give me the memory address of the value this variable is pointing at

## \*pointer

Give me the value this memory address is pointing at

## receiver with pointer

```
func (pointerToPerson *person) updateName() {
    *pointerToPerson
}
```

- \*person: This is a type description - it means we're working with a pointer to a person
- \*pointerToPerson: This is an operator - it means we want to manipulate the value the pointer is referencing.

## shortcut

```
jimPointer := &jim
jimPointer.updateName("jimmy")
// Type of * person, or a pointer to a person
```

equal to

```
jim.updateName("jimmy")
// Type of person
```

go gives this shorcut when the func only receive one argument.

```
func (pointerToPerson *person) updateName()
// *person : Type of *person, or a pointer to a person
```

## Value Types & Reference Types

- Value Types: Use pointers to change these things in a function
  - int
  - float
  - string
  - bool
  - structs
- Reference Types: Don't worry about pointers with these
  - slices
  - maps
  - channels
  - pointers
  - functions
