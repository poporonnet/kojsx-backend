# パッケージ構造について

## 表記の方法
モジュール名::パッケージ名::Struct名

例:
- UserモジュールのModelパッケージのUserモデル: `User::Model::User`
- ContestモジュールのServiceパッケージのCreateContestService: `Contest::Service::CreateContestService`

## 構造について

- モジュールは複数のパッケージを持ちます．
  - **モジュール間で関数呼び出しは行ってはいけません(ToDo: InterModuleパッケージを作る)**
  - UserモジュールとContestモジュールが存在しています
- モジュールは以下のパッケージを持ちます
  - adaptor
    - controller: コントローラー実装
    - handler: ハンドラー実装
    - repository: model/repositoryで定義されたInterfaceの実装
  - model
    - (トップレベル): ドメインモデルの定義
    - repository: RepositoryのInterface定義
    - service: ドメインサービスの実装
  - service: Application Serviceの実装
  - util
    - 必要になったときだけ定義します
