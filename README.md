# GoBooking

## License
*   no licence set

### This is a CLI booking application built using Golang

- Slice is an Abstraction of an Array
- Slices are more flexible and powerful: variable-length or get a sub-array of its own
- Slices are also index-based and have a size, but is Resized when needed

### LOOPS
*	In general, **languages provide various control structures** to control the application flow
*	A loop statement allows us to **execute code multiple times** in a loop.
*	there is only **one** loop in Golang, it doesn't have: 'while-loop' 'do-while-loop' or 'for-each-loop'

### Infinite Loop

### For Loop "strings.Fields()"
*	Splits the string with white space as seperator
*	And returns a slice with the split elements
*	"Nicole Smith" string	["Nicole","Smith"]

### Blank Identifier
*	underscore _ is used to ignore a variable you don't want to use
*	In Go you need to make unused variables explicit

### If Conditional True or False
*	the default condition is set to 'true'
*	there is only one If and one Else statement
*	there can be multiple else-if statements.

### Logical Operator !
*	the exclamation mark ! is called Logical **NOT** operator
*	! reverses the boolean value
*	if a condition is true, the ! now makes the condition false

### Switch Statement
*	Allows a variable to be tested for equality against a list of values
*	'default' statement handles the case, if no match is found

### Functions
*	**Encapsulate code** into own container (=func). Which *logically belong together*
*	Like variable name, the function needs to have a **descriptive name**
*	**Call the func by its name** whenever the block of code needs to be executed

### Functions: more Use cases
*	Group logic that belongs together
*	Reuse logic and so reducing duplication of code

### Parameters
*	Information can be *passed* into functions as parameters
*	Parameters are also called *arguments*

### Return Values
*	A function can **return** data as a result
*	So a function can take an input and return an output
*   A Go funct can return **multiple** values
*	Multiple return values **must** be listed in parentheses() with their *Data Type*

### Map Data Type
*	Has two components: key and value
*	Data values can be retrieved by its **Key** and the respective **Value**
*   All **keys** have the *same* data type
*   All **values** have the *same* data type
*   Give **keys** a descriptive name.

### Struct
*	Struct stands for 'Structure'
*	can hold mixed data types, for key/value pairs

### Concurrency
*   the code gets executed line-by-line in ascending order
*   is used to make the application more efficient
*   *concurrency* in Go is **cheap & easy**
*   in the code example: bookTicketInfo variable will execute in its own thread; and *NOT* interupt the main function process.
*   use **go** keyword infront of variable, this abstracts the information into its own thread.
*   **go** starts a new goroutine
*   A *goroutine* is a lightweight thread managed by the Go runtime.
*   *Syncronizing* the Goroutines
*   the **Main** application will ignore the seperate Goroutine thread, unless explicitly stated otherwise
*   use **WaitGroup** to inform the main function not to exit until the seperate threads have performed their tasks.
*   package *sync* provides basic synchronization functionality
*   **Add** sets the number of goroutines to wait for (increases the counter by the provided number)
*   **Wait** blocks until the WaitGroup counter is 0
*   **Done** decrements the WaitGroup counter by 1. So this is called by the goroutine to indicate that this function has finished its processes.
*   *Go* uses what is called a **Green thread**
*   this is an **Abstraction** of an acutal thread
    *   this abstraction is managed by the Go runtime, which is only interacting with the **high level** goroutines
    *   it is cheaper & lightweight
    *   therefore hundreds or thousands or millions of goroutines can be *run* without affecting the *performance* of the application.
    *   Go uses **Channels** which is a built-in functionality for goroutines to talk to each other
*   Other applications use Operating System Threads
    *   managed by the kernel
    *   hardware dependent
    *   *cost* of threads are higher
    *   Longer start up time
    *   no easy communication between threads

### 'time' functionality for time
*	The **sleep** function stops or blocks the current thread (goroutine) execution for the defined duration.


### Packages in Golang
*	Go programs are organized into *packages*
*	A *package* is a **collection** of Go files
*   *Multiple Packages in the application*
*   organize the app
*   logically goup the code
*   A *Function* needs to be **explicitly** exported
*   Make the fuction available for use with all packages in the application
*   The name of the *function* needs to be **Capitalized**
*   Variables can also be exported, they also need to be **Capitalized**

### Scope: Package level
*	Variables and Functions defined outside of any function, can be accessed in all other files **within the same package**

### Package Level Variables
*	Defined at the top **outside** all functions
*	They can be **accessed** inside any of the functions
*   All in all files, which are in the **same** package level

### Local Variables
*	Defined **inside** a function or a block of code
*	They can be **accessed** inside that function or block of code

