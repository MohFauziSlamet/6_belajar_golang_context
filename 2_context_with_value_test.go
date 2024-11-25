package belajargolangcontext

import (
	"context"
	"fmt"
	"testing"
)

/*
Context With Value
● Pada saat awal membuat context, context tidak memiliki value.
● Kita bisa menambah sebuah value dengan data Pair (key - value)
ke dalam context.
● Saat kita menambah value ke context, secara otomatis akan
tercipta child context baru, artinya original context nya tidak
akan berubah sama sekali.
● Untuk menambahkan value ke context, kita bisa menggunakan
function context.WithValue(parent, key, value).
*/

type contextKey string

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, contextKey("b"), "B")
	contextC := context.WithValue(contextA, contextKey("c"), "C")

	contextD := context.WithValue(contextB, contextKey("d"), "D")
	contextE := context.WithValue(contextB, contextKey("e"), "E")

	contextF := context.WithValue(contextC, contextKey("f"), "F")

	fmt.Printf("contextA: %v\n", contextA)
	fmt.Printf("contextB: %v\n", contextB)
	fmt.Printf("contextC: %v\n", contextC)
	fmt.Printf("contextD: %v\n", contextD)
	fmt.Printf("contextE: %v\n", contextE)
	fmt.Printf("contextF: %v\n", contextF)

	/// get value from context
	fmt.Printf("contextF.Value('f'): %v\n", contextF.Value(contextKey("f")))
	fmt.Printf("contextF.Value('c'): %v\n", contextF.Value(contextKey("c")))
	fmt.Printf("contextF.Value('b'): %v\n", contextF.Value(contextKey("b")))

	fmt.Printf("contextA.Value(\"b\"): %v\n", contextA.Value(contextKey("b")))
}
