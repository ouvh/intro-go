# Bookstore API Documentation 📚

## Overview 🌟

The **Bookstore API** follows a **history-based implementation** approach. This means that every entity is saved with its relationships at the moment of creation. Any updates to a foreign entity will **not trigger cascading updates** to the related entities. This ensures that the historical state of entities is preserved, making the system immutable for past records. ⏳

## Try it Now:
https://intro-go.onrender.com/books

### How to Run it

simply run the main.go in the root of the directory . You Can Also Customize the Period of the SnapShots (default to 5 second) and the Report Period (default 10 Second)






### Design Patterns Used 🎨

- **Facade Pattern**: The facade provides a simplified interface to manage storing and retrieving data. 🏗️
- **Dependency Injection**: Ensures that components remain loosely coupled, making the application easier to test and maintain. 🔌

---

## File Structure 🗂️

The application is structured into multiple directories, each serving a specific purpose:

### 1. **Models** 🏷️
Defines the core data structures (structs) used throughout the application. These structs represent the primary entities of the bookstore, such as `Book`, `Customer`, `Order`, and `SalesReport`. Each struct includes the necessary fields and their relationships with other entities.

**Example:**
```go
type Book struct {
    ID       int64  `json:"id"`
    Title    string `json:"title"`
    Author   string `json:"author"`
    Price    float64 `json:"price"`
} 
```



### 2. **Store** 🛒
Contains the Facade for managing the storage of data. This layer abstracts the complexities of handling data persistence and provides a simplified interface for:

Storing entities. 🗄️
Scheduling snapshots to periodically save data to a local JSON file. 💾
This ensures data persistence and recovery in case of server restarts or crashes.

### 3. **Services** 🔧
Defines the CRUD operations (Create, Read, Update, Delete) for the entities. This layer also incorporates a locking mechanism to ensure consistency of data during concurrent operations.

Key Features:
Thread-Safe Operations: Ensures that multiple users can perform CRUD operations without corrupting the data. 🔒
Consistency Guarantees: Prevents race conditions through proper locking mechanisms. 🏁
Example Service Operation:

``` go
func (s *BookService) CreateBook(book Book) error {
    s.Lock()
    defer s.Unlock()
    // Logic to add the book
}
```

### 4. ***HTTP Handlers*** 🌐
Defines the HTTP endpoints to allow communication with the server. These handlers act as the interface between the client (frontend or external systems) and the backend services.

Responsibilities:
Parsing and validating incoming HTTP requests. 📩
Invoking the appropriate service methods based on the request. 🛠️
Formatting and sending responses back to the client. 💬
Example Endpoints:

POST /books: Adds a new book. 📚
GET /books: Retrieves a list of books. 📖
5. Utilities 🧰
Contains helper functions and modules to support the application, such as:

JSON HTTP Parser 📑
Handles the parsing of JSON payloads from incoming HTTP requests.

Example:

``` go
func LoadJson(r *http.Request, DAO interface{}) error {

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		return fmt.Errorf("unable to read request body")
	}
	errr := json.Unmarshal(body, &DAO)
	if errr != nil {
		return fmt.Errorf("unable to parse json")
	}
	return nil

}
```

HTTP Struct Verifier ✅
Validates the structure of incoming HTTP requests to ensure required fields are present and properly formatted.

## Summary 🎉
This structured approach ensures:

Clear separation of concerns. 🛠️
Maintainable and testable code. 🔍
Data consistency and immutability through a history-based implementation. 🕰️


