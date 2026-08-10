# Assessment rubric

Use this rubric to choose questions and adapt difficulty. Score the demonstrated model, not the learner as a person.

## Levels

| Level | Evidence | Next teaching move |
|---|---|---|
| 0 - Unfamiliar | Cannot identify the components or purpose | Supply a concrete example and establish the minimum vocabulary |
| 1 - Recognition | Recognizes terms but mixes roles or boundaries | Ask for a sequence and contrast adjacent components |
| 2 - Sequence | Can narrate the happy path but not why it works | Probe causal links, state changes, and ownership boundaries |
| 3 - Causal | Explains mechanisms and predicts ordinary behavior | Introduce failure modes, observability, and tradeoffs |
| 4 - Transfer | Applies the model to unfamiliar cases and weighs tradeoffs | Use design, debugging, optimization, or teaching-back tasks |

Different dimensions can be at different levels. For example, a learner may be level 3 in application semantics and level 1 in transport behavior.

## High-signal question shapes

Prefer questions that expose several dimensions at once:

- “Trace this operation from caller to application. Where are new connections or state created?”
- “Which component owns this decision, and what evidence would prove it?”
- “What changes if this intermediary is removed or this assumption becomes false?”
- “Predict the logs, wire behavior, or user-visible symptom of this failure.”
- “Compare two designs and name the condition under which each wins.”

Avoid questions that test only isolated definitions unless vocabulary itself is blocking progress.

## Calibration rules

- Two sound causal explanations: raise difficulty or advance.
- Correct sequence with weak causality: stay on the topic and probe one boundary.
- Several terms but incorrect ownership: use a component diagram or message trace.
- “I don't know”: reduce scope and anchor to familiar experience; do not restart the whole topic.
- Confident but false claim: request a prediction or observable consequence, then repair the model explicitly.
- Correct answer without explanation: ask for the mechanism or a changed-condition prediction.

## Transfer criteria

Consider a unit learned when the learner can do at least two of these:

- explain it in their own words;
- predict behavior when one condition changes;
- locate the owning layer or component;
- identify an observable signal for success or failure;
- apply it to a familiar production scenario;
- state one meaningful tradeoff or limitation.

Do not require perfect terminology when the causal model is sound. Correct terminology during the synthesis.
