# Go Map Initialization and Panic Prevention

This repository demonstrates a common error in Go: panics caused by accessing uninitialized maps.  Uninitialized maps are `nil`, and attempting to access or modify them results in a runtime panic.  The provided code showcases this issue and its solution.

## Bug

The `bug.go` file contains a Go program that attempts to assign a value to a map before it's properly initialized. This leads to a panic.

## Solution

The `bugSolution.go` file presents a corrected version where the map is properly initialized before use, preventing the panic. This is achieved by creating an empty map using the `make` function.