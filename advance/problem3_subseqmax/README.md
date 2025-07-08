# Problem 3: SUBSEQMAX

Cho mảng s = (a₁, ..., aₙ)

Một đoạn s(i, j) = (aᵢ, ..., aⱼ), với 1 ≤ i ≤ j ≤ n có trọng số w(i, j) = aᵢ + aᵢ₊₁ + ... + aⱼ. Hãy tìm một đoạn trong mảng có trọng số lớn nhất, nghĩa là tổng các số trong đoạn là lớn nhất.

### Input
Dòng thứ nhất chứa một số nguyên `(n <= 10^6)`.

Dòng thứ hai chứa `n` số nguyên.
### Output
Ghi ra duy nhất một số nguyên là trọng số lớn nhất tìm được.
### Examples
| Input | Output |
| --- | --- | 
| `6` </br> `-2 11 -4 13 -5 2` | `20` | 