//
// Author: CFR
// Date: create in 2020/10/9 9:45
//
#include <iostream>
#include <cstring>

using namespace std;
const int mod = 1e9 + 7;
int dp[1005][1005];
int main() {
    int n, t;
    cin >> n >> t;
    memset(dp, 0, sizeof(dp));
    int maxJ = 1;
    for (int i = 2; i <= n; ++i) {
        dp[i][0] = 1;
        for (int j = 1; j <= maxJ && j <= t; ++j) {
            dp[i][j] = (dp[i][j - 1] + dp[i - 1][j]) % mod;
            if (j >= i) {
                dp[i][j] = (dp[i][j] - dp[i - 1][j - i] + mod) % mod;
            }
        }
        maxJ += i;
    }
    cout << dp[n][t] << endl;
    return 0;
}
