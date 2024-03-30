import asyncio
from collections import Counter
import json
import time
import requests

counter = Counter()

async def count():
    resp = requests.get("http://10.0.0.42:8009/api/v1/namespaces/default/services/web-api-service:8080/proxy/count")
    data = json.loads(resp.text)
    count_num = data['count']
    app_id = data['app_id']
    counter[app_id] += 1

count_times = 1000

async def main():
    tasks = [count() for _ in range(count_times)]
    await asyncio.gather(*tasks)

# async def main():
#     for i in range(count_times):
#         await count()
#     print(counter)

if __name__ == '__main__': 
    begin = time.time()
    asyncio.run(main())
    end = time.time() 
    print(f"spend {end-begin}")
    print(counter)

    