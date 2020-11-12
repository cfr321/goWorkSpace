//
// Author: CFR
// Date: create in 2020/10/10 9:58
//
#include <iostream>
#include <unordered_map>

using namespace std;
typedef unordered_map<int, int>::const_iterator kv;
int a[3005];
int dp[3005];

int main() {
    int n,add;
    cin >> n;
    unordered_map<int, int> Map;
    for (int i = 1; i <= n; ++i) {
        scanf("%d", &a[i]);
        Map.insert(make_pair(a[i], i));
    }
    dp[1] = 1;
    dp[2] = 2;
    int res = 2;
    for (int i = 3; i <= n; ++i) {
        dp[i] = 2;
        for (int j = i - 1; j >= 2 && a[i] - a[j] < a[j]; --j) {
            kv got = Map.find(a[i] - a[j]);
            if (got != Map.end()) {
                add = 2;
                if (got->second==2 && a[1]+a[2]!=a[j]){
                    add = 1;
                }
                dp[i] = max(dp[i], dp[got->second] + add);
            }
        }
        res = max(res, dp[i]);
    }
    cout << res << endl;
    return 0;
}
