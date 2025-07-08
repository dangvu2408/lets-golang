#include <iostream>
#include <vector>
#include <map>
#include <cmath>
#include <algorithm>
#include <limits>
using namespace std;

const int PLAYER = 5;

struct Player {
    int id;
    int point = 10;
    int guess = 0;
    bool gameover = false;
};

int main() {
    vector<Player> players(PLAYER);
    for (int i = 0; i < PLAYER; i++) {
        players[i].id = i + 1;
    }

    int round = 1;
    bool rule1_active = false;
    bool rule2_active = false;
    bool rule3_active = false;

    while (true) {
        cout << "\n --- ROUND " << round++ << " --- \n";

        int active_player = 0;
        for (auto &p : players) {
            if (!p.gameover) {
                cout << "Player " << p.id << " (" << p.point << " points), choose [0 - 100]:";
                cin >> p.guess;
                active_player++;
            }
        }

        if (active_player <= 2) break;

        if (rule3_active) {
            bool is0 = false, is100 = false;
            int p100 = -1;
            for (auto &p : players) {
                if (!p.gameover) {
                    if (p.guess == 0) is0 = true;
                    if (p.guess == 100) {
                        is100 = true;
                        p100 = p.id;
                    }
                }
            }

            if (is0 && is100) {
                cout << "Rule 3 activated! Player " << p100 << " wins this round! \n";
                for (auto &p : players) {
                    if (!p.gameover && p.id != p100) {
                        p.point--;
                        if (p.point == 0) p.gameover = true;
                    }
                }
                continue;
            }
        }

        double sum = 0;
        int count = 0;
        for (auto &p : players) {
            if (!p.gameover) {
                sum += p.guess;
                count++;
            }
        }
        double avg = (count > 0 ? sum / count : 0);
        int target = std::round(avg * 0.8);
        cout << "Target: " << target << "\n";

        map<int, int> freq;
        for (auto &p : players) {
            if (!p.gameover) freq[p.guess]++;
        }

        if (rule1_active) {
            for (auto &p : players) {
                if (!p.gameover && freq[p.guess] > 1) {
                    p.point--;
                    cout << "Rule 1: Player " << p.id << " lost 1 point due to duplicate. \n";
                    if (p.point == 0) p.gameover = true;
                }
            }
        }

        int winner = -1;
        int mindiff = numeric_limits<int>::max();
        for (auto &p : players) {
            if (!p.gameover && freq[p.guess] == 1) {
                int diff = abs(p.guess - target);
                if (diff < mindiff) {
                    mindiff = diff;
                    winner = p.id;
                }
            }
        }

        if (winner == -1) {
            cout << "No winner this round.\n";
            continue;
        }

        cout << "Winner: Player " << winner << ".\n";

        bool exact = false;
        for (auto &p : players) {
            if (p.id == winner && p.guess == target) exact = true;
        }

        for (auto &p : players) {
            if (!p.gameover && p.id != winner) {
                int loss = (rule2_active && exact) ? 2 : 1;
                p.point -= loss;
                if (p.point <= 0) {
                    p.gameover = true;
                    cout << "Player " << p.id << " has been eliminated!\n";
                }
            }
        }

        int eliminated_count = 0;
        for (auto &p : players) {
            if (p.gameover) eliminated_count++;
        }
        if (eliminated_count >= 1) rule1_active = true;
        if (eliminated_count >= 2) rule2_active = true;
        if (eliminated_count >= 3) rule3_active = true;
    }

    cout << "\n--- GAME OVER ---\n";
    for (auto &p : players) {
        if (!p.gameover) {
            cout << "Player " << p.id << " wins the game with " << p.point << " points!\n";
        }
    }

    return 0;
}