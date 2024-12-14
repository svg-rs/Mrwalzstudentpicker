# Student Selector Application

This is a simple Go application that selects students randomly from a given class period when a key is pressed. The application supports multiple class periods, and each period contains a list of students. Once a student is selected, they are removed from the list, and the process continues until all students are selected.

## Features

- Selects students randomly from a list of class periods.
- Allows for real-time student selection by pressing the "A" key.
- Handles multiple class periods with predefined student lists.
- Provides visual feedback with colored output in the terminal.

## Prerequisites

Before running the application, ensure you have the following installed:

- Go 1.16 or higher
- Git
- The following Go packages:
  - `github.com/eiannone/keyboard` (for handling keyboard inputs)
  - `github.com/fatih/color` (for colored text in the terminal)

Install the necessary Go packages by running:
```bash
go get github.com/eiannone/keyboard
go get github.com/fatih/color
