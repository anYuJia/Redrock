#include <stdio.h>
void talk(int flag);
void talk(int flag)
{
    if(flag == 1)
    {
        printf("gcd is being calculated...\n");
    }
    if(flag == 2)
    {
        printf("lcm is being calculated...\n");
    }
}
