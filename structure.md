# Structure Document
KOJSv6の構造のドキュメントです。  
2023 June 17th Poporon Network

## 大まかな概要
コードの提出が行われた際のフロー
```mermaid
sequenceDiagram
    actor User
    participant Backend
    participant Agent
    participant Worker
    
    User->>Backend: コードの提出
    Backend->>Backend: コードの保存
    Agent -->> Backend: 実行タスクの定期所得
    Note right of Backend: デフォルトでは0.5sごとに取得します 
    Agent ->> Worker: 実行
    Worker -->> Agent: 実行結果
    Agent -->> Backend: 実行結果
    Backend ->> Backend: 採点(正誤判定)
```
詳しくはAgent, Workerのリポジトリをご覧ください.
- [poporonnet/jkojs-agent (Agent)](https://github.com/poporonnet/jkojs-agent)
- [poporonnet/jkojs-worker (Worker)](https://github.com/poporonnet/jkojs-worker)
