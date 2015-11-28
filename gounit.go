package gounit
import (
	"fmt"
)

func main(){
	if isEqual(1,1) {
		fmt.Println("Oldu");
	}
}

func isEqual(v1,v2 int) (bool){
	if v1 == v2 {
		return true
	} else {
		return false
	}
}