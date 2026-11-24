# N-th Root of an Integer

> Striver SDE Sheet 061 · Day 11 Binary Search — binary search over candidate roots; the power check caps out early instead of overflowing. Go, stdlib only, tested.

## Problem

Given n and m, find the integer x with xⁿ = m, or report -1 when m is not a perfect n-th power. First of the "binary search on the answer" problems.

## Approach

xⁿ is strictly increasing in x, so the candidate space 1..m is monotone — binary search territory even though nothing is "an array". The non-obvious hazard is evaluating midⁿ: for mid in the millions and n beyond a few, the power overflows 64 bits long before the comparison happens. The fix is a capped power: multiply step by step, and the moment the running product would exceed m (checked by division, `res > m/x`, which cannot itself overflow), return "too big". The search only ever needs three-way comparison against m, never the true value.

**Complexity:** O(n · log m) time, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
