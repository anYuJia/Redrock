#include <stdio.h>
int main()
{
	unsigned int a;
	scanf("%x",&a);
	bitSwap(a);
	return 0;
}
void bitSwap(unsigned int x)
{
	unsigned int b[4] ={0};
	b[0] = x&0x000000ff;
	b[1] = (x&0x0000ff00)>>8;
	b[2] = (x&0x00ff0000)>>16;
	b[3] = (x&0xff000000)>>24;
	printf("0x%x%x%x%x",b[0],b[1],b[2],b[3]);
}
