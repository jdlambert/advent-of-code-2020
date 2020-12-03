#include <iostream>
#include <fstream>
#include <set>
#include <vector>

using namespace std;

int slope_trees(const vector<vector<bool>> &trees, int d_row, int d_col) {
    int count = 0;
    for (int row = 0, col = 0; row < trees.size(); row += d_row, col = (col + d_col) % trees[0].size()) {
        count += trees[row][col];
    } 
    return count;
}

int part1(const vector<vector<bool>> &trees) {
    return slope_trees(trees, 1, 3);
}

int64_t part2(const vector<vector<bool>> &trees) {
    vector<pair<int, int>> slopes;
    int64_t product = slope_trees(trees, 2, 1);
    for (int i = 1; i < 9; i += 2) {
        product *= slope_trees(trees, 1, i);

    }
    return product;
}

int main() {
    ifstream in("input.txt");
    vector<vector<bool>> trees;
    string line;
    while (getline(in, line)) {
        if (line.empty()) {
            break;
        }
        vector<bool> row;
        for (char item : line) {
            row.push_back(item == '#');
        }
        trees.push_back(row);
    }
    cout << part1(trees) << endl << part2(trees) << endl;
}