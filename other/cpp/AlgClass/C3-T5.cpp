//
// Author: CFR
// Date: create in 2020/10/9 10:41
//
#include <cstdio>
#include <vector>
#include <iostream>

using namespace std;
vector<int> son[301];
int v[301];
int dp[301][301];
int n, m;

void dfs(int x) {

    for (int i = 1; i <= m; ++i) {
        dp[x][i] = v[x];
    }
    for (int i = 0; i < son[x].size(); ++i) { // 子树看成一个物品，进行01背包选择
        int s = son[x][i];
        dfs(s);                               // 计算子树的最大价值，
        for (int j = m; j >= 1; --j) {
            for (int k = 1; k <= m - 1; ++k) {
                dp[x][j] = max(dp[x][j], dp[x][j - k] + dp[s][k]);
            }
        }
    }
}
int main() {
    int f;
    scanf("%d%d", &n, &m);
    for (int i = 1; i <= n; ++i) {
        scanf("%d%d", &f, &v[i]);
        son[f].push_back(i);
    }
    v[0] = 0;
    dfs(0);
    cout << dp[0][m] << endl;
    return 0;
}

