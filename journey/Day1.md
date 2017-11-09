# So this is my journey on Golang...

## Starting with Hello World

Why not?

```
package main

import "fmt"

func main() {
	fmt.Printf("Hello World\n")
}

```

I've notice `Printf` it's kindly remind me of C Programming (which I've abandon it for many years ROFL).

And yes.It is likely as C.

```
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("This is square root of 2: %g\n", math.Sqrt(2))
}
```

## Function
Golang function declaration is simple and it's like strong-type language, we have to define type for each parameter (also return type too).

```
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

And Golang has convenient feature like declare type once for all previous parameters

```
func add(x, y int) int {
  return x + y
}
```

Return value can be multiple value

```
func swap(a, b string) (string, string) {
	return b, a
}

func main() {
	a, b := swap("Hello", "World")
	fmt.Println(a, b) // Easier swap ever!
}

```

And you can also define variable that want to return into the return type of function directly! Like this

```
func split(sum int) (x, y int) { // define return data
	x = sum * 4 / 9
	y = sum - x
	return // same as return x,y because we already DEFINE it
}

func main() {
  fmt.Println(split(18))
}
```

## Variables

One does simply. Variable can define at package level or function level. And also need to declare data type too.

```
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
```

And you can assign initial value to them to.

**OR** you can declare variable with initial value without using `var` keyword by using `:=`

```
func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
```

Type will be consider as it is.

```
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

## What type do we have?

```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

## Zero values

Golang has "Zero values" concept that any varaibles declared without initial value has that *zero* value of it's own type.

`0` for numeric / `false` for boolean / `""` empty string for string

## Type conversion

You can convert type by `T(v)` T is type, v is value

```
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

And one does simply

```
i := 42
f := float64(i)
u := uint(f)
```

## Constants

Likely in JavaScript though. using `const`

```
const Pi = 3.14
```

## For Loop

C like version

```
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```

And you can omit "initial" and "post" loop statement and become somekind of `while` loop in other langauge

```
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

```

And you can easily made *infinite loop*

```
func main() {
	for {
	}
}
```

## If

`if` is same as `for` you don't need parentheses `( )` to surround the condition but the body is require braces `{ }`

```
package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
```

And you can also using short statement `:=` to declare variable then using it in condition

```
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	// fmt.Println(v) // This produce error because v var not in this scope
	return lim
}
```

## If-else

Same as another language

```
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```

## Switch case

Similar to most of language (C/C++/JavaScript)

```
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
```

And `switch` without condition equals `switch true`

## Defer

And this is new for me. Defer is like **"HOLD ON. LET'S GET YOUR JOB DONE THEN COMEBACK TO ME"**

`defer` will execute after that function has return value or finish execution (Even it doesn't execute but it will be evaluated)

```
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
```

What about multiple defer? Think that each time `defer` has been evaluated it will put into **STACK** wait for the end of the time then execution from that stack

**FYI:** Stack is "last-in-first-out order".

```
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
```
