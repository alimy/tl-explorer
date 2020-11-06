# [TL-Schema Explorer](https://schema.horner.tj)
Telegram tl-schema explorer is An AngularJS app to search and view the Telegram API TL-schema in a beautiful way.

![](https://i.imgur.com/akrelfR.png)

#### Usage
```bash
% go get github.com/alimy/tl-explorer
% tl-explorer
Listening in [:8080]. Please open http://localhost:8080 in browser to enjoy yourself.
```

#### Deploy by Docker
```bash
% docker pull bitriory/tl-explorer:latest
% docker run -d --restart=always -p 8080:8080 bitriory/tl-explorer
```
Then you can browser at http://localhost:8080.

**Notes**: The assets copy from [schema.tl](https://github.com/tjhorner/schema.tl). Thanks TJ Horner.


