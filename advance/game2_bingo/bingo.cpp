#include <iostream>
#include <cstdlib>
#include <ctime>
#include <iomanip>
#include <algorithm>
using namespace std;

const int N = 5;
const string COLUMNS = "BINGO";

int randomInRange(int min, int max) {
    return rand() % (max - min + 1) + min;
}

void printBoard(int board[N][N], bool selected[N][N]) {
    cout << "\n    B   I   N   G   O\n";
    for (int i = 0; i < N; ++i) {
        cout << " " << i + 1 << " ";
        for (int j = 0; j < N; ++j) {
            if (i == 2 && j == 2) {
                cout << setw(3) << " * ";
            } else if (selected[i][j]) {
                cout << "[" << setw(2) << board[i][j] << "]";
            } else {
                cout << " " << setw(2) << board[i][j] << " ";
            }
        }
        cout << endl;
    }
}

int columnIndex(char c) {
    c = toupper(c);
    return COLUMNS.find(c);
}

int main() {
    srand(time(0));

    int board[N][N];
    bool selected[N][N] = {false};
    int lineSums[12];           
    bool lineCompleted[12] = {false};

    for (int col = 0; col < N; ++col) {
        bool used[16] = {false};
        for (int row = 0; row < N; ++row) {
            if (row == 2 && col == 2) {
                board[row][col] = 0;
                continue;
            }
            int num;
            do {
                num = randomInRange(col * 15 + 1, col * 15 + 15);
            } while (used[num - col * 15 - 1]);
            used[num - col * 15 - 1] = true;
            board[row][col] = num;
        }
    }

    selected[2][2] = true;

    for (int i = 0; i < 5; ++i) {
        int rowSum = 0, colSum = 0;
        for (int j = 0; j < 5; ++j) {
            rowSum += board[i][j];
            colSum += board[j][i];
        }
        lineSums[i] = rowSum;
        lineSums[i + 5] = colSum;
    }

    int diag1 = 0, diag2 = 0;
    for (int i = 0; i < 5; ++i) {
        diag1 += board[i][i];
        diag2 += board[i][4 - i];
    }
    lineSums[10] = diag1;
    lineSums[11] = diag2;

    int minSum = *min_element(lineSums, lineSums + 12);
    int linesCleared = 0;

    while (linesCleared < 5) {
        printBoard(board, selected);

        cout << "\nNhập ô để chọn (ví dụ: B 2), hoặc 'exit' để thoát: ";
        string inputCol;
        int inputRow;
        cin >> inputCol;

        if (inputCol == "exit" || inputCol == "EXIT") break;
        cin >> inputRow;

        int col = columnIndex(inputCol[0]);
        int row = inputRow - 1;

        if (col < 0 || col >= N || row < 0 || row >= N) {
            cout << "Nhập không hợp lệ. Vui lòng thử lại.\n";
            continue;
        }

        if (selected[row][col]) {
            cout << "Ô này đã được chọn rồi.\n";
            continue;
        }

        selected[row][col] = true;

        for (int i = 0; i < 5; ++i) {
            if (lineCompleted[i]) continue;
            bool full = true;
            for (int j = 0; j < 5; ++j)
                if (!selected[i][j]) full = false;
            if (full) {
                lineCompleted[i] = true;
                linesCleared++;
                if (lineSums[i] == minSum)
                    cout << "Bạn đã hoàn thành hàng " << i + 1 << " với tổng nhỏ nhất!\n";
                else
                    cout << "Bạn đã hoàn thành hàng " << i + 1 << ", nhưng không phải tổng nhỏ nhất.\n";
            }
        }

        for (int j = 0; j < 5; ++j) {
            if (lineCompleted[5 + j]) continue;
            bool full = true;
            for (int i = 0; i < 5; ++i)
                if (!selected[i][j]) full = false;
            if (full) {
                lineCompleted[5 + j] = true;
                linesCleared++;
                if (lineSums[5 + j] == minSum)
                    cout << "Bạn đã hoàn thành cột " << COLUMNS[j] << " với tổng nhỏ nhất!\n";
                else
                    cout << "Bạn đã hoàn thành cột " << COLUMNS[j] << ", nhưng không phải tổng nhỏ nhất.\n";
            }
        }

        if (!lineCompleted[10]) {
            bool diagMain = true;
            for (int i = 0; i < 5; ++i)
                if (!selected[i][i]) diagMain = false;
            if (diagMain) {
                lineCompleted[10] = true;
                linesCleared++;
                if (lineSums[10] == minSum)
                    cout << "Bạn đã hoàn thành đường chéo chính với tổng nhỏ nhất!\n";
                else
                    cout << "Bạn đã hoàn thành đường chéo chính, nhưng không phải tổng nhỏ nhất.\n";
            }
        }

        if (!lineCompleted[11]) {
            bool diagAnti = true;
            for (int i = 0; i < 5; ++i)
                if (!selected[i][4 - i]) diagAnti = false;
            if (diagAnti) {
                lineCompleted[11] = true;
                linesCleared++;
                if (lineSums[11] == minSum)
                    cout << "Bạn đã hoàn thành đường chéo phụ với tổng nhỏ nhất!\n";
                else
                    cout << "Bạn đã hoàn thành đường chéo phụ, nhưng không phải tổng nhỏ nhất.\n";
            }
        }
    }

    if (linesCleared >= 5) {
        cout << "\nBạn đã hoàn thành 5 dòng! Chúc mừng phá đảo trò chơi!\n";
    } else {
        cout << "\nBạn đã thoát trò chơi.\n";
    }

    return 0;
}
