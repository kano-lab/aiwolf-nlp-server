# aiwolf-nlp-server

人狼知能コンテスト2024冬季 国内大会（自然言語部門） のゲームサーバです。  
従来の対戦接続システムでは、対戦接続システムが実行中のサーバに対してSSH接続を確立した後、TCP通信を行う必要がありましたが、新しいゲームサーバでは、WebSocketを使用して直接ゲームサーバと通信を行います。  

## ドキュメント

- [プロトコルの実装について](./doc/protocol.md)
- [ゲームロジックの実装について](./doc/logic.md)

## 実行方法

[Release](https://github.com/kano-lab/aiwolf-nlp-server/releases/latest) から環境に合ったバイナリをダウンロードしてください。  
[default.yml](https://github.com/kano-lab/aiwolf-nlp-server/releases/latest/download/default.yml) をダウンロードしてください。  
必要に応じて設定を変更したうえで、以下のコマンドを実行してください。  

```bash
./aiwolf-nlp-server -c default.yml
```
