import { postApi } from "../../service/api";
import { auth, modal } from "../../stores/store.svelte";

export async function record(title: string, category: number, amount: number) {
  try {
    await postApi({
      path: "/transaction",
      data: { title, category, amount },
      token: auth.token
    })
    modal.title = "成功"
    modal.message = "登録しました"

  } catch (error: any) {
    const status = error.response?.status;
    modal.title = "エラー"
    switch (status) {
      case 400:
        modal.message = "入力形式が正しくありません"
        break;
      case 401:
        modal.message = "ログインしてください"
      default:
        modal.message = "システムエラーが発生しました"
    }
  }
}