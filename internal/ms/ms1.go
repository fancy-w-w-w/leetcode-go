package ms

// 构成交替字符串
// 构成交替字符串需要的最小交换次数
// public int minSwaps(String s) {
//     int n = s.length();

//     //符合题目要求的只有两种
//     //情况1(1在前，1多)："1010"（偶数） "10101"（奇数）
//     //情况2(0在前，0多)："0101"（偶数） "01010"（奇数）
//     int n1 = help(s, '1');
//     int n0 = help(s, '0');
//     if (Math.abs(n1 - n0) > 1)
//         return -1;
//     int res = Integer.MAX_VALUE;
//     // 这两个if不能并排写
//     if (n1 == (n + 1) / 2 && n0 == n / 2) {
//         int diff1 = 0;
//         for (int i = 0; i < s.length(); i++) {
//             if (s.charAt(i) - '0' == i % 2) {
//                 diff1++;
//             }
//         }
//         res = diff1 / 2;
//     }
//     if (n0 == (n + 1) / 2 && n1 == n / 2) {
//         int diff0 = 0;
//         for (int i = 0; i < s.length(); i++) {
//             if (s.charAt(i) - '0' != i % 2) {
//                 diff0++;
//             }
//         }
//         res = Math.min(res, diff0 / 2);
//     }
//     return res;
// }

// private int help(String s, char c) {
//     int count = 0;
//     for (int i = 0; i < s.length(); i++) {
//         if (s.charAt(i) == c)
//             count++;
//     }
//     return count;
// }
