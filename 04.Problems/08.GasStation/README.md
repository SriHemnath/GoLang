## Problem Description

There are N gas stations along a circular route, where the amount of gas at station i is arr[i]. You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from station i to its next station (i+1). At the beginning of the journey, the tank is empty at one of the gas stations.

Return the minimum starting gas station’s index if you need to travel around the circuit once, otherwise return -1.

## Problem Note

Completing the circuit means starting at i and ending up at i again.
Both input arrays are non-empty and have the same length.
Each element in the input arrays is a non-negative integer.

## Example 1

Input: 
arr[] = [1,2,3,4,5]
cost[] = [3,4,5,1,2]
Output: 3

Explanation:
Start at station 3 (index 3) and fill up with 4 unit of gas. Gas available in Tank = 0 + 4 = 4
Travel to station 4.Gas available in tank = 4 - 1 + 5 = 8
Travel to station 0.Gas available in tank = 8 - 2 + 1 = 7
Travel to station 1.Gas available in tank = 7 - 3 + 2 = 6
Travel to station 2.Gas available in tank = 6 - 4 + 3 = 5
Travel to station 3.The cost is 5.Your gas is just enough to travel back to station 3.Therefore, return 3 as the starting index.

Hints:
If the total gas larger equal to or larger than cost then there is a solution.
We need to return the 1st position of the gas station which has higher capacity than cost.
The cost to travel between stations doesn't change no matter where you start, so at the end if the total capacity is less than zero then there is no way
