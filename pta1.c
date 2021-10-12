#include<stdio.h>
int main(){
int n;
scanf("%d",&n);
int count=0;
if (n==1){
printf("0");
}else{
while(true){
	if (n%2==0){
		n/=2;
		count++;
		if (n==1) break;
		else{
			n=3*n+1;
			n/=2;
			count++;
			if (n==1){
			break;}
		}
	}
}
}
printf(count);
return 0;
}

