# gounit
A unit test tool for Go

![Image of SS]
(http://i.imgur.com/MdgGWcG.png)

## Setup
    $ go install github.com/tolgabulca/gounit
    import "github.com/tolgabulca/gounit"

## Usage
Usage is really simple just write gounit, put a point, write assert name and set the parameters

    gounit.AssertName(v1,v2)
## Assert List
##### EqualAssert(v1,v2 int) bool
Compares two values and returns true if two values are equal, otherwise returns false

##### SameAssert(v1,v2 string) bool
Compares two values and returns true if two values are same, otherwise returns false

##### TrueAssert(v bool) bool
Compare value and returns true if value is TRUE,otherwise returns false

##### FalseAssert(v1,v2 string) bool
Compare value and returns true if value is TRUE,otherwise returns false

##### GreaterThanAssert(v1,v2 int) bool
Compares two values and returns true if first value greater than second value, otherwise returns false

##### LessThanAssert(v1,v2 int) bool
Compares two values and returns true if first value less than second value, otherwise returns false

##### CountStringArrayAssert(lenght int,array []string) bool
Compares two values and returns true if first param and arrays lenght are equal, otherwise returns false

##### CountIntArrayAssert(lenght int,array []int) bool
Compares two values and returns true if first param and arrays lenght are equal, otherwise returns false

##### CountFloatArrayAssert(lenght int,array []float32) bool
Compares two values and returns true if first param and arrays lenght are equal, otherwise returns false

##### CountBoolArrayAssert(lenght int,array []bool) bool
Compares two values and returns true if first param and arrays lenght are equal, otherwise returns false

##### RegexpAssert(str string,regex string) bool
Returns true if str matched with regex, otherwise returns false

