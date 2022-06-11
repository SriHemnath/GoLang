Problem Description

There are N children standing in a line with some rating value. You want to distribute a minimum number of candies to these children such that:

1. Each child must have at least one candy.
2. The children with higher ratings will have more candies than their neighbours.

You need to write a program to calculate the minimum candies you must give.

Input: An array arr[] of N integers representing the ratings of each student
Example 1

Input: arr[] = [1, 2, 2]
Output: 4
Explanation: You can distribute to the first, second and third child 1, 2 and 1 candies respectively. The third child gets 1 candy because it satisfies the above two conditions.
Example 2

Input: arr[] = [1, 5, 2, 1]
Output: 7
Explanation: Candies given = [1, 3, 2, 1]