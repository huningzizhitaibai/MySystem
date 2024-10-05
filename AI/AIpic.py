from sparkai.llm.llm import ChatSparkLLM, ChunkPrintHandler
from sparkai.core.messages import ChatMessage
import json
import sys

#星火认知大模型Spark Max的URL值，其他版本大模型URL值请前往文档（https://www.xfyun.cn/doc/spark/Web.html）查看
SPARKAI_URL = 'wss://spark-api.xf-yun.com/v3.5/chat'
#星火认知大模型调用秘钥信息，请前往讯飞开放平台控制台（https://console.xfyun.cn/services/bm35）查看
SPARKAI_APP_ID = '35cb75e8'
SPARKAI_API_SECRET = 'NjJiMDRlZDYxMjdmZmIxYzg2MDY4NTUx'
SPARKAI_API_KEY = 'b2d2d87c48314430324cf32fe6d3c3a0'
#星火认知大模型Spark Max的domain值，其他版本大模型domain值请前往文档（https://www.xfyun.cn/doc/spark/Web.html）查看
SPARKAI_DOMAIN = 'generalv3.5'

if __name__ == '__main__':
    spark = ChatSparkLLM(
        spark_api_url=SPARKAI_URL,
        spark_app_id=SPARKAI_APP_ID,
        spark_api_key=SPARKAI_API_KEY,
        spark_api_secret=SPARKAI_API_SECRET,
        spark_llm_domain=SPARKAI_DOMAIN,
        streaming=False,
    )
    messages = [ChatMessage(
        role="user",
        content=sys.argv[1]
    )]
    handler = ChunkPrintHandler()
    a = spark.generate([messages], callbacks=[handler])
    print(a.generations[0][0].text.encode('utf-8'))
    

    # print(sys.argv[1])
            
    
    
    # generations=[[ChatGeneration(text='你好！很高兴收到你的消息。有什么我可以帮助你的吗？', message=AIMessage(content='你好！很高兴收到你的消息。有什么我可以帮助你的吗？'))]] llm_output={'token_usage': {'question_tokens': 2, 'prompt_tokens': 2, 'completion_tokens': 13, 'total_tokens': 15}} run=[RunInfo(run_id=UUID('cd93f2ed-f04e-4487-82b9-d171d2afa15a'))]