# 🧪 Template with Pipelining

This example demonstrates how to use **function pipelines** in Go's `text/template` package to transform data before rendering it in an HTML template.

## 💡 What It Does

- Defines custom functions in Go:
  - `functionDouble`: Doubles an integer
  - `functionSquare`: Squares a number
  - `functionSqRoot`: Calculates the square root
- Parses an HTML template and registers these functions
- Uses **template pipelines** to apply these functions in sequence
- Passes an integer (`64`) to the template and outputs the transformed results

## Template Explanation

#### 🧾 1. `{{.}}`

This prints the original input value passed from the Go program.

**Input:** `64`  
**Explanation:** No transformation — just displays the raw value.  
**Output:**  
64

#### 🧾 2. `{{. | functionDouble}}`

This applies the `functionDouble` function to the input.

**Input:** `64`  
**Operation:** `functionDouble(64)` → `64 * 2`  
**Result:** `128`  
**Output:**  
128

#### 🧾 3. `{{. | functionDouble | functionSquare}}`

This chains two functions together:

1. `functionDouble(64)` → `64 * 2 = 128`  
2. `functionSquare(128)` → `128^2 =

**Output:**  
16384

#### 🧾 4. `{{. | functionDouble | functionSquare | functionSqRoot}}`

This chains three functions together:

1. `functionDouble(64)` → `64 * 2 = 128`  
2. `functionSquare(128)` → `128^2 = 16384`  
3. `functionSqRoot(16384)` → `√16384 = 128`

**Output:**  
128

## Output

<img width="385" alt="Template with Pipelining Output Screenshot" src="https://github.com/user-attachments/assets/db0b4a35-54d6-405b-b4b6-76f08d5386a4" />
