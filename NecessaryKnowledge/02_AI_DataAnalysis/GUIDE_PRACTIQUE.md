# Guide Pratique – IA & Data Analysis

Ce guide propose des tutoriels et bonnes pratiques pour appliquer l’IA et l’analyse de données à la gouvernance décentralisée.

## 1. Pipeline d’analyse de propositions (Python)
```python
import pandas as pd
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.ensemble import RandomForestClassifier
# Charger les propositions
proposals = pd.read_csv('propositions.csv')
# Vectorisation
X = TfidfVectorizer().fit_transform(proposals['texte'])
# Prédiction (exemple)
model = RandomForestClassifier().fit(X, proposals['label'])
pred = model.predict(X)
```

## 2. Détection de fraudes dans les votes
- Utiliser des modèles d’anomalie (Isolation Forest, AutoEncoder) sur les patterns de vote.
- Visualiser les résultats avec matplotlib ou seaborn.

## 3. Intégration d’un modèle NLP (spaCy)
```python
import spacy
nlp = spacy.load('fr_core_news_md')
doc = nlp("Analyse automatique des propositions citoyennes.")
for ent in doc.ents:
    print(ent.text, ent.label_)
```

## Bonnes pratiques
- Toujours documenter et versionner les jeux de données.
- Privilégier la transparence des modèles (expliquer les décisions IA).
- Respecter la vie privée (anonymisation, ZK-KYC).

## Aller plus loin
- Voir [Resources.md](./Resources.md) pour frameworks, datasets et benchmarks.
