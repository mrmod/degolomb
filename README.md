De Golomb by Chanel

## Exponential Golomb

Decode a exponential Golomb coded byte

```
b := byte(23)
fieldWidths := []int{4, 2, 2}
vals := Degolomb(b, fieldWidths)
// vals == []int{1, 1, 3}
```

Also includes `BitArray` function
```
bits := bitArray(b)
// bits == []int{0,0,0,1,0,1,1,1}
```



