# RestAPI - Item

## Get
| URI | Description | Attributes | Curl Example |
| --- | --- | --- | --- |
| /api/item | 아이템 가지고 오기 | id | `$ curl "http://192.168.219.104/api/item?id=5e24742f901da0498519f7a7"` |


## Post
| URI | Description | Attributes | Curl Example |
| --- | --- | --- | --- |
| /api/item | asset 등록하기 | itemtype, title, author, description, tags | `$ curl -H "Authorization: Basic <TOKEN>" -X POST \`<br>`-F "file1=@thumbnail.jpg;type=image/jpeg" \` <br>`-F "file2=@thumbnail.mov;type=video/quicktime" \`<br>`-F "file3=@data.abc;type=application/octet-stream" \` <br>`-F "itemtype=alembic" \` <br>`-F "title=abc restapi test" \` <br>`-F "author=dchecheb" \` <br>`-F "description=3" \` <br>`-F "tags=test" \` <br>`-F "attribute=key1:value1,key2:value2" \` <br>`http://192.168.219.104/api/item` |
| /api/usingrate | Using Rate 올리기 | id | `$ curl -X POST -d "id=5eaa5758eafdfd2dae3bb050" http://192.168.219.104/api/usingrate`


## Delete
| URI | Description | Attributes | Curl Example |
| --- | --- | --- | --- |
| /api/item | 삭제하기 | id | `curl -H "Authorization: Basic <Token>" -X DELETE "http://192.168.219.104/api/item?id=5ec37a67e048d951ee46a45a"`

## Python example

### GET

#### asset 가지고 오기 

```python
#!/usr/bin/python
#coding:utf-8
import urllib2
import json

request = urllib2.Request("http://192.168.219.104/api/item?id=5e24742f901da0498519f7a7")
result = urllib2.urlopen(request)
data = json.load(result)
print(data)
```

### POST

#### Asset 등록하기
```python
#!/usr/bin/python
#coding:utf-8
import requests, mimetypes, os

token="example.blar-blar"               
fileList=[  # Upload 할 File list                                                
    '/home/chaeyun.bae/cheche/dotori/examples/abc/abc_thumbnail.jpg',
    '/home/chaeyun.bae/cheche/dotori/examples/abc/abc_thumbnail.mov',
    '/home/chaeyun.bae/cheche/dotori/examples/abc/data.abc'
]
data = {    # 어셋 정보 입력
    'itemtype': (None, 'alembic'),
    'title': (None, 'train test'),
    'author': (None, 'dchecheb'),
    'description': (None, '3'),
    'tags': (None, 'test'),
    'attribute': (None, 'key1:value1,key2:value')
}

session = requests.Session()
session.headers.update({'Authorization': 'Basic ' + token })
i = 0
for file in fileList:
    key = 'file[{}]'.format(i)
    mimetype = mimetypes.guess_type(file)[0]    # mimetype 지정
    if not mimetype:
        mimetype = 'application/octet-stream'   # mimetype 인식 못할 경우 application/octet-stream을 default로 보냄
    data[key] = (os.path.basename(file), open(file, 'rb'), mimetype)
    i += 1

response = session.post('http://192.168.219.104/api/item', files=data)    # 전송
print(response.text)
```

#### 검색하기
```python
#!/usr/bin/python
#coding:utf-8
import urllib2
import urllib
import json

data = urllib.urlencode({"searchword":"나무"}) # 쿼리스트링 파라미터를 Encoding
request = urllib2.Request("http://192.168.0.9/api/search",data) 
result = urllib2.urlopen(request)
data = json.load(result)
print(data)
```