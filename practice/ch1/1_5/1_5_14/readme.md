### Answer

- see `cases/ch1/1_5/dynamic_connectivity/unionfind/extra/height_quick_union.go`

- two conditions:
	- if a less deep tree merged to a deeper one, the max tree depth will keep to the bigger one, which will never make the depth longer
	- if equal deep tree merged to another one, the max tree depth will add 1
	- so the wrost scene will be always meet the second condition

| height | size   |
| ------ | ------ |
| 1+1=2  | 2+2=4  |
| 2+1=3  | 4+4=8  |
 3+1=4  | 8+8=16 |
| ...    | ...    |
| n      | 2^n    |

- size = 2^height => height=lg(size)
