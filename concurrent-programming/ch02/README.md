### Processes vs. Threads

* A **process** is a running program. Each process has its own memory space.
* A **thread** is a "path of execution" inside a process. Multiple threads in the same process share memory but can run independently.


### Kernel-level threads (KLTs)

* These are the “real” threads the **operating system (OS) manages directly**.
* The OS decides:

    * when to pause a thread,
    * when to run it on a CPU,
    * and where to store its state (registers, stack, etc.).
* Example: if you write a multithreaded program in C or Java (modern versions), those threads are kernel-level threads.

> Every kernel-level thread can be mapped to a CPU core (so multiple can run in parallel if multiple cores exist).


### User-level threads (ULTs)

* Instead of the OS, **the application itself manages these threads**.
* They run "on top of" one kernel-level thread.
* That means:

    * The OS only sees **one thread** (the kernel thread),
    * but inside it, the app is juggling multiple "fake threads" (user threads).

> Limitation: If one user-level thread gets blocked (e.g., waiting for network data), the whole kernel-level thread is blocked, and all other user-level threads inside it are stuck too.


### Goroutines (Go’s solution)

Goroutines are **not OS threads** and **not exactly user-level threads either**. They’re lighter and more flexible.

Go uses a **hybrid model**:

* Go runtime creates a pool of **kernel-level threads** (say one per CPU core).
* On top of these, it schedules **goroutines** (which behave like super-lightweight user-level threads).
* Each kernel-level thread has a queue of goroutines to run.

* Benefits:
  * Multiple CPU cores can be used (because goroutines are spread across kernel threads).
  * If a goroutine gets blocked (like waiting for a file or network), Go can "move" other goroutines to another kernel-level thread, so they keep running.
  * This is called *work stealing*: If one kernel thread runs out of goroutines, it can “steal” some from another busy thread’s queue, keeping all CPUs busy.
