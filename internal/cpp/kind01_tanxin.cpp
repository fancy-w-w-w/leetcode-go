#include <iostream>
#include <vector>
#include <climits>
#include <queue>
#include <unordered_map>
using namespace std;

class Solution
{
public:
    // 334 递增的三元子序列
    bool increasingTriplet(vector<int> &nums)
    {
        if (nums.size() <= 2)
        {
            return false;
        }
        int first = nums[0], second = int(INT_MAX);
        for (int i = 1; i < nums.size(); i++)
        {
            if (nums[i] > second)
            {
                return true;
            }
            else if (nums[i] > first)
            {
                second = nums[i];
            }
            else
            {
                first = nums[i];
                // second = MY_VALUE;
            }
        }
        return false;
    }

    // 376 摆动序列
    int wiggleMaxLength(vector<int> &nums)
    {
        if (nums.size() <= 1)
        {
            return nums.size();
        }
        int up = 1, down = 1;
        int res;

        for (int i = 1; i < nums.size(); i++)
        {
            if (nums[i] > nums[i - 1])
            {
                up = down + 1;
            }
            else if (nums[i] < nums[i - 1])
            {
                down = up + 1;
            }
        }
        res = max(up, down);
        return res;
    }

    // 358 K距离间隔重排字符串
    // 给你一个非空的字符串 s 和一个整数 k ，你要将这个字符串 s 中的字母进行重新排列，使得重排后的字符串中相同字母的位置间隔距离 至少 为 k
    // 尽可能构造字符串————》贪心   优先选择数量最多的字母，因为字母数量越多，在字符串中保存间距越难
    string rearrangeString(string s, int k)
    {
        if (!k)
        {
            return s;
        }

        //记录每个字母的出现次数
        unordered_map<char, int> map;
        for (auto c : s)
        {
            map[c]++;
        }

        // 按照次数从小到大排序
        priority_queue<pair<int, char>> pq;
        for (auto [c, i] : map)
        {
            pq.emplace(i, c);
        }

        // 暂存队列
        queue<pair<int, char>> q;
        string res;

        while (!pq.empty())
        {
            // 每次放一个数量最多的字母，放入暂存队列
            auto &[i, c] = pq.top();
            res += c;
            q.emplace(i - 1, c);
            pq.pop();

            if (q.size() == k)
            {
                // 如果还有剩余次数
                if (q.front().first > 0)
                {
                    pq.emplace(q.front().first, q.front().second);
                }
                q.pop();
            }
        }

        return res.size() < s.size() ? "" : res;
    }
};