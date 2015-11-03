# gotslist

Golang list with thread safe implementation

# Installing

```bash
go get github.com/prsolucoes/gotslist
```

# Importing into your project

```golang
import "github.com/prsolucoes/gotslist"
```

# How to use

```golang
func ExampleHowToUse() {
	// new
	tslist := NewGoTSList()

	// add
	tslist.Push("New element")

	// remove
	for e := tslist.Front(); e != nil; e = e.Next() {
		tslist.Remove(e)
	}

	// len
	_ = tslist.Len()

	// is empty
	_ = tslist.IsEmpty()

	// lock and unlock
	tslist.Lock()
	tslist.Unlock()

	fmt.Println("ok")

	// Output: ok
}
```

There is a test method to use all methods
