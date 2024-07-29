package stringfunctions

import (
	"fmt"
	s "strings"
)

func StringTest() {
	name := "Selim"
	fmt.Println(s.Count(name, "e"))      // 1
	fmt.Println(s.Contains(name, "k"))   // false
	fmt.Println(s.Index(name, "o"))      // -1
	fmt.Println(s.ToLower(name))         //selim
	fmt.Println(s.ToUpper(name))         //SELIM
	fmt.Println(s.HasPrefix(name, "Se")) //true
	fmt.Println(s.HasSuffix(name, "m"))  //true
	letters := []string{"f", "e", "y", "z", "a"}
	result := s.Join(letters, "*")
	fmt.Println(result)                         //f*e*y*z*a
	fmt.Println(s.Replace(result, "*", "+", 3)) //f+e+y+z*a
	fmt.Println(s.Split(result, "*"))           //[f e y z a]
	fmt.Println(s.Split(result, "-"))           //[f*e*y*z*a]
	fmt.Println(s.Repeat(result, 5))            //f*e*y*z*af*e*y*z*af*e*y*z*af*e*y*z*af*e*y*z*a
}
