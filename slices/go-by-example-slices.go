package slices

import (
	"fmt"
	"slices"
)

func Go_by_example() {
	var s []string
	fmt.Println("uninit:",s, s == nil, len(s)==0)

   // make - To create an empty slice with non-zero length

   s = make([]string, 3)
   fmt.Println("emp:",s, "len:", len(s)==0, "cap: ", cap(s))

   // get and set 

   s[0] = "a"
   s[1] = "b"
   s[2] = "c"

   fmt.Println("set:" ,s)
   fmt.Println("get:", s[2])

   // append - return a slice with one or more new values

   s = append(s, "d")
   s = append(s, "e","f","g")
   fmt.Println("apd:",s)

   // copy - copies elements from a source slice into a destination slice.

   c := make([]string, len(s))
   copy(c,s)
   fmt.Println("cpy",c)
   
   // slice operator 
   
   l := s[2:5]
   fmt.Println("sli",l) 

   l = s[:5]
   fmt.Println("sl2", l)

   l = s[2:]
   fmt.Println("sl3", l)

   // single-line declaration

   t := []string{"g", "h", "i"}
   fmt.Println("dcl:", t)

   t2 := []string{"g", "h", "i"}

   // slices package
   // https://pkg.go.dev/slices
   
   check := slices.Equal(t, t2)
   if check {
     fmt.Println("t == t2")
   }

   if slices.Contains(t2, "i") {
      fmt.Println("t2 contain t")
   }else{
      fmt.Println("do not contain")
   }

}


