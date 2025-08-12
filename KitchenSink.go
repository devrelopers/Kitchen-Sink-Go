// KitchenSink.go - A comprehensive demonstration of Go language features
// This file showcases most Go language constructs with detailed explanations

package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"math/big"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// ============================================================================
// BASIC TYPES AND VARIABLES
// ============================================================================

// Constants - immutable values determined at compile time
const (
	Pi         = 3.14159265359 // Untyped constant
	MaxInt     int = 1<<31 - 1 // Typed constant
	greeting   = "Hello"        // Private constant (lowercase)
	PublicConst = "World"       // Public constant (uppercase)
)

// Iota - Go's enum-like constant generator
const (
	Sunday = iota // 0
	Monday        // 1
	Tuesday       // 2
	Wednesday     // 3
	Thursday      // 4
	Friday        // 5
	Saturday      // 6
)

// Bit flags using iota
const (
	ReadPermission = 1 << iota // 1 << 0 = 1
	WritePermission             // 1 << 1 = 2
	ExecutePermission           // 1 << 2 = 4
)

// Global variables (package-level)
var (
	globalString   string = "I'm a global variable"
	globalInt      int    // Zero-valued to 0
	globalBool     bool   // Zero-valued to false
	globalFloat    float64
	globalComplex  complex128 = 3 + 4i // Complex number
)

// ============================================================================
// CUSTOM TYPES AND STRUCTS
// ============================================================================

// Type alias - creates a new name for existing type
type MyInt int
type StringAlias = string // True alias (Go 1.9+)

// Struct - composite type grouping data
type Person struct {
	Name       string    `json:"name" xml:"name"`           // Struct tags for metadata
	Age        int       `json:"age,omitempty"`             // Omit if empty in JSON
	Email      string    `json:"email" validate:"required"` // Custom validation tag
	internal   string    // Unexported field (private)
	Birthday   time.Time `json:"-"`                    // Exclude from JSON
	Nicknames  []string  `json:"nicknames,omitempty"`  // Slice field
	Attributes map[string]interface{} `json:"attrs"` // Map field
}

// Embedded struct (composition)
type Employee struct {
	Person     // Embedded/anonymous field - promotes Person's fields
	EmployeeID string
	Department string
	Salary     float64 `json:"-"` // Sensitive data, exclude from JSON
}

// Struct with anonymous fields
type Data struct {
	int     // Anonymous field
	string  // Anonymous field
	*Person // Anonymous pointer field
}

// ============================================================================
// INTERFACES
// ============================================================================

// Interface - defines method signatures
type Speaker interface {
	Speak() string
	Listen(words string) error
}

// Empty interface - can hold any type
type Any interface{} // Same as interface{} or any (Go 1.18+)

// Interface embedding
type ReadWriter interface {
	io.Reader // Embedded interface
	io.Writer // Embedded interface
	Close() error
}

// Type assertion interface
type Stringer interface {
	String() string
}

// ============================================================================
// METHODS AND RECEIVERS
// ============================================================================

// Value receiver method - operates on copy
func (p Person) GetAge() int {
	return p.Age
}

// Pointer receiver method - can modify the receiver
func (p *Person) SetAge(age int) {
	p.Age = age
}

// Method with multiple return values
func (p Person) Validate() (bool, error) {
	if p.Age < 0 {
		return false, errors.New("age cannot be negative")
	}
	if p.Name == "" {
		return false, errors.New("name is required")
	}
	return true, nil
}

// Implement Speaker interface for Person
func (p Person) Speak() string {
	return fmt.Sprintf("Hi, I'm %s and I'm %d years old", p.Name, p.Age)
}

func (p Person) Listen(words string) error {
	if words == "" {
		return errors.New("no words to listen to")
	}
	fmt.Printf("%s heard: %s\n", p.Name, words)
	return nil
}

// Implement Stringer interface (like ToString in other languages)
func (p Person) String() string {
	return fmt.Sprintf("Person{Name: %s, Age: %d}", p.Name, p.Age)
}

// Method on custom type
func (m MyInt) Double() MyInt {
	return m * 2
}

// ============================================================================
// FUNCTIONS
// ============================================================================

// Basic function with parameters and return value
func add(a, b int) int {
	return a + b
}

// Multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Named return values
func rectangleArea(length, width float64) (area float64, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return // Naked return - returns named values
}

// Variadic function - accepts variable number of arguments
func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

// Function as parameter (higher-order function)
func apply(fn func(int) int, value int) int {
	return fn(value)
}

// Function returning function (closure)
func multiplier(factor int) func(int) int {
	// Closure captures 'factor' from outer scope
	return func(x int) int {
		return x * factor
	}
}

// Recursive function
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Defer - delays execution until surrounding function returns
func deferExample() {
	defer fmt.Println("This prints last")
	defer fmt.Println("This prints second to last") // LIFO order
	fmt.Println("This prints first")
}

// Panic and recover - error handling mechanism
func panicRecoverExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()
	
	panic("Something went wrong!")
	fmt.Println("This won't execute")
}

// ============================================================================
// POINTERS
// ============================================================================

func pointerDemo() {
	var x int = 42
	var ptr *int = &x    // ptr holds address of x
	
	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Address of x: %p\n", &x)
	fmt.Printf("Value of ptr: %p\n", ptr)
	fmt.Printf("Value ptr points to: %d\n", *ptr) // Dereferencing
	
	*ptr = 100 // Modify value through pointer
	fmt.Printf("New value of x: %d\n", x)
	
	// Pointer to pointer
	var ptrToPtr **int = &ptr
	fmt.Printf("Pointer to pointer: %p\n", ptrToPtr)
	
	// nil pointer
	var nilPtr *string
	if nilPtr == nil {
		fmt.Println("Pointer is nil")
	}
}

// ============================================================================
// ARRAYS AND SLICES
// ============================================================================

func arraySliceDemo() {
	// Arrays - fixed size
	var arr1 [5]int                    // Zero-valued array
	arr2 := [5]int{1, 2, 3, 4, 5}     // Array literal
	arr3 := [...]int{1, 2, 3}         // Compiler counts elements
	arr4 := [5]int{1: 10, 3: 30}      // Sparse initialization
	
	fmt.Printf("Arrays: %v, %v, %v, %v\n", arr1, arr2, arr3, arr4)
	
	// Slices - dynamic arrays
	var slice1 []int                           // nil slice
	slice2 := []int{1, 2, 3}                  // Slice literal
	slice3 := make([]int, 5)                  // Length 5, capacity 5
	slice4 := make([]int, 5, 10)              // Length 5, capacity 10
	slice5 := arr2[1:4]                       // Slice from array
	
	// Slice operations
	slice2 = append(slice2, 4, 5, 6)          // Append elements
	slice6 := append(slice2, slice3...)        // Append slice to slice
	
	// Slice slicing
	subSlice := slice2[1:3]                   // Elements at index 1 and 2
	fullSlice := slice2[:]                    // All elements
	fromStart := slice2[:3]                   // First 3 elements
	toEnd := slice2[2:]                       // From index 2 to end
	
	// Capacity and length
	fmt.Printf("Length: %d, Capacity: %d\n", len(slice4), cap(slice4))
	
	// Copy slices
	dest := make([]int, len(slice2))
	copied := copy(dest, slice2)
	fmt.Printf("Copied %d elements\n", copied)
	
	// 2D slice
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Printf("2D slice: %v\n", matrix)
	
	_ = slice1
	_ = slice5
	_ = slice6
	_ = subSlice
	_ = fullSlice
	_ = fromStart
	_ = toEnd
}

// ============================================================================
// MAPS
// ============================================================================

func mapDemo() {
	// Map creation
	var map1 map[string]int                        // nil map
	map2 := make(map[string]int)                   // Empty map
	map3 := map[string]int{                        // Map literal
		"one":   1,
		"two":   2,
		"three": 3,
	}
	
	// Map operations
	map2["key"] = 42                              // Insert/update
	value := map3["two"]                          // Get value
	value, exists := map3["four"]                 // Check existence
	delete(map3, "one")                           // Delete key
	
	fmt.Printf("Value: %d, Exists: %v\n", value, exists)
	
	// Iterate over map
	for key, val := range map3 {
		fmt.Printf("%s: %d\n", key, val)
	}
	
	// Map of structs
	people := map[string]Person{
		"p1": {Name: "Alice", Age: 30},
		"p2": {Name: "Bob", Age: 25},
	}
	
	// Nested maps
	nested := map[string]map[string]int{
		"group1": {"a": 1, "b": 2},
		"group2": {"c": 3, "d": 4},
	}
	
	_ = map1
	_ = people
	_ = nested
}

// ============================================================================
// CHANNELS AND GOROUTINES
// ============================================================================

func concurrencyDemo() {
	// Goroutines - lightweight threads
	go func() {
		fmt.Println("Running in goroutine")
	}()
	
	// Channels - communication between goroutines
	ch1 := make(chan int)          // Unbuffered channel
	ch2 := make(chan string, 10)   // Buffered channel
	
	// Send and receive
	go func() {
		ch1 <- 42  // Send
	}()
	value := <-ch1  // Receive
	
	// Channel directions
	sendOnly := make(chan<- int)    // Send-only channel
	receiveOnly := make(<-chan int) // Receive-only channel
	
	// Select statement - multiplex channels
	ch3 := make(chan int)
	ch4 := make(chan string)
	
	go func() {
		ch3 <- 1
	}()
	
	go func() {
		ch4 <- "hello"
	}()
	
	select {
	case v1 := <-ch3:
		fmt.Printf("Received from ch3: %d\n", v1)
	case v2 := <-ch4:
		fmt.Printf("Received from ch4: %s\n", v2)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")
	default:
		fmt.Println("No channel ready")
	}
	
	// Close channel
	ch5 := make(chan int, 3)
	ch5 <- 1
	ch5 <- 2
	ch5 <- 3
	close(ch5)
	
	// Range over channel
	for val := range ch5 {
		fmt.Printf("Received: %d\n", val)
	}
	
	// Worker pool pattern
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	
	// Start workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	
	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	
	// Collect results
	for a := 1; a <= 5; a++ {
		<-results
	}
	
	_ = value
	_ = ch2
	_ = sendOnly
	_ = receiveOnly
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(100 * time.Millisecond)
		results <- j * 2
	}
}

// ============================================================================
// SYNCHRONIZATION PRIMITIVES
// ============================================================================

func syncDemo() {
	// Mutex - mutual exclusion
	var mu sync.Mutex
	var counter int
	
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	
	// RWMutex - readers-writers lock
	var rwMu sync.RWMutex
	var data map[string]string = make(map[string]string)
	
	// Multiple readers
	go func() {
		rwMu.RLock()
		_ = data["key"]
		rwMu.RUnlock()
	}()
	
	// Single writer
	go func() {
		rwMu.Lock()
		data["key"] = "value"
		rwMu.Unlock()
	}()
	
	// WaitGroup - wait for goroutines to finish
	var wg sync.WaitGroup
	
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d working\n", id)
		}(i)
	}
	wg.Wait()
	
	// Once - run function only once
	var once sync.Once
	initialize := func() {
		fmt.Println("Initialized only once")
	}
	
	for i := 0; i < 3; i++ {
		once.Do(initialize)
	}
	
	// Atomic operations
	var atomicCounter int64
	atomic.AddInt64(&atomicCounter, 1)
	atomic.LoadInt64(&atomicCounter)
	atomic.StoreInt64(&atomicCounter, 100)
	atomic.CompareAndSwapInt64(&atomicCounter, 100, 200)
	
	// Condition variable
	cond := sync.NewCond(&mu)
	ready := false
	
	go func() {
		mu.Lock()
		for !ready {
			cond.Wait()
		}
		mu.Unlock()
		fmt.Println("Condition met!")
	}()
	
	mu.Lock()
	ready = true
	cond.Signal()
	mu.Unlock()
	
	// Pool - reusable object pool
	pool := &sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
	
	buf := pool.Get().(*bytes.Buffer)
	buf.WriteString("Hello")
	pool.Put(buf)
}

// ============================================================================
// ERROR HANDLING
// ============================================================================

// Custom error type
type CustomError struct {
	Code    int
	Message string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// Wrapping errors (Go 1.13+)
func errorHandlingDemo() error {
	// Creating errors
	err1 := errors.New("simple error")
	err2 := fmt.Errorf("formatted error: %d", 42)
	
	// Custom error
	err3 := CustomError{Code: 404, Message: "Not Found"}
	
	// Error wrapping
	originalErr := errors.New("original error")
	wrappedErr := fmt.Errorf("wrapped: %w", originalErr)
	
	// Check wrapped errors
	if errors.Is(wrappedErr, originalErr) {
		fmt.Println("Error matches original")
	}
	
	// Type assertion for errors
	var customErr CustomError
	if errors.As(err3, &customErr) {
		fmt.Printf("Custom error code: %d\n", customErr.Code)
	}
	
	// Multiple error handling
	errs := []error{err1, err2, err3}
	for _, err := range errs {
		if err != nil {
			log.Printf("Error occurred: %v", err)
		}
	}
	
	return nil
}

// ============================================================================
// GENERICS (Go 1.18+)
// ============================================================================

// Generic function
func Min[T comparable](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Generic type constraint
type Number interface {
	int | int32 | int64 | float32 | float64
}

func Sum[T Number](values []T) T {
	var sum T
	for _, v := range values {
		sum += v
	}
	return sum
}

// Generic struct
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// ============================================================================
// REFLECTION
// ============================================================================

func reflectionDemo() {
	p := Person{Name: "Alice", Age: 30}
	
	// Get type information
	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)
	
	fmt.Printf("Type: %v, Kind: %v\n", t, t.Kind())
	
	// Iterate struct fields
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fmt.Printf("Field: %s, Type: %s, Value: %v, Tag: %s\n",
			field.Name, field.Type, value, field.Tag)
	}
	
	// Modify value through reflection (requires pointer)
	ptrValue := reflect.ValueOf(&p).Elem()
	ageField := ptrValue.FieldByName("Age")
	if ageField.CanSet() {
		ageField.SetInt(31)
	}
	
	// Call method through reflection
	method := v.MethodByName("Speak")
	if method.IsValid() {
		results := method.Call(nil)
		fmt.Printf("Method result: %v\n", results[0])
	}
	
	// Check if implements interface
	speakerType := reflect.TypeOf((*Speaker)(nil)).Elem()
	implements := t.Implements(speakerType)
	fmt.Printf("Implements Speaker: %v\n", implements)
}

// ============================================================================
// CONTEXT
// ============================================================================

func contextDemo() {
	// Create contexts
	ctx := context.Background()                            // Root context
	ctxWithValue := context.WithValue(ctx, "key", "value") // With value
	
	// With cancellation
	ctxWithCancel, cancel := context.WithCancel(ctx)
	defer cancel() // Ensure cleanup
	
	// With deadline
	deadline := time.Now().Add(5 * time.Second)
	ctxWithDeadline, cancel2 := context.WithDeadline(ctx, deadline)
	defer cancel2()
	
	// With timeout
	ctxWithTimeout, cancel3 := context.WithTimeout(ctx, 2*time.Second)
	defer cancel3()
	
	// Use context in goroutine
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Printf("Context cancelled: %v\n", ctx.Err())
		case <-time.After(1 * time.Second):
			fmt.Println("Operation completed")
		}
	}(ctxWithTimeout)
	
	// Get value from context
	if value := ctxWithValue.Value("key"); value != nil {
		fmt.Printf("Context value: %v\n", value)
	}
	
	_ = ctxWithCancel
	_ = ctxWithDeadline
}

// ============================================================================
// FILE I/O
// ============================================================================

func fileIODemo() error {
	// Write to file
	data := []byte("Hello, Go!\n")
	err := os.WriteFile("test.txt", data, 0644)
	if err != nil {
		return err
	}
	
	// Read entire file
	content, err := os.ReadFile("test.txt")
	if err != nil {
		return err
	}
	fmt.Printf("File content: %s", content)
	
	// Open file for reading
	file, err := os.Open("test.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	
	// Read with buffer
	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return err
	}
	fmt.Printf("First line: %s", line)
	
	// Create/open file for writing
	writeFile, err := os.Create("output.txt")
	if err != nil {
		return err
	}
	defer writeFile.Close()
	
	// Write with buffer
	writer := bufio.NewWriter(writeFile)
	writer.WriteString("Buffered write\n")
	writer.Flush()
	
	// File info
	info, err := os.Stat("test.txt")
	if err != nil {
		return err
	}
	fmt.Printf("File: %s, Size: %d, Modified: %v\n",
		info.Name(), info.Size(), info.ModTime())
	
	// Directory operations
	err = os.Mkdir("testdir", 0755)
	if err != nil && !os.IsExist(err) {
		return err
	}
	
	// List directory
	entries, err := os.ReadDir(".")
	if err != nil {
		return err
	}
	for _, entry := range entries {
		fmt.Printf("Entry: %s, IsDir: %v\n", entry.Name(), entry.IsDir())
	}
	
	// Clean up
	os.Remove("test.txt")
	os.Remove("output.txt")
	os.RemoveAll("testdir")
	
	return nil
}

// ============================================================================
// JSON HANDLING
// ============================================================================

func jsonDemo() error {
	// Struct to JSON
	p := Person{
		Name:      "Alice",
		Age:       30,
		Email:     "alice@example.com",
		Nicknames: []string{"Ally", "Al"},
		Attributes: map[string]interface{}{
			"height": 170,
			"city":   "New York",
		},
	}
	
	// Marshal to JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
		return err
	}
	fmt.Printf("JSON: %s\n", jsonData)
	
	// Pretty print JSON
	prettyJSON, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	fmt.Printf("Pretty JSON:\n%s\n", prettyJSON)
	
	// Unmarshal from JSON
	var p2 Person
	err = json.Unmarshal(jsonData, &p2)
	if err != nil {
		return err
	}
	fmt.Printf("Unmarshaled: %+v\n", p2)
	
	// JSON with interface{}
	var data interface{}
	jsonStr := `{"name":"Bob","age":25,"hobbies":["reading","gaming"]}`
	err = json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return err
	}
	
	// Type assertion for dynamic JSON
	if m, ok := data.(map[string]interface{}); ok {
		fmt.Printf("Name: %v\n", m["name"])
		if hobbies, ok := m["hobbies"].([]interface{}); ok {
			for _, hobby := range hobbies {
				fmt.Printf("Hobby: %v\n", hobby)
			}
		}
	}
	
	// JSON encoder/decoder for streams
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.Encode(p)
	
	decoder := json.NewDecoder(&buf)
	var p3 Person
	decoder.Decode(&p3)
	
	return nil
}

// ============================================================================
// HTTP CLIENT/SERVER
// ============================================================================

func httpDemo() {
	// HTTP Server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})
	
	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Hello, JSON!",
		})
	})
	
	// Start server in goroutine
	go func() {
		log.Println("Server starting on :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}()
	
	// Give server time to start
	time.Sleep(100 * time.Millisecond)
	
	// HTTP Client
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	
	// GET request
	resp, err := client.Get("http://localhost:8080/World")
	if err != nil {
		log.Printf("GET error: %v", err)
		return
	}
	defer resp.Body.Close()
	
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Response: %s\n", body)
	
	// POST request with JSON
	jsonData := bytes.NewBuffer([]byte(`{"name":"test"}`))
	resp2, err := client.Post("http://localhost:8080/json", 
		"application/json", jsonData)
	if err != nil {
		log.Printf("POST error: %v", err)
		return
	}
	defer resp2.Body.Close()
}

// ============================================================================
// REGULAR EXPRESSIONS
// ============================================================================

func regexDemo() {
	// Compile regex
	re, err := regexp.Compile(`\d+`)
	if err != nil {
		log.Fatal(err)
	}
	
	text := "The year 2024 has 365 days"
	
	// Find matches
	match := re.FindString(text)
	fmt.Printf("First match: %s\n", match)
	
	allMatches := re.FindAllString(text, -1)
	fmt.Printf("All matches: %v\n", allMatches)
	
	// Replace
	replaced := re.ReplaceAllString(text, "XXX")
	fmt.Printf("Replaced: %s\n", replaced)
	
	// Named groups
	re2 := regexp.MustCompile(`(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})`)
	match2 := re2.FindStringSubmatch("2024-03-15")
	fmt.Printf("Date parts: %v\n", match2)
	
	// Check if matches
	matched := re.MatchString("123")
	fmt.Printf("Matches: %v\n", matched)
}

// ============================================================================
// STRING MANIPULATION
// ============================================================================

func stringDemo() {
	s := "Hello, Go World!"
	
	// String operations
	fmt.Printf("Length: %d\n", len(s))
	fmt.Printf("Uppercase: %s\n", strings.ToUpper(s))
	fmt.Printf("Lowercase: %s\n", strings.ToLower(s))
	fmt.Printf("Title: %s\n", strings.Title(s))
	
	// Substring operations
	fmt.Printf("Contains 'Go': %v\n", strings.Contains(s, "Go"))
	fmt.Printf("Starts with 'Hello': %v\n", strings.HasPrefix(s, "Hello"))
	fmt.Printf("Ends with '!': %v\n", strings.HasSuffix(s, "!"))
	fmt.Printf("Index of 'Go': %d\n", strings.Index(s, "Go"))
	
	// Split and join
	parts := strings.Split(s, " ")
	fmt.Printf("Split: %v\n", parts)
	joined := strings.Join(parts, "-")
	fmt.Printf("Joined: %s\n", joined)
	
	// Trim operations
	s2 := "  trimmed  "
	fmt.Printf("Trim: '%s'\n", strings.TrimSpace(s2))
	fmt.Printf("TrimLeft: '%s'\n", strings.TrimLeft(s2, " "))
	fmt.Printf("TrimRight: '%s'\n", strings.TrimRight(s2, " "))
	
	// Replace
	replaced := strings.Replace(s, "Go", "Golang", 1)
	fmt.Printf("Replaced: %s\n", replaced)
	replacedAll := strings.ReplaceAll(s, "o", "0")
	fmt.Printf("Replaced all: %s\n", replacedAll)
	
	// String builder for efficient concatenation
	var builder strings.Builder
	builder.WriteString("Building")
	builder.WriteString(" a")
	builder.WriteString(" string")
	fmt.Printf("Built: %s\n", builder.String())
	
	// Runes (Unicode code points)
	s3 := "Hello 世界"
	runes := []rune(s3)
	fmt.Printf("Runes: %v\n", runes)
	fmt.Printf("Rune count: %d\n", len(runes))
	
	// Iterate over runes
	for i, r := range s3 {
		fmt.Printf("Position %d: %c (U+%04X)\n", i, r, r)
	}
}

// ============================================================================
// TIME AND DATE
// ============================================================================

func timeDemo() {
	// Current time
	now := time.Now()
	fmt.Printf("Now: %v\n", now)
	
	// Time components
	fmt.Printf("Year: %d, Month: %s, Day: %d\n", 
		now.Year(), now.Month(), now.Day())
	fmt.Printf("Hour: %d, Minute: %d, Second: %d\n",
		now.Hour(), now.Minute(), now.Second())
	
	// Time formatting
	fmt.Printf("RFC3339: %s\n", now.Format(time.RFC3339))
	fmt.Printf("Custom: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("Kitchen: %s\n", now.Format(time.Kitchen))
	
	// Parse time
	parsed, err := time.Parse("2006-01-02", "2024-03-15")
	if err != nil {
		log.Printf("Parse error: %v", err)
	}
	fmt.Printf("Parsed: %v\n", parsed)
	
	// Duration
	duration := 2*time.Hour + 30*time.Minute + 45*time.Second
	fmt.Printf("Duration: %v\n", duration)
	fmt.Printf("In seconds: %.0f\n", duration.Seconds())
	
	// Add/subtract time
	future := now.Add(24 * time.Hour)
	past := now.Add(-7 * 24 * time.Hour)
	fmt.Printf("Tomorrow: %v\n", future)
	fmt.Printf("Week ago: %v\n", past)
	
	// Time comparison
	if future.After(now) {
		fmt.Println("Future is after now")
	}
	if past.Before(now) {
		fmt.Println("Past is before now")
	}
	
	// Sleep
	fmt.Println("Sleeping for 100ms...")
	time.Sleep(100 * time.Millisecond)
	
	// Timer
	timer := time.NewTimer(200 * time.Millisecond)
	<-timer.C
	fmt.Println("Timer expired")
	
	// Ticker
	ticker := time.NewTicker(100 * time.Millisecond)
	count := 0
	for range ticker.C {
		count++
		fmt.Printf("Tick %d\n", count)
		if count >= 3 {
			ticker.Stop()
			break
		}
	}
}

// ============================================================================
// MATH OPERATIONS
// ============================================================================

func mathDemo() {
	// Basic math functions
	fmt.Printf("Abs(-5): %.0f\n", math.Abs(-5))
	fmt.Printf("Ceil(4.3): %.0f\n", math.Ceil(4.3))
	fmt.Printf("Floor(4.7): %.0f\n", math.Floor(4.7))
	fmt.Printf("Round(4.5): %.0f\n", math.Round(4.5))
	fmt.Printf("Max(3, 7): %.0f\n", math.Max(3, 7))
	fmt.Printf("Min(3, 7): %.0f\n", math.Min(3, 7))
	
	// Power and roots
	fmt.Printf("Pow(2, 8): %.0f\n", math.Pow(2, 8))
	fmt.Printf("Sqrt(16): %.0f\n", math.Sqrt(16))
	fmt.Printf("Cbrt(27): %.0f\n", math.Cbrt(27))
	
	// Trigonometry
	angle := math.Pi / 4 // 45 degrees in radians
	fmt.Printf("Sin(45°): %.2f\n", math.Sin(angle))
	fmt.Printf("Cos(45°): %.2f\n", math.Cos(angle))
	fmt.Printf("Tan(45°): %.2f\n", math.Tan(angle))
	
	// Logarithms
	fmt.Printf("Log(e): %.2f\n", math.Log(math.E))
	fmt.Printf("Log10(100): %.0f\n", math.Log10(100))
	fmt.Printf("Log2(8): %.0f\n", math.Log2(8))
	
	// Random numbers
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	fmt.Printf("Random bytes: %x\n", randBytes)
	
	// Big numbers
	big1 := big.NewInt(1234567890123456789)
	big2 := big.NewInt(9876543210987654321)
	bigSum := new(big.Int).Add(big1, big2)
	fmt.Printf("Big sum: %s\n", bigSum.String())
}

// ============================================================================
// CRYPTO OPERATIONS
// ============================================================================

func cryptoDemo() {
	// MD5 hash
	data := []byte("Hello, Go!")
	hash := md5.Sum(data)
	fmt.Printf("MD5: %x\n", hash)
	
	// Using hasher
	h := md5.New()
	h.Write([]byte("Hello"))
	h.Write([]byte(", Go!"))
	hashSum := h.Sum(nil)
	fmt.Printf("MD5 (hasher): %x\n", hashSum)
	
	// Hex encoding
	encoded := hex.EncodeToString(hashSum)
	fmt.Printf("Hex encoded: %s\n", encoded)
	
	decoded, _ := hex.DecodeString(encoded)
	fmt.Printf("Hex decoded: %x\n", decoded)
}

// ============================================================================
// TYPE CONVERSIONS
// ============================================================================

func conversionDemo() {
	// Numeric conversions
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(i)
	var i8 int8 = int8(i)
	
	fmt.Printf("int: %d, float64: %.2f, uint: %d, int8: %d\n", i, f, u, i8)
	
	// String conversions
	s1 := strconv.Itoa(42)                    // int to string
	s2 := strconv.FormatFloat(3.14, 'f', 2, 64) // float to string
	s3 := strconv.FormatBool(true)            // bool to string
	
	fmt.Printf("Strings: %s, %s, %s\n", s1, s2, s3)
	
	// Parse strings
	i1, _ := strconv.Atoi("42")               // string to int
	f1, _ := strconv.ParseFloat("3.14", 64)   // string to float
	b1, _ := strconv.ParseBool("true")        // string to bool
	
	fmt.Printf("Parsed: %d, %.2f, %v\n", i1, f1, b1)
	
	// Byte/string conversions
	bytes := []byte("Hello")
	str := string(bytes)
	fmt.Printf("Bytes: %v, String: %s\n", bytes, str)
	
	// Interface conversions (type assertions)
	var iface interface{} = "Hello"
	s, ok := iface.(string)
	if ok {
		fmt.Printf("Interface to string: %s\n", s)
	}
	
	// Type switches
	checkType := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("Integer: %d\n", v)
		case string:
			fmt.Printf("String: %s\n", v)
		case bool:
			fmt.Printf("Boolean: %v\n", v)
		default:
			fmt.Printf("Unknown type: %T\n", v)
		}
	}
	
	checkType(42)
	checkType("hello")
	checkType(true)
	checkType(3.14)
}

// ============================================================================
// UNSAFE OPERATIONS
// ============================================================================

func unsafeDemo() {
	// Sizeof
	var x int32 = 42
	size := unsafe.Sizeof(x)
	fmt.Printf("Size of int32: %d bytes\n", size)
	
	// Alignof
	type S struct {
		a bool
		b int32
		c int64
	}
	var s S
	fmt.Printf("Alignment of S: %d\n", unsafe.Alignof(s))
	fmt.Printf("Offset of b: %d\n", unsafe.Offsetof(s.b))
	
	// Pointer arithmetic (use with caution!)
	arr := [3]int{1, 2, 3}
	ptr := &arr[0]
	
	// Convert to unsafe.Pointer then to uintptr for arithmetic
	unsafePtr := unsafe.Pointer(ptr)
	uptr := uintptr(unsafePtr)
	
	// Move to next element
	nextPtr := unsafe.Pointer(uptr + unsafe.Sizeof(arr[0]))
	nextValue := *(*int)(nextPtr)
	fmt.Printf("Next value: %d\n", nextValue)
}

// ============================================================================
// TESTING HELPERS
// ============================================================================

// Benchmark helper function
func benchmarkFunction(fn func()) time.Duration {
	start := time.Now()
	fn()
	return time.Since(start)
}

// Memory usage helper
func printMemStats(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s - Memory: Alloc=%v MB, Sys=%v MB, NumGC=%v\n",
		label,
		m.Alloc/1024/1024,
		m.Sys/1024/1024,
		m.NumGC)
}

// ============================================================================
// SORTING
// ============================================================================

func sortingDemo() {
	// Sort basic types
	ints := []int{3, 1, 4, 1, 5, 9, 2, 6}
	sort.Ints(ints)
	fmt.Printf("Sorted ints: %v\n", ints)
	
	strings := []string{"banana", "apple", "cherry"}
	sort.Strings(strings)
	fmt.Printf("Sorted strings: %v\n", strings)
	
	// Custom sorting
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}
	
	// Sort by age
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Printf("Sorted by age: %v\n", people)
	
	// Sort by name
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	fmt.Printf("Sorted by name: %v\n", people)
	
	// Check if sorted
	isSorted := sort.IntsAreSorted(ints)
	fmt.Printf("Ints are sorted: %v\n", isSorted)
	
	// Binary search (on sorted slice)
	index := sort.SearchInts(ints, 5)
	fmt.Printf("Index of 5: %d\n", index)
}

// ============================================================================
// RUNTIME OPERATIONS
// ============================================================================

func runtimeDemo() {
	// Runtime information
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Arch: %s\n", runtime.GOARCH)
	fmt.Printf("CPUs: %d\n", runtime.NumCPU())
	fmt.Printf("Goroutines: %d\n", runtime.NumGoroutine())
	fmt.Printf("Go version: %s\n", runtime.Version())
	
	// Caller information
	pc, file, line, ok := runtime.Caller(0)
	if ok {
		fn := runtime.FuncForPC(pc)
		fmt.Printf("Called from: %s:%d in %s\n", file, line, fn.Name())
	}
	
	// Stack trace
	buf := make([]byte, 1024)
	n := runtime.Stack(buf, false)
	fmt.Printf("Stack trace:\n%s\n", buf[:n])
	
	// Garbage collection
	runtime.GC() // Force garbage collection
	printMemStats("After GC")
	
	// Set max procs
	runtime.GOMAXPROCS(2)
	
	// Yield processor
	runtime.Gosched()
}

// ============================================================================
// INIT FUNCTION
// ============================================================================

// init functions run before main
func init() {
	fmt.Println("First init function")
}

func init() {
	fmt.Println("Second init function - multiple init functions are allowed")
}

// ============================================================================
// MAIN FUNCTION
// ============================================================================

func main() {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("KITCHEN SINK GO - Comprehensive Go Language Demo")
	fmt.Println(strings.Repeat("=", 80))
	
	// Basic demos
	fmt.Println("\n>>> Variables and Constants Demo")
	fmt.Printf("Global string: %s\n", globalString)
	fmt.Printf("Pi constant: %f\n", Pi)
	fmt.Printf("Day of week: %d\n", Wednesday)
	
	// Function demos
	fmt.Println("\n>>> Function Demos")
	fmt.Printf("Add(3, 4) = %d\n", add(3, 4))
	result, err := divide(10, 2)
	if err == nil {
		fmt.Printf("Divide(10, 2) = %.2f\n", result)
	}
	fmt.Printf("Sum(1,2,3,4,5) = %d\n", sum(1, 2, 3, 4, 5))
	
	double := multiplier(2)
	fmt.Printf("Double(5) = %d\n", double(5))
	
	deferExample()
	panicRecoverExample()
	
	// Struct and method demos
	fmt.Println("\n>>> Struct and Method Demos")
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
	}
	fmt.Println(person.Speak())
	person.SetAge(31)
	fmt.Printf("New age: %d\n", person.GetAge())
	
	// Type demos
	fmt.Println("\n>>> Type Demos")
	var myNum MyInt = 42
	fmt.Printf("MyInt doubled: %d\n", myNum.Double())
	
	// Collection demos
	fmt.Println("\n>>> Collections")
	arraySliceDemo()
	mapDemo()
	
	// Concurrency
	fmt.Println("\n>>> Concurrency")
	concurrencyDemo()
	syncDemo()
	
	// Context
	fmt.Println("\n>>> Context")
	contextDemo()
	
	// Error handling
	fmt.Println("\n>>> Error Handling")
	errorHandlingDemo()
	
	// Generics
	fmt.Println("\n>>> Generics")
	fmt.Printf("Min(3, 5) = %d\n", Min(3, 5))
	fmt.Printf("Sum([1,2,3,4,5]) = %d\n", Sum([]int{1, 2, 3, 4, 5}))
	
	stack := Stack[string]{}
	stack.Push("first")
	stack.Push("second")
	if val, ok := stack.Pop(); ok {
		fmt.Printf("Popped: %s\n", val)
	}
	
	// Reflection
	fmt.Println("\n>>> Reflection")
	reflectionDemo()
	
	// Pointers
	fmt.Println("\n>>> Pointers")
	pointerDemo()
	
	// I/O Operations
	fmt.Println("\n>>> File I/O")
	if err := fileIODemo(); err != nil {
		log.Printf("File I/O error: %v", err)
	}
	
	// JSON
	fmt.Println("\n>>> JSON")
	if err := jsonDemo(); err != nil {
		log.Printf("JSON error: %v", err)
	}
	
	// HTTP (commented out to avoid starting server)
	// fmt.Println("\n>>> HTTP")
	// httpDemo()
	
	// Regular expressions
	fmt.Println("\n>>> Regular Expressions")
	regexDemo()
	
	// Strings
	fmt.Println("\n>>> String Operations")
	stringDemo()
	
	// Time
	fmt.Println("\n>>> Time and Date")
	timeDemo()
	
	// Math
	fmt.Println("\n>>> Math Operations")
	mathDemo()
	
	// Crypto
	fmt.Println("\n>>> Crypto Operations")
	cryptoDemo()
	
	// Type conversions
	fmt.Println("\n>>> Type Conversions")
	conversionDemo()
	
	// Unsafe
	fmt.Println("\n>>> Unsafe Operations")
	unsafeDemo()
	
	// Sorting
	fmt.Println("\n>>> Sorting")
	sortingDemo()
	
	// Runtime
	fmt.Println("\n>>> Runtime Information")
	runtimeDemo()
	
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("Demo completed successfully!")
	fmt.Println(strings.Repeat("=", 80))
}