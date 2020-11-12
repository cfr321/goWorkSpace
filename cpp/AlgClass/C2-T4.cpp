//
// Author: CFR
// Date: create in 2020/10/2 12:08
//
#include <vector>
#include <algorithm>
#include <cstdio>

using namespace std;
typedef struct Pos {
    int x;
    int y;
} Pos;

vector<Pos> v;
int pos[100010];
int dis(Pos &p1, Pos &p2) {
    return (p1.x - p2.x) * (p1.x - p2.x) + (p1.y - p2.y) * (p1.y - p2.y);
}

bool cmp(Pos &p1, Pos &p2) {
    if (p1.x == p2.x)
        return p1.y < p2.y;
    return p1.x < p2.x;
}
bool cmpy(int p1, int p2){
    return v[p1].y<v[p2].y;
}

int findMinLen(int l, int r) {
    if (l >= r) {
        return 0;
    }
    if (r == l + 1) {
        return dis(v[l], v[r]);
    }
    int mid = (l + r) / 2;
    int d = min(findMinLen(l, mid), findMinLen(mid, r));
    int num = 0;
    for (int i = l; i <=r ; ++i) {
        if(((v[mid].x-v[i].x)*(v[mid].x-v[i].x))<=d){
            pos[num++]=i;
        }
    }
    sort(pos,pos+num,cmpy);
    for (int i = 0; i < num; ++i) {
        for (int j = i+1; j < num && ((v[pos[j]].y-v[pos[i]].y)*(v[pos[j]].y-v[pos[i]].y)) < d; ++j) {
            d = min(d,dis(v[pos[j]],v[pos[i]]));
        }
    }
    return d;
}

int main() {
    int n;
    scanf("%d", &n);
    v.clear();
    for (int i = 0; i < n; ++i) {
        Pos p;
        scanf("%d", &p.x);
        scanf("%d", &p.y);
        p.x = abs(p.x);
        p.y = abs(p.y);
        v.push_back(p);
    }
    sort(v.begin(), v.end(), cmp);
    int d = findMinLen(0, n - 1);
    printf("%d\n", d);
    return 0;
}
