//TolgaBulca
package gounit
import (
	"fmt"
	"regexp"
)

var show bool = true

// 'show' variable setter
func Show(sh bool) {
	show = sh
}

// Compares two values and returns true if two values are equal
// Otherwise returns false
func EqualAssert(v1,v2 int) bool{
	if v1 == v2 {
		return true
	} else {
		if show == true {
			fmt.Printf("Failed! %d and %d are not equal : \n",v1,v2)
		}
		return false
	}
}

// Compares two values and returns true if two values are same
// Otherwise returns false
func SameAssert(v1,v2 string) bool{
	if v1 == v2 {
		return true
	} else {
		if show == true {
			fmt.Printf("Failed! %s and %s are not same : \n",v1,v2)
		}
		return false
	}
}

// Compare value and returns true if value is TRUE
// Otherwise returns false
func TrueAssert(v bool) bool {
	if v == true {
		return true
	} else {
		if show == true {
			fmt.Printf("Failed! %v is not TRUE : \n",v)
		}
		return false

	}
}

// Compare value and returns true if value is FALSE
// Otherwise returns false
func FalseAssert(v bool) bool {
	if v == true {
		return true
	} else {
		if show == true {
			fmt.Printf("Failed! %v is not FALSE : \n",v)
		}
		return false

	}
}

// Compares two values and returns true if first value greater than...
// second value, otherwise returns false
func GreaterThanAssert(v1,v2 int) bool{
	if v1 > v2 {
		return true
	} else {
		if show == true {
			fmt.Printf("Failed! %d isn't greater than %d : \n",v1,v2)
		}
		return false
	}
}

// Compares two values and returns true if first value less than...
// second value, otherwise returns false
func LessThanAssert(v1,v2 int) bool{
	if v1 < v2 {
		return true
	} else {
		if show == true {
			fmt.Printf("Failed! %d isn't greater than %d : \n",v1,v2)
		}
		return false
	}
}

// Compares two values and returns true if first param
// and arrays lenght are equal, otherwise returns false
func CountStringArrayAssert(lenght int,array []string) bool{
	if lenght == len(array) {
		return true
	} else {
		if show == true {
			fmt.Printf("Failed! %v and %v are not equal : \n",lenght,array)
		}
		return false
	}
}

// Compares two values and returns true if first param
// and arrays lenght are equal, otherwise returns false
func CountIntArrayAssert(lenght int,array []int) bool{
	if lenght == len(array) {
		return true
	} else {
		if show == true {
			fmt.Printf("Failed! %v and %v are not equal : \n",lenght,array)
		}
		return false
	}
}

// Compares two values and returns true if first param
// and arrays lenght are equal, otherwise returns false
func CountFloatArrayAssert(lenght int,array []float32) bool{
	if lenght == len(array) {
		return true
	} else {
		if show == true {
			fmt.Printf("Failed! %v and %v are not equal : \n",lenght,array)
		}
		return false
	}
}

// Compares two values and returns true if first param
// and arrays lenght are equal, otherwise returns false
func CountBoolArrayAssert(lenght int,array []bool) bool{
	if lenght == len(array) {
		return true
	} else {
		if show == true {
			fmt.Printf("Failed! %v and %v are not equal : \n",lenght,array)
		}
		return false
	}
}

// Returns true if str matched with regex
// Otherwise returns false
func RegexpAssert(str string,regex string) bool{
	var Regex_x = regexp.MustCompile(regex)
	if Regex_x.MatchString(str) == true {
		return true
	} else {
		if show == true {
			fmt.Printf("Failed! %v and %v are not equal : \n")
		}
		return false
	}
}
