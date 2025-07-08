# Problem 3: SUBSEQMAX

Cho mảng ![eq](https://latex.codecogs.com/svg.image?\bg{white}&space;s=\left(a_1,a_2,...a_n\right))

Một đoạn ![eq](https://latex.codecogs.com/svg.image?\bg{white}&space;s(i,j)=\left(a_i,...,a_j\right),1\leq&space;i\leq&space;j\leq&space;n) có trọng số ![eq](https://latex.codecogs.com/svg.image?\bg{white}w(i,j)=a_i&plus;a_{i&plus;1}&plus;...&plus;a_j). Hãy tìm một đoạn trong mảng có trọng số lớn nhất, nghĩa là tổng các số trong đoạn là lớn nhất.

### Input
Dòng thứ nhất chứa một số nguyên `(n <= 10^6)`.

Dòng thứ hai chứa `n` số nguyên.
### Output
Ghi ra duy nhất một số nguyên là trọng số lớn nhất tìm được.
### Examples
| Input | Output |
| --- | --- | 
| `6` </br> `-2 11 -4 13 -5 2` | `20` | 