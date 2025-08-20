
### Cache updates and stale data

* When one thread updates a variable, the change first goes into its local CPU cache.
* Another thread may still read an old value if it looks at main memory or a stale cache copy.

### Cache write-through

* One solution is **write-through**, where updates in cache are immediately mirrored to main memory.
* But this doesn’t fix the problem if another CPU still has an outdated copy cached.


### Bus listening and invalidation

* Caches can **listen to memory update messages** on the system bus.
* When they detect an update to a memory block they hold:

    * They either update their cache copy, or
    * Invalidate the cache line so the next access forces a fresh fetch from memory.

### Cache-coherency protocols

* These techniques are part of **cache-coherency protocols**, which ensure consistent views of memory across caches.
* A common protocol is **write-back with invalidation**, though modern CPUs usually combine several methods.


### The coherency wall

* As the number of processor cores grows, keeping caches coherent becomes more complex and expensive.
* Engineers warn that this scaling difficulty will eventually hit a limit, known as the **coherency wall**.
