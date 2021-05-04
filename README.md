# Standard-Library-Golang

__GO STANDARD LIBRARY__

**io.TeeReader**

- Is inspired by the **tee** linux command to duplicate the input whenever we read.
- Its like a multiplexer of a data stream.
- Copies input data to 1 or more files.
- The reader is a stream of data and it accepts 1 writer. (You can use a multiWriter to write to multiple files)
- Depending on the use-case we can opt to discard the file after working on it using 
```go
io.Copy(ioutil.Discard, io.TeeReader(res.Body, &counter{}))
```
-   This is like a `dev/null` way of discarding the data.

***

**io.Writer Interface**

- It only has 1 method wich is the `write` method
- The method takes a byte slice and returns an `int` which is usually the amount of bytes written and an `error`.
- You can see this implemented in the `os.File` package.
- Another io.writer implementation is `json.Encoder`
- Even the `http.ResponseWriter` is an implementation of the io.writer
- What kind of data are we going to pass into the writer?

_Bytes.Buffer_
- This is a generic and abstract buffer which can be used in a number of scenarios
- Its an implementation of the `io.Reader` and the `io.Writer` interfaces.
- This can be thought of as a **memory** buffer that keeps a file in memory.
- It can grow dynamically.
- `buf.Bytes()` returns the contents of the buffer.

***

**io.Reader Interface**

- Implements just one method which is the `Read` method.
- This method is defined over a stream of data that can be a file, network connection, string, any type of data
- It accepts a byte slice.
- Return `n` which is number of successfully read bytes and `error`. Whenever there is no more data to be read in the file you receive an `io.EOF` error to indicate end of file.
- You can see this implemented in the `os.File` package.
- Implementation of the reader interface on a file
    - We copy data from the file into a `buffer`
  
  ```go
  // This creates a read buffer which dictates the amount of bytes to read on each cycle.
  bf := make([]byte, bs) 
	for {
		n, err := reader.Read(bf)
  
  // This statement signifies that we have reached the end of file
		if err == io.EOF{
			break
		}
		if err != nil{
			fmt.Println(err)
			break
		}
        
  //Indicates that we do have data in the buffer and can read according to the buffer size allocated.
		if n > 0 {
			fmt.Println(string(bf[:n]))
		}
	}
  ```
- You can also use the `strings.Reader`
- One other implementation is using the `net.Connection` package
- We can read the data from the connection as if its a file.

***

**Go Maps (Hashmaps)**
- We need to use the map keyword to define a map
- We need to define the `key type` along with the `value type`
```go
map[keyType]valueType
m := make(map[string][]string)
//Or
m := map[string][]string{}
//Or
m := map[string][]string{
	"string1": []string{"new", "string"}
}
```
- You cannot use a `map`, `func`, or `slice` as a keyValue in a map.
- You can use a `bool` or `struct` if required
- You can `add`, `delete`, and `get`
- The delete does not return any error or panic if key does not exist.
- declaring a map using the var statement initialises a nil map and adding items will not work
```go
var m map[string][]string //this will panic when adding data to m.
```
- Deleting a nil map does not give an error either. Always use `make` to create a map.
***

**Context Package**
- This is a type mainly used for deadlines and cancellation signals.
- It propergates cancellation signals across API boundaries ( different parts of your application)
- Also used for storing some type of values which are request scoped.
- Very loose way of storing data.
- You always pass the context to the funtion that you want to control the signal.
```go
//Parent Context that can be passed onto children.
rootCtx := context.Background()
ctx, cancel := context.WithTimeout(rootCtx, time.Duration(timeout) * time.Millisecond)
```
- The cancel function is used to release all the resources associated with the context.
- Usually called with defer to be called at the end of the function.
- The cancellation can be function based
- This is done by calling the cancel function.
- We could also use `context.WithCancel` only waits for cancel signal.

***

**Json Encoding**

- We have the `marshal` and the `unmarshal` as well as the `encode` and `decode` when we have stream data.
- Encode and decode use a buffered memory.
- Marshal creates json from any resource
- HTML is automatically escaped through the Marshal method but it can be disabled if required.

***

**io.MultiWriter**
- Allows us to duplicate writes to multiple files
- We could write data to stdOut and json file at the same time.

***

**io.MultiReader**

- Pretty similar to the multiWriter
- Allows to provide multiple readers into the read function
- When we want to read multiple files of the same type and structure.





