# Money Craft
お金に関する教育用のWebアプリケーションです

# 目次
[概要](#概要)  
[構成図](#構成図)  
[主な機能](#主な機能)
[使用技術](#使用技術)  
[開発環境の構築手順](#開発環境の構築手順)
[ディレクトリ構成](#ディレクトリ構成)  
[参考](#参考)  

# 概要
このアプリケーションでは、家庭内で通貨を使えるようにし、為替や投資などの金融教育を提供することを目的とします。  
また、使用する通貨には金銭的価値はないため、金銭トラブルを回避できます。

# 構成図
![構成図](https://github.com/y-watagashi/MoneyCraft/assets/78391723/afeed9d4-50d1-46a3-b356-f220d3fcdbbd)  

# 主な機能
### 家庭内通貨
ユーザーは、家庭内で使える専用の通貨を作成し、家族や共同生活者とのお金のやり取りをアプリ上で簡単に管理できます

### 仮想の金融体験
為替変動や投資シミュレーションなどを通じて、ユーザーは仮想の状況で金融取引を体験できます。
そのため、現実での金銭を失うリスクをとることなく、教育コンテンツとして使用できます。

  
# 使用技術
- Vue
- Vuetify
- Go
- gin
- MySQL


# 開発環境の構築手順
```shell
$ git clone 
$ cd MoneyCraft
$ docker-compose up --build -d
```
注：ポートが重複しているさいは、起動しない可能性があります。
その際は、他のコンテナを停止するもしくは、`docker-compose.yml`を編集してポートを変更してください。


# ディレクトリ構成
```plain text
.
├── README.md
├── docker-compose.yml
├── backend
│   ├── Dockerfile
│   ├── entrypoint.sh
│   └── src                   // バックエンドソース
├── database
│   ├── Dockerfile
│   ├── db                    // DBの内容を保管
│   └── my.cnf
└── frontend
    ├── Dockerfile
    ├── package-lock.json
    ├── package.json
    ├── public
    │   └── favicon.ico
    ├── src
    │   ├── App.vue
    │   ├── assets           // アセット
    │   ├── components       // コンポーネント
    │   ├── main.js
    │   ├── router
    │   │   └── index.js
    │   ├── stores
    │   │   └── counter.js
    │   └── views            // 各ページ
    └── vite.config.js

```

# 参考
[Docker ドキュメント（日本語）](https://matsuand.github.io/docs.docker.jp.onthefly/): Dockerの公式ドキュメントは、初学者から上級者まで包括的な情報が提供されています。
