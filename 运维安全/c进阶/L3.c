#include <stdio.h>
void talk(int flag);
void pre();
int gcd(int var1, int var2);
int lcm(int var1, int var2);
int main()
{
	pre();
	int a, b;
	int ans_gcd, ans_lcm;
	scanf("%d %d", &a, &b);
	ans_gcd = gcd(a, b);
	ans_lcm = lcm(a, b);
	printf("gcd is :%d\n", ans_gcd);
	printf("lcm is :%d\n", ans_lcm);
	return 0;
}
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
int lcm(int var1, int var2){
	int flag = 2;
	talk(flag);
	
	int temp_gcd, ans;
	temp_gcd = gcd(var1, var2);
	ans = var1 * var2 / temp_gcd;
	return ans;
}
void pre()
{
	printf("Welcom!\n");
	printf("please input two numbers:\n");
}
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
