# create React structure

## What is the project for?

This program was build to improve the creation steps of a **react** app. The program takes the
structure of the given .yaml file and creates directories, .js Files and .module.css files.
Additionally the .js files are filled with the base instructions to make them into a component.
This program was build for my own purpose and it's target audience is me.

## How to install this project?

A current go installation is required.

## How to use the project?

To use the program either you can run the program or build the binary.

```bash As program
go run your/path/main.go -t /path/to/target -s /path/to/structure
```

```bash As binary
go build your/path/
./createReactStructure -t /path/to/target -s /path/to/structure
```

The program can be run without flags. The default values are -t:./src/components and -s:./componentStructure.yaml.

The provided yaml file should be the dictionary names as keys and the component names as arrays.

```yaml Sample structure
Meal:
    - MealItem
    - MealForm
    - MealSummary
UI:
    - Card
    - Modal
Utils:
    - Input
```

## License

MIT License

Copyright (c) [2023] [Patrick Volpert]

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
