**NOTE**: The following shows two other (equivalent) ways of declaring multiple variables:

- **Using the `var` keyword**
  - The `var` keyword is used to declare variables in and outside of functions.
  - Allows the declaration and initialization of a variable to be done separately
  - It is mandatory to include the type with the declaration 
```go
var investmentAmount float64 = 1000
var expectedReturnRate = float64 5.5
var years float64 = 10
```

- **Declare multiple variables in a single statement**
  - Is used in situations where more than one variable of the same type is declared
  - Uses `var`
```go
var investmentAmount, expectedReturnRate, years float64 = 1000, 5.5, 10
```

- **Short variable declaration** (recommended, if possible)
  - Short Variable Declaration Operator (`:=`) is used to create variables having a proper name and initial value.
  - The main purpose of using this operator is to declare and initialize the *local variables* inside of functions and for narrowing the scope of the variable(s).
```go
investmentAmount := 1000
expectedReturnRate := 5.5
years := 10
```