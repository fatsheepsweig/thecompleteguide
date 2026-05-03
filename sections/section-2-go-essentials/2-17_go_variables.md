## Variable Declaration in Go

#### Difference between var keyword and short declaration operator

| var keyword | short declaration operator |
| --- | --- |
| var is a **lexical keyword** present in Golang. | **:=** is known as the short declaration operator. |
| It is used to declare and initialize the variables inside and outside the functions. | It is used to declare and initialize the variables only inside the functions. |
| Using this, variables have generally package level or global level scope. It can also have local scope. | Here, variables has only local scope as they are only declared inside functions. |
| Declaration and initialization of the variables can be done separately. | Declaration and initialization of the variables must be done at the same time. |
| It is mandatory to put type along with the variable declaration. | There is no need to put type. If you, it will give error. |

### Summary of Syntax Rules 

| Method | Syntax | Use Case |
| --- | --- | --- |
| **Standard** | `var x, y int = 1, 2` | Explicitly typed global or local vars. |
| **Inferred** | `var x, y = 1, "a"` | Different types on one line. |
| **Short** | `x, y := 1, "a"` | Local function variables (no `var` keyword). |
| **Block** | `var (...)` | Grouping related variables for clarity. |



