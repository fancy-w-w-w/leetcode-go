#include <iostream>

#include <queue>
#include <vector>

using std::priority_queue;
using std::vector;
using std::greater;

// k个有序数组归并排序
class Solution
{
public:
    int kthSmallest(vector<vector<int>> &matrix, int k)
    {
        struct point
        {
            int val, x, y;
            point(int _val, int _x, int _y) : val(_val), x(_x), y(_y) {}
            bool operator>(const point &a) const { return this->val > a.val; }
        };

        priority_queue<point, vector<point>, greater<point>> que;
        int n = matrix.size();
        for (int i = 0; i < n; i++){
            // point *tmp = new point(matrix[i][0], i, 0);
            // que.emplace(tmp);
            que.emplace(matrix[i][0], i, 0);
        }

        for(int i = 0; i < k-1; i++){
            point now = que.top();
            que.pop();
            if (now.y != n - 1){
                que.emplace(matrix[now.x][now.y+1], now.x, now.y+1);
            }
        }
        return que.top().val;
    }
};