#include <iostream>
#include <fstream>
#include <tuple>
#include <vector>

using namespace std;

typedef tuple<int, int, char, string> Password;

int part1(const vector<Password> &passwords) {
    int valid = 0;
    for (auto pass_tuple : passwords) {
        int low, high;
        char target;
        string password;

        tie(low, high, target, password) = pass_tuple;

        int count = 0;
        for (char c : password) {
            if (c == target) {
                count++;
            }
        }

        if (low <= count && count <= high) {
            valid++;
        }
    }
    return valid;
}

int part2(const vector<Password> &passwords) {
    int valid = 0;
    for (auto pass_tuple : passwords) {
        int first, second;
        char target;
        string password;

        tie(first, second, target, password) = pass_tuple;

        if ((password[first - 1] == target) != (password[second - 1] == target)) {
            valid++;
        }
    }
    return valid;
}

int main() {
    ifstream in("input.txt");

    vector<Password> passwords;
    while (!in.eof()) {
        int num1, num2;
        char dash, colon;
        char character;
        string password;

        in >> num1 >> dash >> num2 >> character >> colon >> password;
        if (!password.empty()) {
            passwords.push_back(make_tuple(num1, num2, character, password));
        }
    }

    cout << part1(passwords) << endl << part2(passwords) << endl;
}