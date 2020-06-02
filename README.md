# Learning Go

This repository contains programs and materials used for Go training.

# Repo structure
Broadly, based on the target audience, programs in this repository are categorized into 3 levels as follows:

| Level | Target Audience | Assumptions |
| ----- | --------------- | ---------- |
| Fundamentals | Learners who do not have any background in Go |Learners have some exposure to other programming langauages|
| Intermediate | Learners who have some understanding about Go |Learners are familiar with basic concepts in Go|
| Advanced | Learners who have good understanding about Go |Learners are familiar with Go and have written good amount of programs in Go|

A few additional points to note:

 Each level is furthur divided into topics and examples. 
- Topics and examples within each topic are arranged in a sequential order. It is highly recommanded to follow the order as it appears. 
- Each example assumes that the learner has gone through previous topics/examples. 
- Each example program (*.go) can be executed as a standalone program using `go run file-name.go`
    - All of the example programs are part of `main` package and contains `main` function

*(Note: In a real life Go project, only one program should contain `main` function. But, in this repository, each program contains a `main` function. This approach is followed intentionally to make it simple to execute the program. __This is not a recommanded practice for a real life project__.)*

At the moment, following topics/exmples are covered:

| Level | Topic | Examples |
| ----- | ----- | ------- |
| Fundamentals | Basics | <ul><li>Hello World</li><li>Package, Imports, Exports</li><li>Print</li></ul> |
| Fundamentals | Variables | <ul><li>Default variable values</li><li>Global variables</li><li>Variables with initializers</li><li>Short variable declarations</li></ul> |
| Fundamentals | Flow control | <ul><li>If</li><li>If/Else</li><li>If with a short statement</li><li>For loop</li><li>Using for as while</li><li>Switch</li><li>Defer</li></ul> |
| Fundamentals | Arrays | <ul><li>Basics</li></ul> |
| Fundamentals | Slices | <ul><li>Basics</li><li>Appending to slices</li><li>Making slices</li><li>Slices pointers</li></ul> |
| Fundamentals | Maps | - |
| Fundamentals | Functions | - |
| Fundamentals | Methods | - |
| Fundamentals | Types | - |
| Fundamentals | Structs | - |
| Fundamentals | Pointers | - |
| Fundamentals | Interfaces | - |
| Fundamentals | Concurrency | - |
| Fundamentals | Testing | - |
| Fundamentals | Logging | - |
| Intermediate | I/O Util | <ul><li>File reading</li><li>Directory reading</li></ul> |
| Advanced | OS package | <ul><li>os.Open() function</li></ul> |

*__More examples would be added soon__*

# Prerequisites
Before you get started, make sure to go through the prerquisites mentioned [here](https://github.com/bijeshos/learning-go/blob/master/prerequisites.md)


# New project setup
If you would like to setup a brand new project, follow the instructions from [here](https://github.com/bijeshos/learning-go/blob/master/new-project-setup.md)

# Go commands reference
A quick reference for Go commands can be found [here](https://github.com/bijeshos/learning-go/blob/master/go-commands-reference.md)

# Reference
If you would like to explore more on Go, following are a few useful URLs:
- https://go.dev/
- https://tour.golang.org/
- https://golang.org/pkg
- https://golang.org/


__** Please note that this repo is *work-in-progress* **__
