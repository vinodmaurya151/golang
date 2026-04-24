//1. Constants (Values that never change)

const pi = 3.14159
const (
	StatusActive   = "active"
	StatusInactive = "inactive"
)

// 2. Type Conversion (Converting between types)

var x int = 42
var y float64 = float64(x)     // Convert int to float
var z string = strconv.Itoa(x) // Convert int to string

// 3. Short Declaration := (Most common in Go)

// Instead of: var name string = "Vinod"
name := "Vinod" // So much cleaner!

// Multiple at once:
age, city, score := 30, "Mumbai", 95.5

// 4. Pointers (Memory addresses - very important)

x := 42
ptr := &x  // ptr holds address of x
*ptr = 100 // Now x becomes 100

// 5. Arrays vs Slices (Lists of data)

// Array (fixed size)
var arr [3]int = [3]int{1, 2, 3}

// Slice (dynamic size - USE THIS)
slice := []int{1, 2, 3}
slice = append(slice, 4) // Grows automatically

// 6. Maps (Key-value pairs - like dictionaries)

scores := map[string]int{
	"Vinod": 95,
	"Raj":   87,
}
scores["NewUser"] = 100

// 7. Structs (Group related data)

type Person struct {
	Name string
	Age  int
	City string
}

p := Person{Name: "Vinod", Age: 30, City: "Mumbai"}
fmt.Println(p.Name)

