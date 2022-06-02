### Answer

- In the weighted quick-union algorithm, suppose that we set `id[find(p)]=q` instead of to `id[find(p)]=find(q)`. Would the resulting algorithm be correct?
> Answer: Yes, but it would increase the tree height, so the performance guarantee would be invalid

- Another question: In the weighted quick-union algorithm, suppose that we `id[p]=q` instead of `qu.id[find(p)]=find(q)`. Would the resulting algorithm be correct?
> Nop. that makes set not a tree but a graph in some scenarios
