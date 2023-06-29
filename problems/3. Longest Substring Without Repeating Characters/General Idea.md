# 3. Longest Substring Without Repeating Characters

## Speed Solution

In this case the fastest solution is to iterate only once through the string and keep track  of all the apperances of the characters. We can use a histogram en take advantage of two key facts:

1. We know that a substring can only with a repeated element or the end of the input string s.

2. We also know that the characters must be ASCII (given by the description).

Therefore, the histogram can have a **limited** size of 256 positions.

Then, we just have to count the substrings until we reach an already seen position, and then we have to start again keeping track of the maximum lenght achieved.

**Speed Complexity**

In this case we are just going to iterate once throught the string s, and the rest of operations (incrementing/accessing elements of a slice) are constant $O(1)$. Therefore, having $n$ as the length of the string s, we have:

$$
O(n)
$$

There is no worst case as we will always have to read all the string to be sure there are no bigger sequences. So it can be represented as:

$$
\theta(n)
$$

**Space Complexity**

The space complexity in this solution is asymptotically small. However, there are some solutions that use even less space, at the cost of a much worse performance. So I think those can be overlooked as in the limit the space they need is almost the same. In my test I've reported mroe less $-0.1MB$ of space required by those solutions. All in all, in the provided solution the space needed is:

$$
O(1)
$$

As it is constant there are no worst cases.

**Testing**

I would recommend you to use the following tests in order to create the program, as the base tests are quite incomplete.

```
"abcabcbb"
"bbbbb"
"pwwkew"
"dvdf"
""
" "
```

## Space Solution

In this problem the provided solution is optimal in space too, so there is no need for another solution.

# 
