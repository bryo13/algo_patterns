const std = @import("std");
const print = std.debug.print;

pub fn main() void {
    const arraylist = [5]i32{ 10, -20, -10, 30, 5 };

    const mSum: i32 = maxSum(arraylist, 3);
    print("{} \n", .{mSum});
}

fn maxSum(arraylist: [5]i32, window: usize) i32 {
    var windowSum: i32 = 0;
    var max: i32 = 0;
    var size: usize = 0;

    for (0..arraylist.len) |i| {
        windowSum += arraylist[i];

        if (windowSum > max) {
            max = windowSum;
        }

        if (i >= window - 1) {
            size = window - 1;
            windowSum -= arraylist[i - size];
        }
    }

    return max;
}
