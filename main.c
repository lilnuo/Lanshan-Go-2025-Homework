#include <stdio.h>
#include <math.h>
#include <stdlib.h>
typedef long long ll;
int main()
{
    ll D;
    scanf("%lld", &D);
    ll min_diff = D;
    ll max_x = (ll)sqrt(D);
    for (ll x = 0; x * x <= D; x++)
    {
        ll remaining = D - x * x;
        ll y = (ll)sqrt(remaining);
        ll diff1 = llabs(x * x + y * y - D);
        ll diff2 = llabs(x * x + (y + 1) * (y + 1) - D);
        if (diff1 < min_diff)
            min_diff = diff1;
        if (diff2 < min_diff)
            min_diff = diff2;
        if (min_diff == 0)
            break;
    }
    printf("%lld\n", min_diff);
    return 0;
}