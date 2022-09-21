#include <iostream>
#include <vector>

using namespace std;

class Solution
{
public:
    int canCompleteCircuit(vector<int> &gas, vector<int> &cost)
    {
        int n = gas.size();
        for (int i = 0; i < n;)
        {
            int sumOfGas = 0, sumOfCost = 0, cnt = 0;
            for (; cnt < n;)
            {
                int j = (i + cnt) % n;
                sumOfGas += gas[j];
                sumOfCost += cost[j];
                if (sumOfCost > sumOfGas)
                {
                    break;
                }
                cnt++;
            }
            if (cnt == n)
            {
                return i;
            }
            else
            {
                i += cnt + 1;
            }
        }
        return -1;
    }
};