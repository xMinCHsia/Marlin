
# Design notes

## Why weight-split the budget

A single draft model has a ceiling: the tokens it predicts well are a
fixed distribution. Multiple drafts with different strengths (one tuned
for prose, one for code) push the accepted-set intersection wider than
either draft alone. The weight split is a simple proportional allocator;
the autotuner adjusts effective lengths, not weights, because lengths are
cheaper to probe safely.

## Why exponential backoff on admission

