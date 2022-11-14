#include <stdio.h>
#include "talk.h"
void pre();
void talk(int flag);
int gcd(int var1, int var2)
{
    int flag = 1;
    talk(flag);

    int ans;
    ans = var2;
    while(var1 % var2 != 0)
    {
        ans = var1 % var2;
        var1 = var2;
        var2 = ans;
    }
    return ans;
}
