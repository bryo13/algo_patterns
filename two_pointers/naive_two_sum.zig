const std = @import("std");

// list has to be sorted -> to be checked
fn NaiveTwoSUm(list: [3]usize, sum: usize) bool {
    for (list) |i| {
        for (list) |j| {
            if (i == j) {
                continue;
            }

            if (list[i] + list[j] == sum) {
                return true;
            }

            if (list[i] + list[j] > sum) {
                break;
            }
        }
    }

    return false;
}

test "naivly test which two numbers add to a given sum" {
    try std.testing.expect(NaiveTwoSUm([3]usize{ 1, 2, 3 }, 5) == true);
}

// checkout slices!
