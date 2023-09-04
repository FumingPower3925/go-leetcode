# 4. Median of Two Sorted Arrays

## Speed Solution

The objective of this problem is to find the median (not confuse with the mean). 

There are many possible aproaches but almost all of them have time complexities of O(m+n) and that sort. This solution is a correct one (the problem states that the solution must be O(log(n+m)), this solution is O(log(n+m).

The general idea behind this problem is that what we are tying to do is a sort of a huge modification to the classic binary search. So, what we know that we will end up doing is "discarding" halves each time and continuing with the half that has the solution.

This is quite complex solution, feel free to take your time to understand it.

**Speed Complexity**

In this case the complexity

$$
O(log(n+m))
$$

The worst case happens when by the structure of the input we don't end up getting the correct one in less than log(n+m) iterations.

**Space Complexity**

The space needed for this solution is almost neglegible, as we only store a fixed amount of variables and functions, thus there are no structures that depend on the input size:

$$
O(1)
$$

1. As it is constant there are no worst cases.

## Space Solution

In this problem the provided solution is optimal in space too, so there is no need for another solution.
