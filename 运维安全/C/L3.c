#include <stdio.h>
int main()
{
    int a[3][3];
    int i,j;
    for(i = 0 ; i < 3 ;i++)
    {
        scanf("%d %d %d",a[i]+0,a[i]+1,a[i]+2);
    }
    printf("\n主对角线：\n");
    for(i = 0 ; i < 3 ; i++)
    {
        printf("%d " ,a[i][i]);
    }
    printf("\n附对角线\n");
    for(i = 0 ;i < 3 ; i++ )
    {
        printf("%d ",a[i][2-i]);
    }
    printf("\n元素和：");
    int sum = 0;
    for(i = 0 ; i < 3 ; i++)
    {
        for(j = 0 ; j < 3 ;j++)
        {
            sum+=a[i][j];
        }
    }
    printf("sum = %d",sum);
}