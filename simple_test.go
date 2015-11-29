package main
import (
	_"fmt"
	"github.com/tolgabulca/gounit"
)
var boolean bool = true
func main() {
	Test1()
}
func Test1(){
	gounit.RegexpAssert("hey go",`^[a-z]`)
	gounit.EqualAssert(1,1)
	gounit.TrueAssert(boolean)
}
