# tslice

This is type-script(java-script) array like slice for GoLang

```bash
go get -u github.com/poteto-go/tslice@latest
```


## Examples

### Filter

```go
array := []int{0, 1, -1, 2}
filtered := tslice.Filter(array, func(data int) bool {
	return data >= 0
})
// => filtered = []int{0, 1, 2}
```

### Reduce

```go
type User struct {
	city string
	age  int
}

users := []User{
	{"Kawasaki", 10}, {"Yokohama", 11}, {"Kawasaki", 22},
}

result := tslice.Reduce(users, func(acc int, cur User) int {
	if cur.city == "Kawasaki" {
		return acc + cur.age
	}
	
	return acc + 0
})
// => 32
```

## Func
- [x] `At`
- [x] `Concat`
- [x] `CopyWithin`
- [x] `Entries`
- [x] `Every`
- [x] `Fill`
- [x] `Filter`
- [x] `Find`: return (0 value, false) if not found
- [x] `FindIndex`
- [x] `FindLast`: return (0 value, false) if not found
- [x] `FindLastIndex`
- [ ] `Flat`
- [ ] `FlatMap`
- [x] `ForEach`
- [x] `Includes`
- [x] `IndexOf`
- [ ] `Join`
- [ ] `Keys`
- [ ] `LastIndexOf`
- [x] `Map`
- [x] `Pop`
- [x] `Push`
- [x] `Reduce`
- [x] `ReduceRight`
- [ ] `Reverse`
- [ ] `Shift`
- [ ] `Slice`
- [ ] `Some`
- [ ] `Sort`
- [ ] `Splice`
- [ ] `ToLocalString`
- [ ] `ToReversed`
- [ ] `ToSorted`
- [ ] `ToSpliced`
- [ ] `ToString`
- [ ] `UnShift`
- [ ] `Values`
- [ ] `With`