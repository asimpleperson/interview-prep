# Riot Games Onsite — Predicted Questions & Follow-ups

Based on 11 interview experiences (2019–2023), mapped to your 6 confirmed interview sections.

---

## 1. Player & Mission Focus

**Evidence from:** Experiences 4, 5, 7, 10, 11

### Likely Questions

1. **"How do you connect your technical work to the player experience?"**
   - Follow-up: Give a specific example where you chose a harder technical path because it was better for users
   - Follow-up: How did you measure the player/user impact?

2. **"Why Riot? Why now?"**
   - Follow-up: What do you know about our mission?
   - Follow-up: How does your career trajectory lead here?
   - Follow-up: What games do you play? What do you notice about the player experience?

3. **"Tell me about a time you balanced user needs against business goals or technical constraints"**
   - Follow-up: Who pushed back? How did you navigate it?
   - Follow-up: What was the outcome for the user?

4. **"What do you know about your current company's culture, and what have you done to embody it?"**
   - Follow-up: How would those values translate to Riot?
   - Follow-up: What's different about how Riot operates vs your current company?
   - *(Experience 11 — this is a UNIQUE Riot question, confirmed by recruiter)*

5. **"Describe a decision you made where player/user empathy was the deciding factor"**
   - Follow-up: How did you build that empathy? Data, research, personal experience?
   - Follow-up: What would you have done differently?

### Prep Strategy

- Study https://www.riotgames.com/en/who-we-are/values — map each value to a personal story
- Reframe all past work as "player" work (your users ARE players)
- Have a genuine "why Riot" answer that goes beyond "I like games"
- Know your current company's culture/values AND how you've embodied them
- Expect a relaxed, conversational tone (Experience 7: director chat about Tencent-Riot relationship) — but it IS evaluated

---

## 2. Deep Tech

**Evidence from:** Experiences 1, 2, 5, 6, 7, 8, 9, 10, 11

### Likely Questions

1. **"Design a service-oriented architecture for [LoL feature]"**
   - Likely topics: LoL end-to-end flow, new feature addition, toxic content filtering, matchmaking
   - Follow-up: Walk me through the API contracts
   - Follow-up: What's your data model? Why that schema?
   - Follow-up: How does this scale to 150M+ players globally?
   - Follow-up: What happens when [component] fails?

2. **"How would you design a content/profanity filter service?"** *(Experiences 6, 7, 11 — recurring)*
   - Follow-up: How do you handle evasion? ("h-e-l-l-o", leetspeak, Unicode tricks)
   - Follow-up: "hello" vs "hellokitty" — how do you avoid false positives?
   - Follow-up: What algorithm would you use? *(They expect Rabin-Karp)*
   - Follow-up: Where in the system architecture does this service sit?
   - Follow-up: How do you update the blacklist without redeploying?
   - Follow-up: What about multilingual support?

3. **"Tell me about the hardest backend issue you've debugged in a live system"**
   - Follow-up: How did you diagnose it? What tools/observability did you use?
   - Follow-up: What was the blast radius? How did you mitigate while fixing?
   - Follow-up: What did you change to prevent recurrence?
   - Follow-up: How did you communicate status to stakeholders during the incident?

4. **"Design [security-adjacent system]"** — password reset, anti-cheat, rate limiter
   - Follow-up: Token generation — how do you ensure uniqueness and security?
   - Follow-up: What's your rate limiting strategy? (fixed window → sliding window → token bucket)
   - Follow-up: How do you handle distributed rate limiting across regions?
   - Follow-up: How do you prevent bad actors from gaming the system?

5. **Language/runtime internals** *(Experience 6 — depth round probes what you claim to know)*
   - If you say Go: goroutine scheduler, GC (tricolor mark-and-sweep), memory model, channel internals
   - Follow-up: How does compilation work in your language?
   - Follow-up: What's the memory footprint of [data structure you're using]?
   - Follow-up: How does concurrency work under the hood?

### Prep Strategy

- **Must-know algorithm: Rabin-Karp rolling hash** — they use it in production for content filtering
- Know Go internals cold: compilation, GC, goroutines, memory model (since that's your language)
- Prepare to go API-level deep: concrete endpoints, request/response schemas, error handling
- Practice the "conversational drill-down" format: they follow up on whatever you say, so don't mention anything you can't defend
- Two interviewers will pull you in different directions simultaneously

---

## 3. Broad Tech (Game of Life Review)

**Evidence from:** Experiences 10, 11 + your actual implementation

### Likely Questions About YOUR Implementation

1. **"Walk me through your overall design and why you structured it this way"**
   - Follow-up: Why a hashmap (`map[Position]Cell`) for the board? What are the trade-offs vs a 2D array?
   - Follow-up: What's the time complexity of a single `Tick()`?
   - Follow-up: What's the space complexity? How does it grow with live cells?

2. **"How do you handle integer overflow / boundary conditions?"**
   - Your `safeAdd()` function will be discussed
   - Follow-up: Why not use `math.MaxInt64` checks instead?
   - Follow-up: What happens to gliders approaching int64 boundaries?
   - Follow-up: Could you use unsigned integers instead? Trade-offs?

3. **"How does your solution scale?"**
   - Follow-up: What if the board has 10 million live cells? Bottlenecks?
   - Follow-up: How would you parallelize `Tick()`? *(map iteration isn't concurrent-safe in Go)*
   - Follow-up: What if the input is infinite/unbounded? *(Experience 11 — distributed sharding was the winning answer)*
   - Follow-up: Could you partition the board for multi-machine processing?

4. **"Tell me about your testing strategy"**
   - Follow-up: How do you verify correctness? Known patterns (blinker, glider)?
   - Follow-up: How would you test the overflow edge cases?
   - Follow-up: What's missing from your tests?

5. **"If you had more time, what would you improve?"**
   - Follow-up: How would you add visualization?
   - Follow-up: Could you support different rule sets (e.g., HighLife)?
   - Follow-up: How would you optimize for sparse vs dense boards?
   - Follow-up: Could you make this a distributed system? What's the coordination challenge?

6. **"Why Go?"**
   - Follow-up: How does Go's map implementation affect performance here?
   - Follow-up: Any garbage collection concerns with your approach? (many small map allocations per tick)

### Known Edge Cases They'll Probe

- Integer overflow at int64 boundaries (you handle this — be ready to explain)
- Sparse board performance (your hashmap approach is already good for this)
- Dense board performance (hashmap overhead vs array)
- Cells at coordinate (0,0) — boundary between negative and positive space
- Empty board / single cell / oscillators / still lifes

### Prep Strategy

- **Be fluent in your own code** — re-read `game_of_life.go` and `entity.go` before the interview
- Have a clear 2-minute walkthrough of the architecture ready
- Know the trade-off: hashmap (sparse-friendly, O(live_cells)) vs 2D array (dense-friendly, fixed memory)
- Prepare the "infinite board" extension answer: sharding by spatial partitioning, ghost cells at shard boundaries
- Know Go map internals since your implementation heavily depends on it

---

## 4. Innovative Problem Solving

**Evidence from:** Experiences 5, 6

### Likely Questions

1. **"Describe a time you had to dig into a root cause that wasn't obvious"**
   - Follow-up: What was your initial hypothesis? Why was it wrong?
   - Follow-up: What tools or methods helped you find the real cause?
   - Follow-up: How long did it take? What did you try that didn't work?
   - Follow-up: What did you learn about your assumptions?

2. **"Tell me about a time you solved a problem in an unconventional way"**
   - Follow-up: What was the standard approach? Why didn't you take it?
   - Follow-up: How did others react to your approach?
   - Follow-up: What was the risk? How did you validate it would work?
   - Follow-up: Would you do it again?

3. **"How do you approach a problem when you don't have all the information?"**
   - Follow-up: Give me a specific example
   - Follow-up: How did you decide what information was critical vs nice-to-have?
   - Follow-up: How did you adapt when new information contradicted your plan?

4. **"Tell me about a time you challenged an assumption that everyone else accepted"**
   - Follow-up: Why did you question it? What made you skeptical?
   - Follow-up: How did you convince others?
   - Follow-up: What happened as a result?

5. **"Describe an experiment or prototype you ran to validate an idea"**
   - Follow-up: What was your hypothesis?
   - Follow-up: How did you define success/failure criteria upfront?
   - Follow-up: What did you learn? Did it change your direction?

### Prep Strategy

- This round is about HOW you think, not WHAT you built
- Stories should show: curiosity → investigation → creative solution → practical result
- Include moments where you were wrong, adapted, and learned — they value growth mindset
- Show collaboration in problem-solving: "I talked to X to get a different perspective"
- Balance creativity with pragmatism — wild ideas are fine IF you can explain how you grounded them

---

## 5. Communicate & Collaborate

**Evidence from:** Experiences 3, 5, 6, 7, 8, 11

### Likely Questions

1. **"How do you work effectively with product managers / cross-functional partners?"**
   - Follow-up: Give a specific example of aligning on a technical decision with a PM
   - Follow-up: How did you handle it when you disagreed with their priorities?
   - Follow-up: How do you translate technical constraints into business terms?

2. **"Describe a time you had to give someone difficult/negative feedback"** *(Experience 6 — confirmed question)*
   - Follow-up: How did you prepare for the conversation?
   - Follow-up: How did they react?
   - Follow-up: What was the outcome? Did the behavior change?
   - Follow-up: What would you do differently?

3. **"How do you receive and act on feedback?"** *(Experience 5 — confirmed topic)*
   - Follow-up: Give an example of feedback that was hard to hear
   - Follow-up: What did you change as a result?
   - Follow-up: How do you solicit feedback proactively?

4. **"Tell me about a cross-team collaboration that was challenging"**
   - Follow-up: What made it difficult? (different priorities, timezone, culture?)
   - Follow-up: How did you build trust?
   - Follow-up: How did you ensure alignment when goals conflicted?
   - Follow-up: What was the outcome?

5. **"How do you explain complex technical concepts to non-technical stakeholders?"** *(Experience 7 — this is evaluated LIVE)*
   - Follow-up: The interviewer (possibly a TPM) will watch HOW you explain things during the entire round
   - Follow-up: Can you read the room and adjust if they don't follow?

6. **"What do you do when you can't finish a task on time?"** *(Experience 5 — confirmed)*
   - Follow-up: How early do you raise the flag?
   - Follow-up: How do you communicate scope/timeline trade-offs?
   - Follow-up: Example of re-scoping a project mid-flight

7. **"Have you had to hand off work to someone else? How did it go?"** *(Experience 6 — confirmed)*
   - Follow-up: How did you ensure continuity?
   - Follow-up: What documentation or knowledge transfer did you do?

### Prep Strategy

- This round evaluates the skill IN REAL TIME, not just your stories
- Speak clearly, check if interviewers are following, adjust your level
- For PM collaboration stories: show empathy for their perspective, not just "I convinced them"
- Feedback stories: show vulnerability and growth, not defensiveness
- Have a "failure/missed deadline" story ready — they want to see how you handle it gracefully

---

## 6. Engineering Leadership

**Evidence from:** Experiences 3, 5, 6, 8, 10, 11

### Likely Questions

1. **"How have you influenced technical direction beyond your immediate team?"**
   - Follow-up: What was the scope? How many teams/people were affected?
   - Follow-up: How did you get buy-in from engineers who disagreed?
   - Follow-up: What's still in place today? What didn't survive?

2. **"Describe how you've mentored or grown engineers around you"** *(Experience 6 — "Have you helped anyone at work?")*
   - Follow-up: What specific approaches did you use?
   - Follow-up: How did you measure their growth?
   - Follow-up: Give an example of someone who was struggling — what did you do?
   - Follow-up: How do you balance mentoring with your own IC work?

3. **"How do you make architectural decisions that outlast you?"**
   - Follow-up: Give a specific example of an architectural decision you drove
   - Follow-up: What trade-offs did you consider?
   - Follow-up: How did you document it and get alignment?
   - Follow-up: How has it held up? What would you change now?

4. **"Tell me about a time you led a team to success"** *(Experience 3 — confirmed, with two manager interviewers)*
   - Follow-up: What was your specific plan? Timeline?
   - Follow-up: How did you handle conflicts within the team?
   - Follow-up: What did you specifically do vs delegate?
   - Follow-up: What was the measurable outcome?
   - **Warning: They drill for EXTREME specifics. Vague = fail.**

5. **"How have you contributed to engineering culture?"**
   - Follow-up: What practices or processes did you introduce?
   - Follow-up: How did you drive adoption?
   - Follow-up: What pushback did you get?

6. **"What's the most recent thing you learned at work?"** *(Experience 6 — confirmed)*
   - Follow-up: How did you learn it? Why?
   - Follow-up: How has it changed how you work?
   - *(Tests growth mindset — even at staff level, you should be actively learning)*

7. **"Describe a time you worked under pressure"** *(Experience 6 — confirmed)*
   - Follow-up: How did you prioritize?
   - Follow-up: How did you keep the team calm?
   - Follow-up: What was the outcome?

### Prep Strategy

- **Two managers will interview you** (Experience 3) — they drill for specifics. Every story needs: WHO, WHAT, WHEN, concrete numbers, specific actions YOU took
- Frame as team success, not personal heroics (Riot's collaborative culture)
- Have 3-4 leadership stories with enough depth to survive 10+ minutes of follow-ups each
- Include conflict resolution stories with: the conflict, your role, specific actions, resolution, relationship after
- "What did you learn recently?" — have a genuine, specific answer ready (shows humility and growth)

---

## Cross-Cutting Patterns

### Two Interviewers in Most Rounds
Experiences consistently report two interviewers per round. This means:
- You'll be pulled in multiple directions — practice managing two threads
- One may play devil's advocate while the other probes depth
- You need to engage both, not just answer one while ignoring the other

### The "Drill-Down" Format
Riot interviews are conversational, not checklist-style. Whatever you say becomes the next question. This means:
- Never mention something you can't defend in depth
- Have 3 levels of detail prepared for every story/design choice
- It's OK to say "I don't know" — better than being exposed as bluffing

### Riot-Specific Culture Signals
- Player-first framing (even for infrastructure work)
- Humility and collaboration over individual heroism
- Know YOUR OWN company's culture (unique to Riot)
- Genuine passion for games/players (doesn't have to be LoL specifically)
- Growth mindset — what are you still learning?
