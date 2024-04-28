from openai import OpenAI

client = OpenAI(
    api_key="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"
)


def openai_response(prompt):
    response =client.chat.completions.create(
        model="gpt-3.5-turbo",
        messages=[
            {
                "role": "user",
                "content": prompt
            }
        ]
    )
    response_dict = response.get("choices")
    if response_dict and len(response_dict) > 0:
        prompt_response = response_dict[0].get("text")
    return prompt_response