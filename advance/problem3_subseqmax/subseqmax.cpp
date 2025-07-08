// Kadance's algorithm
#include <iostream>
using namespace std;

int subArrMax(int arr[], int n) {
    int current = arr[0];
    int maximum = arr[0];
    for (int i = 1; i < n; i++) {
        current = max(arr[i], current + arr[i]);
        maximum = max(maximum, current);
    }
    return maximum;
}

int main() {
    int n;
    cin >> n;
    int arr[n];
    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }
    cout << "Sub Array Max Sum equals: " << subArrMax(arr, n);
    return 0;
}