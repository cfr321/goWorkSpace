//
// Author: CFR
// Date: create in 2020/11/11 18:43
//
#include<cstdio>
#include<algorithm>
#include<iostream>
#include <string.h>
#include <vector>

using namespace std;
struct point {
    long long x;
    long long y;
} p[1000005];
long long sum = 0, maxx = 0, maxy = 0, res = 0, minx = INT32_MAX;

bool cmp(point p1, point p2) {
    return p1.y < p2.y;
}

void add(int i) {
    res++;
    sum += p[i].y;
    minx = min(minx, p[i].x);
    maxx = max(maxx, p[i].x);
    maxy = max(maxy, sum + minx);
    maxy = max(maxy, p[i].x + p[i].y);
}


int dp[105];
int findRotateSteps(string ring, string key) {
    memset(dp, 0, sizeof(dp));
    int len = ring.length();
    int res = INT32_MAX;
    vector<vector<int >> v(26);
    vector<int> tmp,bef;
    for (int i = 0; i < ring.length(); ++i) {
        v[ring[i] - 'a'].push_back(i);
    }
    tmp = v[key[0] - 'a'];
    for (int i = 0; i < tmp.size(); ++i) {
        dp[tmp[i]]= min(tmp[i],len-tmp[i])+1;
    }
    for (int i = 1; i < key.length(); ++i) {
        bef = tmp;
        tmp = v[key[i] - 'a'];
        for (int j = 0; j < tmp.size(); ++j) {
            int m = INT32_MAX;
            for (int k = 0; k < bef.size(); ++k) {
                int l = abs(tmp[j] - bef[k]);
                m = min(m, dp[bef[k]] + min(l, len - l) + 1);
            }
            dp[tmp[j]] = m;
            if (i == key.length() - 1) {
                res = min(res, dp[tmp[j]]);
            }
        }
    }
    return res;
}

int main() {
    long long n, s;
    scanf("%lld%lld", &n, &s);
    for (long long i = 0; i < n; ++i)
        scanf("%lld%lld", &p[i].x, &p[i].y);
    sort(p, p + n, cmp);
    for (long long i = 0; i < n; ++i)
        if (p[i].x + p[i].y <= s) {
            if (p[i].x >= maxx && p[i].y + maxy <= s) {
                add(i);
            }
            if (p[i].x < maxx && p[i].y + sum + min(p[i].x, minx) <= s) {
                add(i);
            }
        }
    printf("%lld", res);
    return 0;
}
