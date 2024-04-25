import json
import random
import fuzzywuzzy
from fuzzywuzzy import process


with open('chatbot/intents.json', 'r') as f:
    data = json.load(f)
def respond(message):
    """
    Fonction qui analyse le message de l'utilisateur et retourne une réponse appropriée.
    """
    
    # Convertir le message en minuscule pour une meilleure comparaison
    message = message.lower()

    # Parcourir les intentions pour trouver une correspondance
    for intent in data['intents']:
        patterns = intent['patterns']
        responses = intent['responses']
        
        # Calculer le score de similarité pour chaque modèle
        for pattern in patterns:
            similarity = fuzzywuzzy.fuzz.ratio(message, pattern)
            
            # Si la similarité est supérieure à un seuil prédéfini (par exemple, 80 %), considérer comme une correspondance
            if similarity > 90:
                # Si une correspondance est trouvée, sélectionner une réponse aléatoire
                response = random.choice(responses)
                return response

    # Si aucune correspondance n'est trouvée, renvoyer une réponse par défaut
    return "Désolé, je ne comprends pas ce que vous voulez dire."


