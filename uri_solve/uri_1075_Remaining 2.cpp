#include<bits/stdc++.h>
using namespace std;

int main()
{
    int n;

    while(cin>>n)
    {
        for(int i=2;i<=10000;i++)
        {
            if(i%n==2)
                cout<<i<<endl;
        }
    }

    return 0;
}
