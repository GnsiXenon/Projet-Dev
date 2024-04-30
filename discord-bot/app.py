import discord
from discord.ext import commands
from chatbot.Chatbot import respond
from gemini_api import generate_text
import base64
# from openai_api import openai_response

Tokenb64 = "TVRJeU9UY3lOalUyTnprMk5qVTNNall5TlEuR3ByRzAxLkVQVnNhWjR4Sk9vTE5pUXRXS2FYSl92ZDBzRnBRMWY5TWlkRG5r"


TOKEN = base64.b64decode(Tokenb64).decode("utf-8")


user = []

class MyBot(commands.Bot):
    def __init__(self):
        super().__init__(command_prefix="!", intents=discord.Intents.all())

    async def setup_hook(self) -> None:
        await self.tree.sync()

    async def on_message(self, message) :
        if message.author == self.user:
            return
        if message.author not in user:
            await message.channel.send("Pour l'utilisation du bot. Faite /start")
            return
        else :
            await message.channel.send(generate_text(message.content))


    async def on_ready(self) -> None:
        print("Ready")


bot = MyBot()


@bot.tree.command(name="start", description="start the conversation with the bot")
async def start(interaction: discord.Interaction):
    user.append(interaction.user)
    await interaction.response.send_message("Bonjour, je suis un bot de chat. Comment puis-je vous aider?")


@bot.tree.command(name="stop", description="stop the conversation with the bot")
async def stop(interaction: discord.Interaction):
    user.remove(interaction.user)
    await interaction.response.send_message("Au revoir!")

bot.run(TOKEN)






