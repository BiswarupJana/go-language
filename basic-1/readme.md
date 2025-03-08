
# Go Types & Null Values

## Basic Types

Go comes with a couple of built-in basic types:

- **int**: A number WITHOUT decimal places (e.g., -5, 10, 12 etc)
- **float64**: A number WITH decimal places (e.g., -5.2, 10.123, 12.9 etc)
- **string**: A text value (created via double quotes or backticks: "Hello World", `Hi everyone`)
- **bool**: true or false

## Niche Basic Types

There are some noteworthy "niche" basic types which you'll typically not need that often but which you should still know about:

- **uint**: An unsigned integer which means a strictly non-negative integer (e.g., 0, 10, 255 etc)
- **int32**: A 32-bit signed integer, which is an integer with a specific range from -2,147,483,648 to 2,147,483,647 (e.g., -1234567890, 1234567890)
- **rune**: An alias for int32, represents a Unicode code point (i.e., a single character), and is used when dealing with Unicode characters (e.g., 'a', 'ñ', '世')
- **uint32**: A 32-bit unsigned integer, an integer that can represent values from 0 to 4,294,967,295
- **int64**: A 64-bit signed integer, an integer that can represent a larger range of values, specifically from -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807

There also are more types like int8 or uint8 which work in the same way (=> integer with smaller number range)

## Null Values

All Go value types come with a so-called "null value" which is the value stored in a variable if no other value is explicitly set.

For example, the following int variable would have a default value of 0 (because 0 is the null value of int, int32 etc):

```go
var age int // age is 0
```

Here's a list of the null values for the different types:

- **int**: 0
- **float64**: 0.0
- **string**: "" (i.e., an empty string)
- **bool**: false

## Value Declaration in Go

Overview

In Go, values can be declared using var, const, or := (short declaration). Each method has specific use cases based on whether the value is mutable, immutable, or inferred from the assigned value.

Default Values in Go

When variables are declared without an explicit value, they take on the following default (zero) values:

**int**

int: 0

**float64**

float64: 0.0

**string**

string: "" (i.e., an empty string)

**bool**

bool: false

Example:

```go
var age int // age is 0
```

Value Declaration Methods

1. Using var

The var keyword is used for declaring variables explicitly. You can specify the type or let Go infer it.

```go
var a int = 10     // Explicit type declaration
var b = "Hello"   // Type inferred from assigned value
var c float64     // Type inferred from assigned value
```

2.Using const

The const keyword is used for declaring constants, which are immutable.

```go
const a = 10     // Constant declaration
const b = "Hello"   // Constant declaration
const c float64     // Constant declaration
```

3.Using :=

The := operator is used for declaring variables and constants in a single statement. It is a shorthand for var and const.

```go
a := 10     // Variable declaration
b := "Hello"   // Variable declaration
c := float64     // Variable declaration
```
### Notes

Variables declared using var at the package level cannot use :=.

Constants cannot be declared using :=.

Unused variables cause compilation errors in Go.

Conclusion

Go provides multiple ways to declare values, each suited for different use cases. Understanding when to use var, const, or := helps in writing clean and efficient Go code.