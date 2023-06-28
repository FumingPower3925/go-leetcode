# 2. Add Two Numbers

## Speed Solution

In this problem maximizing the idea to maximize the speed is to compute the sum while traversing the two lists. A common approach could be to try to read both numbers, then summing them, and then putting it in a list, but jsut thinking about it seems like a lot of work to do. In this problem the right path as I've already said is to compute each digit sum using the following "algorithm":

$$
(x+y)\%10
$$

Doing the mod operation returns the "first" digit of the sum, if the sum is less than 10 (only 1 digit) it returns that digit and if the sum is greater or equal than 9 (two digits) it returns the first one (e.g.: (9+8)%10 = 7). This number can be inserted in the current position in the list, and if the number has more than 2 digits we have to "carry one".

$$
(x+y)/10
$$

This operation can be 0 if the result has 1 digit or 1 if the result has 2 digits. (as the numbers can only go up to 9).

**Speed Complexity**

The speed complexity in this solution is quite trivial, if $n$ is the length of $l1$ and $m$ is the length of $l2$:

$$
O(max(m,n))
$$

This happens as we traverse both lists at the same time.

> Note: It can be also represented as $O(m+n)$

**Space Complexity**

As we haven't got any structure other than the list the space required is:

$$
O(max(m,n))
$$

The list will have max(m,n) nodes.

## Space Solution

In this problem the space solution would not have sense as there is no further way to reduce the space needed.
