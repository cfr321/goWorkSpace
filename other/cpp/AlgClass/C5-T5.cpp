//
// Author: CFR
// Date: create in 2020/11/12 10:29
//
#include <iostream>
#include <vector>
#include <queue>
#include <string.h>

using namespace std;
struct E {
    int u, v, w;
};
priority_queue<int> q;
vector<E> ev;
int usedu[21], usedv[21];
int n, m, k;

void dfs(int t, int sum) {
    if (t==m){
        return;
    }
    if (!usedv[ev[t].v] && !usedu[ev[t].u]) {
        if (q.size() < k || sum + ev[t].w < q.top()) {
            q.push(sum + ev[t].w);
            if (q.size() > k) q.pop();
            usedv[ev[t].v] = 1;
            usedu[ev[t].u] = 1;
            dfs(t+1,sum+ev[t].w);
            usedv[ev[t].v] = 0;
            usedu[ev[t].u] = 0;
        }
    }
    dfs(t+1,sum);
}

int main() {
    cin >> n >> m >> k;
    memset(usedu, 0, sizeof(usedu));
    memset(usedv, 0, sizeof(usedv));
    for (int i = 0; i < m; ++i) {
        E e;
        cin >> e.u >> e.v >> e.w;
        ev.push_back(e);
    }
    dfs(0,0);
    cout<<q.top()<<endl;
    return 0;
}
