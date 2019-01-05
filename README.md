De Golomb by Chanel

## Exponential Golomb

Decode a exponential Golomb coded byte [Wikipedia](https://en.wikipedia.org/wiki/Exponential-Golomb_coding)

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
### Aside

Exp-Golomb coded fields are denotes as N+1 padded with the count of 1-bits as 0's. Eg: 6-base10 = 00111. This program returns 7, which is the encoded value, not 6, which is described as the value of N. This small library was conceived to decode H.264 packets where the goal is not the original value but the encoded value. Thus, the library operates as a software engineer or computer scientist would imagine it to, not as a mathematician would. May Knuth save you if you're both.
