#include <iostream>
#include <vector>
#include <algorithm>
#include <memory>
#include <regex>
#include <thread>
#include <mutex>



int main(){
    return 0;
}

// int v = 1;

// void critical_section(int change_v) {
//     static std::mutex mtx;
//     std::unique_lock<std::mutex> lock(mtx);

//     // 执行竞争操作
//     v = change_v;

//     lock.unlock();
//     // 离开此作用域后 mtx 会被释放
// }

// int main() {
//     std::thread t1(critical_section, 2), t2(critical_section, 3);
//     t1.join();
//     t2.join();

//     std::cout << v << std::endl;
//     return 0;
// }

// struct A;
// struct B;

// struct A {
//     std::shared_ptr<B> pointer;
//     ~A() {
//         std::cout << "A 被销毁" << std::endl;
//     }
// };
// struct B {
//     std::weak_ptr<A> pointer;
//     ~B() {
//         std::cout << "B 被销毁" << std::endl;
//     }
// };
// int main() {
//     auto a = std::make_shared<A>();
//     auto b = std::make_shared<B>();
//     a->pointer = b;
//     b->pointer = a;

//     std::string fnames[] = {"foo.txt", "bar.txt", "test", "a0.txt"};
//     std::regex base_regex("([a-z]+)\\.txt");
//     std::smatch base_match;
//     for (const auto &fname : fnames){
//         if (std::regex_match(fname, base_match, base_regex)){
//             if (base_match.size() == 2){
//                 std::string base = base_match[1].str();
//             std::cout << "sub-match[0]: " << base_match[0].str() << std::endl;
//             std::cout << fname << " sub-match[1]: " << base << std::endl;
//             }
//         }
//     }
// }