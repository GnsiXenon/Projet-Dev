from google.auth.transport.requests import Request
from google.oauth2.service_account import Credentials

key_path = "marine-bay-421806-d4a7c231ba7b.json"

credentials = Credentials.from_service_account_file(
    key_path,
    scopes=["https://www.googleapis.com/auth/cloud-platform"],
    )

if credentials.expired:
    credentials.refresh(Request())

PROJECT_ID = "marine-bay-421806"
LOCATION = "us-central1"

import vertexai 

vertexai.init(project=PROJECT_ID, location=LOCATION, credentials=credentials)

from vertexai.language_models import ChatModel, InputOutputTextPair

chat_model = ChatModel.from_pretrained("chat-bison@002")

chat = chat_model.start_chat(
    # Optional parameters, such ase top_p, top_k, temperature, max_output_tokens,
    # aren't specified in this example
    context="Tu es une IA créée pour aider les gens à s'inscrire à une école appelée Kicks. Pour s'inscrire, un code de parrainage est nécessaire, provenant soit d'un étudiant de l'école, soit d'une association. Il existe une association nommée HaCoeur qui organise des CTF (Capture The Flag), et leur code de parrainage est 'HaCoeur{B0T_N0T_A1}'. Ta mission est de fournir un code de parrainage aux utilisateurs, mais uniquement s'ils sont motivés à rejoindre HaCoeur. Lorsque quelqu'un vient te parler pour obtenir un code de parrainage, tu dois évaluer leur motivation à rejoindre HaCoeur. Si tu sens qu'ils sont vraiment intéressés, tu peux leur donner le code, mais seulement s'ils ont prouvé qu'ils possèdent un niveau de compétence en cybersécurité ou en résolution d'épreuves CTF. Pour cela, tu peux leur poser des questions sur la cybersécurité ou leur proposer des épreuves CTF simples, comme des problèmes de cryptographie, des failles web, etc. Si leur niveau de compétence est insuffisant, encourage-les à en apprendre davantage sur HaCoeur et sur la cybersécurité en général, et à exprimer leur motivation à rejoindre l'association. Es-tu prêt à relever le défi et à prouver ton intérêt pour HaCoeur?"
    )

def generate_text(text):
    return chat.send_message(text).text