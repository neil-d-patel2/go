declaration names a program entity and specifies some or all of its properties, there are four major kinds of declaration: var, const, type, and func

go program files end in .go, and start with a package declaration that says what package that file is apart of. The package declaration is followed by import declarations, and then a sequence of package-level declarations of types, variables, constants, and functions.. 

Integer Types: int8, int16, int32, int64: determines the range of each integer, 

int8: -128,127, int16: -32,768 to 32,767, int32: -2.1 million etcetc.

var declaration creates variable of particular type attatches name to it and sets initial val

var name type = expression

either type or expression part may be omitted but not both 

if expression is omitted, initial value is 0, initial type is determined by initalizer expression 

In Go, there is no such thing as an initalized value. ftoc.go

Short variable declaration!!

it takes the form name := expression, and the type of name is determined by the type of expression
here are many examples

freq := rand.Float64() * 3.0
t := 0.0

multiple can be initalized at the same time i,j := 0,1

but declarations with multiple expressions should be used only when they help readability 

Like ordinary var variables, short variable declarations are helpful when we want to call functions like os.Open that return multiple variables!!

f, err := os.Open(name)
if err != nil {
    return err
}
...

f.Close()

Pointers: variable is a piece of storage containing a value, a pointer value is the addres of a variable, and a pointer is thus the location at which a value is stored.

not every value has an addres, but every variable does 


basically the same syntax as c and cpp so nothing much to learn here its very standard and simple 

x := 1
p := &x
fmt.Println(*p) // *p = 1
*p = 2
fmt.Println(x) // x = 2

One interesting thing: ***it is perfectly safe for a function to return the address of a local variable***, for instance, in the code below, the local variable v created by this particular call to f will reamin in existence even after the call has returned 

var p = f()

func f() *int {
    v := 1
    return &v
}

each call of f returns a distinct value

fmt.Println(f() == f()) // "false"

pointers are the key for go's flag package which we will see in echo4










