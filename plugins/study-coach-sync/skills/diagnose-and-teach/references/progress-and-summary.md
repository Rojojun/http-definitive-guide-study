# Progress, summaries, and consistency

Use this protocol to convert an open-ended teaching conversation into a durable N-day course.

## Establish the schedule

Collect or infer these values:

- source scope and desired outcome;
- total study days `N`;
- minutes available per study day;
- preferred cadence or fixed deadline, if any;
- whether diagnostic Day 0 counts toward `N`.

If `N` or daily minutes are missing, finish any safe current unit and ask for the missing values before finalizing the roadmap. Keep the plan provisional until then.

Allocate days by conceptual weight and learner evidence, not equal page counts. For a large source:

- use about 60-70% of days for new concepts;
- use about 20-30% for retrieval, integration, and practical exercises;
- keep about 10% as adaptive buffer or final synthesis.

Adjust these ratios when `N` is small. Combine familiar chapters; split difficult or foundational chapters. State what was compressed or deferred.

## Start a study day

1. Load the learning ledger when available.
2. Show `Day X/N`, today's outcome, and the time budget.
3. Ask one short no-notes retrieval question from a due review item.
4. Continue the planned unit or replan based on the answer and available time.

Do not make attendance bookkeeping longer than the lesson.

## Define completion boundaries

- **Learning unit:** one durable question whose mechanism the learner can explain or transfer.
- **Study day:** the planned time box or explicitly ended session, even when a unit remains partial.
- **Chapter:** all selected learning outcomes for that chapter, not necessarily every paragraph.
- **Course:** the agreed scope and final integration task are complete.

Mark work as `complete`, `partial`, or `deferred`. Do not call a unit complete merely because five reasoning turns elapsed.

## Emit summaries

### Unit Summary

Emit immediately after a unit is complete:

- one-sentence outcome;
- causal mental model;
- corrected misconceptions;
- practical mapping;
- one retrieval prompt and its review date.

### Daily Summary

Emit when the study day ends:

- Day X/N and status;
- completed and partial units;
- three to seven durable takeaways;
- evidence demonstrated by the learner;
- unresolved misconceptions or questions;
- due review items;
- next day's target.

### Chapter Summary

Emit when a chapter is complete:

- concept map or dependency chain;
- key invariants and tradeoffs;
- source-version caveats and current extensions;
- production checklist or application exercise;
- cumulative retrieval questions.

### Course Summary

Emit when the course is complete:

- before/after mental-model comparison;
- chapter-level synthesis;
- remaining weak areas;
- final transfer task;
- maintenance review plan.

## Maintain the learning ledger

When the workspace is available and the user wants continuity, create one stable Markdown ledger outside the skill folder. When [github-sync.md](github-sync.md) is configured, use the GitHub copy as the source of truth and treat local files as caches. Preserve this schema:

```markdown
# Learning Ledger: <course>

## Plan
- Source:
- Goal:
- Total days:
- Minutes/day:
- Started:
- Current day:
- Status:

## Progress
| Day | Date | Unit | Status | Evidence | Next review |

## Knowledge
### Established
### Corrected misconceptions
### Needs reinforcement

## Review queue
| Due | Prompt | Last result |

## Next session
```

Update existing sections instead of appending duplicate summaries. Record evidence, not just confidence labels. Never store sensitive personal information that is unnecessary for learning.

## Schedule reviews

Default to retrieval after roughly 1, 3, 7, and 14 study days when the course length allows. Ask the learner to predict or explain before showing the note. Retire an item after repeated successful transfer; bring it forward after a weak answer.

## Recover from gaps

- Keep the streak informational and non-punitive.
- Start with a two-minute retrieval check, not a full restart.
- Recalculate remaining scope against the remaining days and time.
- Prefer deferring optional extensions over rushing foundational concepts.
- Report any scope tradeoff clearly.
