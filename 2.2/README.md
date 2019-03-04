## Exercises

### 2.2-1

`n³/1000 - 100n² - 100n + 3` expressed in 𝛩-notation is n³.

### 2.2-2

```plain
SELECTION-SORT(A)
for i = 1 to (A.length - 1)
	mv, mi = A[i], i
	for j = i to A.length
		if mv < A[j]
			mv = A[j]
			mi = j
		A[[i], A[mi] = A[mi], A[i]
```

**Loop Invariant:** At the start of each iteration of the outer for loop, the sub-array `A[1..i-1]` is sorted. All elements in the sub-array `A[i..A.length]` are larger than all elements in the sub-array `A[1..i-1]`

**Initialisation:** At initialisation sub-array `A[1..j]` contains no elements and so is sorted.

**Maintenance:** Each iteration of the loop appends the lowest value in the sub-array `A[i..A.length]` to the sub-array `A[1..i-1]`.

**Termination:** The loop terminates when `i = A.length`. At this point there is only one value left in the sub-array `A[1..i-1]`. There is no need to sort this final remaining value and so the whole array is now sorted.

Best case running time is when the array is sorted is 𝛩(n²). The same is true for the worst case running time.

### 2.2-3

We define the average case to be when the element being sought is located in the middle of the array, at index `n/2`. The worst case occurs when the element being sought is located at the end of the array, at index `n`.

Worst case running time requires us to search all `n` locations in the array. This gives a worst case running time of 𝛩(n).

Average case running time requires us to search `n/2` locations in the array. This gives a worst case running time of 𝛩(n). The constant term `1/2` is ignored as it becomes insignificant as `n` tends towards a large value.

### 2.2-4

To improve the running time of an algorithm we are looking to minimise work. In the best case the work required is zero. To improve the running time of the best case input, we need to identify the input and return immediately (with a pre-determined answer if necessary).
