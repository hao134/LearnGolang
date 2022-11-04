## Understand interface

```
type bot interface{
    getGreeting(string, int) (string, error)
}
```

- bot : interface name
- getGreeting : Function name
- (string, int) : List of argument types
- (string, error) : List of reutrn types

---

- Two structs

```
type englishBot struct
func (englishBot) getGreeting() string

type spanishBot struct
func (spanishBot) getGreeting() string
```

To whom it may concern

```
type bot interface
```

Our program has a new type called 'bot'

```
getGreeting() string
```

If you are a type in this program with a function called 'getGreeting' and you return a string then you are now an honorary member of type "bot"

Now that you're also an honorary member of type "bot", you can now call this function called 'printGreeting'

---

- Concrete Type:
  map, struct, int, string, englishBot
- Interface Type:
  bot
