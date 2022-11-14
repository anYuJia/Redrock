#include <stdio.h>
//10进制转任意进制(1~36)
#define MAX 20
int main()
{
	int dec,base;
	scanf("%d,%d",&dec,&base);
	baseConverter(dec,base);
	return 0;
}

void baseConverter(int dec,int base){
	int a[MAX],i = 0;
	while(dec!=0)
	{
		i++;
		a[i] = dec%base;
		dec/=base;
	}
	for(;i>0;i--){
		if(a[i]<10){
			printf("%d",a[i]);
		}
		else if(a[i]<36)
		{
			printf("%c",a[i]+55);
		}
	}
}
