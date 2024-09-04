# MetaLife

macOS用のWebView版MetaLife。

## ビルド

```sh
go build -o bin/MetaLife.app/Contents/MacOS/MetaLife
```

## インストール

```sh
cp -R bin/MetaLife.app ~/Applications
```

## コード署名 (通知機能のため必要)

```sh
codesign -s - ~/Applications/MetaLife.app
```
