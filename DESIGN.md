
# Design notes

## Why weight-split the budget

A single draft model has a ceiling: the tokens it predicts well are a
fixed distribution. Multiple drafts with different strengths (one tuned
for prose, one for code) push the accepted-set intersection wider than
either draft alone. The weight split is a simple proportional allocator;
the autotuner adjusts effective lengths, not weights, because lengths are
cheaper to probe safely.

## Why exponential backoff on admission

Cache pressure is bursty. A linear probe re-admits a draft the moment
pressure dips; exponential backoff keeps rejected drafts out until the
pressure regime actually changes. The base is `probe_interval_s`, the
exponent is the number of elapsed intervals - so the wait doubles per
interval, bounded by the request timeout.

