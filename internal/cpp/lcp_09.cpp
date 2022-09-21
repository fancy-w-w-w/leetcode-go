#include <iostream>
#include <vector>

using namespace std;

// lcp 09 最少跳跃次数
// 可以向左跳任意步数或者向右跳指定jump[i]次数
class Solution{
    public:
        int minJump(vector<int>& jump){
            vector<int> dp(jump.size(), 0);
            for (int i = dp.size()-1; i >= 0; i--){
                if (i + jump[i] >= jump.size()){
                    // 反向能到达边界
                    dp[i] += 1;
                }else{
                    dp[i] = dp[i+jump[i]] + 1;
                }
                for (int j = i + 1; j < jump.size() && j < i + jump[i] && dp[j] > dp[i]; j++){
                    dp[j] = dp[i] + 1;
                }
            }
            return dp[0];
        }
};