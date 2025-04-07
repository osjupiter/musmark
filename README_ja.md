# musmark

**musmark**は、テンプレートとデータ、およびその出力を1つのファイルで管理できるシンプルなテンプレートツールです。コマンド、SQL、コードなど、さまざまな用途で利用でき、日常的な細かい定型作業の省力化に非常に効果的です。


**musmark**は、テンプレートとデータをMarkdown互換の単一ファイルで管理し、MustacheテンプレートとYAMLデータを使用してコマンドを生成するGoのCLIツールです。

---

## ✨ 特徴

- 🗂️ テンプレートとデータを1つのファイルで管理
- 🧠 Mustacheテンプレートエンジンのサポート
- 📄 YAMLフォーマットによる柔軟なデータ構造
- ⚡ Goによる単一バイナリの利便性

---

## 📦 インストール

```bash
go install github.com/osjupiter/musmark@latest
```

開発用にローカルビルドする場合:

```bash
go build -o musmark
```

---

## 🧪 使用方法

### 入力ファイルの例 (`collect_logs.md`)

入力ファイルは、テンプレートとデータをMarkdown形式で組み合わせます：

````markdown
template
===
```mustache
#!/bin/bash
{{#servers}}
rsync -az {{user}}@{{host}}:{{log_path}} /logs/{{name}}/
{{/servers}}
...
```

data
===
```yaml
env: production
servers:
  - name: app1
    host: app1.prod.example.com
    user: syslog
...
```
````

このテンプレートから以下のようなコマンドが生成されます：

```bash
#!/bin/bash
rsync -az syslog@app1.prod.example.com:/var/log/app/*.log /logs/app1/
rsync -az syslog@app2.prod.example.com:/var/log/app/*.log /logs/app2/
rsync -az dbadmin@db.prod.example.com:/var/log/mysql/*.log /logs/db/
tar czf /logs/all_production_20250407.tar.gz /logs/*/
```

ツールは既存の結果ブロックを賢く処理します：
- 結果ブロックが存在しない場合：末尾に新しく追加
- 結果ブロックが存在する場合：そのセクションのみを更新

### コマンドの実行

```bash
# テンプレートと結果を表示
musmark collect_logs.md

# コマンドを生成して直接実行
musmark -r collect_logs.md | bash

# 生成されるコマンドをプレビュー
musmark -r collect_logs.md
```


---

## 🛠 今後の展望

- [ ] JSONサポート
- [ ] Goテンプレートサポート
- [ ] 複数テンプレートセクションのサポート
- [ ] エラーハンドリングの強化

---

## 📄 ライセンス

MIT License
