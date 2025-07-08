# Exemples – IA & Data Analysis

## Exemple 1 : Analyse de sentiment d’une proposition
```python
from textblob import TextBlob
texte = "Cette proposition est innovante et bénéfique."
blob = TextBlob(texte)
print(blob.sentiment)
```

## Exemple 2 : Détection d’anomalies dans les votes
```python
from sklearn.ensemble import IsolationForest
import numpy as np
votes = np.array([[1, 0, 1], [1, 1, 1], [0, 0, 0], [10, 10, 10]])
model = IsolationForest().fit(votes)
print(model.predict(votes))
```

## Exemple 3 : Visualisation des résultats de vote
```python
import matplotlib.pyplot as plt
votes = [120, 80, 30]
labels = ['Pour', 'Contre', 'Abstention']
plt.pie(votes, labels=labels, autopct='%1.1f%%')
plt.show()
```

## Exemple 4 : Pipeline complet (prétraitement, scoring, visualisation)
Voir le GUIDE_PRACTIQUE.md pour un pipeline détaillé.
