## Go Routins

```
go checkLink(link)
```

- go : Create a new thread go routine...
- checkLink(link) : ...And run this function with it

## Concurrency

We can have multiple threads executing code. If one thread blocks, another one is picked up and worked on.

## Parllelism

Multiple threads executed at the exact same time. Requires multiple CPU's

## Sending Data with Channels

- channel <- 5 : Send the value '5' into this channel
- myNumber <- channel : Wait for a value to be sent into the channel. When we get one assign the value to 'myNumber'
- fmt.Println(<-channel) : Wait for a value to be sent into the channel. When we get one, log it out immediately
