### Amdahl’s Law 

* Suppose a program has two parts:

    * **Sequential part**: fraction `S` of total execution time that cannot be parallelized.
    * **Parallel part**: fraction `P = 1 - S` that can run on multiple processors.

* The **speedup** using `N` processors is given by:

$$
\text{Speedup}(N) = \frac{1}{S + \frac{P}{N}}
$$

**Implications:**

* No matter how many processors you add, the sequential part limits overall speedup.
* If `S` is small, you get close to linear speedup. If `S` is large, adding more processors gives diminishing returns.
* This explains why holding a mutex (making a section of code sequential) reduces concurrency gains—only the parallelizable parts benefit from multiple cores.
