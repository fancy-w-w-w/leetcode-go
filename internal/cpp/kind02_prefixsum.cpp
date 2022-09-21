#include <iostream>
#include <vector>
#include <stack>
#include "string.h"

using namespace std;

class Solution
{
public:
    // 1124 表现良好的最长时间段
    // 前缀和 + 单调栈
    int logestWPI(vector<int> &hours)
    {

        int n = hours.size();
        for (int i = 0; i < n; ++i)
        {
            if (hours[i] > 8)
            {
                hours[i] = 1;
            }
            else
            {
                hours[i] = -1;
            }
        }

        // 求前缀和
        int pre[n + 1];
        memset(pre, 0, sizeof(pre));
        for (int i = 1; i <= n; i++)
        {
            pre[i] = pre[i - 1] + hours[i - 1];
        }

        stack<int> s;
        for (int i = 0; i <= n; i++)
        {
            // 单调递减的栈
            if (s.empty() || pre[s.top() > pre[i]])
            {
                s.push(i);
            }
        }

        int res = 0;
        // 枚举右边界
        for (int i = n ; i >= 0 ; i--){
            if (s.empty()) break;
            while(!s.empty() && s.top() >= i) s.pop();
            while(!s.empty() && pre[s.top()] < pre[i]){
                res = max(res, i - s.top());
                s.pop();
            }
        }
        return res;
    }
};
