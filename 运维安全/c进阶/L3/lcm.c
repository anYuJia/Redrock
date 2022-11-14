#include <stdio.h>
#include "talk.h"
#include "gcd.h"

int lcm(int var1, int var2){
    int flag = 2;
    talk(flag);

    int temp_gcd, ans;
    temp_gcd = gcd(var1, var2);
    ans = var1 * var2 / temp_gcd;
    return ans;
}