import discord
from discord.ext import commands
from chatbot.Chatbot import respond
# from openai_api import openai_response

TOKEN = "MTIyOTcyNjU2Nzk2NjU3MjYyNQ.GmsRxV.xxQXy3delgfZmVMzRYH1ksntkb3iyudtX-ZSvc"



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
            await message.channel.send("For use the bot you need to start the conversation with the bot by typing !start")
            return
        else :
            await message.channel.send(respond(message.content))


    async def on_ready(self) -> None:
        print("Ready")


bot = MyBot()


@bot.tree.command(name="start", description="start the conversation with the bot")
async def start(interaction: discord.Interaction):
    user.append(interaction.user)
    await interaction.response.send_message("Hello, I am a bot. How can I help you?")


@bot.tree.command(name="stop", description="stop the conversation with the bot")
async def stop(interaction: discord.Interaction):
    user.remove(interaction.user)
    await interaction.response.send_message("Goodbye!")


        




bot.run(TOKEN)






