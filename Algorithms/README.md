## Johnson and Trotter algorithm

The idea of Johnson and Trotter algorithm doesn’t require to store all permutations of size n-1 and doesn’t require going through all shorter permutations.

Assign a key value to each element in the string. Key would be index of each element.
Assign a direction to each index
  L for left and R for Right
To begin with let's mark all the elements with direction L
Check whether any element is movable:
a.IF direction is L and index is 0 element is not movable
b.If direction is R and index is n-1, n being the length of the input string, element is not movable.
c.If direction is R and the next element is smaller than the current element, current element can switch position with next element. If  next element is greater than current element element can not move.
d.If direction is L and previous element is smaller than the current element, current element can swap postion with previous one and hence movable. If previous element is greater than current element, element is not movable.

1. If no element is movable no more permutation is possible.
2. Always find the largest element in the set of movable elements and swap with the corresponding element.
3. Once an element is moved, the elements(in our case indices) greater than the current element would change direction.
4. Again the movability would be calculated and the same process would  reapeat()Step-1 until no element is movable.

