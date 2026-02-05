# Rate Limiting Algorithms 

A comprehensive collection and implementation of the most prominent rate limiting algorithms used in modern system design to ensure service stability, fairness, and security.

---

##  What is Rate Limiting?

**Rate limiting** is a strategy used to control the amount of incoming or outgoing traffic to a network or service. It defines a threshold—for example, "100 requests per minute"—and ensures that any client exceeding this limit is either throttled, delayed, or blocked.

### Why is it important?
* **Preventing Abuse:** Protects against DoS (Denial of Service) and brute-force attacks.
* **Cost Control:** Limits the usage of expensive resources (like third-party APIs or auto-scaling cloud infra).
* **Service Stability:** Ensures that a single heavy user doesn't degrade the experience for everyone else.
* **Fairness:** Distributes resources equitably across all users.

---

##  Prominent Algorithms

This repository includes implementations and explanations for the following algorithms:

### 1. Token Bucket 
The most widely used algorithm (used by companies like Stripe and Amazon). 
* **How it works:** A "bucket" holds tokens. Every request consumes one token. Tokens are refilled at a fixed rate. If the bucket is empty, the request is rejected.
* **Pros:** Allows for **bursts** of traffic while maintaining a steady average rate.
* **Best for:** User-facing APIs where occasional spikes are expected.

### 2. Leaky Bucket 
Similar to the Token Bucket but focuses on a steady output rate.
* **How it works:** Requests enter a "bucket" (queue). They "leak" out of the bottom at a constant, fixed rate for processing. If the bucket overflows, new requests are dropped.
* **Pros:** Perfectly smooths out traffic into a **constant flow**.
* **Best for:** Background processing and scenarios requiring a predictable processing speed.

### 3. Fixed Window Counter 
The simplest approach to implement.
* **How it works:** Time is divided into fixed intervals (e.g., 1-minute blocks). Each block has a counter. Once the counter hits the limit, requests are blocked until the next interval starts.
* **Cons:** Can be "gamed" at the edges (a user could send their full limit at the end of Window A and another full limit at the start of Window B).
* **Best for:** Simple use cases where precision isn't critical.

### 4. Sliding Window Log 
A highly accurate but memory-intensive method.
* **How it works:** Instead of a counter, it stores a timestamp for every single request. When a new request arrives, it looks back at the log and counts requests within the last $T$ seconds.
* **Pros:** Extremely precise; no "edge" issues.
* **Cons:** High memory usage because every request's timestamp must be stored.

### 5. Sliding Window Counter 
A hybrid approach that provides the precision of a sliding log with the efficiency of a fixed window.
* **How it works:** It uses a weighted average of the current and previous windows to estimate the request count.
* **Formula:** $Count = \text{Requests in current window} + (\text{Requests in previous window} \times \text{Overlap percentage})$
* **Pros:** Memory efficient and much smoother than the Fixed Window approach.

---

