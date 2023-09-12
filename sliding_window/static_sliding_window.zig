const std = @import("std");
const print = std.debug.print;

pub fn main() void {
    const arraylist = [5]u32{10, 20, 10, 30, 5};
    const mSum: u32 = maxSum(arraylist,3);
    print("{any} \n", .{mSum});
}

fn maxSum(arraylist: [5]u32,window: u8) u32 {
    var i: u8 = 0;
    var windowSum: u32 = 0;
    var total: u32 = 0; 

    while(i < arraylist.len) : (i = i + 1) {
        windowSum += arraylist[i];

        if (windowSum > total) {
            total = windowSum;
        }

        if (i >= window-1) {
            windowSum -= arraylist[i-window+1];
        }
    }
    return total;
}