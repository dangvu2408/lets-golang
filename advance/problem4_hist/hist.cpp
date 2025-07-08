#include <iostream>
#include <stack>
#include <algorithm>
#include <vector>
using namespace std;

long long largest(vector<int> height) {
    stack<int> s;
    long long maxArea = 0;
    height.push_back(0);
    int n = height.size();

    for (int i = 0; i < n; i++) {
        while (!s.empty() && height[i] < height[s.top()]) {
            int h = height[s.top()]; 
            s.pop();
            int width = s.empty() ? i : i - s.top() - 1;
            maxArea = max(maxArea, (long long)h * width);
        }
        s.push(i);
    }
    return maxArea;
}

int main() {
    while (true) {
        int n;
        cin >> n;
        if (n == 0) break;

        vector<int> heights(n);
        for (int i = 0; i < n; ++i) cin >> heights[i];
        cout << largest(heights) << "\n";
    }
    return 0;
}