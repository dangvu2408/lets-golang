#include <iostream>
#include <string>
#include <algorithm>

using namespace std;

string add(string a, string b) {
    if (a.length() < b.length()) { swap(a, b); }
    reverse(a.begin(), a.end());
    reverse(b.begin(), b.end());

    string result = "";
    int carry = 0;

    for (size_t i = 0; i < a.length(); ++i) {
        int da = a[i] - '0';
        int db = (i < b.length()) ? b[i] - '0' : 0;

        int sum = da + db + carry;
        result += (sum % 10) + '0';
        carry = sum/10;
    }

    if (carry > 0) {
        result += carry + '0';
    }

    reverse(result.begin(), result.end());
    return result;
} //sử dụng hàm add(...) ở Problem 1: ADD

int main() {
    int T;
    cin >> T;


    while (T--) {
        string a, b;
        cin >> a >> b;
        cout << add(a, b) << "\n";
    }
    return 0;
}