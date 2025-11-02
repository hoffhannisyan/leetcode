# 🧩 LeetCode Solutions in Go

This repository contains my personal solutions to [LeetCode](https://leetcode.com/) problems, implemented in **Go (Golang)**.

Each problem is organized as a separate file inside the `tasks/` directory and follows a simple, consistent structure.

The main program (`main.go`) runs all solved tasks sequentially and prints their outputs in a clear, readable format.

---

## 📁 Project Structure

```
leetcode/
├── go.mod
├── main.go
├── README.md
├── LICENSE
└── tasks/
    ├── 001_example_problem.go
    └── 001_example_problem_test.go
```

- **`main.go`** → Entry point to run all tasks
- **`tasks/`** → Contains individual LeetCode problem solutions
- **`*_test.go` files** → Contain unit tests for each task

---

## 🚀 How to Run

Run all implemented tasks:

```bash
go run main.go
```

Example output:

```
Running LeetCode tasks...

=== Task 001: Example Problem ===
✓ Passed

=== Task 002: Another Problem ===
✓ Passed

✅ All tasks executed successfully.
```

---

## 🧪 How to Run Tests

Run all tests in the tasks folder:

```bash
go test ./tasks/
```

Run tests with verbose output:

```bash
go test -v ./tasks/
```

Run a specific test by name:

```bash
go test ./tasks/ -run TestExampleProblem
```

Run tests and show coverage:

```bash
go test -cover ./tasks/
```

---

## 💡 Notes

- All tasks are self-contained and written in clean, idiomatic Go
- The goal of this project is to practice algorithms, learn Go, and maintain a neat and consistent structure

---

## 📜 License

This repository is open-source and available under the MIT License.

Made with ❤️ and Go.
