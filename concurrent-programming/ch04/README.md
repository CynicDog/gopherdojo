**Mutex Implementation Overview**

* Mutexes prevent multiple goroutines or threads from accessing shared resources simultaneously, avoiding race conditions.
* On single-processor systems, a mutex could theoretically be implemented by disabling interrupts while a thread holds the lock. However, this is risky—badly written code can block the entire system, and multi-CPU systems require a more robust approach.
* Modern mutexes rely on hardware-supported atomic operations (like test-and-set) to safely acquire a lock. If a thread finds the mutex already locked, the OS blocks it until the lock becomes available.
* Using a mutex effectively makes the code between `Lock` and `Unlock` sequential. This ensures safety but can reduce parallel performance if the lock is held for too long, so it’s important to minimize the critical section.

