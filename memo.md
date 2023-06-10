## ステータス/結果

### 待機中
- WE: Waiting for Execute
  - ワーカーにジョブが取得されていない状態
- WJ: Waiting for Judge
  - ワーカーがコードを実行中の状態

### ノンペナルティー(減点なし)
- IE: Internal Error
  - サーバー側のエラー
  - 問題の標準点数の1割が得点になる

### ペナルティー対象
- CE: Compile Error
  - コンパイルエラー
  - 標準点数の1割が引かれる
- MLE: Memory Limit Exceeded
  - メモリ制限超過
  - 標準点数の1割が引かれる
- TLE: Time Limit Exceeded
  - 実行時間超過
  - 標準点数の1割が引かれる
- RE: Runtime Error
  - 実行時エラー
  - 標準点数の1割が引かれる
- WA: Wrong Answer
  - 誤答
  - 標準点数の1割が引かれる
- ELE: Execution Limit Exceeded
  - 実行回数制限超過
  - その問題は提出できなくなる(テスト実行を除く, コンテスト開催時間中のみ)

