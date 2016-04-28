# ivq
ivq, an inverted-index(term) query engine.


### Possible Queries:

1. **Term Query**: matches a precise term.
2. **Wildcard Query**: portions of term left unspecified, e.g: c*t = cat, cot, cant
3. **Prefix Query**: end of term left unspecified, e.g: ca* = cat, camp, cad
4. **Fuzzy Query**: matches an imprecise term.
    - Accommodate typo or spellings.
    - Generally more costly to perform
5. **Phrase Query**: matches multiple terms in sequence/proximity (need term position)
    - Occurred in certain order. (phrases)
    - Within proximity with each other. (proximity)
6. **Range Query**: matches an alphabetical or numerical range of terms
7. **Boolean Query**: logically composite other queries.
