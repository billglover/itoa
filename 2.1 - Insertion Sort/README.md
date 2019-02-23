## Exercises

### 2.1-1

![image showing the contents of the array during sorting][ex211]

### 2.1-2

```go
// Desc take a slice of int and returns a new slice with the
// elements sorted in descending numerical order.
func Desc(a []int) []int {
	if len(a) == 1 {
		return a
	}

	for j := len(a) - 2; j >= 0; j-- {
		key := a[j]

		i := j + 1
		for i < len(a) && a[i] > key {
			a[i-1] = a[i]
			i = i + 1
		}
		a[i-1] = key
	}

	return a
}
```

### 2.1-3

```plain
LINEAR-SEARCH(A, v)
for j = 1 to A.length
	if A[j] == v
		return j
return nil
```

**Loop Invariant:** At the start of each iteration of the for loop, the subarray `A[1..j]` does not contain the value `v`. If it exists in the array `A` then it must exist in the subarray `A[j..]`.

**Initialisation:** At initialisation subarray `A[1..j]` contains no elements and so it does not contain `v`.

**Maintenance:** Each iteration of the loop increases `j` by `1`. If `A[j+1] == v` the loop terminates (see termination). If `A[j+1] != v` then the subarray `A[1..j+1]` does not contain the value `v` and the loop invariant holds true.

**Termination:** There are two termination conditions:
   - `A[j] == v` in which case the value has been found and the index `j` is returned.
   - `j > A.length` in which case the subarray `A[1..A.length]` does not contain the value `v` and `nil` is returned

### 2.1-4

Consider the *binary addition* problem:

**Input:** Two n-element arrays, representing an n-bit binary numbers, A and B.

**Output:** An (n+1)-element array representing an (n+1)-bit binary number that is the sum of A and B.

**Pseudocode:**
```plain
BINARY-ADD(A, B)
carry = 0
for i = A.length down to 1
  C[i+1] = A[i] + B[i] + carry
  if C[i+1] > 1
    C[i+1] = 0
    carry = 1
  else
    carry = 0

C[0] = carry
return C
```

[ex211]: ex211.png "Exercise 2.1-1"
