#include <stdio.h>
int main()
{
    for (int i = 9; i >= 1; i--)
    {
        for (int j = i; j >= 1; j--)
        {
            printf("* ");
        }
        printf("\n");
    }
}