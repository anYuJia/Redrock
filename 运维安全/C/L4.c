#include <stdio.h>
int main()
{
    printf(" ‰»Î∞…£°\n");
    int a,q = 2;
    int flag = 0;
    int m = 0;
    scanf("%d", &a);
    int sum = 0;
    for (int i = 1; i <= 100;i++ )
    {
        for(int j = 1 ; j <= i;j++)
        {
        	flag++;
            sum += i;
            m++;
            if(m%3 == 0)
            {
            	sum = sum - 2;
			}
            if(flag == a)
            {
                goto re;
            }
        }
    }
    re:printf("%d", sum);
    return 0;
}
