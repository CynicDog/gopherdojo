
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

### The Process Memory Map 

When you start a program, the operating system gives your process its own virtual address space. Think of this as a private “map” of memory — it looks continuous to your program, but underneath, the OS maps pieces of it onto real physical RAM (or swap, or shared memory).

Inside that map, memory is typically divided into regions:

```
[ High memory addresses ]
-------------------------
|        Stack(s)       |  (grows downward)
-------------------------
|         ...           |
| Shared libs / mmap    |
|         ...           |
-------------------------
|         Heap          |  (grows upward)
-------------------------
|   Static data / bss   |
-------------------------
|   Program code (text) |
-------------------------
[  Low memory addresses ]
```

* Stacks and heap are logical regions in the process’s virtual memory.

* The OS kernel maps these logical addresses onto physical RAM pages.

* The CPU caches may temporarily hold values from either the heap or a stack, but cache coherency ensures that updates propagate so other threads/goroutines see consistent values.

### Inline Optimization in Go

Inlining is a compiler optimization where a function call is replaced with the body of the function itself. Instead of generating code to jump into another function (and push arguments/return values on the stack), the compiler just copies the function’s code directly into the caller.


### Error Control Flow: `defer`, `panic` and `recover`

#### `defer`
* `defer` schedules a function call to run at the end of the current function’s execution.
* Deferred calls are stored in a list tied to the function’s stack frame and are executed in **LIFO order** (last deferred, first executed).

#### `panic`

* A **`panic`** is Go’s way of signaling a serious error — it immediately begins **unwinding the stack**.
* As the stack unwinds, Go executes any **deferred calls** registered in each stack frame.
* If no code recovers from the panic, the program crashes and prints a stack trace.

#### `recover`

* **`recover`** is a built-in function that regains control of a panicking goroutine.
* It only works when called **inside a deferred function**.
* If `recover` is used in a defer:

    * It stops the panic from propagating further.
    * It returns the error value passed to `panic`.

    
### Race Conditions and Critical Sections

* A **race condition** occurs when multiple goroutines access the same memory concurrently and at least one of them writes to it. The outcome can vary depending on timing, CPU scheduling, and cache behavior.
* Even if individual instructions are atomic, **CPU caches and registers** can delay memory visibility. Each core may operate on local cached values before periodically flushing to main memory, so other goroutines might not see changes immediately.

#### Critical Section

* A **critical section** is code that must be executed by only one goroutine at a time to prevent conflicts on shared resources.
* Accessing shared data without proper protection can lead to inconsistent results, especially when goroutines run on multiple cores.

#### Synchronization Mechanisms

* **Mutexes** (`sync.Mutex`) lock a critical section, ensuring only one goroutine executes it at a time.
* **Atomic operations** (`sync/atomic`) safely update single variables without full locks.
* Proper synchronization **eliminates race conditions**, ensuring all goroutines see a consistent view of memory.

#### Parallel Execution Warnings

* On a single processor, user-level scheduling is mostly non-preemptive, so races may be less likely—but **never rely on this**.
* With multiple kernel threads (`runtime.GOMAXPROCS(n)`), the OS can interrupt execution anytime, increasing the chance of race conditions.

#### Best Practices

* Question whether **shared memory is necessary**—sometimes goroutines can communicate without sharing state (e.g., via channels).
* Identify **critical code sections** and protect them using proper synchronization.
* Good concurrent programming involves **coordination and communication**, like marking resources in use, to avoid overlapping operations.
