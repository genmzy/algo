### Answer

| index | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 |
| ----- | - | - | - | - | - | - | - | - | - | - |
| id[i] | 1 | 1 | 3 | 1 | 5 | 6 | 1 | 3 | 4 | 5 |

- tree:
```
    (1)
   / | \
  /  |  \
(0) (3) (6) 
    / \   \
  (2) (7) (5)
          / \
        (4) (9)
         |
        (8)
```

- impossible; lg10 < 4 = depth; if weight-quick-union, depth will less than lg10

- A possible operation to make this happen is:
	- quick-union algorithm
	- {{8, 4}, {4, 5}, {9, 5}, {5, 6}, {2, 3},{7, 3}, {3, 1}, {6, 1}, {0, 1}}
