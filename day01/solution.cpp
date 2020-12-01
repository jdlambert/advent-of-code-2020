#include <unordered_set>
#include <vector>
#include <algorithm>
#include <iostream>
#include <fstream>

using namespace std;

int main() {
    
    int num;
    vector<int> nums;
    ifstream in("input.txt");
    while (in >> num) {
        nums.push_back(num);
    }
    
    unordered_set<int> seen;
    for (auto num : nums) {
        if (seen.find(2020 - num) != seen.end()) {
            cout << (2020 - num) * num << endl;
            break;
        }
        seen.insert(num);
    }
    
    sort(nums.begin(), nums.end());
    for (int i = 0; i < nums.size() - 2; i++) {
        int j = i + 1;
        int k = nums.size() - 1;
        while (j < k) {
            int total = nums[i] + nums[j] + nums[k];
            if (total == 2020) {
                cout << nums[i] * nums[j] * nums[k] << endl;
                return 0;
            } else if (total < 2020) {
                ++j;
            } else {
                --k;
            }
        }
    }
}