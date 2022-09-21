#include <iostream>
#include <unordered_map>
#include <list>

using namespace std;

// 缓存的节点信息
struct Node
{
    int key, val, freq;
    // 构造函数
    Node(int _key, int _val, int _freq) : key(_key), val(_val), freq(_freq) {}
};

class LFUCache
{
    int minfreq, capacity;
    // 指向freq_table
    unordered_map<int, list<Node>::iterator> key_table;
    // // map + 双向链表
    unordered_map<int, list<Node>> freq_table;

public:
    LFUCache(int _capacity)
    {
        minfreq = 0;
        capacity = _capacity;
        key_table.clear();
        freq_table.clear();
    }

    int get(int key)
    {
        if (capacity == 0)
            return -1;
        //      当 find() 方法成功找到以指定元素作为键的键值对时，其返回的迭代器就指向该键值对；
        //      当 find() 方法查找失败时，其返回的迭代器和 end() 方法返回的迭代器一样，指向容器中最后一个键值对之后的位置。
        auto it = key_table.find(key);
        if (it == key_table.end())
            return -1;
        // it为正向pair类型的迭代器
        list<Node>::iterator node = it->second;

        int val = node->val, freq = node->freq;
        freq_table[freq].erase(node);
        // 如果链表为空，需要在哈希表中删除该freq，更新minfreq
        if (freq_table[freq].size() == 0)
        {
            freq_table.erase(freq);
            if (minfreq == freq)
            {
                minfreq += 1;
            }
        }
        //插入到freq+1的哈希和链表中
        freq_table[freq + 1].push_front(Node(key, val, freq + 1));
        key_table[key] = freq_table[freq + 1].begin();
        return val;
    }

    void put(int key, int value)
    {
        if (capacity == 0)
            return;
        auto it = key_table.find(key);
        // val之前没存在
        if (it == key_table.end())
        {
            // 缓存已满，需要进行删除
            if (key_table.size() == capacity)
            {
                // minfreq链表末尾节点
                auto it2 = freq_table[minfreq].back();
                key_table.erase(it2.key);
                freq_table[minfreq].pop_back();
                // 直接把链表删除
                if (freq_table[minfreq].size() == 0)
                {
                    freq_table.erase(minfreq);
                }
            }
            freq_table[1].push_front(Node(key, value, 1));
            key_table[key] = freq_table[1].begin();
            minfreq = 1;
        }
        else
        {
            // 之前已经存在
            // 与get操作一致， 需要额外更新缓存的值
            list<Node>::iterator node = it->second;
            int freq = node->freq;
            freq_table[freq].erase(node);
            if (freq_table[freq].size() == 0)
            {
                freq_table.erase(freq);
                if (minfreq == freq)
                {
                    minfreq += 1;
                }
            }
            freq_table[freq + 1].push_front(Node(key, value, freq + 1));
            key_table[key] = freq_table[freq + 1].begin();
        }
    }
};