//
// Author: CFR
// Date: create in 2020/9/26 11:56
//
#include <cstdio>
using namespace std;


int numbers[10];
int dp[10][10][10];
int help[10] = {0, 9, 99, 351, 927, 2151, 4671, 9783, 20079, 40743};

int dfs(int k, bool istop, int s1, int s2) {
    if (k <= 0) {
        return 1;
    }
    int maxN = istop ? numbers[k] : 9;
    int res = 0;
    if (!istop && s2 != -1 && dp[k][s1][s2] != 0) {
        return dp[k][s1][s2];
    }
    for (int i = 0; i <= maxN; ++i) {
        if (s2 == -1) {
            if (i != s1)
                res += dfs(k - 1, istop && i == numbers[k], s1, i);
            else {
                res += dfs(k - 1, istop && i == numbers[k], s1, -1);
            }
        } else {
            if (i == s1 || i == s2) {
                res += dfs(k - 1, istop && i == numbers[k], s1, s2);
            }
        }
    }
    if (!istop && s2 != -1) dp[k][s1][s2] = res;
    return res;
}
int main() {
    int n;
    scanf("%d", &n);

    int k = 0;
    while (n) {
        numbers[++k] = n % 10;
        n /= 10;
    }
    int res = 0;
    res += help[k - 1];
    for (int i = 1; i <= numbers[k]; ++i) {
        res += dfs(k - 1, i == numbers[k], i, -1);
    }
    printf("%d\n", res);
    return 0;
}