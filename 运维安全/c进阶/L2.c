#include <stdio.h>
int main()
{
	int x[3][3];
	for(int i = 0 ; i < 3 ;i++)
	{
			scanf("%d %d %d",*(x+i), *(x+i)+1, *(x+i)+2);
	}
	for(int i=0;i<3;i++){
		for(int j=0;j<=i;j++){
			int temp = *(*(x+i)+j);
			*(*(x+i)+j) = *(*(x+j)+i);
			*(*(x+j)+i) = temp;
		}
	}
	for(int i = 0 ; i < 3 ;i++)
	{
			printf("%d %d %d\n",*(*(x+i)), *(*(x+i)+1), *(*(x+i)+2));
	}
	return 0;
}
