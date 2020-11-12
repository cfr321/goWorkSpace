//
// Author: CFR
// Date: create in 2020/11/12 9:39
//
#include <iostream>
#include<set>
#include <hash_set>

using namespace std;
using namespace __gnu_cxx;
int d[30];
int m, res = INT32_MAX;
void dfs(int k, set<int> s, int p) {
    if (k == m) {
        res = min(res, (int) s.size());
        return;
    }
    s.insert(p+d[k]);
    if (s.size()<res)
         dfs(k+1,s,p+d[k]);
    s.erase(p+d[k]);
    s.insert(p-d[k]);
    if (s.size()<res)
        dfs(k+1,s,p-d[k]);
}
int main() {
    cin >> m;
    for (int i = 0; i < m; ++i) {
        cin >> d[i];
    }
    set<int> s;
    s.insert(0);
    s.insert(d[0]);
    dfs(1,s,d[0]);
    cout << res << endl;
    return 0;
}
