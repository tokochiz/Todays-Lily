ToDoアプリ → 写真ギャラリーアプリ
- Todo一覧 → 写真一覧表示
- Todo追加 → 写真アップロード
- Todo編集 → 写真情報編集
- Todo削除 → 写真削除

ａ． 基本機能（ハンズオンベース）：

写真一覧表示（Todo一覧の実装を参考）
写真のCRUD操作（TodoのCRUD操作を参考）

b. 追加機能:

スライドショー機能
プロフィール画面
ランダム写真選択機能

dog-gallery/
├── frontend/     # Next.jsプロジェクト
└── backend/      # Goプロジェクト
src/
  components/
    photo/
      PhotoList.tsx        // 写真一覧表示
      PhotoSlideshow.tsx   // スライドショー
      PhotoCard.tsx        // 個別写真表示
      PhotoUpload.tsx      // 写真アップロード
    profile/
      ProfileView.tsx      // プロフィール表示
      ProfileEdit.tsx      // プロフィール編集
    common/
      RandomPhotoButton.tsx // ランダム写真選択
      Layout.tsx           // 共通レイアウト

/api/photos:
  get:    # 写真一覧取得
  post:   # 写真アップロード
  
/api/photos/{id}:
  get:    # 個別写真取得
  put:    # 写真情報更新
  delete: # 写真削除

/api/photos/random:
  get:    # ランダム写真取得

/api/profile:
  get:    # プロフィール取得
  put:    # プロフィール更新