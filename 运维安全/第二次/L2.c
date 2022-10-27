#include <stdio.h>
int main()
{
    int a[10];
    for(int i = 0;i<10;i++)
    {
        scanf("%d",a+i);
    }
    sort(a,10);
    for(int i = 0;i<10;i++)
    {
        printf("%d ",a[i]);
    }
}
void swap(int *x , int *y)
{
	int temp;
	temp = *x;
	*x = *y;
	*y = temp;
}
void sort(int x[],int n)
{
	int i,j;
	for(i = 0; i<n-1;i++)
	{
		for(j = 0 ;j<n-1-i;j++)
		{
			if(x[j] < x[j+1]) swap(x+j,x+j+1);
		}
	}
}
