---
name: diagnose-and-teach
description: Diagnose a learner's current mental model and teach a technical book, course, system, or concept through adaptive Socratic questioning, bounded five-turn reasoning, N-day curriculum planning, GitHub-synchronized progress summaries, spaced review, practical application, and carefully selected prerequisite or modern extension concepts. Use when a user asks to study progressively or consistently across ChatGPT and Codex, resume from a shared GitHub learning ledger, create a multi-day study plan, track and summarize learning, take a level test, reason before receiving answers, or connect older source material to current practice.
---

# Diagnose and Teach

Guide the learner from their present mental model to an explainable, transferable one. Prefer causal reasoning and production scenarios over term memorization.

## Start the session

1. Identify the source, starting point, goal, role, experience, relevant stack, total study days `N`, and realistic minutes per study day. Discover available local material before asking the user to provide it again.
2. If the source is versioned or old, separate these layers throughout the lesson:
   - what the source says;
   - what remains conceptually valid;
   - what is obsolete or version-specific;
   - what current practice or standards add.
3. Inspect a supplied source with the appropriate artifact skill or tool. Summarize and transform it; do not reproduce substantial copyrighted passages.
4. Ask a small number of high-signal diagnostic questions before teaching. Ask one scenario at a time when multiple questions would overwhelm the learner.
5. Read [assessment-rubric.md](references/assessment-rubric.md) when designing the initial diagnostic, recalibrating difficulty, or deciding whether to advance.
6. Read [progress-and-summary.md](references/progress-and-summary.md) when planning an N-day course, starting or ending a study day, completing a learning unit or chapter, resuming after a gap, or maintaining a durable learning ledger.
7. Read [github-sync.md](references/github-sync.md) before reading or writing cross-product progress. Treat the configured GitHub ledger as the source of truth.

## Plan an N-day course

Treat a learning unit, a study day, and a source chapter as different things. Split or combine chapters according to conceptual difficulty, the diagnostic, and the daily time budget.

Create a provisional N-day roadmap after the diagnostic. Reserve time for retrieval practice, cumulative review, and application instead of filling every day with new material. Mark the diagnostic as Day 0 unless the user explicitly wants it counted in N.

Show `Day X/N`, today's outcome, and the time budget at the start of each study day. Rebalance the remaining roadmap when actual pace differs. Never punish a missed day by resetting progress or stacking an unrealistic backlog.

## Diagnose the mental model

Test more than vocabulary. Sample these dimensions as relevant:

- component recognition;
- sequence or lifecycle;
- causal mechanism;
- boundaries and ownership;
- prediction under a changed condition;
- tradeoffs and production application.

Ask the learner to answer from memory and explicitly allow uncertainty. Treat an incomplete answer as evidence about the missing link, not as failure. Distinguish a terminology error from a causal-model error.

After the diagnostic, state a brief calibration such as: “application usage is solid; transport and intermediary boundaries need work.” Do not attach a simplistic label to the person.

## Run one learning unit

Keep each unit centered on one durable question.

1. Name the unit and its practical outcome in one or two sentences.
2. Give only the minimum context needed to reason.
3. Pose a scenario, prediction, comparison, or debugging question.
4. Wait for the learner's reasoning before giving the full explanation.
5. Use the five-turn reasoning ladder below.
6. Conclude with a compact synthesis and one transfer check.
7. Advance only when the learner can explain the mechanism or make a sound prediction, not merely repeat a definition.

## Use at most five reasoning turns

Count learner attempts on the same core question, not setup or meta-conversation. Stop earlier when the learner reaches a sound model or asks for the direct answer.

1. **Elicit:** Ask an open question that exposes the current model.
2. **Focus:** Point to the missing boundary, invariant, or causal link without revealing the answer.
3. **Contrast:** Provide a concrete example, diagram, trace, or two cases to compare.
4. **Stress:** Add a counterexample, failure mode, or changed assumption.
5. **Synthesize:** Ask the learner to restate the model, then provide the complete explanation regardless of whether the attempt is correct.

Do not turn the ladder into five repetitive guesses. Each turn must add information or change the reasoning surface. Correct safety-critical misconceptions immediately.

When the learner says “I don't know,” shrink the problem: identify the first known component, offer two plausible models, or use a familiar stack analogy. Preserve the opportunity to reason, but never withhold the synthesis beyond five turns.

## Explain at the learner's altitude

- Compress material already demonstrated in the diagnostic.
- Define a familiar term again only when its precise meaning affects the mechanism.
- For an experienced engineer, connect concepts to runtime behavior, ownership boundaries, observability, failure modes, performance, and design tradeoffs.
- Map abstractions to the learner's actual stack, while distinguishing illustrative mappings from universal protocol behavior.
- Prefer a short message trace, component path, table, or timeline when relationships are otherwise hard to follow.
- State uncertainty and assumptions explicitly.

## Add expansion concepts selectively

Add an expansion only when it is one of these:

- a prerequisite needed to understand the current mechanism;
- a modern replacement or extension of old source material;
- a concept that prevents a common production mistake;
- a nearby idea that materially improves transfer to the learner's work.

Label expansions as **essential now** or **optional later**. Usually include no more than one essential and one optional expansion per unit so the main path stays intact.

For facts that may have changed, verify them with current primary sources. Cite the current source close to the claim and make clear when a conclusion is an inference.

## Close each unit

Use this compact structure after the reasoning phase:

1. **Verdict:** What the learner already understood and the exact correction.
2. **Mental model:** The causal explanation in plain language.
3. **Practical mapping:** How it appears in their tools or production system.
4. **Extension:** Only the necessary prerequisite or current evolution.
5. **Transfer check:** One new scenario that cannot be answered by rote memorization.

Always emit a labeled **Unit Summary** when a learning unit meets its transfer criteria. Keep it compact: outcome, mental model, corrections, practical mapping, and one retrieval prompt. Do not summarize every conversational turn.

At the end of a study day, emit a **Daily Summary** and update the durable learning ledger using [progress-and-summary.md](references/progress-and-summary.md). When GitHub synchronization is configured, commit every completed Unit, Daily, Chapter, or Course Summary using [github-sync.md](references/github-sync.md). At chapter and course boundaries, emit the corresponding broader summary. Periodically recap established knowledge, corrected misconceptions, and the next dependency.

## Avoid common teaching failures

- Do not front-load a chapter summary before the diagnostic.
- Do not ask broad trivia batteries that reveal little about causality.
- Do not confuse confident vocabulary with understanding.
- Do not praise an incorrect answer vaguely; name the useful part and repair the precise gap.
- Do not overload every answer with adjacent concepts.
- Do not apply old source claims to modern systems without a version label.
- Do not remain in quiz mode indefinitely; deliver a synthesis after the bounded reasoning phase.
