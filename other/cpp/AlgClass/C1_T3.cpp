#include <iostream>
#include <algorithm>
#include <cstdio>

using namespace std;
int N[100001];

int main() {
    int n, m, target, res, minY;
    int l, r, mid;
    cin >> n;
    for (int i = 0; i < n; ++i) {
        scanf("%d",&N[i]);
    }
    cin >> m;
    sort(N, N + n);
    while (m > 0) {
        scanf("%d",&target);

        res = 0;
        m--;


        minY = (target + N[n-1] - 1) / N[n-1];
        l = 0;
        r = n-1;
        while (l <= r) {
            mid = (l + r) / 2;
            if (N[mid] >= minY) {
                r = mid - 1;
            } else {
                l = mid + 1;
            }
        }

        int YY = l;
        int ZZ = n - 1;

        while (YY < ZZ) {
            minY = (target + N[YY] - 1) / N[YY];
            l = YY;
            r = ZZ;
            while (l <= r) {
                mid = (l + r) / 2;
                if (N[mid] >= minY) {
                    r = mid - 1;
                } else {
                    l = mid + 1;
                }
            }
            if (l == YY) {
                l++;
            }
            res += 2 * (n - l);
            YY++;
            ZZ = l - 1;
        }
        res += (n - YY) * (n - YY - 1);
        printf("%d\n",res);
    }
}