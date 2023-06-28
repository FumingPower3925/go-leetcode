# 1. Two Sum

## Speed Solution

In this case the solution to maximise the speed is by using the following idea.

As the description states, we need to find **two** numbers that add to target, this also means that if you have one of the numbers, with a very simple operation you can get the other one.

$$
num+x=target
$$

So if we have a number (num), for example the one that we are dealing with in that iteration of the loop, and the target number (as we are provided it by the second argument of the function) we can always know $x$ by doing $x=target-num$.

**Speed Complexity**

In this case we are just going to iterate once throught the array nums, and the rest of operations (adding and finding elements in a map) are constant $O(1)$. Therefore, having $n$ as the length of the array "nums", we have:

$$
O(n)
$$

The worst case happens when one of the numbers we are searching for is the very last one of the array nums.

**Space Complexity**

The space complexity in this solution is "relatively high", as we have to keep in memory a hashmap with all the unique copies of numbers that we've seen so far. Therefore, having $n$ as the size of the map "hist", we have:

$$
O(n)
$$

The worst case happens when one of the numbers we are searching for is the very last one of the array nums and they are all different.

**FAQS**

**What happens when we find a replicate of a number in "nums"?**

> In this case, as we store the numbers in a map having them as the keys we can only store a single copy of each number. However, we only care about duplicates if they are both part of the solution in which case we don't need to save the second one in the map as we can end the function just there.

## Space Solution

In this case the solution to minimize the space is by using the following idea.

The only way to minimize the space is to avoid using all the structures/variables possible.

In this case to achieve this we need to get rid of the map used in the speed solution, so we have no other option to solve the problem than by brute force.

There is only a small optimitzation that can be done here, and that one is only checking  a number with all the following ones, as the previous ones have already been checked in previous iterations.

**Speed Complexity**

In this case as we have $n$ iterations in the first for, and for each one we do $n-i$ iterations in the inner for, this may not seem obvious at first glance but it can be proven that is:

$$
O(n^2)
$$

The worst case happens when the solution elements are the last two of nums.

**Space Complexity**

As we haven't got any structure the space is:

$$
O(1)
$$

Therefore, no worst case are possible.

# 
